// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"materia/ent/categoria"
	"materia/ent/materia"
	"materia/ent/predicate"
	"materia/ent/profissionalmaterias"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MateriaUpdate is the builder for updating Materia entities.
type MateriaUpdate struct {
	config
	hooks    []Hook
	mutation *MateriaMutation
}

// Where appends a list predicates to the MateriaUpdate builder.
func (mu *MateriaUpdate) Where(ps ...predicate.Materia) *MateriaUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUnidadeID sets the "unidade_id" field.
func (mu *MateriaUpdate) SetUnidadeID(i int) *MateriaUpdate {
	mu.mutation.ResetUnidadeID()
	mu.mutation.SetUnidadeID(i)
	return mu
}

// AddUnidadeID adds i to the "unidade_id" field.
func (mu *MateriaUpdate) AddUnidadeID(i int) *MateriaUpdate {
	mu.mutation.AddUnidadeID(i)
	return mu
}

// SetTitulo sets the "titulo" field.
func (mu *MateriaUpdate) SetTitulo(s string) *MateriaUpdate {
	mu.mutation.SetTitulo(s)
	return mu
}

// SetURLAmigavel sets the "url_amigavel" field.
func (mu *MateriaUpdate) SetURLAmigavel(s string) *MateriaUpdate {
	mu.mutation.SetURLAmigavel(s)
	return mu
}

// SetDataAgendamento sets the "data_agendamento" field.
func (mu *MateriaUpdate) SetDataAgendamento(t time.Time) *MateriaUpdate {
	mu.mutation.SetDataAgendamento(t)
	return mu
}

// SetNillableDataAgendamento sets the "data_agendamento" field if the given value is not nil.
func (mu *MateriaUpdate) SetNillableDataAgendamento(t *time.Time) *MateriaUpdate {
	if t != nil {
		mu.SetDataAgendamento(*t)
	}
	return mu
}

// ClearDataAgendamento clears the value of the "data_agendamento" field.
func (mu *MateriaUpdate) ClearDataAgendamento() *MateriaUpdate {
	mu.mutation.ClearDataAgendamento()
	return mu
}

// SetFonte sets the "fonte" field.
func (mu *MateriaUpdate) SetFonte(s string) *MateriaUpdate {
	mu.mutation.SetFonte(s)
	return mu
}

// SetNillableFonte sets the "fonte" field if the given value is not nil.
func (mu *MateriaUpdate) SetNillableFonte(s *string) *MateriaUpdate {
	if s != nil {
		mu.SetFonte(*s)
	}
	return mu
}

// ClearFonte clears the value of the "fonte" field.
func (mu *MateriaUpdate) ClearFonte() *MateriaUpdate {
	mu.mutation.ClearFonte()
	return mu
}

// SetLinkFonte sets the "link_fonte" field.
func (mu *MateriaUpdate) SetLinkFonte(s string) *MateriaUpdate {
	mu.mutation.SetLinkFonte(s)
	return mu
}

// SetNillableLinkFonte sets the "link_fonte" field if the given value is not nil.
func (mu *MateriaUpdate) SetNillableLinkFonte(s *string) *MateriaUpdate {
	if s != nil {
		mu.SetLinkFonte(*s)
	}
	return mu
}

// ClearLinkFonte clears the value of the "link_fonte" field.
func (mu *MateriaUpdate) ClearLinkFonte() *MateriaUpdate {
	mu.mutation.ClearLinkFonte()
	return mu
}

// SetTexto sets the "texto" field.
func (mu *MateriaUpdate) SetTexto(s string) *MateriaUpdate {
	mu.mutation.SetTexto(s)
	return mu
}

// SetImagemURL sets the "imagem_url" field.
func (mu *MateriaUpdate) SetImagemURL(s string) *MateriaUpdate {
	mu.mutation.SetImagemURL(s)
	return mu
}

// SetNillableImagemURL sets the "imagem_url" field if the given value is not nil.
func (mu *MateriaUpdate) SetNillableImagemURL(s *string) *MateriaUpdate {
	if s != nil {
		mu.SetImagemURL(*s)
	}
	return mu
}

