// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mmesh

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ModelRuntimeClient is the client API for ModelRuntime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelRuntimeClient interface {
	// Load a model, return when model is fully loaded.
	// Include size of loaded model in response if no additional cost.
	// A gRPC error code of PRECONDITION_FAILED or INVALID_ARGUMENT
	// should be returned if no attempt to load the model was made
	// (so can be sure that no space remains used).
	// Note that the RPC may be cancelled by model-mesh prior to completion,
	// after which an unloadModel call will immediately be sent for the same model.
	// To avoid state inconsistency and "leaking" memory, implementors should
	// ensure that this case is properly handled, i.e. that the model doesn't
	// remain loaded after returning successfully from this unloadModel call.
	LoadModel(ctx context.Context, in *LoadModelRequest, opts ...grpc.CallOption) (*LoadModelResponse, error)
	// Unload a previously loaded (or failed) model. Return when model
	// is fully unloaded, or immediately if not found/loaded.
	UnloadModel(ctx context.Context, in *UnloadModelRequest, opts ...grpc.CallOption) (*UnloadModelResponse, error)
	// Predict size of not-yet-loaded model - must return almost immediately.
	// Should not perform expensive computation or remote lookups.
	// Should be a conservative estimate.
	PredictModelSize(ctx context.Context, in *PredictModelSizeRequest, opts ...grpc.CallOption) (*PredictModelSizeResponse, error)
	// Calculate size (memory consumption) of currently-loaded model
	ModelSize(ctx context.Context, in *ModelSizeRequest, opts ...grpc.CallOption) (*ModelSizeResponse, error)
	// Provide basic runtime status and parameters; called only during startup.
	// Before returning a READY status, implementations should check for and
	// purge any/all currently-loaded models. Since this is only called during
	// startup, there should very rarely be any, but if there are it implies
	// the model-mesh container restarted unexpectedly and such a purge must
	// be done to ensure continued consistency of state and avoid over-committing
	// resources.
	RuntimeStatus(ctx context.Context, in *RuntimeStatusRequest, opts ...grpc.CallOption) (*RuntimeStatusResponse, error)
}

type modelRuntimeClient struct {
	cc grpc.ClientConnInterface
}

func NewModelRuntimeClient(cc grpc.ClientConnInterface) ModelRuntimeClient {
	return &modelRuntimeClient{cc}
}

