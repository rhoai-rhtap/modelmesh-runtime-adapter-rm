// Copyright 2021 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-logr/logr"
	"github.com/kserve/modelmesh-runtime-adapter/internal/proto/mmesh"
)

const (
	StateManagerChannelLength = 25
)

type iPullerServer interface {
	loadModel(ctx context.Context, req *mmesh.LoadModelRequest) (*mmesh.LoadModelResponse, error)
	unloadModel(ctx context.Context, req *mmesh.UnloadModelRequest) (*mmesh.UnloadModelResponse, error)
}

type modelStateManager struct {
	log      logr.Logger
	s        iPullerServer
	data     map[string]*modelData
	requests chan *request
	results  chan *result
}

type modelData struct {
	m        iPullerServer
	refCount int
	requests chan *request
	results  chan *result
}

type grpcRequest interface {
	GetModelId() string
}

type request struct {
	modelId string
	ctx     context.Context

	//One of LoadModelRequest or UnloadModelRequest
	grpcRequest grpcRequest
	c           chan *result
}

type result struct {
	modelId string

	//One of LoadModelResponse or UnloadModelResponse
	grpcResponse interface{}

	//The error returned from executing the request
	err error

	// The channel for communicating the status of this individual
	// req/res pair
	c chan *result
}

func newModelStateManager(log logr.Logger, s *PullerServer) (*modelStateManager, error) {
	m := &modelStateManager{
		s: s, data: make(map[string]*modelData),
		log:      log,
		requests: make(chan *request, StateManagerChannelLength),
		results:  make(chan *result, StateManagerChannelLength),
	}
	go m.execute()
	return m, nil
}

func (m *modelStateManager) submitRequest(ctx context.Context, req grpcRequest) (interface{}, error) {
	c := make(chan *result)
	select {
	case m.requests <- &request{req.GetModelId(), ctx, req, c}:
		res := <-c
		return res.grpcResponse, res.err
	default:
		return nil, fmt.Errorf("Unable to send load/unload model request")
	}
}

func (m *modelStateManager) loadModel(ctx context.Context, req *mmesh.LoadModelRequest) (*mmesh.LoadModelResponse, error) {
	resp, err := m.submitRequest(ctx, req)
	return resp.(*mmesh.LoadModelResponse), err
}

func (m *modelStateManager) unloadModel(ctx context.Context, req *mmesh.UnloadModelRequest) (*mmesh.UnloadModelResponse, error) {
	resp, err := m.submitRequest(ctx, req)
	return resp.(*mmesh.UnloadModelResponse), err
}

func (m *modelStateManager) execute() {
	reqChan := m.requests
	for {
		select {
		case req, ok := <-reqChan:
			if !ok {
				reqChan = nil
				break // from select statement
			}
			var data *modelData
			if d, ok := m.data[req.modelId]; ok {
				data = d
			} else {
				data = &modelData{
					m:        m.s,
					refCount: 0,
					requests: make(chan *request, StateManagerChannelLength),
					results:  m.results,
				}
				m.data[req.modelId] = data
				go data.execute(&m.log)
			}

			data.refCount += 1
			data.requests <- req

		case result := <-m.results:
			result.c <- result
			if data, ok := m.data[result.modelId]; ok {
				data.refCount -= 1

				if data.refCount <= 0 {
					close(data.requests)
					delete(m.data, result.modelId)
				}
			}
		}
		if reqChan == nil && len(m.data) == 0 {
			close(m.results)
			break // exit goroutine
		}
	}
}

func (d *modelData) execute(log *logr.Logger) {
	req, ok := <-d.requests
	for ; ok; req, ok = <-d.requests {
		var res interface{}
		var err error
		switch request := req.grpcRequest.(type) {
		case *mmesh.LoadModelRequest:
			res, err = d.m.loadModel(req.ctx, request)
		case *mmesh.UnloadModelRequest:
			res, err = d.m.unloadModel(req.ctx, request)
		default:
			(*log).Error(nil, "Unrecognized request type", "type", fmt.Sprintf("%T", req.grpcRequest))
			continue
		}
		d.results <- &result{req.modelId, res, err, req.c}
	}
}

func (m *modelStateManager) unloadAll() error {
	pullerServer := m.s.(*PullerServer)
	modelids, err := pullerServer.puller.ListModels()
	if err != nil {
		m.log.Error(err, "Unable to list the models for unloading")
		return err
	}
	for _, modelid := range modelids {
		excluded := false
		for _, exclude := range PurgeExcludePrefixes {
			if strings.HasPrefix(modelid, exclude) {
				excluded = true
				break
			}
		}

		if !excluded {
			unloadReq := &mmesh.UnloadModelRequest{ModelId: modelid}
			ctx := context.Background()
			_, err := pullerServer.modelRuntimeClient.UnloadModel(ctx, unloadReq)
			if err != nil {
				// When an error occurs unloading a model, abort all model
				// unloading with the error
				m.log.Error(err, "Error requesting unload of model")
				return err
			}

			filePath := path.Join(pullerServer.puller.PullerConfig.RootModelDir, modelid)
			m.log.Info("Purging", "filepath", filePath)
			os.RemoveAll(filePath)
		} else {
			m.log.Info("Skipping purge because it is excluded from deletion", "filename", modelid)
		}
	}

	return nil
}
