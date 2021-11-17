package constant

import (
	"fmt"
	"time"
)

const (
	ServiceName   = "sphinx.npool.top"
	TradingDomain = "sphinx.npool.top"
	AgentDomain   = "sphinx.npool.top"
	AdminDomain   = "sphinx.npool.top"
	GrpcTimeout   = time.Second * 10
)

func GetMQChannel() string {
	return fmt.Sprintf("%s::channel", ServiceName)
}
