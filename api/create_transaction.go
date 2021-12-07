package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/sphinxservice"
	ccontant "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	scontant "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateTransaction ..
// TODO invoke review service
func (s *Server) CreateTransaction(ctx context.Context, in *sphinxservice.CreateTransactionRequest) (resp *sphinxservice.CreateTransactionResponse, err error) {
	if in.GetTransactionID() == "" {
		logger.Sugar().Errorf("CreateTransaction TransactionID empty", err)
		return &sphinxservice.CreateTransactionResponse{}, status.Error(codes.InvalidArgument, "TransactionID empty")
	}
	if in.GetName() == "" {
		logger.Sugar().Errorf("CreateTransaction Name empty", err)
		return &sphinxservice.CreateTransactionResponse{}, status.Error(codes.InvalidArgument, "Name empty")
	}
	if in.GetAmount() <= 0 {
		logger.Sugar().Errorf("CreateTransaction Amount less than 0", err)
		return &sphinxservice.CreateTransactionResponse{}, status.Error(codes.InvalidArgument, "Amount must than 0")
	}
	if in.GetFrom() == "" {
		logger.Sugar().Errorf("CreateTransaction From empty", err)
		return &sphinxservice.CreateTransactionResponse{}, status.Error(codes.InvalidArgument, "From empty")
	}
	if in.GetTo() == "" {
		logger.Sugar().Errorf("CreateTransaction To empty", err)
		return &sphinxservice.CreateTransactionResponse{}, status.Error(codes.InvalidArgument, "To empty")
	}
	conn, err := grpc.GetGRPCConn(ccontant.ServiceName, grpc.GRPCTAG)
	if err != nil {
		logger.Sugar().Errorf("CreateTransaction get coininfo service conn error: %v", err)
		return &sphinxservice.CreateTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}

	coinClient := coininfo.NewSphinxCoinInfoClient(conn)

	ctx, cancel := context.WithTimeout(ctx, scontant.GrpcTimeout)
	defer cancel()

	coinInfo, err := coinClient.GetCoinInfo(ctx, &coininfo.GetCoinInfoRequest{
		Name: in.GetName(),
	})
	if ent.IsNotFound(err) {
		logger.Sugar().Errorf("CreateTransaction get coininfo Name: %v not found", in.GetName())
		return &sphinxservice.CreateTransactionResponse{}, status.Errorf(codes.InvalidArgument, "Name: %v not support", in.GetName())
	}
	if err != nil {
		logger.Sugar().Errorf("CreateTransaction call GetCoinInfo Name: %v error: %v", in.GetName(), err)
		return &sphinxservice.CreateTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}

	if coinInfo.GetInfo().GetName() != in.GetName() {
		logger.Sugar().Errorf("CreateTransaction call GetCoinInfo Name: %v not support", in.GetName(), err)
		return &sphinxservice.CreateTransactionResponse{}, status.Errorf(codes.InvalidArgument, "Name: %v not support", in.GetName())
	}

	params := crud.CreateTransactionParams{
		TransactionID: in.GetTransactionID(),
		Name:          in.GetName(),
		Amount:        in.GetAmount(),
		From:          in.GetFrom(),
		To:            in.GetTo(),
	}
	if _, err := crud.CreateTransaction(ctx, params); err != nil {
		logger.Sugar().Errorf("CreateTransaction call db CreateTransaction params: %v error: %v", params, err)
		return &sphinxservice.CreateTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &sphinxservice.CreateTransactionResponse{}, nil
}
