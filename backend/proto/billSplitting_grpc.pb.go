// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: proto/billSplitting.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BillSplitting_CreateExpenseSummaryChart_FullMethodName = "/proto.BillSplitting/CreateExpenseSummaryChart"
	BillSplitting_GetAuthToken_FullMethodName              = "/proto.BillSplitting/GetAuthToken"
	BillSplitting_CreateLineGroup_FullMethodName           = "/proto.BillSplitting/CreateLineGroup"
	BillSplitting_AddMembership_FullMethodName             = "/proto.BillSplitting/AddMembership"
	BillSplitting_GetLineGroup_FullMethodName              = "/proto.BillSplitting/GetLineGroup"
	BillSplitting_GetMembership_FullMethodName             = "/proto.BillSplitting/GetMembership"
	BillSplitting_CreateExpense_FullMethodName             = "/proto.BillSplitting/CreateExpense"
	BillSplitting_CreateTrendingImage_FullMethodName       = "/proto.BillSplitting/CreateTrendingImage"
)

// BillSplittingClient is the client API for BillSplitting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillSplittingClient interface {
	CreateExpenseSummaryChart(ctx context.Context, in *CreateExpenseSummaryChartRequest, opts ...grpc.CallOption) (*CreateExpenseSummaryChartResponse, error)
	GetAuthToken(ctx context.Context, in *GetAuthTokenRequest, opts ...grpc.CallOption) (*GetAuthTokenResponse, error)
	CreateLineGroup(ctx context.Context, in *CreateLineGroupRequest, opts ...grpc.CallOption) (*CreateLineGroupResponse, error)
	AddMembership(ctx context.Context, in *AddMembershipRequest, opts ...grpc.CallOption) (*AddMembershipResponse, error)
	GetLineGroup(ctx context.Context, in *GetLineGroupRequest, opts ...grpc.CallOption) (*GetLineGroupResponse, error)
	GetMembership(ctx context.Context, in *GetMembershipRequest, opts ...grpc.CallOption) (*GetMembershipResponse, error)
	CreateExpense(ctx context.Context, in *CreateExpenseRequest, opts ...grpc.CallOption) (*CreateExpenseResponse, error)
	CreateTrendingImage(ctx context.Context, in *CreateTrendingImageRequest, opts ...grpc.CallOption) (*CreateTrendingImageResponse, error)
}

type billSplittingClient struct {
	cc grpc.ClientConnInterface
}

func NewBillSplittingClient(cc grpc.ClientConnInterface) BillSplittingClient {
	return &billSplittingClient{cc}
}

