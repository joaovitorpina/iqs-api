// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"profissional/ent/predicate"
	"profissional/ent/profissional"
	"profissional/ent/whatsapp"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WhatsAppUpdate is the builder for updating WhatsApp entities.
type WhatsAppUpdate struct {
	config
	hooks    []Hook
	mutation *WhatsAppMutation
}

// Where appends a list predicates to the WhatsAppUpdate builder.
func (wau *WhatsAppUpdate) Where(ps ...predicate.WhatsApp) *WhatsAppUpdate {
	wau.mutation.Where(ps...)
	return wau
}

// SetNumero sets the "numero" field.
func (wau *WhatsAppUpdate) SetNumero(i int64) *WhatsAppUpdate {
	wau.mutation.ResetNumero()
	wau.mutation.SetNumero(i)
	return wau
}

// AddNumero adds i to the "numero" field.
func (wau *WhatsAppUpdate) AddNumero(i int64) *WhatsAppUpdate {
	wau.mutation.AddNumero(i)
	return wau
}

// SetPrincipal sets the "principal" field.
func (wau *WhatsAppUpdate) SetPrincipal(b bool) *WhatsAppUpdate {
	wau.mutation.SetPrincipal(b)
	return wau
}

// SetNillablePrincipal sets the "principal" field if the given value is not nil.
func (wau *WhatsAppUpdate) SetNillablePrincipal(b *bool) *WhatsAppUpdate {
	if b != nil {
		wau.SetPrincipal(*b)
	}
	return wau
}

// SetProfissionalID sets the "profissional" edge to the Profissional entity by ID.
func (wau *WhatsAppUpdate) SetProfissionalID(id int) *WhatsAppUpdate {
	wau.mutation.SetProfissionalID(id)
	return wau
}

// SetNillableProfissionalID sets the "profissional" edge to the Profissional entity by ID if the given value is not nil.
func (wau *WhatsAppUpdate) SetNillableProfissionalID(id *int) *WhatsAppUpdate {
	if id != nil {
		wau = wau.SetProfissionalID(*id)
	}
	return wau
}

// SetProfissional sets the "profissional" edge to the Profissional entity.
func (wau *WhatsAppUpdate) SetProfissional(p *Profissional) *WhatsAppUpdate {
	return wau.SetProfissionalID(p.ID)
}

// Mutation returns the WhatsAppMutation object of the builder.
func (wau *WhatsAppUpdate) Mutation() *WhatsAppMutation {
	return wau.mutation
}

// ClearProfissional clears the "profissional" edge to the Profissional entity.
func (wau *WhatsAppUpdate) ClearProfissional() *WhatsAppUpdate {
	wau.mutation.ClearProfissional()
	return wau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wau *WhatsAppUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wau.hooks) == 0 {
		affected, err = wau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WhatsAppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wau.mutation = mutation
			affected, err = wau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wau.hooks) - 1; i >= 0; i-- {
			if wau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wau *WhatsAppUpdate) SaveX(ctx context.Context) int {
	affected, err := wau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wau *WhatsAppUpdate) Exec(ctx context.Context) error {
	_, err := wau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wau *WhatsAppUpdate) ExecX(ctx context.Context) {
	if err := wau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wau *WhatsAppUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   whatsapp.Table,
			Columns: whatsapp.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: whatsapp.FieldID,
			},
		},
	}
	if ps := wau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wau.mutation.Numero(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: whatsapp.FieldNumero,
		})
	}
	if value, ok := wau.mutation.AddedNumero(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: whatsapp.FieldNumero,
		})
	}
	if value, ok := wau.mutation.Principal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: whatsapp.FieldPrincipal,
		})
	}
	if wau.mutation.ProfissionalCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wau.mutation.ProfissionalIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{whatsapp.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WhatsAppUpdateOne is the builder for updating a single WhatsApp entity.
type WhatsAppUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WhatsAppMutation
}

