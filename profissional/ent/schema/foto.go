package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Foto holds the schema definition for the Foto entity.
type Foto struct {
	ent.Schema
}

// Fields of the Foto.
func (Foto) Fields() []ent.Field {
	return []ent.Field{
		field.String("titulo"),
		field.String("url"),
	}
}

// Edges of the Foto.
func (Foto) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profissional", Profissional.Type).
			Ref("fotos").
			Unique().
			Required(),
	}
}
