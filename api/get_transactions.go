package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/sphinxservice"
	ccontant "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-service/pkg/check"
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	scontant "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTransactions(ctx context.Context, in *sphinxservice.GetTransactionsRequest) (resp *sphinxservice.GetTransactionsResponse, err error) {
	if in.GetState() != "" && !check.State(transaction.Status(in.GetState())) {
		logger.Sugar().Errorf("GetTransactions check State invalid")
		return &sphinxservice.GetTransactionsResponse{}, status.Errorf(codes.InvalidArgument, "State: %v invalid", in.GetState())
	}

	ctx, cancel := context.WithTimeout(ctx, scontant.GrpcTimeout)
	defer cancel()

	if in.GetName() != "" {
		conn, err := grpc.GetGRPCConn(ccontant.ServiceName, grpc.GRPCTAG)
		if err != nil {
			logger.Sugar().Errorf("GetTransactions get coininfo service conn error: %v", err)
			return &sphinxservice.GetTransactionsResponse{}, status.Error(codes.Internal, "internal server error")
		}

		coinClient := coininfo.NewSphinxCoinInfoClient(conn)

		coinInfo, err := coinClient.GetCoinInfo(ctx, &coininfo.GetCoinInfoRequest{
			Name: in.GetName(),
		})
		if ent.IsNotFound(err) {
			logger.Sugar().Errorf("GetTransactions call GetCoinInfo Name: %v not found", in.GetName())
			return &sphinxservice.GetTransactionsResponse{}, status.Errorf(codes.Internal, "Name: %v not found", in.GetName())
		}
		if err != nil {
			logger.Sugar().Errorf("GetTransactions call GetCoinInfo Name: %v error: %v", in.GetName(), err)
			return &sphinxservice.GetTransactionsResponse{}, status.Error(codes.Internal, "internal server error")
		}

		if coinInfo.GetInfo().GetName() != in.GetName() {
			logger.Sugar().Errorf("GetTransactions call GetCoinInfo Name: %v not support", in.GetName(), err)
			return &sphinxservice.GetTransactionsResponse{}, status.Errorf(codes.InvalidArgument, "Name: %v not support", in.GetName())
		}
	}

	transactionInfos, total, err := crud.GetTransactions(ctx, &crud.GetTransactionsParams{
		Name:   in.GetName(),
		State:  transaction.Status(in.GetState()),
		From:   in.GetFrom(),
		To:     in.GetTo(),
		Offset: int(in.GetOffset()),
		Limit:  int(in.GetLimit()),
	})
	if err != nil {
		logger.Sugar().Errorf("GetTransactions call GetTransactions error: %v", err)
		return &sphinxservice.GetTransactionsResponse{}, status.Error(codes.Internal, "internal server error")
	}

	infos := make([]*sphinxservice.TransactionInfo, len(transactionInfos))
	for idx, transInfo := range transactionInfos {
		infos[idx] = &sphinxservice.TransactionInfo{
			TransactionID: transInfo.TransactionID,
			Cid:           transInfo.Cid,
			Name:          transInfo.Name,
			Amount:        price.DBPriceToVisualPrice(transInfo.Amount),
			From:          transInfo.From,
			To:            transInfo.To,
			State:         string(transInfo.Status),
			CreatedAt:     transInfo.CreatedAt,
			UpdatedAt:     transInfo.UpdatedAt,
		}
	}

	return &sphinxservice.GetTransactionsResponse{
		Total: int32(total),
		Infos: infos,
	}, nil
}
