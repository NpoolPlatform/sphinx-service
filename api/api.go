package api

import (
	"context"
	"github.com/NpoolPlatform/sphinx-service/message/npool"
	"github.com/NpoolPlatform/sphinx-service/pkg/core"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterTradingServer(server, core.GetTradingServer())
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterTradingHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	// return npool.RegisterServiceExampleHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
