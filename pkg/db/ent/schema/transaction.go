package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Transaction struct {
	ent.Schema
}

func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").NotEmpty().Default(""),
		field.Uint64("amount").Default(0),
		field.String("from").NotEmpty().Default(""),
		field.String("to").NotEmpty().Default(""),
		field.String("transaction_id").NotEmpty().
			Unique(),
		field.String("cid").Default(""),
		field.Enum("status").
			Values("pending_review", "confirm", "rejected", "pending_transaction", "done"),
		field.Uint32("created_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("updated_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("deleted_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

func (Transaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("transaction_id").
			Unique(),
		index.Fields("from"),
		index.Fields("to"),
		index.Fields("status"),
		index.Fields("created_at"),
		index.Fields("updated_at"),
	}
}
