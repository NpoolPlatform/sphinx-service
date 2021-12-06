package tasks

import (
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	constant "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
)

func init() {
	tasks["syncTransaction"] = syncTransaction
}

// syncTransaction ..
func syncTransaction() {
	for range time.NewTicker(constant.TaskDuration).C {
		func() {
			logger.Sugar().Infof("syncTransaction")
		}()
	}
}
