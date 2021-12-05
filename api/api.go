package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	trading.UnimplementedTradingServer
}

func Register(server grpc.ServiceRegistrar) {
	trading.RegisterTradingServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return trading.RegisterTradingHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
