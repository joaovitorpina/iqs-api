// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"profissional/ent/podcast"
	"profissional/ent/profissional"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Podcast is the model entity for the Podcast schema.
type Podcast struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Titulo holds the value of the "titulo" field.
	Titulo string `json:"titulo,omitempty"`
	// Codigo holds the value of the "codigo" field.
	Codigo string `json:"codigo,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PodcastQuery when eager-loading is set.
	Edges                 PodcastEdges `json:"edges"`
	profissional_podcasts *int
}

// PodcastEdges holds the relations/edges for other nodes in the graph.
type PodcastEdges struct {
	// Profissional holds the value of the profissional edge.
	Profissional *Profissional `json:"profissional,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProfissionalOrErr returns the Profissional value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PodcastEdges) ProfissionalOrErr() (*Profissional, error) {
	if e.loadedTypes[0] {
		if e.Profissional == nil {
			// The edge profissional was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: profissional.Label}
		}
		return e.Profissional, nil
	}
	return nil, &NotLoadedError{edge: "profissional"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Podcast) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case podcast.FieldID:
			values[i] = new(sql.NullInt64)
		case podcast.FieldTitulo, podcast.FieldCodigo:
			values[i] = new(sql.NullString)
		case podcast.ForeignKeys[0]: // profissional_podcasts
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Podcast", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Podcast fields.
func (po *Podcast) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case podcast.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case podcast.FieldTitulo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field titulo", values[i])
			} else if value.Valid {
				po.Titulo = value.String
			}
		case podcast.FieldCodigo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field codigo", values[i])
			} else if value.Valid {
				po.Codigo = value.String
			}
		case podcast.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field profissional_podcasts", value)
			} else if value.Valid {
				po.profissional_podcasts = new(int)
				*po.profissional_podcasts = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryProfissional queries the "profissional" edge of the Podcast entity.
func (po *Podcast) QueryProfissional() *ProfissionalQuery {
	return (&PodcastClient{config: po.config}).QueryProfissional(po)
}

// Update returns a builder for updating this Podcast.
// Note that you need to call Podcast.Unwrap() before calling this method if this Podcast
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Podcast) Update() *PodcastUpdateOne {
	return (&PodcastClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Podcast entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Podcast) Unwrap() *Podcast {
	tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Podcast is not a transactional entity")
	}
	po.config.driver = tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Podcast) String() string {
	var builder strings.Builder
	builder.WriteString("Podcast(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteString(", titulo=")
	builder.WriteString(po.Titulo)
	builder.WriteString(", codigo=")
	builder.WriteString(po.Codigo)
	builder.WriteByte(')')
	return builder.String()
}

// Podcasts is a parsable slice of Podcast.
type Podcasts []*Podcast

func (po Podcasts) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
