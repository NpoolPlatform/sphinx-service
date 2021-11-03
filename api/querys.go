package echo

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/core"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedTradingServer
}

func (s *Server) GetCoinInfos(ctx context.Context, req *npool.GetCoinInfosRequest) (cilist *npool.CoinInfoList, err error) {
	cilist, err = core.GetCoinInfos(ctx, req)
}

func (s *Server) GetCoinInfo(ctx context.Context, req *npool.GetCoinInfoRequest) (cilist *npool.CoinInfoRow, err error) {
	cilist, err = core.GetCoinInfo(ctx, req)
}
