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
// Code generated by MockGen. DO NOT EDIT.
// Source: model-mesh-mlserver-adapter/generated/mlserver/dataplane.pb.go

// Package mock_mlserver is a generated GoMock package.
package mock_mlserver

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	inference "github.com/kserve/modelmesh-runtime-adapter/internal/proto/mlserver/dataplane"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockisInferParameter_ParameterChoice is a mock of isInferParameter_ParameterChoice interface
type MockisInferParameter_ParameterChoice struct {
	ctrl     *gomock.Controller
	recorder *MockisInferParameter_ParameterChoiceMockRecorder
}

// MockisInferParameter_ParameterChoiceMockRecorder is the mock recorder for MockisInferParameter_ParameterChoice
type MockisInferParameter_ParameterChoiceMockRecorder struct {
	mock *MockisInferParameter_ParameterChoice
}

// NewMockisInferParameter_ParameterChoice creates a new mock instance
func NewMockisInferParameter_ParameterChoice(ctrl *gomock.Controller) *MockisInferParameter_ParameterChoice {
	mock := &MockisInferParameter_ParameterChoice{ctrl: ctrl}
	mock.recorder = &MockisInferParameter_ParameterChoiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisInferParameter_ParameterChoice) EXPECT() *MockisInferParameter_ParameterChoiceMockRecorder {
	return m.recorder
}

// isInferParameter_ParameterChoice mocks base method
func (m *MockisInferParameter_ParameterChoice) isInferParameter_ParameterChoice() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isInferParameter_ParameterChoice")
}

// isInferParameter_ParameterChoice indicates an expected call of isInferParameter_ParameterChoice
func (mr *MockisInferParameter_ParameterChoiceMockRecorder) isInferParameter_ParameterChoice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isInferParameter_ParameterChoice", reflect.TypeOf((*MockisInferParameter_ParameterChoice)(nil).isInferParameter_ParameterChoice))
}

// MockGRPCInferenceServiceClient is a mock of GRPCInferenceServiceClient interface
type MockGRPCInferenceServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCInferenceServiceClientMockRecorder
}

// MockGRPCInferenceServiceClientMockRecorder is the mock recorder for MockGRPCInferenceServiceClient
type MockGRPCInferenceServiceClientMockRecorder struct {
	mock *MockGRPCInferenceServiceClient
}

// NewMockGRPCInferenceServiceClient creates a new mock instance
func NewMockGRPCInferenceServiceClient(ctrl *gomock.Controller) *MockGRPCInferenceServiceClient {
	mock := &MockGRPCInferenceServiceClient{ctrl: ctrl}
	mock.recorder = &MockGRPCInferenceServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGRPCInferenceServiceClient) EXPECT() *MockGRPCInferenceServiceClientMockRecorder {
	return m.recorder
}

