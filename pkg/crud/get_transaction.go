package crud

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxservice"
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	sconstant "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
)

func GetTransaction(ctx context.Context, in *sphinxservice.GetTransactionRequest) (resp *ent.Transaction, err error) {
	resp, err = db.Client().Transaction.Query().Where(
		transaction.TransactionID(in.TransactionID),
	).Only(ctx)
	return
}

type GetTransactionsParams struct {
	Name   string
	State  transaction.Status
	From   string
	To     string
	Offset int
	Limit  int
}

func GetTransactions(ctx context.Context, params *GetTransactionsParams) ([]*ent.Transaction, int, error) {
	if params.Limit == 0 {
		params.Limit = sconstant.PageSize
	}

	stm := db.Client().
		Transaction.
		Query()

	if params.Name != "" {
		stm.Where(transaction.NameEQ(params.Name))
	}

	if params.From != "" {
		stm.Where(transaction.FromEQ(params.From))
	}

	if params.To != "" {
		stm.Where(transaction.ToEQ(params.To))
	}

	if len(params.State) > 0 {
		stm.Where(transaction.StatusEQ(params.State))
	}

	// total
	total, err := stm.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// infos
	trans, err := stm.
		Order(ent.Asc(transaction.FieldCreatedAt)).
		Offset(params.Offset).
		Limit(params.Limit).
		All(ctx)

	return trans, total, err
}
