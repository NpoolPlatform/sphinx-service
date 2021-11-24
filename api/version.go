package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ping pong
func (s *Server) Version(ctx context.Context, in *emptypb.Empty) (*trading.VersionResponse, error) {
	resp, err := app.Version()
	if err != nil {
		logger.Sugar().Errorw("[Version] get service version error: %w", err)
		return &trading.VersionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