func (c *billSplittingClient) CreateExpenseSummaryChart(ctx context.Context, in *CreateExpenseSummaryChartRequest, opts ...grpc.CallOption) (*CreateExpenseSummaryChartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateExpenseSummaryChartResponse)
	err := c.cc.Invoke(ctx, BillSplitting_CreateExpenseSummaryChart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billSplittingClient) GetAuthToken(ctx context.Context, in *GetAuthTokenRequest, opts ...grpc.CallOption) (*GetAuthTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAuthTokenResponse)
	err := c.cc.Invoke(ctx, BillSplitting_GetAuthToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billSplittingClient) CreateLineGroup(ctx context.Context, in *CreateLineGroupRequest, opts ...grpc.CallOption) (*CreateLineGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLineGroupResponse)
	err := c.cc.Invoke(ctx, BillSplitting_CreateLineGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billSplittingClient) AddMembership(ctx context.Context, in *AddMembershipRequest, opts ...grpc.CallOption) (*AddMembershipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddMembershipResponse)
	err := c.cc.Invoke(ctx, BillSplitting_AddMembership_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billSplittingClient) GetLineGroup(ctx context.Context, in *GetLineGroupRequest, opts ...grpc.CallOption) (*GetLineGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLineGroupResponse)
	err := c.cc.Invoke(ctx, BillSplitting_GetLineGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billSplittingClient) GetMembership(ctx context.Context, in *GetMembershipRequest, opts ...grpc.CallOption) (*GetMembershipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMembershipResponse)
	err := c.cc.Invoke(ctx, BillSplitting_GetMembership_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billSplittingClient) CreateExpense(ctx context.Context, in *CreateExpenseRequest, opts ...grpc.CallOption) (*CreateExpenseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateExpenseResponse)
	err := c.cc.Invoke(ctx, BillSplitting_CreateExpense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billSplittingClient) CreateTrendingImage(ctx context.Context, in *CreateTrendingImageRequest, opts ...grpc.CallOption) (*CreateTrendingImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTrendingImageResponse)
	err := c.cc.Invoke(ctx, BillSplitting_CreateTrendingImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillSplittingServer is the server API for BillSplitting service.
// All implementations must embed UnimplementedBillSplittingServer
// for forward compatibility.
type BillSplittingServer interface {
	CreateExpenseSummaryChart(context.Context, *CreateExpenseSummaryChartRequest) (*CreateExpenseSummaryChartResponse, error)
	GetAuthToken(context.Context, *GetAuthTokenRequest) (*GetAuthTokenResponse, error)
	CreateLineGroup(context.Context, *CreateLineGroupRequest) (*CreateLineGroupResponse, error)
	AddMembership(context.Context, *AddMembershipRequest) (*AddMembershipResponse, error)
	GetLineGroup(context.Context, *GetLineGroupRequest) (*GetLineGroupResponse, error)
	GetMembership(context.Context, *GetMembershipRequest) (*GetMembershipResponse, error)
	CreateExpense(context.Context, *CreateExpenseRequest) (*CreateExpenseResponse, error)
	CreateTrendingImage(context.Context, *CreateTrendingImageRequest) (*CreateTrendingImageResponse, error)
	mustEmbedUnimplementedBillSplittingServer()
}

// UnimplementedBillSplittingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBillSplittingServer struct{}

func (UnimplementedBillSplittingServer) CreateExpenseSummaryChart(context.Context, *CreateExpenseSummaryChartRequest) (*CreateExpenseSummaryChartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExpenseSummaryChart not implemented")
}
func (UnimplementedBillSplittingServer) GetAuthToken(context.Context, *GetAuthTokenRequest) (*GetAuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthToken not implemented")
}
func (UnimplementedBillSplittingServer) CreateLineGroup(context.Context, *CreateLineGroupRequest) (*CreateLineGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLineGroup not implemented")
}
func (UnimplementedBillSplittingServer) AddMembership(context.Context, *AddMembershipRequest) (*AddMembershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMembership not implemented")
}
func (UnimplementedBillSplittingServer) GetLineGroup(context.Context, *GetLineGroupRequest) (*GetLineGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLineGroup not implemented")
}
func (UnimplementedBillSplittingServer) GetMembership(context.Context, *GetMembershipRequest) (*GetMembershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMembership not implemented")
}
func (UnimplementedBillSplittingServer) CreateExpense(context.Context, *CreateExpenseRequest) (*CreateExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExpense not implemented")
}
func (UnimplementedBillSplittingServer) CreateTrendingImage(context.Context, *CreateTrendingImageRequest) (*CreateTrendingImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrendingImage not implemented")
}
func (UnimplementedBillSplittingServer) mustEmbedUnimplementedBillSplittingServer() {}
func (UnimplementedBillSplittingServer) testEmbeddedByValue()                       {}

// UnsafeBillSplittingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillSplittingServer will
// result in compilation errors.
type UnsafeBillSplittingServer interface {
	mustEmbedUnimplementedBillSplittingServer()
}

func RegisterBillSplittingServer(s grpc.ServiceRegistrar, srv BillSplittingServer) {
	// If the following call pancis, it indicates UnimplementedBillSplittingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BillSplitting_ServiceDesc, srv)
}

func _BillSplitting_CreateExpenseSummaryChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExpenseSummaryChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillSplittingServer).CreateExpenseSummaryChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillSplitting_CreateExpenseSummaryChart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillSplittingServer).CreateExpenseSummaryChart(ctx, req.(*CreateExpenseSummaryChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillSplitting_GetAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillSplittingServer).GetAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillSplitting_GetAuthToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillSplittingServer).GetAuthToken(ctx, req.(*GetAuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillSplitting_CreateLineGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLineGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillSplittingServer).CreateLineGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillSplitting_CreateLineGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillSplittingServer).CreateLineGroup(ctx, req.(*CreateLineGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillSplitting_AddMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillSplittingServer).AddMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillSplitting_AddMembership_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillSplittingServer).AddMembership(ctx, req.(*AddMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillSplitting_GetLineGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLineGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillSplittingServer).GetLineGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillSplitting_GetLineGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillSplittingServer).GetLineGroup(ctx, req.(*GetLineGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillSplitting_GetMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillSplittingServer).GetMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillSplitting_GetMembership_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillSplittingServer).GetMembership(ctx, req.(*GetMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillSplitting_CreateExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillSplittingServer).CreateExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillSplitting_CreateExpense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillSplittingServer).CreateExpense(ctx, req.(*CreateExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillSplitting_CreateTrendingImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTrendingImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillSplittingServer).CreateTrendingImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillSplitting_CreateTrendingImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillSplittingServer).CreateTrendingImage(ctx, req.(*CreateTrendingImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BillSplitting_ServiceDesc is the grpc.ServiceDesc for BillSplitting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BillSplitting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BillSplitting",
	HandlerType: (*BillSplittingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExpenseSummaryChart",
			Handler:    _BillSplitting_CreateExpenseSummaryChart_Handler,
		},
		{
			MethodName: "GetAuthToken",
			Handler:    _BillSplitting_GetAuthToken_Handler,
		},
		{
			MethodName: "CreateLineGroup",
			Handler:    _BillSplitting_CreateLineGroup_Handler,
		},
		{
			MethodName: "AddMembership",
			Handler:    _BillSplitting_AddMembership_Handler,
		},
		{
			MethodName: "GetLineGroup",
			Handler:    _BillSplitting_GetLineGroup_Handler,
		},
		{
			MethodName: "GetMembership",
			Handler:    _BillSplitting_GetMembership_Handler,
		},
		{
			MethodName: "CreateExpense",
			Handler:    _BillSplitting_CreateExpense_Handler,
		},
		{
			MethodName: "CreateTrendingImage",
			Handler:    _BillSplitting_CreateTrendingImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/billSplitting.proto",
}
