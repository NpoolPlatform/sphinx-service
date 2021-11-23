package main

import (
	"time"

	"github.com/NpoolPlatform/sphinx-service/api"
	db "github.com/NpoolPlatform/sphinx-service/pkg/db"
	msgcli "github.com/NpoolPlatform/sphinx-service/pkg/message/client"
	msglistener "github.com/NpoolPlatform/sphinx-service/pkg/message/listener"
	msg "github.com/NpoolPlatform/sphinx-service/pkg/message/message"
	msgsrv "github.com/NpoolPlatform/sphinx-service/pkg/message/server"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	cli "github.com/urfave/cli/v2"

	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		if err := db.Init(); err != nil {
			return err
		}

		go func() {
			if err := grpc2.RunGRPC(rpcRegister); err != nil {
				logger.Sugar().Errorf("fail to run grpc server: %v", err)
			}
		}()

		if err := logger.Init(logger.DebugLevel, "/tmp/sphinx-service.log"); err != nil {
			return err
		}

		if err := msgsrv.Init(); err != nil {
			return err
		}
		if err := msgcli.Init(); err != nil {
			return err
		}

		go msglistener.Listen(false)
		go msgSender()

		return grpc2.RunGRPCGateWay(rpcGatewayRegister)
	},
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)
	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return api.RegisterGateway(mux, endpoint, opts)
}

func msgSender() {
	id := 0
	for false {
		logger.Sugar().Infof("send example")
		err := msgsrv.PublishDefaultNotification(&msg.NotificationTransaction{
			CoinType:            0,
			TransactionType:     0,
			TransactionIDInsite: "",
			AmountFloat64:       0,
			AddressFrom:         "",
			AddressTo:           "",
			TransactionIDChain:  "",
			SignatureUser:       "",
			SignaturePlatform:   "",
			CreatetimeUtc:       0,
			UpdatetimeUtc:       0,
			IsSuccess:           false,
			IsFailed:            false,
		})
		if err != nil {
			logger.Sugar().Errorf("fail to send example: %v", err)
			return
		}
		id++
		time.Sleep(3 * time.Second)
	}
}
