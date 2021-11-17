package client

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/signproxy"
	grpcBase "google.golang.org/grpc"
)

var (
	ClientApproval *grpcBase.ClientConn
	ClientProxy    signproxy.SignProxyClient
	ClientCoinInfo coininfo.SphinxCoininfoClient
)

// 审核服务 grpc远程调用

func InitApprovalClient() (err error) {
	_, err = grpc.GetGRPCConn("TO-BE-IMPLEMENTED")
	if err != nil {
		logger.Sugar().Errorf("get grpc connection failure, err: %v", err)
		return
	}
	// MARK
	return
}

// 钱包代理服务 grpc远程调用 其余在rabbitmq做消息通知

func InitProxyClient() (err error) {
	conn, err := grpc.GetGRPCConn("TO-BE-IMPLEMENTED")
	if err != nil {
		logger.Sugar().Errorf("get grpc connection failure, err: %v", err)
		return
	}
	ClientProxy = signproxy.NewSignProxyClient(conn)
	return
}
