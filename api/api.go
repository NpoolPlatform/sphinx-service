package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedTradingServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterTradingServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterTradingHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
