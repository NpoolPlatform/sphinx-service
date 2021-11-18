package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type CoinInfo struct {
	ent.Schema
}

func (CoinInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("name").NotEmpty().MaxLen(16).Unique(),
		field.String("unit").NotEmpty().MaxLen(4),
		field.Bool("is_presale").Default(false),
	}
}

func (CoinInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("keys", KeyStore.Type),
		edge.To("transactions", Transaction.Type),
		edge.To("reviews", Review.Type),
		edge.To("wallet_nodes", WalletNode.Type),
	}
}

func (CoinInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
		index.Fields("unit"),
	}
}
