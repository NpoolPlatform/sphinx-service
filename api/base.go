package api

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const FlagPrintError = true

// only for api layer
// in apps please use xerrors instead
func PatchGRPCError(errOri error, msg string) (err error) {
	if err != nil {
		logger.Sugar().Warn(msg, errOri)
		if !FlagPrintError {
			err = status.Error(codes.Internal, "[api] internal server error")
		} else {
			err = status.Errorf(codes.Internal, "[api] %v, err: %v", msg, errOri)
		}
	}
	return
}
