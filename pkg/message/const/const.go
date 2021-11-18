package constant

import (
	"fmt"
	"time"
)

const (
	ServiceName   = "sphinx-service.npool.top"
	TradingDomain = "sphinx-service.npool.top"
	AgentDomain   = "sphinx-service.npool.top"
	AdminDomain   = "sphinx-service.npool.top"
	GrpcTimeout   = time.Second * 10
)

func GetMQChannel() string {
	return fmt.Sprintf("%s::channel", ServiceName)
}
