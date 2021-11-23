package message

import (
	"fmt"
	"reflect"

	msgsrv "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/server"
	"github.com/NpoolPlatform/message/npool/signproxy"    //nolint
	"github.com/NpoolPlatform/message/npool/sphinxplugin" //nolint
	constant "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
)

const (
	QueueAdminApprove = "admin-approve"
)

func InitQueues() (err error) {
	err = msgsrv.DeclareQueue(GetQueueName())
	return
}

type Example struct {
	ID      int    `json:"id"`
	Example string `json:"example"`
}

func GetQueueName() string {
	return fmt.Sprintf("%s::%s", constant.ServiceName, reflect.TypeOf(NotificationTransaction{}).String())
}

type NotificationTransaction struct {
	CoinType            sphinxplugin.CoinType     `json:"coin_type"`             // when: always
	TransactionType     signproxy.TransactionType `json:"transaction_type"`      // when: always
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
