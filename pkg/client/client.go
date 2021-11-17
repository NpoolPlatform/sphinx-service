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

func Init() (err error) {
	err = InitProxyClient()
	if err != nil {
		logger.Sugar().Errorf("init proxy client error, %w", err)
	}
	return
}

// 审核服务 init grpc
func InitApprovalClient() (err error) {
	_, err = grpc.GetGRPCConn("TO-BE-IMPLEMENTED")
	if err != nil {
		logger.Sugar().Errorf("get grpc connection failure, err: %v", err)
	}
	return
}

// 钱包代理服务 init grpc
func InitProxyClient() (err error) {
	conn, err := grpc.GetGRPCConn("TO-BE-IMPLEMENTED")
	if err != nil {
		logger.Sugar().Errorf("get grpc connection failure, err: %v", err)
	} else {
		ClientProxy = signproxy.NewSignProxyClient(conn)
	}
	return
}
