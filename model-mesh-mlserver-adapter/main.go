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
package main

import (
	"fmt"
	"net"
	"os"

	"github.com/kserve/modelmesh-runtime-adapter/internal/proto/mmesh"
	"github.com/kserve/modelmesh-runtime-adapter/model-mesh-mlserver-adapter/server"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func main() {
	log := zap.New(zap.UseDevMode(true))
	log = log.WithName("MLServer Adapter")

	adapterConfig, err := server.GetAdapterConfigurationFromEnv(log)
	if err != nil {
		log.Error(err, "Error reading configuration")
		os.Exit(1)
	}
	log.Info("Starting MLServer Adapter", "adapter_config", adapterConfig)

	MLServer := server.NewMLServerAdapterServer(adapterConfig.MLServerPort, adapterConfig, log)
	defer MLServer.Conn.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", adapterConfig.Port))
	if err != nil {
		log.Error(err, "*** Adapter failed to listen port")
		os.Exit(1)
	}

	log.Info("Adapter will run at port", "port", adapterConfig.Port, "MLServer port", adapterConfig.MLServerPort)

	var opts []grpc.ServerOption

	log.Info("Adapter gRPC Server Registered...")
	grpcServer := grpc.NewServer(opts...)
	mmesh.RegisterModelRuntimeServer(grpcServer, MLServer)
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Error(err, "*** Adapter terminated with error ")
	} else {
		log.Info("*** Adapter terminated")
	}
}
