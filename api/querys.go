package api

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/npool"

	"github.com/NpoolPlatform/sphinx-service/pkg/core"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedTradingServer
}

func (s *Server) GetCoinInfos(ctx context.Context, req *npool.GetCoinInfosRequest) (cilist *npool.CoinInfoList, err error) {
	cilist, err = core.GetCoinInfos(ctx, req)
	return
}

func (s *Server) GetCoinInfo(ctx context.Context, req *npool.GetCoinInfoRequest) (cilist *npool.CoinInfoRow, err error) {
	cilist, err = core.GetCoinInfo(ctx, req)
	return
}
