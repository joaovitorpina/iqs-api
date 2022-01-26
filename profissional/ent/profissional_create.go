// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"profissional/ent/convenio"
	"profissional/ent/especializacao"
	"profissional/ent/foto"
	"profissional/ent/podcast"
	"profissional/ent/profissional"
	"profissional/ent/tratamento"
	"profissional/ent/video"
	"profissional/ent/whatsapp"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfissionalCreate is the builder for creating a Profissional entity.
type ProfissionalCreate struct {
	config
	mutation *ProfissionalMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *ProfissionalCreate) SetCreateTime(t time.Time) *ProfissionalCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *ProfissionalCreate) SetNillableCreateTime(t *time.Time) *ProfissionalCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *ProfissionalCreate) SetUpdateTime(t time.Time) *ProfissionalCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *ProfissionalCreate) SetNillableUpdateTime(t *time.Time) *ProfissionalCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetNome sets the "nome" field.
func (pc *ProfissionalCreate) SetNome(s string) *ProfissionalCreate {
	pc.mutation.SetNome(s)
	return pc
}

// SetURLAmigavel sets the "url_amigavel" field.
func (pc *ProfissionalCreate) SetURLAmigavel(s string) *ProfissionalCreate {
	pc.mutation.SetURLAmigavel(s)
	return pc
}

// SetRecomendado sets the "recomendado" field.
func (pc *ProfissionalCreate) SetRecomendado(b bool) *ProfissionalCreate {
	pc.mutation.SetRecomendado(b)
	return pc
}

// SetNillableRecomendado sets the "recomendado" field if the given value is not nil.
func (pc *ProfissionalCreate) SetNillableRecomendado(b *bool) *ProfissionalCreate {
	if b != nil {
		pc.SetRecomendado(*b)
	}
	return pc
}

// SetAtivo sets the "ativo" field.
func (pc *ProfissionalCreate) SetAtivo(b bool) *ProfissionalCreate {
	pc.mutation.SetAtivo(b)
	return pc
}

// SetNillableAtivo sets the "ativo" field if the given value is not nil.
func (pc *ProfissionalCreate) SetNillableAtivo(b *bool) *ProfissionalCreate {
	if b != nil {
		pc.SetAtivo(*b)
	}
	return pc
}

// SetSobre sets the "sobre" field.
func (pc *ProfissionalCreate) SetSobre(s string) *ProfissionalCreate {
	pc.mutation.SetSobre(s)
	return pc
}

// SetNillableSobre sets the "sobre" field if the given value is not nil.
func (pc *ProfissionalCreate) SetNillableSobre(s *string) *ProfissionalCreate {
	if s != nil {
		pc.SetSobre(*s)
	}
	return pc
}

// SetConselho sets the "conselho" field.
func (pc *ProfissionalCreate) SetConselho(s string) *ProfissionalCreate {
	pc.mutation.SetConselho(s)
	return pc
}

// SetNillableConselho sets the "conselho" field if the given value is not nil.
func (pc *ProfissionalCreate) SetNillableConselho(s *string) *ProfissionalCreate {
	if s != nil {
		pc.SetConselho(*s)
	}
	return pc
}

// SetNumeroIdentificacao sets the "numero_identificacao" field.
func (pc *ProfissionalCreate) SetNumeroIdentificacao(s string) *ProfissionalCreate {
	pc.mutation.SetNumeroIdentificacao(s)
	return pc
}

// SetNillableNumeroIdentificacao sets the "numero_identificacao" field if the given value is not nil.
func (pc *ProfissionalCreate) SetNillableNumeroIdentificacao(s *string) *ProfissionalCreate {
	if s != nil {
		pc.SetNumeroIdentificacao(*s)
	}
	return pc
}

// SetTelefone sets the "telefone" field.
func (pc *ProfissionalCreate) SetTelefone(i int32) *ProfissionalCreate {
	pc.mutation.SetTelefone(i)
	return pc
}

// SetCelular sets the "celular" field.
func (pc *ProfissionalCreate) SetCelular(i int32) *ProfissionalCreate {
	pc.mutation.SetCelular(i)
	return pc
}

// SetEmail sets the "email" field.
func (pc *ProfissionalCreate) SetEmail(s string) *ProfissionalCreate {
	pc.mutation.SetEmail(s)
	return pc
}

// SetSite sets the "site" field.
func (pc *ProfissionalCreate) SetSite(s string) *ProfissionalCreate {
	pc.mutation.SetSite(s)
	return pc
}

