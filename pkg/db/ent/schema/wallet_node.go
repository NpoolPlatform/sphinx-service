package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type WalletNode struct {
	ent.Schema
}

func (WalletNode) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("uuid").Unique(),
		field.String("location"),
		field.String("host_vendor"),
		field.String("public_ip"),
		field.String("local_ip"),
		field.Int64("createtime_utc"),
		field.Int64("last_online_time_utc"),
	}
}

func (WalletNode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("coin", CoinInfo.Type).Ref("wallet_nodes").
			Unique().Required(),
	}
}

func (WalletNode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid"),
		index.Fields("location"),
		index.Fields("host_vendor"),
		index.Fields("last_online_time_utc"),
	}
}
