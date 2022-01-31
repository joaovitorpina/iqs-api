// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"profissional/ent/areasaude"
	"profissional/ent/especializacao"
	"profissional/ent/profissional"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EspecializacaoCreate is the builder for creating a Especializacao entity.
type EspecializacaoCreate struct {
	config
	mutation *EspecializacaoMutation
	hooks    []Hook
}

// SetDescricao sets the "descricao" field.
func (ec *EspecializacaoCreate) SetDescricao(s string) *EspecializacaoCreate {
	ec.mutation.SetDescricao(s)
	return ec
}

// SetAreasaudeID sets the "areasaude" edge to the AreaSaude entity by ID.
func (ec *EspecializacaoCreate) SetAreasaudeID(id int) *EspecializacaoCreate {
	ec.mutation.SetAreasaudeID(id)
	return ec
}

// SetAreasaude sets the "areasaude" edge to the AreaSaude entity.
func (ec *EspecializacaoCreate) SetAreasaude(a *AreaSaude) *EspecializacaoCreate {
	return ec.SetAreasaudeID(a.ID)
}

// AddProfissionaiIDs adds the "profissionais" edge to the Profissional entity by IDs.
func (ec *EspecializacaoCreate) AddProfissionaiIDs(ids ...int) *EspecializacaoCreate {
	ec.mutation.AddProfissionaiIDs(ids...)
	return ec
}

// AddProfissionais adds the "profissionais" edges to the Profissional entity.
func (ec *EspecializacaoCreate) AddProfissionais(p ...*Profissional) *EspecializacaoCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ec.AddProfissionaiIDs(ids...)
}

// Mutation returns the EspecializacaoMutation object of the builder.
func (ec *EspecializacaoCreate) Mutation() *EspecializacaoMutation {
	return ec.mutation
}

// Save creates the Especializacao in the database.
func (ec *EspecializacaoCreate) Save(ctx context.Context) (*Especializacao, error) {
	var (
		err  error
		node *Especializacao
	)
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EspecializacaoMutation)
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
func (ec *EspecializacaoCreate) SaveX(ctx context.Context) *Especializacao {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EspecializacaoCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EspecializacaoCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EspecializacaoCreate) check() error {
	if _, ok := ec.mutation.Descricao(); !ok {
		return &ValidationError{Name: "descricao", err: errors.New(`ent: missing required field "Especializacao.descricao"`)}
	}
	if _, ok := ec.mutation.AreasaudeID(); !ok {
		return &ValidationError{Name: "areasaude", err: errors.New(`ent: missing required edge "Especializacao.areasaude"`)}
	}
	return nil
}

func (ec *EspecializacaoCreate) sqlSave(ctx context.Context) (*Especializacao, error) {
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

func (ec *EspecializacaoCreate) createSpec() (*Especializacao, *sqlgraph.CreateSpec) {
	var (
		_node = &Especializacao{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: especializacao.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: especializacao.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.Descricao(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: especializacao.FieldDescricao,
		})
		_node.Descricao = value
	}
	if nodes := ec.mutation.AreasaudeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   especializacao.AreasaudeTable,
			Columns: []string{especializacao.AreasaudeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: areasaude.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.area_saude_especializacoes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ProfissionaisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   especializacao.ProfissionaisTable,
			Columns: especializacao.ProfissionaisPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EspecializacaoCreateBulk is the builder for creating many Especializacao entities in bulk.
type EspecializacaoCreateBulk struct {
	config
	builders []*EspecializacaoCreate
}

// Save creates the Especializacao entities in the database.
func (ecb *EspecializacaoCreateBulk) Save(ctx context.Context) ([]*Especializacao, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Especializacao, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EspecializacaoMutation)
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
func (ecb *EspecializacaoCreateBulk) SaveX(ctx context.Context) []*Especializacao {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EspecializacaoCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EspecializacaoCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
