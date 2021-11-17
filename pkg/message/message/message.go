package message

import (
	"fmt"
	"reflect"

	msgsrv "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/server"
	constant "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
)

const (
	QueueExample      = "example"
	QueueAdminApprove = "admin-approve"
	QueueAgent        = "agent"
	QueueTrading      = "trading"
)

func InitQueues() (err error) {
	if err = msgsrv.DeclareQueue(QueueExample); err != nil {
		return
	}
	if err = msgsrv.DeclareQueue(QueueAdminApprove); err != nil {
		return
	}
	if err = msgsrv.DeclareQueue(QueueAgent); err != nil {
		return
	}
	if err = msgsrv.DeclareQueue(QueueTrading); err != nil {
		return
	}
	return
}

type Example struct {
	ID      int    `json:"id"`
	Example string `json:"example"`
}

type NotificationTransaction struct {
	ID                  int32  `json:"id"`
	TransactionIDInsite string `json:"transaction_id_insite"`
	// TODO
}

func GetQueueName() string {
	return fmt.Sprintf("%s::%s", constant.ServiceName, reflect.TypeOf(NotificationTransaction{}).String())
}
