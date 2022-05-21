// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: basic.proto

package pb

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

// BasicClient is the client API for Basic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BasicClient interface {
	//FileDetail //文件详情
	FileDetail(ctx context.Context, in *FileDetailReq, opts ...grpc.CallOption) (*FileDetailResp, error)
	//FileList 文件列表
	FileList(ctx context.Context, in *FileListReq, opts ...grpc.CallOption) (*FileListResp, error)
	//addFile 新增文件
	AddFile(ctx context.Context, in *AddFileReq, opts ...grpc.CallOption) (*AddFileResp, error)
}

type basicClient struct {
	cc grpc.ClientConnInterface
}

func NewBasicClient(cc grpc.ClientConnInterface) BasicClient {
	return &basicClient{cc}
}

func (c *basicClient) FileDetail(ctx context.Context, in *FileDetailReq, opts ...grpc.CallOption) (*FileDetailResp, error) {
	out := new(FileDetailResp)
	err := c.cc.Invoke(ctx, "/pb.Basic/FileDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicClient) FileList(ctx context.Context, in *FileListReq, opts ...grpc.CallOption) (*FileListResp, error) {
	out := new(FileListResp)
	err := c.cc.Invoke(ctx, "/pb.Basic/FileList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicClient) AddFile(ctx context.Context, in *AddFileReq, opts ...grpc.CallOption) (*AddFileResp, error) {
	out := new(AddFileResp)
	err := c.cc.Invoke(ctx, "/pb.Basic/addFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasicServer is the server API for Basic service.
// All implementations must embed UnimplementedBasicServer
// for forward compatibility
type BasicServer interface {
	//FileDetail //文件详情
	FileDetail(context.Context, *FileDetailReq) (*FileDetailResp, error)
	//FileList 文件列表
	FileList(context.Context, *FileListReq) (*FileListResp, error)
	//addFile 新增文件
	AddFile(context.Context, *AddFileReq) (*AddFileResp, error)
	mustEmbedUnimplementedBasicServer()
}

// UnimplementedBasicServer must be embedded to have forward compatible implementations.
type UnimplementedBasicServer struct {
}

func (UnimplementedBasicServer) FileDetail(context.Context, *FileDetailReq) (*FileDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileDetail not implemented")
}
func (UnimplementedBasicServer) FileList(context.Context, *FileListReq) (*FileListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileList not implemented")
}
func (UnimplementedBasicServer) AddFile(context.Context, *AddFileReq) (*AddFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFile not implemented")
}
func (UnimplementedBasicServer) mustEmbedUnimplementedBasicServer() {}

// UnsafeBasicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BasicServer will
// result in compilation errors.
type UnsafeBasicServer interface {
	mustEmbedUnimplementedBasicServer()
}

func RegisterBasicServer(s grpc.ServiceRegistrar, srv BasicServer) {
	s.RegisterService(&Basic_ServiceDesc, srv)
}

func _Basic_FileDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).FileDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Basic/FileDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).FileDetail(ctx, req.(*FileDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Basic_FileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).FileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Basic/FileList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).FileList(ctx, req.(*FileListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Basic_AddFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).AddFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Basic/addFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).AddFile(ctx, req.(*AddFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Basic_ServiceDesc is the grpc.ServiceDesc for Basic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Basic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Basic",
	HandlerType: (*BasicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FileDetail",
			Handler:    _Basic_FileDetail_Handler,
		},
		{
			MethodName: "FileList",
			Handler:    _Basic_FileList_Handler,
		},
		{
			MethodName: "addFile",
			Handler:    _Basic_AddFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basic.proto",
}