// SetNumero sets the "numero" field.
func (wauo *WhatsAppUpdateOne) SetNumero(i int64) *WhatsAppUpdateOne {
	wauo.mutation.ResetNumero()
	wauo.mutation.SetNumero(i)
	return wauo
}

// AddNumero adds i to the "numero" field.
func (wauo *WhatsAppUpdateOne) AddNumero(i int64) *WhatsAppUpdateOne {
	wauo.mutation.AddNumero(i)
	return wauo
}

// SetPrincipal sets the "principal" field.
func (wauo *WhatsAppUpdateOne) SetPrincipal(b bool) *WhatsAppUpdateOne {
	wauo.mutation.SetPrincipal(b)
	return wauo
}

// SetNillablePrincipal sets the "principal" field if the given value is not nil.
func (wauo *WhatsAppUpdateOne) SetNillablePrincipal(b *bool) *WhatsAppUpdateOne {
	if b != nil {
		wauo.SetPrincipal(*b)
	}
	return wauo
}

// SetProfissionalID sets the "profissional" edge to the Profissional entity by ID.
func (wauo *WhatsAppUpdateOne) SetProfissionalID(id int) *WhatsAppUpdateOne {
	wauo.mutation.SetProfissionalID(id)
	return wauo
}

// SetNillableProfissionalID sets the "profissional" edge to the Profissional entity by ID if the given value is not nil.
func (wauo *WhatsAppUpdateOne) SetNillableProfissionalID(id *int) *WhatsAppUpdateOne {
	if id != nil {
		wauo = wauo.SetProfissionalID(*id)
	}
	return wauo
}

// SetProfissional sets the "profissional" edge to the Profissional entity.
func (wauo *WhatsAppUpdateOne) SetProfissional(p *Profissional) *WhatsAppUpdateOne {
	return wauo.SetProfissionalID(p.ID)
}

// Mutation returns the WhatsAppMutation object of the builder.
func (wauo *WhatsAppUpdateOne) Mutation() *WhatsAppMutation {
	return wauo.mutation
}

// ClearProfissional clears the "profissional" edge to the Profissional entity.
func (wauo *WhatsAppUpdateOne) ClearProfissional() *WhatsAppUpdateOne {
	wauo.mutation.ClearProfissional()
	return wauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wauo *WhatsAppUpdateOne) Select(field string, fields ...string) *WhatsAppUpdateOne {
	wauo.fields = append([]string{field}, fields...)
	return wauo
}

// Save executes the query and returns the updated WhatsApp entity.
func (wauo *WhatsAppUpdateOne) Save(ctx context.Context) (*WhatsApp, error) {
	var (
		err  error
		node *WhatsApp
	)
	if len(wauo.hooks) == 0 {
		node, err = wauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WhatsAppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wauo.mutation = mutation
			node, err = wauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wauo.hooks) - 1; i >= 0; i-- {
			if wauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wauo *WhatsAppUpdateOne) SaveX(ctx context.Context) *WhatsApp {
	node, err := wauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wauo *WhatsAppUpdateOne) Exec(ctx context.Context) error {
	_, err := wauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wauo *WhatsAppUpdateOne) ExecX(ctx context.Context) {
	if err := wauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wauo *WhatsAppUpdateOne) sqlSave(ctx context.Context) (_node *WhatsApp, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   whatsapp.Table,
			Columns: whatsapp.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: whatsapp.FieldID,
			},
		},
	}
	id, ok := wauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WhatsApp.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, whatsapp.FieldID)
		for _, f := range fields {
			if !whatsapp.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != whatsapp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wauo.mutation.Numero(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: whatsapp.FieldNumero,
		})
	}
	if value, ok := wauo.mutation.AddedNumero(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: whatsapp.FieldNumero,
		})
	}
	if value, ok := wauo.mutation.Principal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: whatsapp.FieldPrincipal,
		})
	}
	if wauo.mutation.ProfissionalCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wauo.mutation.ProfissionalIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WhatsApp{config: wauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{whatsapp.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
