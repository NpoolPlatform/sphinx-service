package tasks

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
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
			defer cancel()
			transInfos, _, err := crud.GetTransactions(ctx, &crud.GetTransactionsParams{})
			if err != nil {
				logger.Sugar().Infof("err:%v", err)
			}

			for _, trans := range transInfos {
				logger.Sugar().Infof("info: %v", trans)
				// conn, err := grpc.GetGRPCConn(pconst.ServiceName, grpc.GRPCTAG)
				// if err != nil {

				// }

				// client := sphinxproxy.NewSphinxProxyClient(conn)
				// tranInfo, err := client.GetTransaction(ctx, &sphinxproxy.GetTransactionRequest{
				// 	TransactionID: trans.TransactionID,
				// })
				// if err != nil {

				// }
			}
		}()
	}
}
