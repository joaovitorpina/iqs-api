package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WhatsApp holds the schema definition for the WhatsApp entity.
type WhatsApp struct {
	ent.Schema
}

// Fields of the WhatsApp.
func (WhatsApp) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("numero"),
		field.Bool("principal").
			Default(false),
	}
}

// Edges of the WhatsApp.
func (WhatsApp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profissional", Profissional.Type).
			Ref("whatsapps").
			Unique().
			Required(),
	}
}