// SetFacebook sets the "facebook" field.
func (pc *ProfissionalCreate) SetFacebook(s string) *ProfissionalCreate {
	pc.mutation.SetFacebook(s)
	return pc
}

// SetNillableFacebook sets the "facebook" field if the given value is not nil.
func (pc *ProfissionalCreate) SetNillableFacebook(s *string) *ProfissionalCreate {
	if s != nil {
		pc.SetFacebook(*s)
	}
	return pc
}

// SetInstagram sets the "instagram" field.
func (pc *ProfissionalCreate) SetInstagram(s string) *ProfissionalCreate {
	pc.mutation.SetInstagram(s)
	return pc
}

// SetNillableInstagram sets the "instagram" field if the given value is not nil.
func (pc *ProfissionalCreate) SetNillableInstagram(s *string) *ProfissionalCreate {
	if s != nil {
		pc.SetInstagram(*s)
	}
	return pc
}

// SetYoutube sets the "youtube" field.
func (pc *ProfissionalCreate) SetYoutube(s string) *ProfissionalCreate {
	pc.mutation.SetYoutube(s)
	return pc
}

// SetNillableYoutube sets the "youtube" field if the given value is not nil.
func (pc *ProfissionalCreate) SetNillableYoutube(s *string) *ProfissionalCreate {
	if s != nil {
		pc.SetYoutube(*s)
	}
	return pc
}

// SetLinkedin sets the "linkedin" field.
func (pc *ProfissionalCreate) SetLinkedin(s string) *ProfissionalCreate {
	pc.mutation.SetLinkedin(s)
	return pc
}

// SetNillableLinkedin sets the "linkedin" field if the given value is not nil.
func (pc *ProfissionalCreate) SetNillableLinkedin(s *string) *ProfissionalCreate {
	if s != nil {
		pc.SetLinkedin(*s)
	}
	return pc
}

// SetUnidadeID sets the "unidade_id" field.
func (pc *ProfissionalCreate) SetUnidadeID(i int) *ProfissionalCreate {
	pc.mutation.SetUnidadeID(i)
	return pc
}

// SetEnderecoID sets the "endereco_id" field.
func (pc *ProfissionalCreate) SetEnderecoID(i int) *ProfissionalCreate {
	pc.mutation.SetEnderecoID(i)
	return pc
}

// SetImagemPerfilURL sets the "imagem_perfil_url" field.
func (pc *ProfissionalCreate) SetImagemPerfilURL(s string) *ProfissionalCreate {
	pc.mutation.SetImagemPerfilURL(s)
	return pc
}

// AddWhatsappIDs adds the "whatsapps" edge to the WhatsApp entity by IDs.
func (pc *ProfissionalCreate) AddWhatsappIDs(ids ...int) *ProfissionalCreate {
	pc.mutation.AddWhatsappIDs(ids...)
	return pc
}

// AddWhatsapps adds the "whatsapps" edges to the WhatsApp entity.
func (pc *ProfissionalCreate) AddWhatsapps(w ...*WhatsApp) *ProfissionalCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pc.AddWhatsappIDs(ids...)
}

// AddVideoIDs adds the "videos" edge to the Video entity by IDs.
func (pc *ProfissionalCreate) AddVideoIDs(ids ...int) *ProfissionalCreate {
	pc.mutation.AddVideoIDs(ids...)
	return pc
}

// AddVideos adds the "videos" edges to the Video entity.
func (pc *ProfissionalCreate) AddVideos(v ...*Video) *ProfissionalCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return pc.AddVideoIDs(ids...)
}

// AddTratamentoIDs adds the "tratamentos" edge to the Tratamento entity by IDs.
func (pc *ProfissionalCreate) AddTratamentoIDs(ids ...int) *ProfissionalCreate {
	pc.mutation.AddTratamentoIDs(ids...)
	return pc
}

// AddTratamentos adds the "tratamentos" edges to the Tratamento entity.
func (pc *ProfissionalCreate) AddTratamentos(t ...*Tratamento) *ProfissionalCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddTratamentoIDs(ids...)
}

// AddPodcastIDs adds the "podcasts" edge to the Podcast entity by IDs.
func (pc *ProfissionalCreate) AddPodcastIDs(ids ...int) *ProfissionalCreate {
	pc.mutation.AddPodcastIDs(ids...)
	return pc
}

// AddPodcasts adds the "podcasts" edges to the Podcast entity.
func (pc *ProfissionalCreate) AddPodcasts(p ...*Podcast) *ProfissionalCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPodcastIDs(ids...)
}

