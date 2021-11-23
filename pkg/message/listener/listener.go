package listener

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	msgcli "github.com/NpoolPlatform/sphinx-service/pkg/message/client"
	msg "github.com/NpoolPlatform/sphinx-service/pkg/message/message"
)

func Listen() {
	if false {
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

func comsumeTransactionSucceeded(notification *msg.NotificationTransaction) error {
	logger.Sugar().Infof("good news everyone: %+w", notification)
	return nil
}
