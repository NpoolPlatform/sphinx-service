// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoinInfosColumns holds the columns for the "coin_infos" table.
	CoinInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeInt32},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "unit", Type: field.TypeString},
		{Name: "is_presale", Type: field.TypeBool, Default: false},
		{Name: "logo_image", Type: field.TypeString, Default: ""},
	}
	// CoinInfosTable holds the schema information for the "coin_infos" table.
	CoinInfosTable = &schema.Table{
		Name:       "coin_infos",
		Columns:    CoinInfosColumns,
		PrimaryKey: []*schema.Column{CoinInfosColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "coininfo_name",
				Unique:  true,
				Columns: []*schema.Column{CoinInfosColumns[2]},
			},
			{
				Name:    "coininfo_unit",
				Unique:  false,
				Columns: []*schema.Column{CoinInfosColumns[3]},
			},
		},
	}
	// EmptiesColumns holds the columns for the "empties" table.
	EmptiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// EmptiesTable holds the schema information for the "empties" table.
	EmptiesTable = &schema.Table{
		Name:       "empties",
		Columns:    EmptiesColumns,
		PrimaryKey: []*schema.Column{EmptiesColumns[0]},
	}
	// ReviewsColumns holds the columns for the "reviews" table.
	ReviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "is_approved", Type: field.TypeBool, Default: false},
		{Name: "operator_note", Type: field.TypeString, Size: 70},
		{Name: "createtime_utc", Type: field.TypeInt64},
		{Name: "updatetime_utc", Type: field.TypeInt64},
	}
	// ReviewsTable holds the schema information for the "reviews" table.
	ReviewsTable = &schema.Table{
		Name:       "reviews",
		Columns:    ReviewsColumns,
		PrimaryKey: []*schema.Column{ReviewsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "review_is_approved",
				Unique:  false,
				Columns: []*schema.Column{ReviewsColumns[1]},
			},
			{
				Name:    "review_createtime_utc",
				Unique:  false,
				Columns: []*schema.Column{ReviewsColumns[3]},
			},
		},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "amount_uint64", Type: field.TypeUint64},
		{Name: "amount_float64", Type: field.TypeFloat64},
		{Name: "address_from", Type: field.TypeString, Size: 64},
		{Name: "address_to", Type: field.TypeString, Size: 64},
		{Name: "need_manual_review", Type: field.TypeBool, Default: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"recharge", "payment", "withdraw", "unknown"}},
		{Name: "transaction_id_insite", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "transaction_id_chain", Type: field.TypeString, Size: 80},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending_review", "pending_process", "pending_signinfo", "pending_sign", "pending_broadcast", "pending_confirm", "done", "rejected", "error", "error_expected"}},
		{Name: "mutex", Type: field.TypeBool, Default: false},
		{Name: "signature_user", Type: field.TypeString, Size: 16},
		{Name: "signature_platform", Type: field.TypeString, Size: 64},
		{Name: "createtime_utc", Type: field.TypeInt64},
		{Name: "updatetime_utc", Type: field.TypeInt64},
		{Name: "coin_info_transactions", Type: field.TypeUUID, Nullable: true},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_coin_infos_transactions",
				Columns:    []*schema.Column{TransactionsColumns[15]},
				RefColumns: []*schema.Column{CoinInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "transaction_address_from",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[3]},
			},
			{
				Name:    "transaction_address_to",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[4]},
			},
			{
				Name:    "transaction_type",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[6]},
			},
			{
				Name:    "transaction_status",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[9]},
			},
			{
				Name:    "transaction_createtime_utc",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[13]},
			},
			{
				Name:    "transaction_transaction_id_insite",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[7]},
			},
			{
				Name:    "transaction_transaction_id_chain",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[8]},
			},
		},
	}
	// WalletNodesColumns holds the columns for the "wallet_nodes" table.
	WalletNodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "uuid", Type: field.TypeString, Unique: true},
		{Name: "location", Type: field.TypeString},
		{Name: "host_vendor", Type: field.TypeString},
		{Name: "public_ip", Type: field.TypeString},
		{Name: "local_ip", Type: field.TypeString},
		{Name: "createtime_utc", Type: field.TypeInt64},
		{Name: "last_online_time_utc", Type: field.TypeInt64},
	}
	// WalletNodesTable holds the schema information for the "wallet_nodes" table.
	WalletNodesTable = &schema.Table{
		Name:       "wallet_nodes",
		Columns:    WalletNodesColumns,
		PrimaryKey: []*schema.Column{WalletNodesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "walletnode_uuid",
				Unique:  false,
				Columns: []*schema.Column{WalletNodesColumns[1]},
			},
			{
				Name:    "walletnode_location",
				Unique:  false,
				Columns: []*schema.Column{WalletNodesColumns[2]},
			},
			{
				Name:    "walletnode_host_vendor",
				Unique:  false,
				Columns: []*schema.Column{WalletNodesColumns[3]},
			},
			{
				Name:    "walletnode_last_online_time_utc",
				Unique:  false,
				Columns: []*schema.Column{WalletNodesColumns[7]},
			},
		},
	}
	// CoinInfoReviewsColumns holds the columns for the "coin_info_reviews" table.
	CoinInfoReviewsColumns = []*schema.Column{
		{Name: "coin_info_id", Type: field.TypeUUID},
		{Name: "review_id", Type: field.TypeInt32},
	}
	// CoinInfoReviewsTable holds the schema information for the "coin_info_reviews" table.
	CoinInfoReviewsTable = &schema.Table{
		Name:       "coin_info_reviews",
		Columns:    CoinInfoReviewsColumns,
		PrimaryKey: []*schema.Column{CoinInfoReviewsColumns[0], CoinInfoReviewsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "coin_info_reviews_coin_info_id",
				Columns:    []*schema.Column{CoinInfoReviewsColumns[0]},
				RefColumns: []*schema.Column{CoinInfosColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "coin_info_reviews_review_id",
				Columns:    []*schema.Column{CoinInfoReviewsColumns[1]},
				RefColumns: []*schema.Column{ReviewsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CoinInfoWalletNodesColumns holds the columns for the "coin_info_wallet_nodes" table.
	CoinInfoWalletNodesColumns = []*schema.Column{
		{Name: "coin_info_id", Type: field.TypeUUID},
		{Name: "wallet_node_id", Type: field.TypeInt32},
	}
	// CoinInfoWalletNodesTable holds the schema information for the "coin_info_wallet_nodes" table.
	CoinInfoWalletNodesTable = &schema.Table{
		Name:       "coin_info_wallet_nodes",
		Columns:    CoinInfoWalletNodesColumns,
		PrimaryKey: []*schema.Column{CoinInfoWalletNodesColumns[0], CoinInfoWalletNodesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "coin_info_wallet_nodes_coin_info_id",
				Columns:    []*schema.Column{CoinInfoWalletNodesColumns[0]},
				RefColumns: []*schema.Column{CoinInfosColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "coin_info_wallet_nodes_wallet_node_id",
				Columns:    []*schema.Column{CoinInfoWalletNodesColumns[1]},
				RefColumns: []*schema.Column{WalletNodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TransactionReviewColumns holds the columns for the "transaction_review" table.
	TransactionReviewColumns = []*schema.Column{
		{Name: "transaction_id", Type: field.TypeInt32},
		{Name: "review_id", Type: field.TypeInt32},
	}
	// TransactionReviewTable holds the schema information for the "transaction_review" table.
	TransactionReviewTable = &schema.Table{
		Name:       "transaction_review",
		Columns:    TransactionReviewColumns,
		PrimaryKey: []*schema.Column{TransactionReviewColumns[0], TransactionReviewColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transaction_review_transaction_id",
				Columns:    []*schema.Column{TransactionReviewColumns[0]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "transaction_review_review_id",
				Columns:    []*schema.Column{TransactionReviewColumns[1]},
				RefColumns: []*schema.Column{ReviewsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoinInfosTable,
		EmptiesTable,
		ReviewsTable,
		TransactionsTable,
		WalletNodesTable,
		CoinInfoReviewsTable,
		CoinInfoWalletNodesTable,
		TransactionReviewTable,
	}
)

func init() {
	TransactionsTable.ForeignKeys[0].RefTable = CoinInfosTable
	CoinInfoReviewsTable.ForeignKeys[0].RefTable = CoinInfosTable
	CoinInfoReviewsTable.ForeignKeys[1].RefTable = ReviewsTable
	CoinInfoWalletNodesTable.ForeignKeys[0].RefTable = CoinInfosTable
	CoinInfoWalletNodesTable.ForeignKeys[1].RefTable = WalletNodesTable
	TransactionReviewTable.ForeignKeys[0].RefTable = TransactionsTable
	TransactionReviewTable.ForeignKeys[1].RefTable = ReviewsTable
}
