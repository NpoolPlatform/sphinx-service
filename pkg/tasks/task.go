package tasks

import "github.com/NpoolPlatform/go-service-framework/pkg/logger"

var tasks = make(map[string]func())

func Run() {
	for tname, task := range tasks {
		logger.Sugar().Infof("run task: %v", tname)
		go task()
	}
}
