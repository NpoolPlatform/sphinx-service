package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type KeyStore struct {
	ent.Schema
}

func (KeyStore) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("address").MinLen(8).MaxLen(48),
		field.String("private_key").
			MinLen(8).
			MaxLen(80).
			Sensitive(),
	}
}

func (KeyStore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("coin", CoinInfo.Type).Ref("keys").
			Unique().Required(),
	}
}

func (KeyStore) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("address").
			Edges("coin").
			Unique(),
	}
}