// AddFotoIDs adds the "fotos" edge to the Foto entity by IDs.
func (pc *ProfissionalCreate) AddFotoIDs(ids ...int) *ProfissionalCreate {
	pc.mutation.AddFotoIDs(ids...)
	return pc
}

// AddFotos adds the "fotos" edges to the Foto entity.
func (pc *ProfissionalCreate) AddFotos(f ...*Foto) *ProfissionalCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pc.AddFotoIDs(ids...)
}

// AddConvenioIDs adds the "convenios" edge to the Convenio entity by IDs.
func (pc *ProfissionalCreate) AddConvenioIDs(ids ...int) *ProfissionalCreate {
	pc.mutation.AddConvenioIDs(ids...)
	return pc
}

// AddConvenios adds the "convenios" edges to the Convenio entity.
func (pc *ProfissionalCreate) AddConvenios(c ...*Convenio) *ProfissionalCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddConvenioIDs(ids...)
}

// AddEspecializacoIDs adds the "especializacoes" edge to the Especializacao entity by IDs.
func (pc *ProfissionalCreate) AddEspecializacoIDs(ids ...int) *ProfissionalCreate {
	pc.mutation.AddEspecializacoIDs(ids...)
	return pc
}

// AddEspecializacoes adds the "especializacoes" edges to the Especializacao entity.
func (pc *ProfissionalCreate) AddEspecializacoes(e ...*Especializacao) *ProfissionalCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pc.AddEspecializacoIDs(ids...)
}

// Mutation returns the ProfissionalMutation object of the builder.
func (pc *ProfissionalCreate) Mutation() *ProfissionalMutation {
	return pc.mutation
}

