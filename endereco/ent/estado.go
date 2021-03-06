// Code generated by entc, DO NOT EDIT.

package ent

import (
	"endereco/ent/estado"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Estado is the model entity for the Estado schema.
type Estado struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Nome holds the value of the "nome" field.
	Nome string `json:"nome,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EstadoQuery when eager-loading is set.
	Edges EstadoEdges `json:"edges"`
}

// EstadoEdges holds the relations/edges for other nodes in the graph.
type EstadoEdges struct {
	// Cidades holds the value of the cidades edge.
	Cidades []*Cidade `json:"cidades,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CidadesOrErr returns the Cidades value or an error if the edge
// was not loaded in eager-loading.
func (e EstadoEdges) CidadesOrErr() ([]*Cidade, error) {
	if e.loadedTypes[0] {
		return e.Cidades, nil
	}
	return nil, &NotLoadedError{edge: "cidades"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Estado) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case estado.FieldID:
			values[i] = new(sql.NullInt64)
		case estado.FieldNome:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Estado", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Estado fields.
func (e *Estado) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case estado.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case estado.FieldNome:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nome", values[i])
			} else if value.Valid {
				e.Nome = value.String
			}
		}
	}
	return nil
}

// QueryCidades queries the "cidades" edge of the Estado entity.
func (e *Estado) QueryCidades() *CidadeQuery {
	return (&EstadoClient{config: e.config}).QueryCidades(e)
}

// Update returns a builder for updating this Estado.
// Note that you need to call Estado.Unwrap() before calling this method if this Estado
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Estado) Update() *EstadoUpdateOne {
	return (&EstadoClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Estado entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Estado) Unwrap() *Estado {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Estado is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Estado) String() string {
	var builder strings.Builder
	builder.WriteString("Estado(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", nome=")
	builder.WriteString(e.Nome)
	builder.WriteByte(')')
	return builder.String()
}

// Estados is a parsable slice of Estado.
type Estados []*Estado

func (e Estados) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
