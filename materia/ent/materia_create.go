// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"materia/ent/categoria"
	"materia/ent/materia"
	"materia/ent/profissionalmaterias"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MateriaCreate is the builder for creating a Materia entity.
type MateriaCreate struct {
	config
	mutation *MateriaMutation
	hooks    []Hook
}

// SetUnidadeID sets the "unidade_id" field.
func (mc *MateriaCreate) SetUnidadeID(i int) *MateriaCreate {
	mc.mutation.SetUnidadeID(i)
	return mc
}

// SetTitulo sets the "titulo" field.
func (mc *MateriaCreate) SetTitulo(s string) *MateriaCreate {
	mc.mutation.SetTitulo(s)
	return mc
}

// SetURLAmigavel sets the "url_amigavel" field.
func (mc *MateriaCreate) SetURLAmigavel(s string) *MateriaCreate {
	mc.mutation.SetURLAmigavel(s)
	return mc
}

// SetDataAgendamento sets the "data_agendamento" field.
func (mc *MateriaCreate) SetDataAgendamento(t time.Time) *MateriaCreate {
	mc.mutation.SetDataAgendamento(t)
	return mc
}

// SetNillableDataAgendamento sets the "data_agendamento" field if the given value is not nil.
func (mc *MateriaCreate) SetNillableDataAgendamento(t *time.Time) *MateriaCreate {
	if t != nil {
		mc.SetDataAgendamento(*t)
	}
	return mc
}

// SetFonte sets the "fonte" field.
func (mc *MateriaCreate) SetFonte(s string) *MateriaCreate {
	mc.mutation.SetFonte(s)
	return mc
}

// SetNillableFonte sets the "fonte" field if the given value is not nil.
func (mc *MateriaCreate) SetNillableFonte(s *string) *MateriaCreate {
	if s != nil {
		mc.SetFonte(*s)
	}
	return mc
}

// SetLinkFonte sets the "link_fonte" field.
func (mc *MateriaCreate) SetLinkFonte(s string) *MateriaCreate {
	mc.mutation.SetLinkFonte(s)
	return mc
}

// SetNillableLinkFonte sets the "link_fonte" field if the given value is not nil.
func (mc *MateriaCreate) SetNillableLinkFonte(s *string) *MateriaCreate {
	if s != nil {
		mc.SetLinkFonte(*s)
	}
	return mc
}

// SetTexto sets the "texto" field.
func (mc *MateriaCreate) SetTexto(s string) *MateriaCreate {
	mc.mutation.SetTexto(s)
	return mc
}

// SetImagemURL sets the "imagem_url" field.
func (mc *MateriaCreate) SetImagemURL(s string) *MateriaCreate {
	mc.mutation.SetImagemURL(s)
	return mc
}

// SetNillableImagemURL sets the "imagem_url" field if the given value is not nil.
func (mc *MateriaCreate) SetNillableImagemURL(s *string) *MateriaCreate {
	if s != nil {
		mc.SetImagemURL(*s)
	}
	return mc
}

// SetPatrocinado sets the "patrocinado" field.
func (mc *MateriaCreate) SetPatrocinado(b bool) *MateriaCreate {
	mc.mutation.SetPatrocinado(b)
	return mc
}

// SetNillablePatrocinado sets the "patrocinado" field if the given value is not nil.
func (mc *MateriaCreate) SetNillablePatrocinado(b *bool) *MateriaCreate {
	if b != nil {
		mc.SetPatrocinado(*b)
	}
	return mc
}

// SetStatus sets the "status" field.
func (mc *MateriaCreate) SetStatus(b bool) *MateriaCreate {
	mc.mutation.SetStatus(b)
	return mc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mc *MateriaCreate) SetNillableStatus(b *bool) *MateriaCreate {
	if b != nil {
		mc.SetStatus(*b)
	}
	return mc
}

// SetCategoriaID sets the "categoria" edge to the Categoria entity by ID.
func (mc *MateriaCreate) SetCategoriaID(id int) *MateriaCreate {
	mc.mutation.SetCategoriaID(id)
	return mc
}

// SetCategoria sets the "categoria" edge to the Categoria entity.
func (mc *MateriaCreate) SetCategoria(c *Categoria) *MateriaCreate {
	return mc.SetCategoriaID(c.ID)
}