// ClearImagemURL clears the value of the "imagem_url" field.
func (mu *MateriaUpdate) ClearImagemURL() *MateriaUpdate {
	mu.mutation.ClearImagemURL()
	return mu
}

// SetPatrocinado sets the "patrocinado" field.
func (mu *MateriaUpdate) SetPatrocinado(b bool) *MateriaUpdate {
	mu.mutation.SetPatrocinado(b)
	return mu
}

// SetNillablePatrocinado sets the "patrocinado" field if the given value is not nil.
func (mu *MateriaUpdate) SetNillablePatrocinado(b *bool) *MateriaUpdate {
	if b != nil {
		mu.SetPatrocinado(*b)
	}
	return mu
}

// SetStatus sets the "status" field.
func (mu *MateriaUpdate) SetStatus(b bool) *MateriaUpdate {
	mu.mutation.SetStatus(b)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MateriaUpdate) SetNillableStatus(b *bool) *MateriaUpdate {
	if b != nil {
		mu.SetStatus(*b)
	}
	return mu
}

// SetCategoriaID sets the "categoria" edge to the Categoria entity by ID.
func (mu *MateriaUpdate) SetCategoriaID(id int) *MateriaUpdate {
	mu.mutation.SetCategoriaID(id)
	return mu
}

// SetCategoria sets the "categoria" edge to the Categoria entity.
func (mu *MateriaUpdate) SetCategoria(c *Categoria) *MateriaUpdate {
	return mu.SetCategoriaID(c.ID)
}

// AddProfissionalMateriaIDs adds the "profissional_materias" edge to the ProfissionalMaterias entity by IDs.
func (mu *MateriaUpdate) AddProfissionalMateriaIDs(ids ...int) *MateriaUpdate {
	mu.mutation.AddProfissionalMateriaIDs(ids...)
	return mu
}

// AddProfissionalMaterias adds the "profissional_materias" edges to the ProfissionalMaterias entity.
func (mu *MateriaUpdate) AddProfissionalMaterias(p ...*ProfissionalMaterias) *MateriaUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.AddProfissionalMateriaIDs(ids...)
}

// Mutation returns the MateriaMutation object of the builder.
func (mu *MateriaUpdate) Mutation() *MateriaMutation {
	return mu.mutation
}

// ClearCategoria clears the "categoria" edge to the Categoria entity.
func (mu *MateriaUpdate) ClearCategoria() *MateriaUpdate {
	mu.mutation.ClearCategoria()
	return mu
}

// ClearProfissionalMaterias clears all "profissional_materias" edges to the ProfissionalMaterias entity.
func (mu *MateriaUpdate) ClearProfissionalMaterias() *MateriaUpdate {
	mu.mutation.ClearProfissionalMaterias()
	return mu
}

// RemoveProfissionalMateriaIDs removes the "profissional_materias" edge to ProfissionalMaterias entities by IDs.
func (mu *MateriaUpdate) RemoveProfissionalMateriaIDs(ids ...int) *MateriaUpdate {
	mu.mutation.RemoveProfissionalMateriaIDs(ids...)
	return mu
}

// RemoveProfissionalMaterias removes "profissional_materias" edges to ProfissionalMaterias entities.
func (mu *MateriaUpdate) RemoveProfissionalMaterias(p ...*ProfissionalMaterias) *MateriaUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.RemoveProfissionalMateriaIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MateriaUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MateriaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MateriaUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MateriaUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MateriaUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MateriaUpdate) check() error {
	if _, ok := mu.mutation.CategoriaID(); mu.mutation.CategoriaCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Materia.categoria"`)
	}
	return nil
}

