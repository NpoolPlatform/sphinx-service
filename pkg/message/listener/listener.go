package listener

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/app"
	msgcli "github.com/NpoolPlatform/sphinx-service/pkg/message/client"
	msg "github.com/NpoolPlatform/sphinx-service/pkg/message/message"
)

func Listen() {
	FlagDebug := true
	if FlagDebug {
		go listenTransactionSucceeded()
	}
}

func listenTransactionSucceeded() {
	for {
		logger.Sugar().Infof("listening for transaction success")
		err := msgcli.ComsumerOfAgent(comsumeTransactionSucceeded)
		if err != nil {
			logger.Sugar().Errorf("fail to consume transaction successor: %v", err)
		}
	}
}

func comsumeTransactionSucceeded(notification *msg.NotificationTransaction) (err error) {
	resp, err := app.ACK(context.Background(), &trading.ACKRequest{
		TransactionType:     notification.TransactionType,
		CoinTypeId:          int32(notification.CoinType),
		TransactionIdInsite: notification.TransactionIDInsite,
		TransactionIdChain:  notification.TransactionIDChain,
		Address:             notification.AddressFrom,
		Balance:             notification.AmountFloat64,
		IsOkay:              true,
		ErrorMessage:        "",
	})
	logger.Sugar().Infof("good news everyone: %+w", resp)
	return
}
