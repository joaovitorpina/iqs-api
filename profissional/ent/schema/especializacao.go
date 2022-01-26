package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Especializacao holds the schema definition for the Especializacao entity.
type Especializacao struct {
	ent.Schema
}

// Fields of the Especializacao.
func (Especializacao) Fields() []ent.Field {
	return []ent.Field{
		field.String("descricao"),
	}
}

// Edges of the Especializacao.
func (Especializacao) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("areasaude", AreaSaude.Type).
			Ref("especializacoes").
			Unique(),
		edge.To("profissionais", Profissional.Type),
	}
}
