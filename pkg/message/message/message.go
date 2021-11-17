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
	TransactionIDInsite string `json:"transaction_id_insite"` // unique
	AmountUint64        uint64 `json:"amount_uint64"`         // need to be converted through public repo
	AddressFrom         string `json:"address_from"`
	AddressTo           string `json:"address_to"`
	TransactionIDChain  string `json:"transaction_id_chain"` // empty when created, return when finished
	SignatureUser       string `json:"signature_user"`       // preserved for 2FA verification, implement this in v2
	SignaturePlatform   string `json:"signature_platform"`   // preserved for signproxy to verify host, about v3
	CreatetimeUtc       int    `json:"createtime_utc"`       // for 2FA
	UpdatetimeUtc       int    `json:"updatetime_utc"`       // for return
	IsSuccess           bool   `json:"is_success"`           // return true when completed
	IsFailed            bool   `json:"is_failed"`            // return true when error occurred
}

func GetQueueName() string {
	return fmt.Sprintf("%s::%s", constant.ServiceName, reflect.TypeOf(NotificationTransaction{}).String())
}
