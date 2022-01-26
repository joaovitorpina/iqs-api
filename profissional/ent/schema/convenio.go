package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Convenio holds the schema definition for the Convenio entity.
type Convenio struct {
	ent.Schema
}

// Fields of the Convenio.
func (Convenio) Fields() []ent.Field {
	return []ent.Field{
		field.String("nome").Unique(),
	}
}

// Edges of the Convenio.
func (Convenio) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profissional", Profissional.Type).
			Ref("convenios").
			Unique(),
	}
}
