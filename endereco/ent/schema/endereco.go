package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Endereco holds the schema definition for the Endereco entity.
type Endereco struct {
	ent.Schema
}

// Fields of the Endereco.
func (Endereco) Fields() []ent.Field {
	return []ent.Field{
		field.String("numero").NotEmpty(),
	}
}

func (Endereco) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Endereco.
func (Endereco) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cep", Cep.Type).
			Ref("enderecos").
			Unique().
			Required(),
	}
}
