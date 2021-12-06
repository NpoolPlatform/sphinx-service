package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/sphinxservice"
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	scontant "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTransaction(ctx context.Context, in *sphinxservice.GetTransactionRequest) (resp *sphinxservice.GetTransactionResponse, err error) {
	if in.GetTransactionID() == "" {
		logger.Sugar().Errorf("GetTransaction check TransactionID empty")
		return &sphinxservice.GetTransactionResponse{}, status.Error(codes.InvalidArgument, "TransactionID empty")
	}

	ctx, cancel := context.WithTimeout(ctx, scontant.GrpcTimeout)
	defer cancel()

	transactionInfo, err := crud.GetTransaction(ctx, in)
	if ent.IsNotFound(err) {
		logger.Sugar().Errorf("GetTransaction call GetTransaction TransactionID: %v not found", in.GetTransactionID())
		return &sphinxservice.GetTransactionResponse{}, status.Errorf(codes.NotFound, "TransactionID: %v not found", in.GetTransactionID())
	}
	if err != nil {
		logger.Sugar().Errorf("GetTransaction call GetTransaction error: %v", err)
		return &sphinxservice.GetTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &sphinxservice.GetTransactionResponse{
		Info: &sphinxservice.TransactionInfo{
			TransactionID: transactionInfo.TransactionID,
			Cid:           transactionInfo.Cid,
			Name:          transactionInfo.Name,
			Amount:        price.DBPriceToVisualPrice(transactionInfo.Amount),
			From:          transactionInfo.From,
			To:            transactionInfo.To,
			State:         string(transactionInfo.Status),
			CreatedAt:     transactionInfo.CreatedAt,
			UpdatedAt:     transactionInfo.UpdatedAt,
		},
	}, nil
}
