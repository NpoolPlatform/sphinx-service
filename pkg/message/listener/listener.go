package listener

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/sphinx-service/api"
	msgcli "github.com/NpoolPlatform/sphinx-service/pkg/message/client"
	msg "github.com/NpoolPlatform/sphinx-service/pkg/message/message"
)

func Listen() {
	go listenTemplateExample()
	go listenTransactionSucceeded()
	go listenTransactionApproval()
}

func listenTemplateExample() {
	for {
		logger.Sugar().Infof("consume template example")
		err := msgcli.ConsumeExample(func(example *msg.Example) error {
			logger.Sugar().Infof("go example: %+w", example)
			// Call event handler in api module
			return nil
		})
		if err != nil {
			logger.Sugar().Errorf("fail to consume example: %v", err)
			return
		}
	}
}

func listenTransactionSucceeded() {
	for {
		logger.Sugar().Infof("listening for transaction success")
		err := msgcli.ComsumerOfAgent(comsumeTransactionSucceeded)
		if err != nil {
			logger.Sugar().Errorf("fail to consume example: %v", err)
		}
	}
}

func comsumeTransactionSucceeded(notification *msg.NotificationTransaction) error {
	logger.Sugar().Infof("good news everyone: %+w", notification)
	return nil
}

func listenTransactionApproval() {
	for {
		logger.Sugar().Infof("listening for transaction approvals")
		err := msgcli.ComsumerOfTradingForApproval(consumeApproval)
		if err != nil { // one failure
			logger.Sugar().Errorf("fail to consume transaction, error info: %v", err)
		}
	}
}

func consumeApproval(notification *msg.NotificationTransaction) (err error) {
	logger.Sugar().Infof("gonna auto approve tx: %+w", notification)
	var isApproved bool
	isApproved, err = api.ApproveTransaction(notification.TransactionIDInsite)
	if err == nil && isApproved {
		go api.NotifyProcessTransaction(notification)
	}
	return
}
