package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/app"
)

// 接收异步返回
func (s *Server) ACK(ctx context.Context, in *trading.ACKRequest) (resp *trading.ACKResponse, err error) {
	resp, err = app.ACK(ctx, in)
	if err != nil {
		err = PatchGRPCError(err, "ack response error")
		resp = &trading.ACKResponse{}
	}
	logger.Sugar().Infof("[ack] received req: %+w", in)
	return
}
