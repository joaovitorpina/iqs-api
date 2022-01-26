package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Podcast holds the schema definition for the Podcast entity.
type Podcast struct {
	ent.Schema
}

// Fields of the Podcast.
func (Podcast) Fields() []ent.Field {
	return []ent.Field{
		field.String("titulo"),
		field.String("codigo"),
	}
}

// Edges of the Podcast.
func (Podcast) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profissional", Profissional.Type).
			Ref("podcasts").
			Unique(),
	}
}
