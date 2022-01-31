package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tratamento holds the schema definition for the Tratamento entity.
type Tratamento struct {
	ent.Schema
}

// Fields of the Tratamento.
func (Tratamento) Fields() []ent.Field {
	return []ent.Field{
		field.String("descricao"),
	}
}

// Edges of the Tratamento.
func (Tratamento) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profissional", Profissional.Type).
			Ref("tratamentos").
			Unique().
			Required(),
	}
}
