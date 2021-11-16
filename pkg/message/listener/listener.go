package listener

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	msgcli "github.com/NpoolPlatform/sphinx-service/pkg/message/client"
	msg "github.com/NpoolPlatform/sphinx-service/pkg/message/message"
)

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

func Listen() {
	go listenTemplateExample()
}

func listenTransactionSucceeded() {
	for {
		logger.Sugar().Infof("listening for transaction success")
		err := msgcli.ComsumerOfAgent(func(notification *msg.NotificationTransaction) error {
			logger.Sugar().Infof("good news everyone: %+w", notification)
			// Call event handler in api module
			return nil
		})
		if err != nil {
			logger.Sugar().Errorf("fail to consume example: %v", err)
			return
		}
	}
}

func listenTransactionApprove() {
	for {
		logger.Sugar().Infof("listening for transaction approvals")
		err := msgcli.ComsumerOfAgent(func(notification *msg.NotificationTransaction) error {
			logger.Sugar().Infof("gonna auto approve tx: %+w", notification)
			// Call event handler in api module
			return nil
		})
		if err != nil {
			logger.Sugar().Errorf("fail to consume example: %v", err)
			return
		}
	}
}
