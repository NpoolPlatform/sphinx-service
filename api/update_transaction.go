package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateTransaction(ctx context.Context, in *trading.UpdateTransactionRequest) (resp *trading.UpdateTransactionResponse, err error) {
	if in.GetTransactionID() == "" {
		logger.Sugar().Errorf("UpdateTransaction check TransactionID empty")
		return &trading.UpdateTransactionResponse{}, status.Errorf(codes.InvalidArgument, "TransactionID empty")
	}

	if in.GetState() == "" {
		logger.Sugar().Errorf("UpdateTransaction check State empty")
		return &trading.UpdateTransactionResponse{}, status.Errorf(codes.InvalidArgument, "State empty")
	}

	if _, err := crud.GetTransaction(ctx, &trading.GetTransactionRequest{
		TransactionID: in.GetTransactionID(),
	}); err != nil {
		logger.Sugar().Errorf("UpdateTransaction call GetTransaction error: %v", err)
		return &trading.UpdateTransactionResponse{}, status.Errorf(codes.Internal, "internal server error")
	}

	if err := crud.UpdateTransactionStatus(ctx, crud.UpdateTransactionStatusParams{}); err != nil {
		logger.Sugar().Errorf("UpdateTransaction call UpdateTransactionStatus error: %v", err)
		return &trading.UpdateTransactionResponse{}, status.Errorf(codes.Internal, "internal server error")
	}

	return &trading.UpdateTransactionResponse{}, nil
}
