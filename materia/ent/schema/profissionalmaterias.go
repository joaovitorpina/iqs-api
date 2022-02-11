package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProfissionalMaterias holds the schema definition for the ProfissionalMaterias entity.
type ProfissionalMaterias struct {
	ent.Schema
}

// Fields of the ProfissionalMaterias.
func (ProfissionalMaterias) Fields() []ent.Field {
	return []ent.Field{
		field.Int("profissional_id").
			Positive(),
	}
}

// Edges of the ProfissionalMaterias.
func (ProfissionalMaterias) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("materia", Materia.Type).
			Ref("profissional_materias").
			Unique().
			Required(),
	}
}
