package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Categoria holds the schema definition for the Categoria entity.
type Categoria struct {
	ent.Schema
}

// Fields of the Categoria.
func (Categoria) Fields() []ent.Field {
	return []ent.Field{
		field.String("descricao"),
	}
}

// Edges of the Categoria.
func (Categoria) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("materias", Materia.Type),
	}
}
