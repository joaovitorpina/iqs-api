// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"profissional/ent/profissional"
	"profissional/ent/whatsapp"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WhatsAppCreate is the builder for creating a WhatsApp entity.
type WhatsAppCreate struct {
	config
	mutation *WhatsAppMutation
	hooks    []Hook
}

// SetNumero sets the "numero" field.
func (wac *WhatsAppCreate) SetNumero(i int64) *WhatsAppCreate {
	wac.mutation.SetNumero(i)
	return wac
}

// SetPrincipal sets the "principal" field.
func (wac *WhatsAppCreate) SetPrincipal(b bool) *WhatsAppCreate {
	wac.mutation.SetPrincipal(b)
	return wac
}

// SetNillablePrincipal sets the "principal" field if the given value is not nil.
func (wac *WhatsAppCreate) SetNillablePrincipal(b *bool) *WhatsAppCreate {
	if b != nil {
		wac.SetPrincipal(*b)
	}
	return wac
}

// SetProfissionalID sets the "profissional" edge to the Profissional entity by ID.
func (wac *WhatsAppCreate) SetProfissionalID(id int) *WhatsAppCreate {
	wac.mutation.SetProfissionalID(id)
	return wac
}

// SetNillableProfissionalID sets the "profissional" edge to the Profissional entity by ID if the given value is not nil.
func (wac *WhatsAppCreate) SetNillableProfissionalID(id *int) *WhatsAppCreate {
	if id != nil {
		wac = wac.SetProfissionalID(*id)
	}
	return wac
}

// SetProfissional sets the "profissional" edge to the Profissional entity.
func (wac *WhatsAppCreate) SetProfissional(p *Profissional) *WhatsAppCreate {
	return wac.SetProfissionalID(p.ID)
}

// Mutation returns the WhatsAppMutation object of the builder.
func (wac *WhatsAppCreate) Mutation() *WhatsAppMutation {
	return wac.mutation
}

// Save creates the WhatsApp in the database.
func (wac *WhatsAppCreate) Save(ctx context.Context) (*WhatsApp, error) {
	var (
		err  error
		node *WhatsApp
	)
	wac.defaults()
	if len(wac.hooks) == 0 {
		if err = wac.check(); err != nil {
			return nil, err
		}
		node, err = wac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WhatsAppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wac.check(); err != nil {
				return nil, err
			}
			wac.mutation = mutation
			if node, err = wac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wac.hooks) - 1; i >= 0; i-- {
			if wac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wac *WhatsAppCreate) SaveX(ctx context.Context) *WhatsApp {
	v, err := wac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wac *WhatsAppCreate) Exec(ctx context.Context) error {
	_, err := wac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wac *WhatsAppCreate) ExecX(ctx context.Context) {
	if err := wac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wac *WhatsAppCreate) defaults() {
	if _, ok := wac.mutation.Principal(); !ok {
		v := whatsapp.DefaultPrincipal
		wac.mutation.SetPrincipal(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wac *WhatsAppCreate) check() error {
	if _, ok := wac.mutation.Numero(); !ok {
		return &ValidationError{Name: "numero", err: errors.New(`ent: missing required field "numero"`)}
	}
	if _, ok := wac.mutation.Principal(); !ok {
		return &ValidationError{Name: "principal", err: errors.New(`ent: missing required field "principal"`)}
	}
	return nil
}

func (wac *WhatsAppCreate) sqlSave(ctx context.Context) (*WhatsApp, error) {
	_node, _spec := wac.createSpec()
	if err := sqlgraph.CreateNode(ctx, wac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wac *WhatsAppCreate) createSpec() (*WhatsApp, *sqlgraph.CreateSpec) {
	var (
		_node = &WhatsApp{config: wac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: whatsapp.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: whatsapp.FieldID,
			},
		}
	)
	if value, ok := wac.mutation.Numero(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: whatsapp.FieldNumero,
		})
		_node.Numero = value
	}
	if value, ok := wac.mutation.Principal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: whatsapp.FieldPrincipal,
		})
		_node.Principal = value
	}
	if nodes := wac.mutation.ProfissionalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   whatsapp.ProfissionalTable,
			Columns: []string{whatsapp.ProfissionalColumn},
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
		_node.profissional_whatsapps = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WhatsAppCreateBulk is the builder for creating many WhatsApp entities in bulk.
type WhatsAppCreateBulk struct {
	config
	builders []*WhatsAppCreate
}

// Save creates the WhatsApp entities in the database.
func (wacb *WhatsAppCreateBulk) Save(ctx context.Context) ([]*WhatsApp, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wacb.builders))
	nodes := make([]*WhatsApp, len(wacb.builders))
	mutators := make([]Mutator, len(wacb.builders))
	for i := range wacb.builders {
		func(i int, root context.Context) {
			builder := wacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WhatsAppMutation)
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
					_, err = mutators[i+1].Mutate(root, wacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wacb *WhatsAppCreateBulk) SaveX(ctx context.Context) []*WhatsApp {
	v, err := wacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wacb *WhatsAppCreateBulk) Exec(ctx context.Context) error {
	_, err := wacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wacb *WhatsAppCreateBulk) ExecX(ctx context.Context) {
	if err := wacb.Exec(ctx); err != nil {
		panic(err)
	}
}
