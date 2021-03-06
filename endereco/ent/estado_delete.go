// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"endereco/ent/estado"
	"endereco/ent/predicate"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EstadoDelete is the builder for deleting a Estado entity.
type EstadoDelete struct {
	config
	hooks    []Hook
	mutation *EstadoMutation
}

// Where appends a list predicates to the EstadoDelete builder.
func (ed *EstadoDelete) Where(ps ...predicate.Estado) *EstadoDelete {
	ed.mutation.Where(ps...)
	return ed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ed *EstadoDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ed.hooks) == 0 {
		affected, err = ed.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EstadoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ed.mutation = mutation
			affected, err = ed.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ed.hooks) - 1; i >= 0; i-- {
			if ed.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ed.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ed.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ed *EstadoDelete) ExecX(ctx context.Context) int {
	n, err := ed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ed *EstadoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: estado.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: estado.FieldID,
			},
		},
	}
	if ps := ed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ed.driver, _spec)
}

// EstadoDeleteOne is the builder for deleting a single Estado entity.
type EstadoDeleteOne struct {
	ed *EstadoDelete
}

// Exec executes the deletion query.
func (edo *EstadoDeleteOne) Exec(ctx context.Context) error {
	n, err := edo.ed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{estado.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (edo *EstadoDeleteOne) ExecX(ctx context.Context) {
	edo.ed.ExecX(ctx)
}
