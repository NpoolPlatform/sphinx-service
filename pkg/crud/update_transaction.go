package crud

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
)

type UpdateTransactionStatusParams struct {
	TransactionID string
	State         transaction.Status
}

func UpdateTransactionStatus(ctx context.Context, params UpdateTransactionStatusParams) error {
	return db.Client().
		Transaction.
		Update().
		Where(transaction.TransactionIDEQ(params.TransactionID)).
		SetStatus(params.State).
		Exec(ctx)
}
