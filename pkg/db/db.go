package db

import (
	"context"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
)

var myClient *ent.Client

func Init() error {
	myClient = ent.NewClient(ent.Driver(app.Mysql().Driver))
	return myClient.Schema.Create(context.Background())
}

func Client() *ent.Client {
	return myClient
}
