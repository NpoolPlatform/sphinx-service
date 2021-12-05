package crud

import (
	"context"

	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	sconstant "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
)

func GetTransaction(ctx context.Context, in *trading.GetTransactionRequest) (resp *ent.Transaction, err error) {
	resp, err = db.Client().Transaction.Query().Where(
		transaction.TransactionID(in.TransactionID),
	).Only(ctx)
	return
}

type GetTransactionsParams struct {
	name   string
	state  transaction.Status
	from   string
	to     string
	offset int
	limit  int
}

func GetTransactions(ctx context.Context, params *GetTransactionsParams) ([]*ent.Transaction, int, error) {
	if params.limit == 0 {
		params.limit = sconstant.PageSize
	}

	stm := db.Client().
		Transaction.
		Query()

	if params.name != "" {
		stm.Where(transaction.NameEQ(params.name))
	}

	if params.from != "" {
		stm.Where(transaction.FromEQ(params.from))
	}

	if params.to != "" {
		stm.Where(transaction.ToEQ(params.to))
	}

	if params.state != "" {
		stm.Where(transaction.StatusEQ(params.state))
	}

	// total
	total, err := stm.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// infos
	trans, err := stm.
		Order(ent.Desc(transaction.FieldCreatedAt)).
		Offset(params.offset).
		Limit(params.limit).
		All(ctx)

	return trans, total, err
}
