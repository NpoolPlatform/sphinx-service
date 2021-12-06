package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxservice"
	"github.com/NpoolPlatform/sphinx-service/pkg/check"
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateTransaction(ctx context.Context, in *sphinxservice.UpdateTransactionRequest) (resp *sphinxservice.UpdateTransactionResponse, err error) {
	if in.GetTransactionID() == "" {
		logger.Sugar().Errorf("UpdateTransaction check TransactionID empty")
		return &sphinxservice.UpdateTransactionResponse{}, status.Errorf(codes.InvalidArgument, "TransactionID empty")
	}

	if in.GetState() == "" {
		logger.Sugar().Errorf("UpdateTransaction check State empty")
		return &sphinxservice.UpdateTransactionResponse{}, status.Errorf(codes.InvalidArgument, "State empty")
	}

	if !check.State(transaction.Status(in.GetState())) {
		logger.Sugar().Errorf("UpdateTransaction check State invalid")
		return &sphinxservice.UpdateTransactionResponse{}, status.Errorf(codes.InvalidArgument, "State: %v invalid", in.GetState())
	}

	_, err = crud.GetTransaction(ctx, &sphinxservice.GetTransactionRequest{
		TransactionID: in.GetTransactionID(),
	})
	if ent.IsNotFound(err) {
		logger.Sugar().Errorf("UpdateTransaction call GetTransaction TransactionID: %v not found", in.GetTransactionID())
		return &sphinxservice.UpdateTransactionResponse{}, status.Errorf(codes.NotFound, "TransactionID: %v not found", in.GetTransactionID())
	}
	if err != nil {
		logger.Sugar().Errorf("UpdateTransaction call GetTransaction error: %v", err)
		return &sphinxservice.UpdateTransactionResponse{}, status.Errorf(codes.Internal, "internal server error")
	}

	if err := crud.UpdateTransactionStatus(ctx, crud.UpdateTransactionStatusParams{
		TransactionID: in.GetTransactionID(),
		State:         transaction.Status(in.GetState()),
	}); err != nil {
		logger.Sugar().Errorf("UpdateTransaction call UpdateTransactionStatus error: %v", err)
		return &sphinxservice.UpdateTransactionResponse{}, status.Errorf(codes.Internal, "internal server error")
	}

	return &sphinxservice.UpdateTransactionResponse{}, nil
}
