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
			{
				Name:    "coininfo_id",
				Unique:  true,
				Columns: []*schema.Column{CoinInfosColumns[0]},
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
		{Name: "coin_info_reviews", Type: field.TypeUUID, Nullable: true},
		{Name: "transaction_review", Type: field.TypeInt32, Nullable: true},
	}
	// ReviewsTable holds the schema information for the "reviews" table.
	ReviewsTable = &schema.Table{
		Name:       "reviews",
		Columns:    ReviewsColumns,
		PrimaryKey: []*schema.Column{ReviewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reviews_coin_infos_reviews",
				Columns:    []*schema.Column{ReviewsColumns[5]},
				RefColumns: []*schema.Column{CoinInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reviews_transactions_review",
				Columns:    []*schema.Column{ReviewsColumns[6]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
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
				Name:    "transaction_transaction_id_insite_coin_info_transactions",
				Unique:  true,
				Columns: []*schema.Column{TransactionsColumns[7], TransactionsColumns[15]},
			},
			{
				Name:    "transaction_transaction_id_chain_coin_info_transactions",
				Unique:  true,
				Columns: []*schema.Column{TransactionsColumns[8], TransactionsColumns[15]},
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
		{Name: "coin_info_wallet_nodes", Type: field.TypeUUID, Nullable: true},
	}
	// WalletNodesTable holds the schema information for the "wallet_nodes" table.
	WalletNodesTable = &schema.Table{
		Name:       "wallet_nodes",
		Columns:    WalletNodesColumns,
		PrimaryKey: []*schema.Column{WalletNodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "wallet_nodes_coin_infos_wallet_nodes",
				Columns:    []*schema.Column{WalletNodesColumns[8]},
				RefColumns: []*schema.Column{CoinInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
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
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoinInfosTable,
		EmptiesTable,
		ReviewsTable,
		TransactionsTable,
		WalletNodesTable,
	}
)

func init() {
	ReviewsTable.ForeignKeys[0].RefTable = CoinInfosTable
	ReviewsTable.ForeignKeys[1].RefTable = TransactionsTable
	TransactionsTable.ForeignKeys[0].RefTable = CoinInfosTable
	WalletNodesTable.ForeignKeys[0].RefTable = CoinInfosTable
}
