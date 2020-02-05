package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// OneToMany ´txs´
type Account struct {
	ent.Schema
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("balance"),
		field.Uint8("version"),
		field.Bool("is_verified"),
		field.Int64("user_id").Unique(),
	}
}

func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("txs", Txs.Type), // one-to-many
	}
}