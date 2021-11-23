package client

import (
	"encoding/json"

	"golang.org/x/xerrors"

	msgcli "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/client"
	constant "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
	msg "github.com/NpoolPlatform/sphinx-service/pkg/message/message"

	"github.com/streadway/amqp"
)

type client struct {
	*msgcli.Client
	consumers map[string]<-chan amqp.Delivery
}

var myClients = map[string]*client{}

func Init() error {
	_myClient, err := msgcli.New(constant.ServiceName)
	if err != nil {
		return err
	}

	err = _myClient.DeclareQueue(msg.GetQueueName())
	if err != nil {
		return err
	}

	sampleClient := &client{
		Client:    _myClient,
		consumers: map[string]<-chan amqp.Delivery{},
	}
	examples, err := _myClient.Consume(msg.GetQueueName())
	if err != nil {
		return xerrors.Errorf("fail to construct initial default consume: %v", err)
	}
	sampleClient.consumers[msg.GetQueueName()] = examples

	myClients[constant.ServiceName] = sampleClient

	return nil
}

// Accept transaction status (success/failed) message
func ComsumerOfAgent(h func(*msg.NotificationTransaction) error) error {
	successTxs, ok := myClients[constant.ServiceName].consumers[msg.GetQueueName()]
	if !ok {
		return xerrors.Errorf("agent consumer is not constructed")
	}

	for d := range successTxs {
		tx := msg.NotificationTransaction{}
		err := json.Unmarshal(d.Body, &tx)
		if err != nil {
			return xerrors.Errorf("parse agent message error: %v", err)
		}

		if h != nil {
			err = h(&tx)
			if err != nil {
				return err
			}
		}
	}
	return xerrors.Errorf("WE SHOULD NOT BE HERE")
}
