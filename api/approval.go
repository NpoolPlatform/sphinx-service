package api

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/sphinx-service/pkg/approval"
	"github.com/NpoolPlatform/sphinx-service/pkg/message/message"
)

// deprecated
func ApproveTransaction(transactionIDInsite string) (isApproved bool, err error) {
	isApproved, err = approval.ApproveTransaction(transactionIDInsite)
	if err != nil {
		logger.Sugar().Errorw("approve transaction error: %w", err)
		if DebugFlag {
			err = errInternal
		}
	}
	return
}

// as well
func NotifyProcessTransaction(notification *message.NotificationTransaction) {}
