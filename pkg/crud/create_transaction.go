package crud

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
)

type CreateTransactionParams struct {
	TransactionID string
	Name          string
	Amount        float64
	From          string
	To            string
}

func CreateTransaction(ctx context.Context, params CreateTransactionParams) (info *ent.Transaction, err error) {
	return db.Client().Transaction.Create().
		SetTransactionID(params.TransactionID).
		SetName(params.Name).
		SetAmount(price.VisualPriceToDBPrice(params.Amount)).
		SetFrom(params.From).
		SetTo(params.To).
		SetStatus(transaction.StatusPendingTransaction).
		Save(ctx)
}
