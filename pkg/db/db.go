package db

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/migrate"
)

var myClient *ent.Client

func Migrate() {
	ctx := context.Background()
	err := myClient.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		logger.Sugar().Errorf("failed creating schema resources: %v", err)
	}
}

func Init() (err error) {
	myClient = ent.NewClient(ent.Driver(app.Mysql().Driver))
	err = myClient.Schema.Create(context.Background())
	if err == nil {
		go Migrate()
		logger.Sugar().Info("Database migrating")
	}
	return
}

func Client() *ent.Client {
	return myClient
}
