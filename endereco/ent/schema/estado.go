package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Estado holds the schema definition for the Estado entity.
type Estado struct {
	ent.Schema
}

// Fields of the Estado.
func (Estado) Fields() []ent.Field {
	return []ent.Field{
		field.String("nome").NotEmpty(),
	}
}

// Edges of the Estado.
func (Estado) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cidades", Cidade.Type),
	}
}