// Save creates the Profissional in the database.
func (pc *ProfissionalCreate) Save(ctx context.Context) (*Profissional, error) {
	var (
		err  error
		node *Profissional
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfissionalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfissionalCreate) SaveX(ctx context.Context) *Profissional {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProfissionalCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProfissionalCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProfissionalCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := profissional.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := profissional.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.Recomendado(); !ok {
		v := profissional.DefaultRecomendado
		pc.mutation.SetRecomendado(v)
	}
	if _, ok := pc.mutation.Ativo(); !ok {
		v := profissional.DefaultAtivo
		pc.mutation.SetAtivo(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProfissionalCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "update_time"`)}
	}
	if _, ok := pc.mutation.Nome(); !ok {
		return &ValidationError{Name: "nome", err: errors.New(`ent: missing required field "nome"`)}
	}
	if v, ok := pc.mutation.Nome(); ok {
		if err := profissional.NomeValidator(v); err != nil {
			return &ValidationError{Name: "nome", err: fmt.Errorf(`ent: validator failed for field "nome": %w`, err)}
		}
	}
	if _, ok := pc.mutation.URLAmigavel(); !ok {
		return &ValidationError{Name: "url_amigavel", err: errors.New(`ent: missing required field "url_amigavel"`)}
	}
	if v, ok := pc.mutation.URLAmigavel(); ok {
		if err := profissional.URLAmigavelValidator(v); err != nil {
			return &ValidationError{Name: "url_amigavel", err: fmt.Errorf(`ent: validator failed for field "url_amigavel": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Recomendado(); !ok {
		return &ValidationError{Name: "recomendado", err: errors.New(`ent: missing required field "recomendado"`)}
	}
	if _, ok := pc.mutation.Ativo(); !ok {
		return &ValidationError{Name: "ativo", err: errors.New(`ent: missing required field "ativo"`)}
	}
	if _, ok := pc.mutation.Telefone(); !ok {
		return &ValidationError{Name: "telefone", err: errors.New(`ent: missing required field "telefone"`)}
	}
	if _, ok := pc.mutation.Celular(); !ok {
		return &ValidationError{Name: "celular", err: errors.New(`ent: missing required field "celular"`)}
	}
	if _, ok := pc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "email"`)}
	}
	if _, ok := pc.mutation.Site(); !ok {
		return &ValidationError{Name: "site", err: errors.New(`ent: missing required field "site"`)}
	}
	if _, ok := pc.mutation.UnidadeID(); !ok {
		return &ValidationError{Name: "unidade_id", err: errors.New(`ent: missing required field "unidade_id"`)}
	}
	if v, ok := pc.mutation.UnidadeID(); ok {
		if err := profissional.UnidadeIDValidator(v); err != nil {
			return &ValidationError{Name: "unidade_id", err: fmt.Errorf(`ent: validator failed for field "unidade_id": %w`, err)}
		}
	}
	if _, ok := pc.mutation.EnderecoID(); !ok {
		return &ValidationError{Name: "endereco_id", err: errors.New(`ent: missing required field "endereco_id"`)}
	}
	if v, ok := pc.mutation.EnderecoID(); ok {
		if err := profissional.EnderecoIDValidator(v); err != nil {
			return &ValidationError{Name: "endereco_id", err: fmt.Errorf(`ent: validator failed for field "endereco_id": %w`, err)}
		}
	}
	if _, ok := pc.mutation.ImagemPerfilURL(); !ok {
		return &ValidationError{Name: "imagem_perfil_url", err: errors.New(`ent: missing required field "imagem_perfil_url"`)}
	}
	return nil
}

func (pc *ProfissionalCreate) sqlSave(ctx context.Context) (*Profissional, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *ProfissionalCreate) createSpec() (*Profissional, *sqlgraph.CreateSpec) {
	var (
		_node = &Profissional{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: profissional.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profissional.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profissional.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profissional.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Nome(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profissional.FieldNome,
		})
		_node.Nome = value
	}
	if value, ok := pc.mutation.URLAmigavel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profissional.FieldURLAmigavel,
		})
		_node.URLAmigavel = value
	}
	if value, ok := pc.mutation.Recomendado(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: profissional.FieldRecomendado,
		})
		_node.Recomendado = value
	}
	if value, ok := pc.mutation.Ativo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: profissional.FieldAtivo,
		})
		_node.Ativo = value
	}
	if value, ok := pc.mutation.Sobre(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profissional.FieldSobre,
		})
		_node.Sobre = value
	}
	if value, ok := pc.mutation.Conselho(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profissional.FieldConselho,
		})
		_node.Conselho = value
	}
	if value, ok := pc.mutation.NumeroIdentificacao(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profissional.FieldNumeroIdentificacao,
		})
		_node.NumeroIdentificacao = value
	}
	if value, ok := pc.mutation.Telefone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: profissional.FieldTelefone,
		})
		_node.Telefone = value
	}
	if value, ok := pc.mutation.Celular(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: profissional.FieldCelular,
		})
		_node.Celular = value
	}
	if value, ok := pc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profissional.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := pc.mutation.Site(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profissional.FieldSite,
		})
		_node.Site = value
	}
	if value, ok := pc.mutation.Facebook(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profissional.FieldFacebook,
		})
		_node.Facebook = value
	}
	if value, ok := pc.mutation.Instagram(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profissional.FieldInstagram,
		})
		_node.Instagram = value
	}
	if value, ok := pc.mutation.Youtube(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profissional.FieldYoutube,
		})
		_node.Youtube = value
	}
	if value, ok := pc.mutation.Linkedin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profissional.FieldLinkedin,
		})
		_node.Linkedin = value
	}
	if value, ok := pc.mutation.UnidadeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profissional.FieldUnidadeID,
		})
		_node.UnidadeID = value
	}
	if value, ok := pc.mutation.EnderecoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profissional.FieldEnderecoID,
		})
		_node.EnderecoID = value
	}
	if value, ok := pc.mutation.ImagemPerfilURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profissional.FieldImagemPerfilURL,
		})
		_node.ImagemPerfilURL = value
	}
	if nodes := pc.mutation.WhatsappsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profissional.WhatsappsTable,
			Columns: []string{profissional.WhatsappsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: whatsapp.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.VideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profissional.VideosTable,
			Columns: []string{profissional.VideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TratamentosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profissional.TratamentosTable,
			Columns: []string{profissional.TratamentosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tratamento.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PodcastsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profissional.PodcastsTable,
			Columns: []string{profissional.PodcastsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: podcast.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.FotosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profissional.FotosTable,
			Columns: []string{profissional.FotosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: foto.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ConveniosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profissional.ConveniosTable,
			Columns: []string{profissional.ConveniosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: convenio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.EspecializacoesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   profissional.EspecializacoesTable,
			Columns: profissional.EspecializacoesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: especializacao.FieldID,
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

// ProfissionalCreateBulk is the builder for creating many Profissional entities in bulk.
type ProfissionalCreateBulk struct {
	config
	builders []*ProfissionalCreate
}

// Save creates the Profissional entities in the database.
func (pcb *ProfissionalCreateBulk) Save(ctx context.Context) ([]*Profissional, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Profissional, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfissionalMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProfissionalCreateBulk) SaveX(ctx context.Context) []*Profissional {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProfissionalCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProfissionalCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
