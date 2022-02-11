package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Profissional holds the schema definition for the Profissional entity.
type Profissional struct {
	ent.Schema
}

// Fields of the Profissional.
func (Profissional) Fields() []ent.Field {
	return []ent.Field{
		field.String("nome").NotEmpty(),
		field.String("url_amigavel").
			NotEmpty().
			Unique(),
		field.Bool("recomendado").
			Default(false),
		field.Bool("ativo").
			Default(true),
		field.Text("sobre").
			Optional(),
		field.String("conselho").
			Optional(),
		field.String("numero_identificacao").
			Optional(),
		field.Int64("telefone").
			Optional(),
		field.Int64("celular").
			Optional(),
		field.String("email").
			Optional(),
		field.String("site").
			Optional(),
		field.String("facebook").
			Optional(),
		field.String("instagram").
			Optional(),
		field.String("youtube").
			Optional(),
		field.String("linkedin").
			Optional(),
		field.Int("unidade_id").
			Min(1),
		field.Int("endereco_id").
			Min(1),
		field.String("imagem_perfil_url"),
	}
}

func (Profissional) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Profissional.
func (Profissional) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("whatsapps", WhatsApp.Type),
		edge.To("videos", Video.Type),
		edge.To("tratamentos", Tratamento.Type),
		edge.To("podcasts", Podcast.Type),
		edge.To("fotos", Foto.Type),
		edge.To("convenios", Convenio.Type),
		edge.From("especializacoes", Especializacao.Type).
			Ref("profissionais"),
	}
}
