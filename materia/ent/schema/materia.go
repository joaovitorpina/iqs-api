package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Materia holds the schema definition for the Materia entity.
type Materia struct {
	ent.Schema
}

// Fields of the Materia.
func (Materia) Fields() []ent.Field {
	return []ent.Field{
		field.Int("unidade_id"),
		field.String("titulo"),
		field.String("url_amigavel").
			Unique(),
		field.Time("data_agendamento").
			Optional(),
		field.String("fonte").
			Optional(),
		field.String("link_fonte").
			Optional(),
		field.Text("texto"),
		field.String("imagem_url").
			Optional(),
		field.Bool("patrocinado").
			Default(false),
		field.Bool("status").
			Default(true),
	}
}

// Edges of the Materia.
func (Materia) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("categoria", Categoria.Type).
			Ref("materias").
			Required().
			Unique(),
		edge.To("profissional_materias", ProfissionalMaterias.Type),
	}
}
