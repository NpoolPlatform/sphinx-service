package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-app-template/api/sphinx"
	"github.com/NpoolPlatform/go-service-app-template/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterServiceExampleServer(server, &sphinx.Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterServiceExampleHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
