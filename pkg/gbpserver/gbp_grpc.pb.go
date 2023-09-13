// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: gbp.proto

package gbpserver

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

const (
	GBP_ListObjects_FullMethodName = "/gbpserver.GBP/ListObjects"
	GBP_ListVTEPs_FullMethodName   = "/gbpserver.GBP/ListVTEPs"
	GBP_GetSnapShot_FullMethodName = "/gbpserver.GBP/GetSnapShot"
)

// GBPClient is the client API for GBP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GBPClient interface {
	// Obtains the objects currently in the policy database as a stream
	ListObjects(ctx context.Context, in *Version, opts ...grpc.CallOption) (GBP_ListObjectsClient, error)
	ListVTEPs(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*VTEPList, error)
	GetSnapShot(ctx context.Context, in *VTEP, opts ...grpc.CallOption) (*ObjectList, error)
}

type gBPClient struct {
	cc grpc.ClientConnInterface
}

func NewGBPClient(cc grpc.ClientConnInterface) GBPClient {
	return &gBPClient{cc}
}

func (c *gBPClient) ListObjects(ctx context.Context, in *Version, opts ...grpc.CallOption) (GBP_ListObjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &GBP_ServiceDesc.Streams[0], GBP_ListObjects_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gBPListObjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GBP_ListObjectsClient interface {
	Recv() (*GBPOperation, error)
	grpc.ClientStream
}

type gBPListObjectsClient struct {
	grpc.ClientStream
}

func (x *gBPListObjectsClient) Recv() (*GBPOperation, error) {
	m := new(GBPOperation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gBPClient) ListVTEPs(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*VTEPList, error) {
	out := new(VTEPList)
	err := c.cc.Invoke(ctx, GBP_ListVTEPs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gBPClient) GetSnapShot(ctx context.Context, in *VTEP, opts ...grpc.CallOption) (*ObjectList, error) {
	out := new(ObjectList)
	err := c.cc.Invoke(ctx, GBP_GetSnapShot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GBPServer is the server API for GBP service.
// All implementations must embed UnimplementedGBPServer
// for forward compatibility
type GBPServer interface {
	// Obtains the objects currently in the policy database as a stream
	ListObjects(*Version, GBP_ListObjectsServer) error
	ListVTEPs(context.Context, *EmptyMsg) (*VTEPList, error)
	GetSnapShot(context.Context, *VTEP) (*ObjectList, error)
	mustEmbedUnimplementedGBPServer()
}

// UnimplementedGBPServer must be embedded to have forward compatible implementations.
type UnimplementedGBPServer struct {
}

func (UnimplementedGBPServer) ListObjects(*Version, GBP_ListObjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListObjects not implemented")
}
func (UnimplementedGBPServer) ListVTEPs(context.Context, *EmptyMsg) (*VTEPList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVTEPs not implemented")
}
func (UnimplementedGBPServer) GetSnapShot(context.Context, *VTEP) (*ObjectList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnapShot not implemented")
}
func (UnimplementedGBPServer) mustEmbedUnimplementedGBPServer() {}

// UnsafeGBPServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GBPServer will
// result in compilation errors.
type UnsafeGBPServer interface {
	mustEmbedUnimplementedGBPServer()
}

func RegisterGBPServer(s grpc.ServiceRegistrar, srv GBPServer) {
	s.RegisterService(&GBP_ServiceDesc, srv)
}

func _GBP_ListObjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Version)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GBPServer).ListObjects(m, &gBPListObjectsServer{stream})
}

type GBP_ListObjectsServer interface {
	Send(*GBPOperation) error
	grpc.ServerStream
}

type gBPListObjectsServer struct {
	grpc.ServerStream
}

func (x *gBPListObjectsServer) Send(m *GBPOperation) error {
	return x.ServerStream.SendMsg(m)
}

func _GBP_ListVTEPs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GBPServer).ListVTEPs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GBP_ListVTEPs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GBPServer).ListVTEPs(ctx, req.(*EmptyMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _GBP_GetSnapShot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VTEP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GBPServer).GetSnapShot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GBP_GetSnapShot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GBPServer).GetSnapShot(ctx, req.(*VTEP))
	}
	return interceptor(ctx, in, info, handler)
}

// GBP_ServiceDesc is the grpc.ServiceDesc for GBP service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GBP_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gbpserver.GBP",
	HandlerType: (*GBPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListVTEPs",
			Handler:    _GBP_ListVTEPs_Handler,
		},
		{
			MethodName: "GetSnapShot",
			Handler:    _GBP_GetSnapShot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListObjects",
			Handler:       _GBP_ListObjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gbp.proto",
}
