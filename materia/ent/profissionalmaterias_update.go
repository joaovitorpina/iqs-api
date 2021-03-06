// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"materia/ent/materia"
	"materia/ent/predicate"
	"materia/ent/profissionalmaterias"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfissionalMateriasUpdate is the builder for updating ProfissionalMaterias entities.
type ProfissionalMateriasUpdate struct {
	config
	hooks    []Hook
	mutation *ProfissionalMateriasMutation
}

// Where appends a list predicates to the ProfissionalMateriasUpdate builder.
func (pmu *ProfissionalMateriasUpdate) Where(ps ...predicate.ProfissionalMaterias) *ProfissionalMateriasUpdate {
	pmu.mutation.Where(ps...)
	return pmu
}

// SetProfissionalID sets the "profissional_id" field.
func (pmu *ProfissionalMateriasUpdate) SetProfissionalID(i int) *ProfissionalMateriasUpdate {
	pmu.mutation.ResetProfissionalID()
	pmu.mutation.SetProfissionalID(i)
	return pmu
}

// AddProfissionalID adds i to the "profissional_id" field.
func (pmu *ProfissionalMateriasUpdate) AddProfissionalID(i int) *ProfissionalMateriasUpdate {
	pmu.mutation.AddProfissionalID(i)
	return pmu
}

// SetMateriaID sets the "materia" edge to the Materia entity by ID.
func (pmu *ProfissionalMateriasUpdate) SetMateriaID(id int) *ProfissionalMateriasUpdate {
	pmu.mutation.SetMateriaID(id)
	return pmu
}

// SetMateria sets the "materia" edge to the Materia entity.
func (pmu *ProfissionalMateriasUpdate) SetMateria(m *Materia) *ProfissionalMateriasUpdate {
	return pmu.SetMateriaID(m.ID)
}

// Mutation returns the ProfissionalMateriasMutation object of the builder.
func (pmu *ProfissionalMateriasUpdate) Mutation() *ProfissionalMateriasMutation {
	return pmu.mutation
}

// ClearMateria clears the "materia" edge to the Materia entity.
func (pmu *ProfissionalMateriasUpdate) ClearMateria() *ProfissionalMateriasUpdate {
	pmu.mutation.ClearMateria()
	return pmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pmu *ProfissionalMateriasUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pmu.hooks) == 0 {
		if err = pmu.check(); err != nil {
			return 0, err
		}
		affected, err = pmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfissionalMateriasMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pmu.check(); err != nil {
				return 0, err
			}
			pmu.mutation = mutation
			affected, err = pmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pmu.hooks) - 1; i >= 0; i-- {
			if pmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pmu *ProfissionalMateriasUpdate) SaveX(ctx context.Context) int {
	affected, err := pmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pmu *ProfissionalMateriasUpdate) Exec(ctx context.Context) error {
	_, err := pmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmu *ProfissionalMateriasUpdate) ExecX(ctx context.Context) {
	if err := pmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmu *ProfissionalMateriasUpdate) check() error {
	if v, ok := pmu.mutation.ProfissionalID(); ok {
		if err := profissionalmaterias.ProfissionalIDValidator(v); err != nil {
			return &ValidationError{Name: "profissional_id", err: fmt.Errorf(`ent: validator failed for field "ProfissionalMaterias.profissional_id": %w`, err)}
		}
	}
	if _, ok := pmu.mutation.MateriaID(); pmu.mutation.MateriaCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProfissionalMaterias.materia"`)
	}
	return nil
}

