package constant

import (
	"fmt"
	"time"
)

const (
	ServiceName = "sphinx-service.npool.top"
)

const (
	GrpcTimeout  = time.Second * 10
	TaskDuration = time.Second * 1

	PageSize = 100
)

func GetMQChannel() string {
	return fmt.Sprintf("%s::channel", ServiceName)
}
