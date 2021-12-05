package crud

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
)

type UpdateTransactionStatusParams struct {
	transactionID string
	state         transaction.Status
}

func UpdateTransactionStatus(ctx context.Context, params UpdateTransactionStatusParams) error {
	return db.Client().
		Transaction.
		Update().
		Where(transaction.TransactionIDEQ(params.transactionID)).
		SetStatus(params.state).
		Exec(ctx)
}
