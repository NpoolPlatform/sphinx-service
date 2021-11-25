package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Transaction struct {
	ent.Schema
}

func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.Uint64("amount_uint64"),
		field.Float("amount_float64"),
		field.String("address_from").MaxLen(64).NotEmpty(),
		field.String("address_to").MaxLen(64).NotEmpty(),
		field.Bool("need_manual_review").Default(true),
		field.Enum("type").
			Values("recharge", "payment", "withdraw", "unknown"),
		field.String("transaction_id_insite").
			MaxLen(64).Unique(),
		field.String("transaction_id_chain").
			MaxLen(80),
		field.Enum("status").
			Values("pending_review", "pending_process", "pending_signinfo", "pending_sign", "pending_broadcast", "pending_confirm", "done", "rejected", "error", "error_expected"),
		field.Bool("mutex").
			Default(false),
		field.String("signature_user").MaxLen(16),
		field.String("signature_platform").MaxLen(64),
		field.Int64("createtime_utc"),
		field.Int64("updatetime_utc"),
	}
}

func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("coin", CoinInfo.Type).Ref("transactions").Unique(),
		edge.To("review", Review.Type),
	}
}

func (Transaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("address_from"),
		index.Fields("address_to"),
		index.Fields("type"),
		index.Fields("status"),
		index.Fields("createtime_utc"),
		index.Fields("transaction_id_insite").Unique(),
		index.Fields("transaction_id_chain"),
	}
}
