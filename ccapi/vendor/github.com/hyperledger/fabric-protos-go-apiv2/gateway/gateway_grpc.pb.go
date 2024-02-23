// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gateway/gateway.proto

package gateway

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	// The Endorse service passes a proposed transaction to the gateway in order to
	// obtain sufficient endorsement.
	// The gateway will determine the endorsement plan for the requested chaincode and
	// forward to the appropriate peers for endorsement. It will return to the client a
	// prepared transaction in the form of an Envelope message as defined
	// in common/common.proto. The client must sign the contents of this envelope
	// before invoking the Submit service.
	Endorse(ctx context.Context, in *EndorseRequest, opts ...grpc.CallOption) (*EndorseResponse, error)
	// The Submit service will process the prepared transaction returned from Endorse service
	// once it has been signed by the client. It will wait for the transaction to be submitted to the
	// ordering service but the client must invoke the CommitStatus service to wait for the transaction
	// to be committed.
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
	// The CommitStatus service will indicate whether a prepared transaction previously submitted to
	// the Submit service has been committed. It will wait for the commit to occur if it hasn’t already
	// committed.
	CommitStatus(ctx context.Context, in *SignedCommitStatusRequest, opts ...grpc.CallOption) (*CommitStatusResponse, error)
	// The Evaluate service passes a proposed transaction to the gateway in order to invoke the
	// transaction function and return the result to the client. No ledger updates are made.
	// The gateway will select an appropriate peer to query based on block height and load.
	Evaluate(ctx context.Context, in *EvaluateRequest, opts ...grpc.CallOption) (*EvaluateResponse, error)
	// The ChaincodeEvents service supplies a stream of responses, each containing all the events emitted by the
	// requested chaincode for a specific block. The streamed responses are ordered by ascending block number. Responses
	// are only returned for blocks that contain the requested events, while blocks not containing any of the requested
	// events are skipped.
	ChaincodeEvents(ctx context.Context, in *SignedChaincodeEventsRequest, opts ...grpc.CallOption) (Gateway_ChaincodeEventsClient, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) Endorse(ctx context.Context, in *EndorseRequest, opts ...grpc.CallOption) (*EndorseResponse, error) {
	out := new(EndorseResponse)
	err := c.cc.Invoke(ctx, "/gateway.Gateway/Endorse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, "/gateway.Gateway/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CommitStatus(ctx context.Context, in *SignedCommitStatusRequest, opts ...grpc.CallOption) (*CommitStatusResponse, error) {
	out := new(CommitStatusResponse)
	err := c.cc.Invoke(ctx, "/gateway.Gateway/CommitStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Evaluate(ctx context.Context, in *EvaluateRequest, opts ...grpc.CallOption) (*EvaluateResponse, error) {
	out := new(EvaluateResponse)
	err := c.cc.Invoke(ctx, "/gateway.Gateway/Evaluate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) ChaincodeEvents(ctx context.Context, in *SignedChaincodeEventsRequest, opts ...grpc.CallOption) (Gateway_ChaincodeEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Gateway_ServiceDesc.Streams[0], "/gateway.Gateway/ChaincodeEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayChaincodeEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gateway_ChaincodeEventsClient interface {
	Recv() (*ChaincodeEventsResponse, error)
	grpc.ClientStream
}

type gatewayChaincodeEventsClient struct {
	grpc.ClientStream
}

func (x *gatewayChaincodeEventsClient) Recv() (*ChaincodeEventsResponse, error) {
	m := new(ChaincodeEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations should embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	// The Endorse service passes a proposed transaction to the gateway in order to
	// obtain sufficient endorsement.
	// The gateway will determine the endorsement plan for the requested chaincode and
	// forward to the appropriate peers for endorsement. It will return to the client a
	// prepared transaction in the form of an Envelope message as defined
	// in common/common.proto. The client must sign the contents of this envelope
	// before invoking the Submit service.
	Endorse(context.Context, *EndorseRequest) (*EndorseResponse, error)
	// The Submit service will process the prepared transaction returned from Endorse service
	// once it has been signed by the client. It will wait for the transaction to be submitted to the
	// ordering service but the client must invoke the CommitStatus service to wait for the transaction
	// to be committed.
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	// The CommitStatus service will indicate whether a prepared transaction previously submitted to
	// the Submit service has been committed. It will wait for the commit to occur if it hasn’t already
	// committed.
	CommitStatus(context.Context, *SignedCommitStatusRequest) (*CommitStatusResponse, error)
	// The Evaluate service passes a proposed transaction to the gateway in order to invoke the
	// transaction function and return the result to the client. No ledger updates are made.
	// The gateway will select an appropriate peer to query based on block height and load.
	Evaluate(context.Context, *EvaluateRequest) (*EvaluateResponse, error)
	// The ChaincodeEvents service supplies a stream of responses, each containing all the events emitted by the
	// requested chaincode for a specific block. The streamed responses are ordered by ascending block number. Responses
	// are only returned for blocks that contain the requested events, while blocks not containing any of the requested
	// events are skipped.
	ChaincodeEvents(*SignedChaincodeEventsRequest, Gateway_ChaincodeEventsServer) error
}

// UnimplementedGatewayServer should be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) Endorse(context.Context, *EndorseRequest) (*EndorseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Endorse not implemented")
}
func (UnimplementedGatewayServer) Submit(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedGatewayServer) CommitStatus(context.Context, *SignedCommitStatusRequest) (*CommitStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitStatus not implemented")
}
func (UnimplementedGatewayServer) Evaluate(context.Context, *EvaluateRequest) (*EvaluateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Evaluate not implemented")
}
func (UnimplementedGatewayServer) ChaincodeEvents(*SignedChaincodeEventsRequest, Gateway_ChaincodeEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method ChaincodeEvents not implemented")
}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_Endorse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndorseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Endorse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.Gateway/Endorse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Endorse(ctx, req.(*EndorseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.Gateway/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CommitStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedCommitStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CommitStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.Gateway/CommitStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CommitStatus(ctx, req.(*SignedCommitStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Evaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Evaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.Gateway/Evaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Evaluate(ctx, req.(*EvaluateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_ChaincodeEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SignedChaincodeEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServer).ChaincodeEvents(m, &gatewayChaincodeEventsServer{stream})
}

type Gateway_ChaincodeEventsServer interface {
	Send(*ChaincodeEventsResponse) error
	grpc.ServerStream
}

type gatewayChaincodeEventsServer struct {
	grpc.ServerStream
}

func (x *gatewayChaincodeEventsServer) Send(m *ChaincodeEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Endorse",
			Handler:    _Gateway_Endorse_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _Gateway_Submit_Handler,
		},
		{
			MethodName: "CommitStatus",
			Handler:    _Gateway_CommitStatus_Handler,
		},
		{
			MethodName: "Evaluate",
			Handler:    _Gateway_Evaluate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChaincodeEvents",
			Handler:       _Gateway_ChaincodeEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gateway/gateway.proto",
}