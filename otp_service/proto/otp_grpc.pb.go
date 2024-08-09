// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: otp_service/proto/otp.proto

package proto

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
	OTPService_GenerateOTP_FullMethodName = "/otp.OTPService/GenerateOTP"
	OTPService_VerifyOTP_FullMethodName   = "/otp.OTPService/VerifyOTP"
)

// OTPServiceClient is the client API for OTPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OTPServiceClient interface {
	GenerateOTP(ctx context.Context, in *GenerateOTPRequest, opts ...grpc.CallOption) (*GenerateOTPResponse, error)
	VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error)
}

type oTPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOTPServiceClient(cc grpc.ClientConnInterface) OTPServiceClient {
	return &oTPServiceClient{cc}
}

func (c *oTPServiceClient) GenerateOTP(ctx context.Context, in *GenerateOTPRequest, opts ...grpc.CallOption) (*GenerateOTPResponse, error) {
	out := new(GenerateOTPResponse)
	err := c.cc.Invoke(ctx, OTPService_GenerateOTP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oTPServiceClient) VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error) {
	out := new(VerifyOTPResponse)
	err := c.cc.Invoke(ctx, OTPService_VerifyOTP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OTPServiceServer is the server API for OTPService service.
// All implementations must embed UnimplementedOTPServiceServer
// for forward compatibility
type OTPServiceServer interface {
	GenerateOTP(context.Context, *GenerateOTPRequest) (*GenerateOTPResponse, error)
	VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error)
	mustEmbedUnimplementedOTPServiceServer()
}

// UnimplementedOTPServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOTPServiceServer struct {
}

func (UnimplementedOTPServiceServer) GenerateOTP(context.Context, *GenerateOTPRequest) (*GenerateOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateOTP not implemented")
}
func (UnimplementedOTPServiceServer) VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOTP not implemented")
}
func (UnimplementedOTPServiceServer) mustEmbedUnimplementedOTPServiceServer() {}

// UnsafeOTPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OTPServiceServer will
// result in compilation errors.
type UnsafeOTPServiceServer interface {
	mustEmbedUnimplementedOTPServiceServer()
}

func RegisterOTPServiceServer(s grpc.ServiceRegistrar, srv OTPServiceServer) {
	s.RegisterService(&OTPService_ServiceDesc, srv)
}

func _OTPService_GenerateOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OTPServiceServer).GenerateOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OTPService_GenerateOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OTPServiceServer).GenerateOTP(ctx, req.(*GenerateOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OTPService_VerifyOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OTPServiceServer).VerifyOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OTPService_VerifyOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OTPServiceServer).VerifyOTP(ctx, req.(*VerifyOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OTPService_ServiceDesc is the grpc.ServiceDesc for OTPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OTPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "otp.OTPService",
	HandlerType: (*OTPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateOTP",
			Handler:    _OTPService_GenerateOTP_Handler,
		},
		{
			MethodName: "VerifyOTP",
			Handler:    _OTPService_VerifyOTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "otp_service/proto/otp.proto",
}
