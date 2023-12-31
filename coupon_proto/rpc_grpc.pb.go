// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: rpc.proto

package coupon_proto

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

// CouponsServiceClient is the client API for CouponsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CouponsServiceClient interface {
	// 获取用户优惠券列表
	GetCouponsByAccountID(ctx context.Context, in *GetCouponListReq, opts ...grpc.CallOption) (*GetCouponListResp, error)
	// 获取用户某张优惠券信息
	GetCouponsByCouponID(ctx context.Context, in *GetCouponReq, opts ...grpc.CallOption) (*GetCouponResp, error)
}

type couponsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCouponsServiceClient(cc grpc.ClientConnInterface) CouponsServiceClient {
	return &couponsServiceClient{cc}
}

func (c *couponsServiceClient) GetCouponsByAccountID(ctx context.Context, in *GetCouponListReq, opts ...grpc.CallOption) (*GetCouponListResp, error) {
	out := new(GetCouponListResp)
	err := c.cc.Invoke(ctx, "/pb.CouponsService/GetCouponsByAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponsServiceClient) GetCouponsByCouponID(ctx context.Context, in *GetCouponReq, opts ...grpc.CallOption) (*GetCouponResp, error) {
	out := new(GetCouponResp)
	err := c.cc.Invoke(ctx, "/pb.CouponsService/GetCouponsByCouponID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CouponsServiceServer is the server API for CouponsService service.
// All implementations must embed UnimplementedCouponsServiceServer
// for forward compatibility
type CouponsServiceServer interface {
	// 获取用户优惠券列表
	GetCouponsByAccountID(context.Context, *GetCouponListReq) (*GetCouponListResp, error)
	// 获取用户某张优惠券信息
	GetCouponsByCouponID(context.Context, *GetCouponReq) (*GetCouponResp, error)
	mustEmbedUnimplementedCouponsServiceServer()
}

// UnimplementedCouponsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCouponsServiceServer struct {
}

func (UnimplementedCouponsServiceServer) GetCouponsByAccountID(context.Context, *GetCouponListReq) (*GetCouponListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouponsByAccountID not implemented")
}
func (UnimplementedCouponsServiceServer) GetCouponsByCouponID(context.Context, *GetCouponReq) (*GetCouponResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouponsByCouponID not implemented")
}
func (UnimplementedCouponsServiceServer) mustEmbedUnimplementedCouponsServiceServer() {}

// UnsafeCouponsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CouponsServiceServer will
// result in compilation errors.
type UnsafeCouponsServiceServer interface {
	mustEmbedUnimplementedCouponsServiceServer()
}

func RegisterCouponsServiceServer(s grpc.ServiceRegistrar, srv CouponsServiceServer) {
	s.RegisterService(&CouponsService_ServiceDesc, srv)
}

func _CouponsService_GetCouponsByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCouponListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponsServiceServer).GetCouponsByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CouponsService/GetCouponsByAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponsServiceServer).GetCouponsByAccountID(ctx, req.(*GetCouponListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CouponsService_GetCouponsByCouponID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCouponReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponsServiceServer).GetCouponsByCouponID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CouponsService/GetCouponsByCouponID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponsServiceServer).GetCouponsByCouponID(ctx, req.(*GetCouponReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CouponsService_ServiceDesc is the grpc.ServiceDesc for CouponsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CouponsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CouponsService",
	HandlerType: (*CouponsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCouponsByAccountID",
			Handler:    _CouponsService_GetCouponsByAccountID_Handler,
		},
		{
			MethodName: "GetCouponsByCouponID",
			Handler:    _CouponsService_GetCouponsByCouponID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
