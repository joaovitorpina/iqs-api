// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"endereco/ent/cep"
	"endereco/ent/endereco"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EnderecoCreate is the builder for creating a Endereco entity.
type EnderecoCreate struct {
	config
	mutation *EnderecoMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ec *EnderecoCreate) SetCreateTime(t time.Time) *EnderecoCreate {
	ec.mutation.SetCreateTime(t)
	return ec
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ec *EnderecoCreate) SetNillableCreateTime(t *time.Time) *EnderecoCreate {
	if t != nil {
		ec.SetCreateTime(*t)
	}
	return ec
}

// SetUpdateTime sets the "update_time" field.
func (ec *EnderecoCreate) SetUpdateTime(t time.Time) *EnderecoCreate {
	ec.mutation.SetUpdateTime(t)
	return ec
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ec *EnderecoCreate) SetNillableUpdateTime(t *time.Time) *EnderecoCreate {
	if t != nil {
		ec.SetUpdateTime(*t)
	}
	return ec
}

// SetNumero sets the "numero" field.
func (ec *EnderecoCreate) SetNumero(s string) *EnderecoCreate {
	ec.mutation.SetNumero(s)
	return ec
}

// SetCepID sets the "cep" edge to the Cep entity by ID.
func (ec *EnderecoCreate) SetCepID(id int32) *EnderecoCreate {
	ec.mutation.SetCepID(id)
	return ec
}

// SetCep sets the "cep" edge to the Cep entity.
func (ec *EnderecoCreate) SetCep(c *Cep) *EnderecoCreate {
	return ec.SetCepID(c.ID)
}

// Mutation returns the EnderecoMutation object of the builder.
func (ec *EnderecoCreate) Mutation() *EnderecoMutation {
	return ec.mutation
}

// Save creates the Endereco in the database.
func (ec *EnderecoCreate) Save(ctx context.Context) (*Endereco, error) {
	var (
		err  error
		node *Endereco
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EnderecoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EnderecoCreate) SaveX(ctx context.Context) *Endereco {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EnderecoCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EnderecoCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EnderecoCreate) defaults() {
	if _, ok := ec.mutation.CreateTime(); !ok {
		v := endereco.DefaultCreateTime()
		ec.mutation.SetCreateTime(v)
	}
	if _, ok := ec.mutation.UpdateTime(); !ok {
		v := endereco.DefaultUpdateTime()
		ec.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EnderecoCreate) check() error {
	if _, ok := ec.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Endereco.create_time"`)}
	}
	if _, ok := ec.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Endereco.update_time"`)}
	}
	if _, ok := ec.mutation.Numero(); !ok {
		return &ValidationError{Name: "numero", err: errors.New(`ent: missing required field "Endereco.numero"`)}
	}
	if v, ok := ec.mutation.Numero(); ok {
		if err := endereco.NumeroValidator(v); err != nil {
			return &ValidationError{Name: "numero", err: fmt.Errorf(`ent: validator failed for field "Endereco.numero": %w`, err)}
		}
	}
	if _, ok := ec.mutation.CepID(); !ok {
		return &ValidationError{Name: "cep", err: errors.New(`ent: missing required edge "Endereco.cep"`)}
	}
	return nil
}

func (ec *EnderecoCreate) sqlSave(ctx context.Context) (*Endereco, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ec *EnderecoCreate) createSpec() (*Endereco, *sqlgraph.CreateSpec) {
	var (
		_node = &Endereco{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: endereco.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: endereco.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: endereco.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ec.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: endereco.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ec.mutation.Numero(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: endereco.FieldNumero,
		})
		_node.Numero = value
	}
	if nodes := ec.mutation.CepIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   endereco.CepTable,
			Columns: []string{endereco.CepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: cep.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.cep_enderecos = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EnderecoCreateBulk is the builder for creating many Endereco entities in bulk.
type EnderecoCreateBulk struct {
	config
	builders []*EnderecoCreate
}

// Save creates the Endereco entities in the database.
func (ecb *EnderecoCreateBulk) Save(ctx context.Context) ([]*Endereco, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Endereco, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EnderecoMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EnderecoCreateBulk) SaveX(ctx context.Context) []*Endereco {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EnderecoCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EnderecoCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
