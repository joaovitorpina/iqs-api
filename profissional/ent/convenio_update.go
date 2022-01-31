// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"profissional/ent/convenio"
	"profissional/ent/predicate"
	"profissional/ent/profissional"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConvenioUpdate is the builder for updating Convenio entities.
type ConvenioUpdate struct {
	config
	hooks    []Hook
	mutation *ConvenioMutation
}

// Where appends a list predicates to the ConvenioUpdate builder.
func (cu *ConvenioUpdate) Where(ps ...predicate.Convenio) *ConvenioUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetNome sets the "nome" field.
func (cu *ConvenioUpdate) SetNome(s string) *ConvenioUpdate {
	cu.mutation.SetNome(s)
	return cu
}

// SetProfissionalID sets the "profissional" edge to the Profissional entity by ID.
func (cu *ConvenioUpdate) SetProfissionalID(id int) *ConvenioUpdate {
	cu.mutation.SetProfissionalID(id)
	return cu
}

// SetNillableProfissionalID sets the "profissional" edge to the Profissional entity by ID if the given value is not nil.
func (cu *ConvenioUpdate) SetNillableProfissionalID(id *int) *ConvenioUpdate {
	if id != nil {
		cu = cu.SetProfissionalID(*id)
	}
	return cu
}

// SetProfissional sets the "profissional" edge to the Profissional entity.
func (cu *ConvenioUpdate) SetProfissional(p *Profissional) *ConvenioUpdate {
	return cu.SetProfissionalID(p.ID)
}

// Mutation returns the ConvenioMutation object of the builder.
func (cu *ConvenioUpdate) Mutation() *ConvenioMutation {
	return cu.mutation
}

// ClearProfissional clears the "profissional" edge to the Profissional entity.
func (cu *ConvenioUpdate) ClearProfissional() *ConvenioUpdate {
	cu.mutation.ClearProfissional()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConvenioUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConvenioMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConvenioUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConvenioUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConvenioUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ConvenioUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   convenio.Table,
			Columns: convenio.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: convenio.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Nome(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: convenio.FieldNome,
		})
	}
	if cu.mutation.ProfissionalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   convenio.ProfissionalTable,
			Columns: []string{convenio.ProfissionalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profissional.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ProfissionalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   convenio.ProfissionalTable,
			Columns: []string{convenio.ProfissionalColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{convenio.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ConvenioUpdateOne is the builder for updating a single Convenio entity.
type ConvenioUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConvenioMutation
}

// SetNome sets the "nome" field.
func (cuo *ConvenioUpdateOne) SetNome(s string) *ConvenioUpdateOne {
	cuo.mutation.SetNome(s)
	return cuo
}

// SetProfissionalID sets the "profissional" edge to the Profissional entity by ID.
func (cuo *ConvenioUpdateOne) SetProfissionalID(id int) *ConvenioUpdateOne {
	cuo.mutation.SetProfissionalID(id)
	return cuo
}

// SetNillableProfissionalID sets the "profissional" edge to the Profissional entity by ID if the given value is not nil.
func (cuo *ConvenioUpdateOne) SetNillableProfissionalID(id *int) *ConvenioUpdateOne {
	if id != nil {
		cuo = cuo.SetProfissionalID(*id)
	}
	return cuo
}

// SetProfissional sets the "profissional" edge to the Profissional entity.
func (cuo *ConvenioUpdateOne) SetProfissional(p *Profissional) *ConvenioUpdateOne {
	return cuo.SetProfissionalID(p.ID)
}

// Mutation returns the ConvenioMutation object of the builder.
func (cuo *ConvenioUpdateOne) Mutation() *ConvenioMutation {
	return cuo.mutation
}

// ClearProfissional clears the "profissional" edge to the Profissional entity.
func (cuo *ConvenioUpdateOne) ClearProfissional() *ConvenioUpdateOne {
	cuo.mutation.ClearProfissional()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConvenioUpdateOne) Select(field string, fields ...string) *ConvenioUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Convenio entity.
func (cuo *ConvenioUpdateOne) Save(ctx context.Context) (*Convenio, error) {
	var (
		err  error
		node *Convenio
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConvenioMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConvenioUpdateOne) SaveX(ctx context.Context) *Convenio {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConvenioUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConvenioUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ConvenioUpdateOne) sqlSave(ctx context.Context) (_node *Convenio, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   convenio.Table,
			Columns: convenio.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: convenio.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Convenio.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, convenio.FieldID)
		for _, f := range fields {
			if !convenio.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != convenio.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Nome(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: convenio.FieldNome,
		})
	}
	if cuo.mutation.ProfissionalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   convenio.ProfissionalTable,
			Columns: []string{convenio.ProfissionalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profissional.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ProfissionalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   convenio.ProfissionalTable,
			Columns: []string{convenio.ProfissionalColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Convenio{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{convenio.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
