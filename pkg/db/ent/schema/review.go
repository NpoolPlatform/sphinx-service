package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Review struct {
	ent.Schema
}

func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.Bool("is_approved").Default(false),
		field.String("operator_note").MaxLen(70),
		field.Int64("createtime_utc"),
		field.Int64("updatetime_utc"),
	}
}

func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("transaction", Transaction.Type).
			Ref("review").Unique().Required(),
		edge.From("coin", CoinInfo.Type).Ref("reviews").
			Unique().Required(),
	}
}

func (Review) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_approved"),
		index.Fields("createtime_utc"),
	}
}
