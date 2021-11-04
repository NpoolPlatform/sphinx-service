package core

import (
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
)



var Client *ent.Client

func init() {
	Client = db.Client()
}