func (pmu *ProfissionalMateriasUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profissionalmaterias.Table,
			Columns: profissionalmaterias.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profissionalmaterias.FieldID,
			},
		},
	}
	if ps := pmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmu.mutation.ProfissionalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profissionalmaterias.FieldProfissionalID,
		})
	}
	if value, ok := pmu.mutation.AddedProfissionalID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profissionalmaterias.FieldProfissionalID,
		})
	}
	if pmu.mutation.MateriaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profissionalmaterias.MateriaTable,
			Columns: []string{profissionalmaterias.MateriaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: materia.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmu.mutation.MateriaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profissionalmaterias.MateriaTable,
			Columns: []string{profissionalmaterias.MateriaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: materia.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profissionalmaterias.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProfissionalMateriasUpdateOne is the builder for updating a single ProfissionalMaterias entity.
type ProfissionalMateriasUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProfissionalMateriasMutation
}

// SetProfissionalID sets the "profissional_id" field.
func (pmuo *ProfissionalMateriasUpdateOne) SetProfissionalID(i int) *ProfissionalMateriasUpdateOne {
	pmuo.mutation.ResetProfissionalID()
	pmuo.mutation.SetProfissionalID(i)
	return pmuo
}

// AddProfissionalID adds i to the "profissional_id" field.
func (pmuo *ProfissionalMateriasUpdateOne) AddProfissionalID(i int) *ProfissionalMateriasUpdateOne {
	pmuo.mutation.AddProfissionalID(i)
	return pmuo
}

// SetMateriaID sets the "materia" edge to the Materia entity by ID.
func (pmuo *ProfissionalMateriasUpdateOne) SetMateriaID(id int) *ProfissionalMateriasUpdateOne {
	pmuo.mutation.SetMateriaID(id)
	return pmuo
}

// SetMateria sets the "materia" edge to the Materia entity.
func (pmuo *ProfissionalMateriasUpdateOne) SetMateria(m *Materia) *ProfissionalMateriasUpdateOne {
	return pmuo.SetMateriaID(m.ID)
}

// Mutation returns the ProfissionalMateriasMutation object of the builder.
func (pmuo *ProfissionalMateriasUpdateOne) Mutation() *ProfissionalMateriasMutation {
	return pmuo.mutation
}

// ClearMateria clears the "materia" edge to the Materia entity.
func (pmuo *ProfissionalMateriasUpdateOne) ClearMateria() *ProfissionalMateriasUpdateOne {
	pmuo.mutation.ClearMateria()
	return pmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pmuo *ProfissionalMateriasUpdateOne) Select(field string, fields ...string) *ProfissionalMateriasUpdateOne {
	pmuo.fields = append([]string{field}, fields...)
	return pmuo
}

// Save executes the query and returns the updated ProfissionalMaterias entity.
func (pmuo *ProfissionalMateriasUpdateOne) Save(ctx context.Context) (*ProfissionalMaterias, error) {
	var (
		err  error
		node *ProfissionalMaterias
	)
	if len(pmuo.hooks) == 0 {
		if err = pmuo.check(); err != nil {
			return nil, err
		}
		node, err = pmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfissionalMateriasMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pmuo.check(); err != nil {
				return nil, err
			}
			pmuo.mutation = mutation
			node, err = pmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pmuo.hooks) - 1; i >= 0; i-- {
			if pmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pmuo *ProfissionalMateriasUpdateOne) SaveX(ctx context.Context) *ProfissionalMaterias {
	node, err := pmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pmuo *ProfissionalMateriasUpdateOne) Exec(ctx context.Context) error {
	_, err := pmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmuo *ProfissionalMateriasUpdateOne) ExecX(ctx context.Context) {
	if err := pmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmuo *ProfissionalMateriasUpdateOne) check() error {
	if v, ok := pmuo.mutation.ProfissionalID(); ok {
		if err := profissionalmaterias.ProfissionalIDValidator(v); err != nil {
			return &ValidationError{Name: "profissional_id", err: fmt.Errorf(`ent: validator failed for field "ProfissionalMaterias.profissional_id": %w`, err)}
		}
	}
	if _, ok := pmuo.mutation.MateriaID(); pmuo.mutation.MateriaCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProfissionalMaterias.materia"`)
	}
	return nil
}

func (pmuo *ProfissionalMateriasUpdateOne) sqlSave(ctx context.Context) (_node *ProfissionalMaterias, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profissionalmaterias.Table,
			Columns: profissionalmaterias.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profissionalmaterias.FieldID,
			},
		},
	}
	id, ok := pmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProfissionalMaterias.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profissionalmaterias.FieldID)
		for _, f := range fields {
			if !profissionalmaterias.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != profissionalmaterias.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmuo.mutation.ProfissionalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profissionalmaterias.FieldProfissionalID,
		})
	}
	if value, ok := pmuo.mutation.AddedProfissionalID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profissionalmaterias.FieldProfissionalID,
		})
	}
	if pmuo.mutation.MateriaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profissionalmaterias.MateriaTable,
			Columns: []string{profissionalmaterias.MateriaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: materia.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmuo.mutation.MateriaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profissionalmaterias.MateriaTable,
			Columns: []string{profissionalmaterias.MateriaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: materia.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProfissionalMaterias{config: pmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profissionalmaterias.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
