package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Version(ctx context.Context, in *emptypb.Empty) (*trading.VersionResponse, error) {
	v, err := version.Version()
	if err != nil {
		logger.Sugar().Errorf("Version call Version error: %v", err)
		return &trading.VersionResponse{}, status.Errorf(codes.Internal, "internal server error")
	}
	return v, nil
}
