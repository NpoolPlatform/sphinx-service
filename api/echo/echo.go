package echo

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/npool"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedTradingServer
}

func (s *Server) Echo(ctx context.Context, in *npool.SignScriptRequest) (*npool.SignInfo, error) {
	return nil, nil
}
