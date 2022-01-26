package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AreaSaude holds the schema definition for the AreaSaude entity.
type AreaSaude struct {
	ent.Schema
}

// Fields of the AreaSaude.
func (AreaSaude) Fields() []ent.Field {
	return []ent.Field{
		field.String("descricao"),
	}
}

// Edges of the AreaSaude.
func (AreaSaude) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("especializacoes", Especializacao.Type),
	}
}
