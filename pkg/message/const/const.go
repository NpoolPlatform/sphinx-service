package constant

import (
	"fmt"
	"time"
)

const (
	ServiceName = "sphinx-service.npool.top"
)

const (
	GrpcTimeout = time.Second * 10
	PageSize    = 10
)

func GetMQChannel() string {
	return fmt.Sprintf("%s::channel", ServiceName)
}
