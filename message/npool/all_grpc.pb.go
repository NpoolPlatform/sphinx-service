// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package npool

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

// WalletNodeClient is the client API for WalletNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletNodeClient interface {
	// 获取预签名信息
	GetSignInfo(ctx context.Context, in *GetSignInfoRequest, opts ...grpc.CallOption) (*SignInfo, error)
	// 广播交易
	BroadcastScript(ctx context.Context, in *BroadcastScriptRequest, opts ...grpc.CallOption) (*SuccessCode, error)
	// 查询交易状态
	StatusScript(ctx context.Context, in *StatusScriptRequest, opts ...grpc.CallOption) (*ScriptInfo, error)
}

type walletNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletNodeClient(cc grpc.ClientConnInterface) WalletNodeClient {
	return &walletNodeClient{cc}
}

func (c *walletNodeClient) GetSignInfo(ctx context.Context, in *GetSignInfoRequest, opts ...grpc.CallOption) (*SignInfo, error) {
	out := new(SignInfo)
	err := c.cc.Invoke(ctx, "/sphinx.v1.WalletNode/GetSignInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletNodeClient) BroadcastScript(ctx context.Context, in *BroadcastScriptRequest, opts ...grpc.CallOption) (*SuccessCode, error) {
	out := new(SuccessCode)
	err := c.cc.Invoke(ctx, "/sphinx.v1.WalletNode/BroadcastScript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletNodeClient) StatusScript(ctx context.Context, in *StatusScriptRequest, opts ...grpc.CallOption) (*ScriptInfo, error) {
	out := new(ScriptInfo)
	err := c.cc.Invoke(ctx, "/sphinx.v1.WalletNode/StatusScript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletNodeServer is the server API for WalletNode service.
// All implementations must embed UnimplementedWalletNodeServer
// for forward compatibility
type WalletNodeServer interface {
	// 获取预签名信息
	GetSignInfo(context.Context, *GetSignInfoRequest) (*SignInfo, error)
	// 广播交易
	BroadcastScript(context.Context, *BroadcastScriptRequest) (*SuccessCode, error)
	// 查询交易状态
	StatusScript(context.Context, *StatusScriptRequest) (*ScriptInfo, error)
	mustEmbedUnimplementedWalletNodeServer()
}

// UnimplementedWalletNodeServer must be embedded to have forward compatible implementations.
type UnimplementedWalletNodeServer struct {
}

func (UnimplementedWalletNodeServer) GetSignInfo(context.Context, *GetSignInfoRequest) (*SignInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignInfo not implemented")
}
func (UnimplementedWalletNodeServer) BroadcastScript(context.Context, *BroadcastScriptRequest) (*SuccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastScript not implemented")
}
func (UnimplementedWalletNodeServer) StatusScript(context.Context, *StatusScriptRequest) (*ScriptInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusScript not implemented")
}
func (UnimplementedWalletNodeServer) mustEmbedUnimplementedWalletNodeServer() {}

// UnsafeWalletNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletNodeServer will
// result in compilation errors.
type UnsafeWalletNodeServer interface {
	mustEmbedUnimplementedWalletNodeServer()
}

func RegisterWalletNodeServer(s grpc.ServiceRegistrar, srv WalletNodeServer) {
	s.RegisterService(&WalletNode_ServiceDesc, srv)
}

func _WalletNode_GetSignInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletNodeServer).GetSignInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.WalletNode/GetSignInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletNodeServer).GetSignInfo(ctx, req.(*GetSignInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletNode_BroadcastScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletNodeServer).BroadcastScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.WalletNode/BroadcastScript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletNodeServer).BroadcastScript(ctx, req.(*BroadcastScriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletNode_StatusScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletNodeServer).StatusScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.WalletNode/StatusScript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletNodeServer).StatusScript(ctx, req.(*StatusScriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletNode_ServiceDesc is the grpc.ServiceDesc for WalletNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sphinx.v1.WalletNode",
	HandlerType: (*WalletNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSignInfo",
			Handler:    _WalletNode_GetSignInfo_Handler,
		},
		{
			MethodName: "BroadcastScript",
			Handler:    _WalletNode_BroadcastScript_Handler,
		},
		{
			MethodName: "StatusScript",
			Handler:    _WalletNode_StatusScript_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/all.proto",
}

// WalletAgentClient is the client API for WalletAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletAgentClient interface {
	// 接收健康报告
	AcceptNode(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*SuccessCode, error)
}

type walletAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletAgentClient(cc grpc.ClientConnInterface) WalletAgentClient {
	return &walletAgentClient{cc}
}

func (c *walletAgentClient) AcceptNode(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*SuccessCode, error) {
	out := new(SuccessCode)
	err := c.cc.Invoke(ctx, "/sphinx.v1.WalletAgent/AcceptNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletAgentServer is the server API for WalletAgent service.
// All implementations must embed UnimplementedWalletAgentServer
// for forward compatibility
type WalletAgentServer interface {
	// 接收健康报告
	AcceptNode(context.Context, *NodeInfo) (*SuccessCode, error)
	mustEmbedUnimplementedWalletAgentServer()
}

// UnimplementedWalletAgentServer must be embedded to have forward compatible implementations.
type UnimplementedWalletAgentServer struct {
}

func (UnimplementedWalletAgentServer) AcceptNode(context.Context, *NodeInfo) (*SuccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptNode not implemented")
}
func (UnimplementedWalletAgentServer) mustEmbedUnimplementedWalletAgentServer() {}

// UnsafeWalletAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletAgentServer will
// result in compilation errors.
type UnsafeWalletAgentServer interface {
	mustEmbedUnimplementedWalletAgentServer()
}

func RegisterWalletAgentServer(s grpc.ServiceRegistrar, srv WalletAgentServer) {
	s.RegisterService(&WalletAgent_ServiceDesc, srv)
}

func _WalletAgent_AcceptNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletAgentServer).AcceptNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.WalletAgent/AcceptNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletAgentServer).AcceptNode(ctx, req.(*NodeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletAgent_ServiceDesc is the grpc.ServiceDesc for WalletAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sphinx.v1.WalletAgent",
	HandlerType: (*WalletAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcceptNode",
			Handler:    _WalletAgent_AcceptNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/all.proto",
}

// TradingClient is the client API for Trading service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradingClient interface {
	// 创建账户
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*UserAddress, error)
	// 用户申请提现
	WithdrawApply(ctx context.Context, in *WithdrawApplyRequest, opts ...grpc.CallOption) (*SuccessCode, error)
	// （非主动） 充值购买 - 异步扫描收款地址通知回调
	RechargeUser(ctx context.Context, in *RechargeUserRequest, opts ...grpc.CallOption) (*SuccessCode, error)
	// （非主动） 确认提现 - 从消息队列中获取确认提现的数据并操作
	WithdrawConfirmed(ctx context.Context, in *WithdrawConfirmedRequest, opts ...grpc.CallOption) (*SuccessCode, error)
}

type tradingClient struct {
	cc grpc.ClientConnInterface
}

func NewTradingClient(cc grpc.ClientConnInterface) TradingClient {
	return &tradingClient{cc}
}

func (c *tradingClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*UserAddress, error) {
	out := new(UserAddress)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingClient) WithdrawApply(ctx context.Context, in *WithdrawApplyRequest, opts ...grpc.CallOption) (*SuccessCode, error) {
	out := new(SuccessCode)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/WithdrawApply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingClient) RechargeUser(ctx context.Context, in *RechargeUserRequest, opts ...grpc.CallOption) (*SuccessCode, error) {
	out := new(SuccessCode)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/RechargeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingClient) WithdrawConfirmed(ctx context.Context, in *WithdrawConfirmedRequest, opts ...grpc.CallOption) (*SuccessCode, error) {
	out := new(SuccessCode)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/WithdrawConfirmed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradingServer is the server API for Trading service.
// All implementations must embed UnimplementedTradingServer
// for forward compatibility
type TradingServer interface {
	// 创建账户
	RegisterUser(context.Context, *RegisterUserRequest) (*UserAddress, error)
	// 用户申请提现
	WithdrawApply(context.Context, *WithdrawApplyRequest) (*SuccessCode, error)
	// （非主动） 充值购买 - 异步扫描收款地址通知回调
	RechargeUser(context.Context, *RechargeUserRequest) (*SuccessCode, error)
	// （非主动） 确认提现 - 从消息队列中获取确认提现的数据并操作
	WithdrawConfirmed(context.Context, *WithdrawConfirmedRequest) (*SuccessCode, error)
	mustEmbedUnimplementedTradingServer()
}

// UnimplementedTradingServer must be embedded to have forward compatible implementations.
type UnimplementedTradingServer struct {
}

func (UnimplementedTradingServer) RegisterUser(context.Context, *RegisterUserRequest) (*UserAddress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedTradingServer) WithdrawApply(context.Context, *WithdrawApplyRequest) (*SuccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawApply not implemented")
}
func (UnimplementedTradingServer) RechargeUser(context.Context, *RechargeUserRequest) (*SuccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RechargeUser not implemented")
}
func (UnimplementedTradingServer) WithdrawConfirmed(context.Context, *WithdrawConfirmedRequest) (*SuccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawConfirmed not implemented")
}
func (UnimplementedTradingServer) mustEmbedUnimplementedTradingServer() {}

// UnsafeTradingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TradingServer will
// result in compilation errors.
type UnsafeTradingServer interface {
	mustEmbedUnimplementedTradingServer()
}

func RegisterTradingServer(s grpc.ServiceRegistrar, srv TradingServer) {
	s.RegisterService(&Trading_ServiceDesc, srv)
}

func _Trading_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trading_WithdrawApply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).WithdrawApply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/WithdrawApply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).WithdrawApply(ctx, req.(*WithdrawApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trading_RechargeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RechargeUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).RechargeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/RechargeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).RechargeUser(ctx, req.(*RechargeUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trading_WithdrawConfirmed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawConfirmedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).WithdrawConfirmed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/WithdrawConfirmed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).WithdrawConfirmed(ctx, req.(*WithdrawConfirmedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Trading_ServiceDesc is the grpc.ServiceDesc for Trading service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Trading_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sphinx.v1.Trading",
	HandlerType: (*TradingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _Trading_RegisterUser_Handler,
		},
		{
			MethodName: "WithdrawApply",
			Handler:    _Trading_WithdrawApply_Handler,
		},
		{
			MethodName: "RechargeUser",
			Handler:    _Trading_RechargeUser_Handler,
		},
		{
			MethodName: "WithdrawConfirmed",
			Handler:    _Trading_WithdrawConfirmed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/all.proto",
}

// ReviewClient is the client API for Review service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewClient interface {
	// 获取待审核列表
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*ReviewList, error)
	// 设置管理员的审核权限
	SetAdminPermission(ctx context.Context, in *SetAdminPermissionRequest, opts ...grpc.CallOption) (*SuccessCode, error)
	// 审核交易
	ProcessReview(ctx context.Context, in *ProcessReviewRequest, opts ...grpc.CallOption) (*ReviewList, error)
}

type reviewClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewClient(cc grpc.ClientConnInterface) ReviewClient {
	return &reviewClient{cc}
}

func (c *reviewClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*ReviewList, error) {
	out := new(ReviewList)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Review/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) SetAdminPermission(ctx context.Context, in *SetAdminPermissionRequest, opts ...grpc.CallOption) (*SuccessCode, error) {
	out := new(SuccessCode)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Review/SetAdminPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) ProcessReview(ctx context.Context, in *ProcessReviewRequest, opts ...grpc.CallOption) (*ReviewList, error) {
	out := new(ReviewList)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Review/ProcessReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServer is the server API for Review service.
// All implementations must embed UnimplementedReviewServer
// for forward compatibility
type ReviewServer interface {
	// 获取待审核列表
	GetList(context.Context, *GetListRequest) (*ReviewList, error)
	// 设置管理员的审核权限
	SetAdminPermission(context.Context, *SetAdminPermissionRequest) (*SuccessCode, error)
	// 审核交易
	ProcessReview(context.Context, *ProcessReviewRequest) (*ReviewList, error)
	mustEmbedUnimplementedReviewServer()
}

// UnimplementedReviewServer must be embedded to have forward compatible implementations.
type UnimplementedReviewServer struct {
}

func (UnimplementedReviewServer) GetList(context.Context, *GetListRequest) (*ReviewList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedReviewServer) SetAdminPermission(context.Context, *SetAdminPermissionRequest) (*SuccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAdminPermission not implemented")
}
func (UnimplementedReviewServer) ProcessReview(context.Context, *ProcessReviewRequest) (*ReviewList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessReview not implemented")
}
func (UnimplementedReviewServer) mustEmbedUnimplementedReviewServer() {}

// UnsafeReviewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServer will
// result in compilation errors.
type UnsafeReviewServer interface {
	mustEmbedUnimplementedReviewServer()
}

func RegisterReviewServer(s grpc.ServiceRegistrar, srv ReviewServer) {
	s.RegisterService(&Review_ServiceDesc, srv)
}

func _Review_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Review/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_SetAdminPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAdminPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).SetAdminPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Review/SetAdminPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).SetAdminPermission(ctx, req.(*SetAdminPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_ProcessReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ProcessReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Review/ProcessReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ProcessReview(ctx, req.(*ProcessReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Review_ServiceDesc is the grpc.ServiceDesc for Review service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Review_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sphinx.v1.Review",
	HandlerType: (*ReviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _Review_GetList_Handler,
		},
		{
			MethodName: "SetAdminPermission",
			Handler:    _Review_SetAdminPermission_Handler,
		},
		{
			MethodName: "ProcessReview",
			Handler:    _Review_ProcessReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/all.proto",
}
