package client

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/coininfo"
	review "github.com/NpoolPlatform/message/npool/review"
	"github.com/NpoolPlatform/message/npool/signproxy"
)

var (
	ClientApproval review.ReviewServiceClient
	ClientProxy    signproxy.SignProxyClient
	ClientCoinInfo coininfo.SphinxCoinInfoClient
)

func Init() (err error) {
	err = InitApprovalClient()
	if err != nil {
		logger.Sugar().Errorf("init proxy client error, %w", err)
	}
	return
}

// 审核服务 init grpc TODO
func InitApprovalClient() (err error) {
	conn, err := grpc.GetGRPCConn("TO-BE-IMPLEMENTED")
	if err == nil {
		ClientApproval = review.NewReviewServiceClient(conn)
	}
	return
}
