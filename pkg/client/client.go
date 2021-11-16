package client

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	grpcBase "google.golang.org/grpc"
)

var ClientApproval *grpcBase.ClientConn

// 审核服务 grpc远程调用

func GetApprovalClient() *grpcBase.ClientConn {
	ClientApproval.Connect()
	return ClientApproval
}

func InitApprovalClient() (err error) {
	// get service conn
	_, err = grpc.GetGRPCConn("TO-BE-IMPLEMENTED")
	if err != nil {
		logger.Sugar().Errorf("get grpc connection failure, err: %v", err)
		return
	}
	// get service client
	// ClientApproval = signproxy.NewSphinxSignproxyClient(conn)
	return
}

// MARK
func LetApproveTransaction(transactionIDInsite string) (err error) {
	return nil
}

// 钱包代理服务

func InitProxyClient() (err error) {
	return nil
}
