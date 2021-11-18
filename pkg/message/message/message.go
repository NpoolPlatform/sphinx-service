package message

import (
	"fmt"
	"reflect"

	msgsrv "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/server"
	"github.com/NpoolPlatform/message/npool/signproxy"
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
	if err = msgsrv.DeclareQueue(GetQueueName()); err != nil {
		return
	}
	return
}

type Example struct {
	ID      int    `json:"id"`
	Example string `json:"example"`
}

type NotificationTransaction struct {
	TransactionType     signproxy.TransactionType `json:"transaction_type"`      // when: always
	CoinType            signproxy.CoinType        `json:"coin_type"`             // when: always
	UUID                string                    `json:"uuid"`                  // when: create account; usage: for trading service to locate request when get return
	TransactionIDInsite string                    `json:"transaction_id_insite"` // when: Transaction; unique
	AmountFloat64       float64                   `json:"amount_float64"`        // when: Transaction
	AddressFrom         string                    `json:"address_from"`          // when: Transaction
	AddressTo           string                    `json:"address_to"`            // when: Transaction
	TransactionIDChain  string                    `json:"transaction_id_chain"`  // when: Transaction empty when created, return when finished
	SignatureUser       string                    `json:"signature_user"`        // when: Transaction preserved for 2FA verification, implement this in v2
	SignaturePlatform   string                    `json:"signature_platform"`    // when: Transaction preserved for signproxy to verify host, about v3
	CreatetimeUtc       int64                     `json:"createtime_utc"`        // when: Transaction for 2FA
	UpdatetimeUtc       int64                     `json:"updatetime_utc"`        // when: Transaction for return
	IsSuccess           bool                      `json:"is_success"`            // when: Transaction return true when completed
	IsFailed            bool                      `json:"is_failed"`             // when: Transaction return true when error occurred
}

func GetQueueName() string {
	return fmt.Sprintf("%s::%s", constant.ServiceName, reflect.TypeOf(NotificationTransaction{}).String())
}
