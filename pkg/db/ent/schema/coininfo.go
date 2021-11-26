package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type CoinInfo struct {
	ent.Schema
}

func (CoinInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Int32("coin_type_id"),
		field.String("name").NotEmpty().Unique(),
		field.String("unit").NotEmpty(),
		field.Bool("is_presale").Default(false),
		field.String("logo_image").Default(""),
	}
}

func (CoinInfo) Edges() []ent.Edge {
	return []ent.Edge{
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