func (c *modelRuntimeClient) LoadModel(ctx context.Context, in *LoadModelRequest, opts ...grpc.CallOption) (*LoadModelResponse, error) {
	out := new(LoadModelResponse)
	err := c.cc.Invoke(ctx, "/mmesh.ModelRuntime/loadModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRuntimeClient) UnloadModel(ctx context.Context, in *UnloadModelRequest, opts ...grpc.CallOption) (*UnloadModelResponse, error) {
	out := new(UnloadModelResponse)
	err := c.cc.Invoke(ctx, "/mmesh.ModelRuntime/unloadModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRuntimeClient) PredictModelSize(ctx context.Context, in *PredictModelSizeRequest, opts ...grpc.CallOption) (*PredictModelSizeResponse, error) {
	out := new(PredictModelSizeResponse)
	err := c.cc.Invoke(ctx, "/mmesh.ModelRuntime/predictModelSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRuntimeClient) ModelSize(ctx context.Context, in *ModelSizeRequest, opts ...grpc.CallOption) (*ModelSizeResponse, error) {
	out := new(ModelSizeResponse)
	err := c.cc.Invoke(ctx, "/mmesh.ModelRuntime/modelSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRuntimeClient) RuntimeStatus(ctx context.Context, in *RuntimeStatusRequest, opts ...grpc.CallOption) (*RuntimeStatusResponse, error) {
	out := new(RuntimeStatusResponse)
	err := c.cc.Invoke(ctx, "/mmesh.ModelRuntime/runtimeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelRuntimeServer is the server API for ModelRuntime service.
// All implementations should embed UnimplementedModelRuntimeServer
// for forward compatibility
type ModelRuntimeServer interface {
	// Load a model, return when model is fully loaded.
	// Include size of loaded model in response if no additional cost.
	// A gRPC error code of PRECONDITION_FAILED or INVALID_ARGUMENT
	// should be returned if no attempt to load the model was made
	// (so can be sure that no space remains used).
	// Note that the RPC may be cancelled by model-mesh prior to completion,
	// after which an unloadModel call will immediately be sent for the same model.
	// To avoid state inconsistency and "leaking" memory, implementors should
	// ensure that this case is properly handled, i.e. that the model doesn't
	// remain loaded after returning successfully from this unloadModel call.
	LoadModel(context.Context, *LoadModelRequest) (*LoadModelResponse, error)
	// Unload a previously loaded (or failed) model. Return when model
	// is fully unloaded, or immediately if not found/loaded.
	UnloadModel(context.Context, *UnloadModelRequest) (*UnloadModelResponse, error)
	// Predict size of not-yet-loaded model - must return almost immediately.
	// Should not perform expensive computation or remote lookups.
	// Should be a conservative estimate.
	PredictModelSize(context.Context, *PredictModelSizeRequest) (*PredictModelSizeResponse, error)
	// Calculate size (memory consumption) of currently-loaded model
	ModelSize(context.Context, *ModelSizeRequest) (*ModelSizeResponse, error)
	// Provide basic runtime status and parameters; called only during startup.
	// Before returning a READY status, implementations should check for and
	// purge any/all currently-loaded models. Since this is only called during
	// startup, there should very rarely be any, but if there are it implies
	// the model-mesh container restarted unexpectedly and such a purge must
	// be done to ensure continued consistency of state and avoid over-committing
	// resources.
	RuntimeStatus(context.Context, *RuntimeStatusRequest) (*RuntimeStatusResponse, error)
}

// UnimplementedModelRuntimeServer should be embedded to have forward compatible implementations.
type UnimplementedModelRuntimeServer struct {
}

func (UnimplementedModelRuntimeServer) LoadModel(context.Context, *LoadModelRequest) (*LoadModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadModel not implemented")
}
func (UnimplementedModelRuntimeServer) UnloadModel(context.Context, *UnloadModelRequest) (*UnloadModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnloadModel not implemented")
}
func (UnimplementedModelRuntimeServer) PredictModelSize(context.Context, *PredictModelSizeRequest) (*PredictModelSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PredictModelSize not implemented")
}
func (UnimplementedModelRuntimeServer) ModelSize(context.Context, *ModelSizeRequest) (*ModelSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelSize not implemented")
}
func (UnimplementedModelRuntimeServer) RuntimeStatus(context.Context, *RuntimeStatusRequest) (*RuntimeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RuntimeStatus not implemented")
}

// UnsafeModelRuntimeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelRuntimeServer will
// result in compilation errors.
type UnsafeModelRuntimeServer interface {
	mustEmbedUnimplementedModelRuntimeServer()
}

func RegisterModelRuntimeServer(s grpc.ServiceRegistrar, srv ModelRuntimeServer) {
	s.RegisterService(&ModelRuntime_ServiceDesc, srv)
}

func _ModelRuntime_LoadModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRuntimeServer).LoadModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mmesh.ModelRuntime/loadModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRuntimeServer).LoadModel(ctx, req.(*LoadModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRuntime_UnloadModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnloadModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRuntimeServer).UnloadModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mmesh.ModelRuntime/unloadModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRuntimeServer).UnloadModel(ctx, req.(*UnloadModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRuntime_PredictModelSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictModelSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRuntimeServer).PredictModelSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mmesh.ModelRuntime/predictModelSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRuntimeServer).PredictModelSize(ctx, req.(*PredictModelSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRuntime_ModelSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRuntimeServer).ModelSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mmesh.ModelRuntime/modelSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRuntimeServer).ModelSize(ctx, req.(*ModelSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRuntime_RuntimeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuntimeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRuntimeServer).RuntimeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mmesh.ModelRuntime/runtimeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRuntimeServer).RuntimeStatus(ctx, req.(*RuntimeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelRuntime_ServiceDesc is the grpc.ServiceDesc for ModelRuntime service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelRuntime_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mmesh.ModelRuntime",
	HandlerType: (*ModelRuntimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "loadModel",
			Handler:    _ModelRuntime_LoadModel_Handler,
		},
		{
			MethodName: "unloadModel",
			Handler:    _ModelRuntime_UnloadModel_Handler,
		},
		{
			MethodName: "predictModelSize",
			Handler:    _ModelRuntime_PredictModelSize_Handler,
		},
		{
			MethodName: "modelSize",
			Handler:    _ModelRuntime_ModelSize_Handler,
		},
		{
			MethodName: "runtimeStatus",
			Handler:    _ModelRuntime_RuntimeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model-runtime.proto",
}
