package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxservice"
	"github.com/NpoolPlatform/sphinx-service/pkg/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Version(ctx context.Context, in *emptypb.Empty) (*sphinxservice.VersionResponse, error) {
	v, err := version.Version()
	if err != nil {
		logger.Sugar().Errorf("Version call Version error: %v", err)
		return &sphinxservice.VersionResponse{}, status.Errorf(codes.Internal, "internal server error")
	}
	return v, nil
}
