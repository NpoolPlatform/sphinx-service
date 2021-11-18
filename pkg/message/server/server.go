package server

import (
	msgsrv "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/server"
	msg "github.com/NpoolPlatform/sphinx-service/pkg/message/message"
)

func Init() error {
	return msg.InitQueues()
}

func PublishExample(example *msg.Example) error {
	return msgsrv.PublishToQueue(msg.QueueExample, example)
}

func PublishNotificationTransactionCreate(notification *msg.NotificationTransaction) error {
	return msgsrv.PublishToQueue(msg.QueueAdminApprove, notification)
}

func PublishNotificationTransactionProcess(notification *msg.NotificationTransaction) error {
	return msgsrv.PublishToQueue(msg.QueueAgent, notification)
}

func PublishDefaultNotification(notification *msg.NotificationTransaction) error {
	return msgsrv.PublishToQueue(msg.GetQueueName(), notification)
}
