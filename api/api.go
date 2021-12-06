package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxservice"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	sphinxservice.UnimplementedSphinxServiceServer
}

func Register(server grpc.ServiceRegistrar) {
	sphinxservice.RegisterSphinxServiceServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return sphinxservice.RegisterSphinxServiceHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
