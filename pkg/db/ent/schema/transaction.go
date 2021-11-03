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
		field.Int("amount_int").Positive(),
		field.Int("amount_digits").NonNegative().Default(9),
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
			Values("pending_review", "pending_process", "pending_signinfo", "pending_signaction", "pending_broadcast", "pending_confirm", "done", "rejected", "error", "error_expected"),
		field.Bool("mutex").
			Default(false),
		field.Int("createtime_utc"),
		field.Int("updatetime_utc"),
	}
}

func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("coin", CoinInfo.Type).Ref("transactions").
			Unique().Required(),
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
		index.Fields("transaction_id_insite").
			Edges("coin").Unique(),
		index.Fields("transaction_id_chain").
			Edges("coin").Unique(),
	}
}
