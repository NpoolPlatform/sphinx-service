package api

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const DebugFlag = false

type Server struct {
	npool.UnimplementedTradingServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterTradingServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterTradingHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
