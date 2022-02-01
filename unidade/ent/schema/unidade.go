package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Unidade holds the schema definition for the Unidade entity.
type Unidade struct {
	ent.Schema
}

// Fields of the Unidade.
func (Unidade) Fields() []ent.Field {
	return []ent.Field{
		field.String("descricao"),
		field.String("url_amigavel"),
		field.Int64("endereco_id"),
		field.Int("latitude"),
		field.Int("longitude"),
		field.Int64("telefone").
			Optional(),
		field.Int64("celular").
			Optional(),
		field.String("email"),
		field.String("facebook").
			Optional(),
		field.String("instagram").
			Optional(),
		field.String("youtube").
			Optional(),
		field.Bool("ativo"),
	}
}

// Edges of the Unidade.
func (Unidade) Edges() []ent.Edge {
	return nil
}
