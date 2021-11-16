package client

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	grpcBase "google.golang.org/grpc"
)

var (
	serviceNamespace string
	ClientApproval   *grpcBase.ClientConn
)

func SetServiceName(str string) {
	serviceNamespace = str
}

func InitApprovalClient() (err error) {
	ClientApproval, err = grpc.GetGRPCConn(config.GetStringValueWithNameSpace(serviceNamespace, "grpc_conn_address"))
	if err != nil {
		logger.Sugar().Errorf("get grpc connection failure, err: %v", err)
	}
	return
}

func LetApproveTransaction(transactionIDInsite string) (err error) {
	return nil
}
