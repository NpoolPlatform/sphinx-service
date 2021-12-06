package tasks

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	pconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	constant "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
)

func init() {
	tasks["createTransaction"] = createTransaction
}

// createTransaction ..
func createTransaction() {
	for range time.NewTicker(constant.TaskDuration).C {
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), constant.GrpcTimeout)
			defer cancel()

			// find confirm transaction
			transactionInfos, _, err := crud.GetTransactions(ctx, &crud.GetTransactionsParams{
				Limit: constant.PageSize,
				State: transaction.StatusConfirm,
			})
			if err != nil {
				logger.Sugar().Error("call GetTransactions error: %v", err)
				return
			}

			for _, transInfo := range transactionInfos {
				conn, err := grpc.GetGRPCConn(pconst.ServiceName, grpc.GRPCTAG)
				if err != nil {
					logger.Sugar().Error("call GetGRPCConn error: %v", err)
					continue
				}

				proxyClient := sphinxproxy.NewSphinxProxyClient(conn)
				if _, err := proxyClient.CreateTransaction(ctx, &sphinxproxy.CreateTransactionRequest{
					Name:          transInfo.Name,
					TransactionID: transInfo.TransactionID,
					From:          transInfo.From,
					To:            transInfo.To,
					Value:         price.DBPriceToVisualPrice(transInfo.Amount),
				}); err != nil {
					logger.Sugar().Error("call CreateTransaction error: %v", err)
					continue
				}

				if err := crud.UpdateTransactionStatus(ctx, crud.UpdateTransactionStatusParams{
					TransactionID: transInfo.TransactionID,
					State:         transaction.StatusPendingTransaction,
				}); err != nil {
					logger.Sugar().Error("call UpdateTransactionStatus error: %v", err)
					continue
				}

				logger.Sugar().Infof("transaction: %v done", transInfo.TransactionID)
			}
		}()
	}
}
