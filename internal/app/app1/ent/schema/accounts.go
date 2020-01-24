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
		field.Bool("is_verified"),
		field.Int64("user_id").Unique(),
	}
}

func (Accounts) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("txs", Txs.Type), // one-to-many
	}
}