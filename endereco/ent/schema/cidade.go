package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Cidade holds the schema definition for the Cidade entity.
type Cidade struct {
	ent.Schema
}

// Fields of the Cidade.
func (Cidade) Fields() []ent.Field {
	return []ent.Field{
		field.String("nome").NotEmpty(),
	}
}

// Edges of the Cidade.
func (Cidade) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("estado", Estado.Type).
			Ref("cidades").
			Unique(),
		edge.To("ceps", Cep.Type),
	}
}
