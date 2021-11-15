package message

import (
	msgsrv "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/server"
)

const (
	QueueExample      = "example"
	QueueAdminApprove = "admin-approve"
	QueueAgent        = "agent"
	QueueTrading      = "trading"
)

func InitQueues() error {
	err := msgsrv.DeclareQueue(QueueExample)
	if err != nil {
		return err
	}
	msgsrv.DeclareQueue(QueueAdminApprove)
	msgsrv.DeclareQueue(QueueAgent)
	msgsrv.DeclareQueue(QueueTrading)
	return nil
}

type Example struct {
	ID      int    `json:"id"`
	Example string `json:"example"`
}

type NotificationTransaction struct {
	ID                  int32  `json:"id"`
	TransactionIDInsite string `json:"transaction_id_insite"`
}
