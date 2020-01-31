package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// ManyToOne ´accounts´
type Txs struct {
	ent.Schema
}

func (Txs) Config() ent.Config {
	return ent.Config{Table:"txs"}
}

func (Txs) Fields() []ent.Field {
	return []ent.Field {
		field.String("tx_type").MaxLen(10).NotEmpty(),
		field.Time("created_at"),
		field.Int64("amount"),
		field.String("reference_number").Optional().Nillable().MaxLen(36),
	}
}

func (Txs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("accounts", Accounts.Type).Ref("txs").Unique(), // many-to-one
	}
}