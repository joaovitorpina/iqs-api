package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Cep holds the schema definition for the Cep entity.
type Cep struct {
	ent.Schema
}

// Fields of the Cep.
func (Cep) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").
			Unique().
			Immutable().
			NonNegative().
			Max(99999999),
		field.String("logradouro").
			NotEmpty(),
		field.String("bairro").
			NotEmpty(),
	}
}

// Edges of the Cep.
func (Cep) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cidade", Cidade.Type).
			Ref("ceps").
			Unique(),
		edge.To("enderecos", Endereco.Type),
	}
}