// ServerLive mocks base method
func (m *MockGRPCInferenceServiceClient) ServerLive(ctx context.Context, in *inference.ServerLiveRequest, opts ...grpc.CallOption) (*inference.ServerLiveResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ServerLive", varargs...)
	ret0, _ := ret[0].(*inference.ServerLiveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerLive indicates an expected call of ServerLive
func (mr *MockGRPCInferenceServiceClientMockRecorder) ServerLive(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerLive", reflect.TypeOf((*MockGRPCInferenceServiceClient)(nil).ServerLive), varargs...)
}

// ServerReady mocks base method
func (m *MockGRPCInferenceServiceClient) ServerReady(ctx context.Context, in *inference.ServerReadyRequest, opts ...grpc.CallOption) (*inference.ServerReadyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ServerReady", varargs...)
	ret0, _ := ret[0].(*inference.ServerReadyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerReady indicates an expected call of ServerReady
func (mr *MockGRPCInferenceServiceClientMockRecorder) ServerReady(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerReady", reflect.TypeOf((*MockGRPCInferenceServiceClient)(nil).ServerReady), varargs...)
}

// ModelReady mocks base method
func (m *MockGRPCInferenceServiceClient) ModelReady(ctx context.Context, in *inference.ModelReadyRequest, opts ...grpc.CallOption) (*inference.ModelReadyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ModelReady", varargs...)
	ret0, _ := ret[0].(*inference.ModelReadyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelReady indicates an expected call of ModelReady
func (mr *MockGRPCInferenceServiceClientMockRecorder) ModelReady(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelReady", reflect.TypeOf((*MockGRPCInferenceServiceClient)(nil).ModelReady), varargs...)
}

// ServerMetadata mocks base method
func (m *MockGRPCInferenceServiceClient) ServerMetadata(ctx context.Context, in *inference.ServerMetadataRequest, opts ...grpc.CallOption) (*inference.ServerMetadataResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ServerMetadata", varargs...)
	ret0, _ := ret[0].(*inference.ServerMetadataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerMetadata indicates an expected call of ServerMetadata
func (mr *MockGRPCInferenceServiceClientMockRecorder) ServerMetadata(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerMetadata", reflect.TypeOf((*MockGRPCInferenceServiceClient)(nil).ServerMetadata), varargs...)
}

// ModelMetadata mocks base method
func (m *MockGRPCInferenceServiceClient) ModelMetadata(ctx context.Context, in *inference.ModelMetadataRequest, opts ...grpc.CallOption) (*inference.ModelMetadataResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ModelMetadata", varargs...)
	ret0, _ := ret[0].(*inference.ModelMetadataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelMetadata indicates an expected call of ModelMetadata
func (mr *MockGRPCInferenceServiceClientMockRecorder) ModelMetadata(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelMetadata", reflect.TypeOf((*MockGRPCInferenceServiceClient)(nil).ModelMetadata), varargs...)
}

// ModelInfer mocks base method
func (m *MockGRPCInferenceServiceClient) ModelInfer(ctx context.Context, in *inference.ModelInferRequest, opts ...grpc.CallOption) (*inference.ModelInferResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ModelInfer", varargs...)
	ret0, _ := ret[0].(*inference.ModelInferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelInfer indicates an expected call of ModelInfer
func (mr *MockGRPCInferenceServiceClientMockRecorder) ModelInfer(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelInfer", reflect.TypeOf((*MockGRPCInferenceServiceClient)(nil).ModelInfer), varargs...)
}

// MockGRPCInferenceServiceServer is a mock of GRPCInferenceServiceServer interface
type MockGRPCInferenceServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCInferenceServiceServerMockRecorder
}

// MockGRPCInferenceServiceServerMockRecorder is the mock recorder for MockGRPCInferenceServiceServer
type MockGRPCInferenceServiceServerMockRecorder struct {
	mock *MockGRPCInferenceServiceServer
}

// NewMockGRPCInferenceServiceServer creates a new mock instance
func NewMockGRPCInferenceServiceServer(ctrl *gomock.Controller) *MockGRPCInferenceServiceServer {
	mock := &MockGRPCInferenceServiceServer{ctrl: ctrl}
	mock.recorder = &MockGRPCInferenceServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGRPCInferenceServiceServer) EXPECT() *MockGRPCInferenceServiceServerMockRecorder {
	return m.recorder
}

// ServerLive mocks base method
func (m *MockGRPCInferenceServiceServer) ServerLive(arg0 context.Context, arg1 *inference.ServerLiveRequest) (*inference.ServerLiveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerLive", arg0, arg1)
	ret0, _ := ret[0].(*inference.ServerLiveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerLive indicates an expected call of ServerLive
func (mr *MockGRPCInferenceServiceServerMockRecorder) ServerLive(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerLive", reflect.TypeOf((*MockGRPCInferenceServiceServer)(nil).ServerLive), arg0, arg1)
}

// ServerReady mocks base method
func (m *MockGRPCInferenceServiceServer) ServerReady(arg0 context.Context, arg1 *inference.ServerReadyRequest) (*inference.ServerReadyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerReady", arg0, arg1)
	ret0, _ := ret[0].(*inference.ServerReadyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerReady indicates an expected call of ServerReady
func (mr *MockGRPCInferenceServiceServerMockRecorder) ServerReady(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerReady", reflect.TypeOf((*MockGRPCInferenceServiceServer)(nil).ServerReady), arg0, arg1)
}

// ModelReady mocks base method
func (m *MockGRPCInferenceServiceServer) ModelReady(arg0 context.Context, arg1 *inference.ModelReadyRequest) (*inference.ModelReadyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelReady", arg0, arg1)
	ret0, _ := ret[0].(*inference.ModelReadyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelReady indicates an expected call of ModelReady
func (mr *MockGRPCInferenceServiceServerMockRecorder) ModelReady(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelReady", reflect.TypeOf((*MockGRPCInferenceServiceServer)(nil).ModelReady), arg0, arg1)
}

// ServerMetadata mocks base method
func (m *MockGRPCInferenceServiceServer) ServerMetadata(arg0 context.Context, arg1 *inference.ServerMetadataRequest) (*inference.ServerMetadataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerMetadata", arg0, arg1)
	ret0, _ := ret[0].(*inference.ServerMetadataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerMetadata indicates an expected call of ServerMetadata
func (mr *MockGRPCInferenceServiceServerMockRecorder) ServerMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerMetadata", reflect.TypeOf((*MockGRPCInferenceServiceServer)(nil).ServerMetadata), arg0, arg1)
}

// ModelMetadata mocks base method
func (m *MockGRPCInferenceServiceServer) ModelMetadata(arg0 context.Context, arg1 *inference.ModelMetadataRequest) (*inference.ModelMetadataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelMetadata", arg0, arg1)
	ret0, _ := ret[0].(*inference.ModelMetadataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelMetadata indicates an expected call of ModelMetadata
func (mr *MockGRPCInferenceServiceServerMockRecorder) ModelMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelMetadata", reflect.TypeOf((*MockGRPCInferenceServiceServer)(nil).ModelMetadata), arg0, arg1)
}

// ModelInfer mocks base method
func (m *MockGRPCInferenceServiceServer) ModelInfer(arg0 context.Context, arg1 *inference.ModelInferRequest) (*inference.ModelInferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelInfer", arg0, arg1)
	ret0, _ := ret[0].(*inference.ModelInferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelInfer indicates an expected call of ModelInfer
func (mr *MockGRPCInferenceServiceServerMockRecorder) ModelInfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelInfer", reflect.TypeOf((*MockGRPCInferenceServiceServer)(nil).ModelInfer), arg0, arg1)
}