func (mu *MateriaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   materia.Table,
			Columns: materia.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: materia.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UnidadeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: materia.FieldUnidadeID,
		})
	}
	if value, ok := mu.mutation.AddedUnidadeID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: materia.FieldUnidadeID,
		})
	}
	if value, ok := mu.mutation.Titulo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldTitulo,
		})
	}
	if value, ok := mu.mutation.URLAmigavel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldURLAmigavel,
		})
	}
	if value, ok := mu.mutation.DataAgendamento(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: materia.FieldDataAgendamento,
		})
	}
	if mu.mutation.DataAgendamentoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: materia.FieldDataAgendamento,
		})
	}
	if value, ok := mu.mutation.Fonte(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldFonte,
		})
	}
	if mu.mutation.FonteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: materia.FieldFonte,
		})
	}
	if value, ok := mu.mutation.LinkFonte(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldLinkFonte,
		})
	}
	if mu.mutation.LinkFonteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: materia.FieldLinkFonte,
		})
	}
	if value, ok := mu.mutation.Texto(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldTexto,
		})
	}
	if value, ok := mu.mutation.ImagemURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldImagemURL,
		})
	}
	if mu.mutation.ImagemURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: materia.FieldImagemURL,
		})
	}
	if value, ok := mu.mutation.Patrocinado(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: materia.FieldPatrocinado,
		})
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: materia.FieldStatus,
		})
	}
	if mu.mutation.CategoriaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   materia.CategoriaTable,
			Columns: []string{materia.CategoriaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: categoria.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.CategoriaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   materia.CategoriaTable,
			Columns: []string{materia.CategoriaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: categoria.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ProfissionalMateriasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   materia.ProfissionalMateriasTable,
			Columns: []string{materia.ProfissionalMateriasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profissionalmaterias.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedProfissionalMateriasIDs(); len(nodes) > 0 && !mu.mutation.ProfissionalMateriasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   materia.ProfissionalMateriasTable,
			Columns: []string{materia.ProfissionalMateriasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profissionalmaterias.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ProfissionalMateriasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   materia.ProfissionalMateriasTable,
			Columns: []string{materia.ProfissionalMateriasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profissionalmaterias.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{materia.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MateriaUpdateOne is the builder for updating a single Materia entity.
type MateriaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MateriaMutation
}

// SetUnidadeID sets the "unidade_id" field.
func (muo *MateriaUpdateOne) SetUnidadeID(i int) *MateriaUpdateOne {
	muo.mutation.ResetUnidadeID()
	muo.mutation.SetUnidadeID(i)
	return muo
}

// AddUnidadeID adds i to the "unidade_id" field.
func (muo *MateriaUpdateOne) AddUnidadeID(i int) *MateriaUpdateOne {
	muo.mutation.AddUnidadeID(i)
	return muo
}

// SetTitulo sets the "titulo" field.
func (muo *MateriaUpdateOne) SetTitulo(s string) *MateriaUpdateOne {
	muo.mutation.SetTitulo(s)
	return muo
}

// SetURLAmigavel sets the "url_amigavel" field.
func (muo *MateriaUpdateOne) SetURLAmigavel(s string) *MateriaUpdateOne {
	muo.mutation.SetURLAmigavel(s)
	return muo
}

// SetDataAgendamento sets the "data_agendamento" field.
func (muo *MateriaUpdateOne) SetDataAgendamento(t time.Time) *MateriaUpdateOne {
	muo.mutation.SetDataAgendamento(t)
	return muo
}

// SetNillableDataAgendamento sets the "data_agendamento" field if the given value is not nil.
func (muo *MateriaUpdateOne) SetNillableDataAgendamento(t *time.Time) *MateriaUpdateOne {
	if t != nil {
		muo.SetDataAgendamento(*t)
	}
	return muo
}

// ClearDataAgendamento clears the value of the "data_agendamento" field.
func (muo *MateriaUpdateOne) ClearDataAgendamento() *MateriaUpdateOne {
	muo.mutation.ClearDataAgendamento()
	return muo
}

// SetFonte sets the "fonte" field.
func (muo *MateriaUpdateOne) SetFonte(s string) *MateriaUpdateOne {
	muo.mutation.SetFonte(s)
	return muo
}

// SetNillableFonte sets the "fonte" field if the given value is not nil.
func (muo *MateriaUpdateOne) SetNillableFonte(s *string) *MateriaUpdateOne {
	if s != nil {
		muo.SetFonte(*s)
	}
	return muo
}

// ClearFonte clears the value of the "fonte" field.
func (muo *MateriaUpdateOne) ClearFonte() *MateriaUpdateOne {
	muo.mutation.ClearFonte()
	return muo
}

// SetLinkFonte sets the "link_fonte" field.
func (muo *MateriaUpdateOne) SetLinkFonte(s string) *MateriaUpdateOne {
	muo.mutation.SetLinkFonte(s)
	return muo
}

// SetNillableLinkFonte sets the "link_fonte" field if the given value is not nil.
func (muo *MateriaUpdateOne) SetNillableLinkFonte(s *string) *MateriaUpdateOne {
	if s != nil {
		muo.SetLinkFonte(*s)
	}
	return muo
}

// ClearLinkFonte clears the value of the "link_fonte" field.
func (muo *MateriaUpdateOne) ClearLinkFonte() *MateriaUpdateOne {
	muo.mutation.ClearLinkFonte()
	return muo
}

// SetTexto sets the "texto" field.
func (muo *MateriaUpdateOne) SetTexto(s string) *MateriaUpdateOne {
	muo.mutation.SetTexto(s)
	return muo
}

// SetImagemURL sets the "imagem_url" field.
func (muo *MateriaUpdateOne) SetImagemURL(s string) *MateriaUpdateOne {
	muo.mutation.SetImagemURL(s)
	return muo
}

// SetNillableImagemURL sets the "imagem_url" field if the given value is not nil.
func (muo *MateriaUpdateOne) SetNillableImagemURL(s *string) *MateriaUpdateOne {
	if s != nil {
		muo.SetImagemURL(*s)
	}
	return muo
}

// ClearImagemURL clears the value of the "imagem_url" field.
func (muo *MateriaUpdateOne) ClearImagemURL() *MateriaUpdateOne {
	muo.mutation.ClearImagemURL()
	return muo
}

// SetPatrocinado sets the "patrocinado" field.
func (muo *MateriaUpdateOne) SetPatrocinado(b bool) *MateriaUpdateOne {
	muo.mutation.SetPatrocinado(b)
	return muo
}

// SetNillablePatrocinado sets the "patrocinado" field if the given value is not nil.
func (muo *MateriaUpdateOne) SetNillablePatrocinado(b *bool) *MateriaUpdateOne {
	if b != nil {
		muo.SetPatrocinado(*b)
	}
	return muo
}

// SetStatus sets the "status" field.
func (muo *MateriaUpdateOne) SetStatus(b bool) *MateriaUpdateOne {
	muo.mutation.SetStatus(b)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MateriaUpdateOne) SetNillableStatus(b *bool) *MateriaUpdateOne {
	if b != nil {
		muo.SetStatus(*b)
	}
	return muo
}

// SetCategoriaID sets the "categoria" edge to the Categoria entity by ID.
func (muo *MateriaUpdateOne) SetCategoriaID(id int) *MateriaUpdateOne {
	muo.mutation.SetCategoriaID(id)
	return muo
}

// SetCategoria sets the "categoria" edge to the Categoria entity.
func (muo *MateriaUpdateOne) SetCategoria(c *Categoria) *MateriaUpdateOne {
	return muo.SetCategoriaID(c.ID)
}

// AddProfissionalMateriaIDs adds the "profissional_materias" edge to the ProfissionalMaterias entity by IDs.
func (muo *MateriaUpdateOne) AddProfissionalMateriaIDs(ids ...int) *MateriaUpdateOne {
	muo.mutation.AddProfissionalMateriaIDs(ids...)
	return muo
}

// AddProfissionalMaterias adds the "profissional_materias" edges to the ProfissionalMaterias entity.
func (muo *MateriaUpdateOne) AddProfissionalMaterias(p ...*ProfissionalMaterias) *MateriaUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.AddProfissionalMateriaIDs(ids...)
}

// Mutation returns the MateriaMutation object of the builder.
func (muo *MateriaUpdateOne) Mutation() *MateriaMutation {
	return muo.mutation
}

// ClearCategoria clears the "categoria" edge to the Categoria entity.
func (muo *MateriaUpdateOne) ClearCategoria() *MateriaUpdateOne {
	muo.mutation.ClearCategoria()
	return muo
}

// ClearProfissionalMaterias clears all "profissional_materias" edges to the ProfissionalMaterias entity.
func (muo *MateriaUpdateOne) ClearProfissionalMaterias() *MateriaUpdateOne {
	muo.mutation.ClearProfissionalMaterias()
	return muo
}

// RemoveProfissionalMateriaIDs removes the "profissional_materias" edge to ProfissionalMaterias entities by IDs.
func (muo *MateriaUpdateOne) RemoveProfissionalMateriaIDs(ids ...int) *MateriaUpdateOne {
	muo.mutation.RemoveProfissionalMateriaIDs(ids...)
	return muo
}

// RemoveProfissionalMaterias removes "profissional_materias" edges to ProfissionalMaterias entities.
func (muo *MateriaUpdateOne) RemoveProfissionalMaterias(p ...*ProfissionalMaterias) *MateriaUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.RemoveProfissionalMateriaIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MateriaUpdateOne) Select(field string, fields ...string) *MateriaUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Materia entity.
func (muo *MateriaUpdateOne) Save(ctx context.Context) (*Materia, error) {
	var (
		err  error
		node *Materia
	)
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MateriaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MateriaUpdateOne) SaveX(ctx context.Context) *Materia {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MateriaUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MateriaUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MateriaUpdateOne) check() error {
	if _, ok := muo.mutation.CategoriaID(); muo.mutation.CategoriaCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Materia.categoria"`)
	}
	return nil
}

func (muo *MateriaUpdateOne) sqlSave(ctx context.Context) (_node *Materia, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   materia.Table,
			Columns: materia.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: materia.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Materia.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, materia.FieldID)
		for _, f := range fields {
			if !materia.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != materia.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UnidadeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: materia.FieldUnidadeID,
		})
	}
	if value, ok := muo.mutation.AddedUnidadeID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: materia.FieldUnidadeID,
		})
	}
	if value, ok := muo.mutation.Titulo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldTitulo,
		})
	}
	if value, ok := muo.mutation.URLAmigavel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldURLAmigavel,
		})
	}
	if value, ok := muo.mutation.DataAgendamento(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: materia.FieldDataAgendamento,
		})
	}
	if muo.mutation.DataAgendamentoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: materia.FieldDataAgendamento,
		})
	}
	if value, ok := muo.mutation.Fonte(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldFonte,
		})
	}
	if muo.mutation.FonteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: materia.FieldFonte,
		})
	}
	if value, ok := muo.mutation.LinkFonte(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldLinkFonte,
		})
	}
	if muo.mutation.LinkFonteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: materia.FieldLinkFonte,
		})
	}
	if value, ok := muo.mutation.Texto(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldTexto,
		})
	}
	if value, ok := muo.mutation.ImagemURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldImagemURL,
		})
	}
	if muo.mutation.ImagemURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: materia.FieldImagemURL,
		})
	}
	if value, ok := muo.mutation.Patrocinado(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: materia.FieldPatrocinado,
		})
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: materia.FieldStatus,
		})
	}
	if muo.mutation.CategoriaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   materia.CategoriaTable,
			Columns: []string{materia.CategoriaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: categoria.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.CategoriaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   materia.CategoriaTable,
			Columns: []string{materia.CategoriaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: categoria.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ProfissionalMateriasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   materia.ProfissionalMateriasTable,
			Columns: []string{materia.ProfissionalMateriasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profissionalmaterias.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedProfissionalMateriasIDs(); len(nodes) > 0 && !muo.mutation.ProfissionalMateriasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   materia.ProfissionalMateriasTable,
			Columns: []string{materia.ProfissionalMateriasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profissionalmaterias.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ProfissionalMateriasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   materia.ProfissionalMateriasTable,
			Columns: []string{materia.ProfissionalMateriasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profissionalmaterias.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Materia{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{materia.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
