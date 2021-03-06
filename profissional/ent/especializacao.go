// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"profissional/ent/areasaude"
	"profissional/ent/especializacao"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Especializacao is the model entity for the Especializacao schema.
type Especializacao struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Descricao holds the value of the "descricao" field.
	Descricao string `json:"descricao,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EspecializacaoQuery when eager-loading is set.
	Edges                      EspecializacaoEdges `json:"edges"`
	area_saude_especializacoes *int
}

// EspecializacaoEdges holds the relations/edges for other nodes in the graph.
type EspecializacaoEdges struct {
	// Areasaude holds the value of the areasaude edge.
	Areasaude *AreaSaude `json:"areasaude,omitempty"`
	// Profissionais holds the value of the profissionais edge.
	Profissionais []*Profissional `json:"profissionais,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AreasaudeOrErr returns the Areasaude value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EspecializacaoEdges) AreasaudeOrErr() (*AreaSaude, error) {
	if e.loadedTypes[0] {
		if e.Areasaude == nil {
			// The edge areasaude was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: areasaude.Label}
		}
		return e.Areasaude, nil
	}
	return nil, &NotLoadedError{edge: "areasaude"}
}

// ProfissionaisOrErr returns the Profissionais value or an error if the edge
// was not loaded in eager-loading.
func (e EspecializacaoEdges) ProfissionaisOrErr() ([]*Profissional, error) {
	if e.loadedTypes[1] {
		return e.Profissionais, nil
	}
	return nil, &NotLoadedError{edge: "profissionais"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Especializacao) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case especializacao.FieldID:
			values[i] = new(sql.NullInt64)
		case especializacao.FieldDescricao:
			values[i] = new(sql.NullString)
		case especializacao.ForeignKeys[0]: // area_saude_especializacoes
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Especializacao", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Especializacao fields.
func (e *Especializacao) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case especializacao.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case especializacao.FieldDescricao:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field descricao", values[i])
			} else if value.Valid {
				e.Descricao = value.String
			}
		case especializacao.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field area_saude_especializacoes", value)
			} else if value.Valid {
				e.area_saude_especializacoes = new(int)
				*e.area_saude_especializacoes = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAreasaude queries the "areasaude" edge of the Especializacao entity.
func (e *Especializacao) QueryAreasaude() *AreaSaudeQuery {
	return (&EspecializacaoClient{config: e.config}).QueryAreasaude(e)
}

// QueryProfissionais queries the "profissionais" edge of the Especializacao entity.
func (e *Especializacao) QueryProfissionais() *ProfissionalQuery {
	return (&EspecializacaoClient{config: e.config}).QueryProfissionais(e)
}

// Update returns a builder for updating this Especializacao.
// Note that you need to call Especializacao.Unwrap() before calling this method if this Especializacao
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Especializacao) Update() *EspecializacaoUpdateOne {
	return (&EspecializacaoClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Especializacao entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Especializacao) Unwrap() *Especializacao {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Especializacao is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Especializacao) String() string {
	var builder strings.Builder
	builder.WriteString("Especializacao(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", descricao=")
	builder.WriteString(e.Descricao)
	builder.WriteByte(')')
	return builder.String()
}

// Especializacaos is a parsable slice of Especializacao.
type Especializacaos []*Especializacao

func (e Especializacaos) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
