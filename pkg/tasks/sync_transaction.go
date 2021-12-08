package tasks

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	pt "github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
	pconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	constant "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
)

func init() {
	tasks["syncTransaction"] = syncTransaction
}

// syncTransaction ..
func syncTransaction() {
	for range time.NewTicker(constant.TaskDuration).C {
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), constant.GrpcTimeout)
			transInfos, _, err := crud.GetTransactions(ctx, &crud.GetTransactionsParams{
				Limit: constant.PageSize,
				State: transaction.StatusPendingTransaction,
			})
			cancel()
			if err != nil {
				logger.Sugar().Infof("call GetTransactions error: %v", err)
				return
			}

			for _, trans := range transInfos {
				handle(trans)
			}
		}()
	}
}

func handle(trans *ent.Transaction) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.GrpcTimeout)
	defer cancel()

	conn, err := grpc.GetGRPCConn(pconst.ServiceName, grpc.GRPCTAG)
	if err != nil {
		logger.Sugar().Errorf("call GetGRPCConn error: %v", err)
		return
	}

	client := sphinxproxy.NewSphinxProxyClient(conn)
	tranInfo, err := client.GetTransaction(ctx, &sphinxproxy.GetTransactionRequest{
		TransactionID: trans.TransactionID,
	})
	if err != nil {
		logger.Sugar().Errorf("call GetTransaction error: %v", err)
		return
	}

	info := tranInfo.GetInfo()
	if info.GetState() == pt.StateDone.String() {
		if err := crud.UpdateTransactionStatus(ctx, crud.UpdateTransactionStatusParams{
			TransactionID: trans.TransactionID,
			State:         transaction.StatusDone,
			CID:           info.GetCID(),
			ExitCode:      info.GetExitCode(),
		}); err != nil {
			logger.Sugar().Errorf("call UpdateTransactionStatus error: %v", err)
			return
		}
	}
}
