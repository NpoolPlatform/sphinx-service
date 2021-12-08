package crud

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
)

type UpdateTransactionStatusParams struct {
	TransactionID string
	State         transaction.Status
	CID           string
	ExitCode      int64
}

func UpdateTransactionStatus(ctx context.Context, params UpdateTransactionStatusParams) error {
	stm := db.Client().
		Transaction.
		Update().
		Where(transaction.TransactionIDEQ(params.TransactionID))

	if transaction.StatusDone == params.State {
		stm.
			SetCid(params.CID).
			SetExitCode(params.ExitCode)
	}

	stm.SetStatus(params.State)
	return stm.Exec(ctx)
}
