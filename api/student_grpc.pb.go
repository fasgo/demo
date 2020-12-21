// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	// +POST /demo/students
	Add(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error)
	// +DELETE /demo/students/:sno
	Del(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error)
	// +PUT /demo/students/:sno
	Upd(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error)
	// +websocket /demo/student/ws
	// +GET /demo/students/:sno
	Get(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error)
	// +GET /demo/students
	All(ctx context.Context, in *AllReq, opts ...grpc.CallOption) (*AllRsp, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) Add(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/api.StudentService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) Del(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/api.StudentService/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) Upd(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/api.StudentService/Upd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) Get(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/api.StudentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) All(ctx context.Context, in *AllReq, opts ...grpc.CallOption) (*AllRsp, error) {
	out := new(AllRsp)
	err := c.cc.Invoke(ctx, "/api.StudentService/All", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	// +POST /demo/students
	Add(context.Context, *Student) (*Student, error)
	// +DELETE /demo/students/:sno
	Del(context.Context, *Student) (*Student, error)
	// +PUT /demo/students/:sno
	Upd(context.Context, *Student) (*Student, error)
	// +websocket /demo/student/ws
	// +GET /demo/students/:sno
	Get(context.Context, *Student) (*Student, error)
	// +GET /demo/students
	All(context.Context, *AllReq) (*AllRsp, error)
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) Add(context.Context, *Student) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedStudentServiceServer) Del(context.Context, *Student) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedStudentServiceServer) Upd(context.Context, *Student) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upd not implemented")
}
func (UnimplementedStudentServiceServer) Get(context.Context, *Student) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStudentServiceServer) All(context.Context, *AllReq) (*AllRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	s.RegisterService(&_StudentService_serviceDesc, srv)
}

func _StudentService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Student)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StudentService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Add(ctx, req.(*Student))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Student)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StudentService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Del(ctx, req.(*Student))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_Upd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Student)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Upd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StudentService/Upd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Upd(ctx, req.(*Student))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Student)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StudentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Get(ctx, req.(*Student))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StudentService/All",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).All(ctx, req.(*AllReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _StudentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _StudentService_Add_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _StudentService_Del_Handler,
		},
		{
			MethodName: "Upd",
			Handler:    _StudentService_Upd_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StudentService_Get_Handler,
		},
		{
			MethodName: "All",
			Handler:    _StudentService_All_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student.proto",
}
