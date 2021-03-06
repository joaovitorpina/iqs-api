// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"profissional/ent/profissional"
	"profissional/ent/tratamento"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TratamentoCreate is the builder for creating a Tratamento entity.
type TratamentoCreate struct {
	config
	mutation *TratamentoMutation
	hooks    []Hook
}

// SetDescricao sets the "descricao" field.
func (tc *TratamentoCreate) SetDescricao(s string) *TratamentoCreate {
	tc.mutation.SetDescricao(s)
	return tc
}

// SetProfissionalID sets the "profissional" edge to the Profissional entity by ID.
func (tc *TratamentoCreate) SetProfissionalID(id int) *TratamentoCreate {
	tc.mutation.SetProfissionalID(id)
	return tc
}

// SetProfissional sets the "profissional" edge to the Profissional entity.
func (tc *TratamentoCreate) SetProfissional(p *Profissional) *TratamentoCreate {
	return tc.SetProfissionalID(p.ID)
}

// Mutation returns the TratamentoMutation object of the builder.
func (tc *TratamentoCreate) Mutation() *TratamentoMutation {
	return tc.mutation
}

// Save creates the Tratamento in the database.
func (tc *TratamentoCreate) Save(ctx context.Context) (*Tratamento, error) {
	var (
		err  error
		node *Tratamento
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TratamentoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TratamentoCreate) SaveX(ctx context.Context) *Tratamento {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TratamentoCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TratamentoCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TratamentoCreate) check() error {
	if _, ok := tc.mutation.Descricao(); !ok {
		return &ValidationError{Name: "descricao", err: errors.New(`ent: missing required field "Tratamento.descricao"`)}
	}
	if _, ok := tc.mutation.ProfissionalID(); !ok {
		return &ValidationError{Name: "profissional", err: errors.New(`ent: missing required edge "Tratamento.profissional"`)}
	}
	return nil
}

func (tc *TratamentoCreate) sqlSave(ctx context.Context) (*Tratamento, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TratamentoCreate) createSpec() (*Tratamento, *sqlgraph.CreateSpec) {
	var (
		_node = &Tratamento{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tratamento.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tratamento.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Descricao(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tratamento.FieldDescricao,
		})
		_node.Descricao = value
	}
	if nodes := tc.mutation.ProfissionalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tratamento.ProfissionalTable,
			Columns: []string{tratamento.ProfissionalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profissional.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.profissional_tratamentos = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TratamentoCreateBulk is the builder for creating many Tratamento entities in bulk.
type TratamentoCreateBulk struct {
	config
	builders []*TratamentoCreate
}

// Save creates the Tratamento entities in the database.
func (tcb *TratamentoCreateBulk) Save(ctx context.Context) ([]*Tratamento, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tratamento, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TratamentoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TratamentoCreateBulk) SaveX(ctx context.Context) []*Tratamento {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TratamentoCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TratamentoCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
