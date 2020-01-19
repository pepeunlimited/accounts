package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// OneToMany ´txs´
type Accounts struct {
	ent.Schema
}

func (Accounts) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("balance"),
		field.Uint8("version"),
		field.String("type").MaxLen(4),
		field.Bool("is_withdrawable"),
		field.Int64("user_id"),
	}
}

func (Accounts) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("txs", Txs.Type), // one-to-many
	}
}