// AddProfissionalMateriaIDs adds the "profissional_materias" edge to the ProfissionalMaterias entity by IDs.
func (mc *MateriaCreate) AddProfissionalMateriaIDs(ids ...int) *MateriaCreate {
	mc.mutation.AddProfissionalMateriaIDs(ids...)
	return mc
}

// AddProfissionalMaterias adds the "profissional_materias" edges to the ProfissionalMaterias entity.
func (mc *MateriaCreate) AddProfissionalMaterias(p ...*ProfissionalMaterias) *MateriaCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddProfissionalMateriaIDs(ids...)
}

// Mutation returns the MateriaMutation object of the builder.
func (mc *MateriaCreate) Mutation() *MateriaMutation {
	return mc.mutation
}

// Save creates the Materia in the database.
func (mc *MateriaCreate) Save(ctx context.Context) (*Materia, error) {
	var (
		err  error
		node *Materia
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MateriaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MateriaCreate) SaveX(ctx context.Context) *Materia {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MateriaCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MateriaCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MateriaCreate) defaults() {
	if _, ok := mc.mutation.Patrocinado(); !ok {
		v := materia.DefaultPatrocinado
		mc.mutation.SetPatrocinado(v)
	}
	if _, ok := mc.mutation.Status(); !ok {
		v := materia.DefaultStatus
		mc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MateriaCreate) check() error {
	if _, ok := mc.mutation.UnidadeID(); !ok {
		return &ValidationError{Name: "unidade_id", err: errors.New(`ent: missing required field "Materia.unidade_id"`)}
	}
	if _, ok := mc.mutation.Titulo(); !ok {
		return &ValidationError{Name: "titulo", err: errors.New(`ent: missing required field "Materia.titulo"`)}
	}
	if _, ok := mc.mutation.URLAmigavel(); !ok {
		return &ValidationError{Name: "url_amigavel", err: errors.New(`ent: missing required field "Materia.url_amigavel"`)}
	}
	if _, ok := mc.mutation.Texto(); !ok {
		return &ValidationError{Name: "texto", err: errors.New(`ent: missing required field "Materia.texto"`)}
	}
	if _, ok := mc.mutation.Patrocinado(); !ok {
		return &ValidationError{Name: "patrocinado", err: errors.New(`ent: missing required field "Materia.patrocinado"`)}
	}
	if _, ok := mc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Materia.status"`)}
	}
	if _, ok := mc.mutation.CategoriaID(); !ok {
		return &ValidationError{Name: "categoria", err: errors.New(`ent: missing required edge "Materia.categoria"`)}
	}
	return nil
}

func (mc *MateriaCreate) sqlSave(ctx context.Context) (*Materia, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MateriaCreate) createSpec() (*Materia, *sqlgraph.CreateSpec) {
	var (
		_node = &Materia{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: materia.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: materia.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.UnidadeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: materia.FieldUnidadeID,
		})
		_node.UnidadeID = value
	}
	if value, ok := mc.mutation.Titulo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldTitulo,
		})
		_node.Titulo = value
	}
	if value, ok := mc.mutation.URLAmigavel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldURLAmigavel,
		})
		_node.URLAmigavel = value
	}
	if value, ok := mc.mutation.DataAgendamento(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: materia.FieldDataAgendamento,
		})
		_node.DataAgendamento = value
	}
	if value, ok := mc.mutation.Fonte(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldFonte,
		})
		_node.Fonte = value
	}
	if value, ok := mc.mutation.LinkFonte(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldLinkFonte,
		})
		_node.LinkFonte = value
	}
	if value, ok := mc.mutation.Texto(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldTexto,
		})
		_node.Texto = value
	}
	if value, ok := mc.mutation.ImagemURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: materia.FieldImagemURL,
		})
		_node.ImagemURL = value
	}
	if value, ok := mc.mutation.Patrocinado(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: materia.FieldPatrocinado,
		})
		_node.Patrocinado = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: materia.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := mc.mutation.CategoriaIDs(); len(nodes) > 0 {
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
		_node.categoria_materias = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ProfissionalMateriasIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MateriaCreateBulk is the builder for creating many Materia entities in bulk.
type MateriaCreateBulk struct {
	config
	builders []*MateriaCreate
}

// Save creates the Materia entities in the database.
func (mcb *MateriaCreateBulk) Save(ctx context.Context) ([]*Materia, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Materia, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MateriaMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MateriaCreateBulk) SaveX(ctx context.Context) []*Materia {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MateriaCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MateriaCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
