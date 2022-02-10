// Code generated by entc, DO NOT EDIT.

package materia

import (
	"materia/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UnidadeID applies equality check predicate on the "unidade_id" field. It's identical to UnidadeIDEQ.
func UnidadeID(v int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnidadeID), v))
	})
}

// Titulo applies equality check predicate on the "titulo" field. It's identical to TituloEQ.
func Titulo(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitulo), v))
	})
}

// URLAmigavel applies equality check predicate on the "url_amigavel" field. It's identical to URLAmigavelEQ.
func URLAmigavel(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURLAmigavel), v))
	})
}

// DataAgendamento applies equality check predicate on the "data_agendamento" field. It's identical to DataAgendamentoEQ.
func DataAgendamento(v time.Time) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDataAgendamento), v))
	})
}

// Fonte applies equality check predicate on the "fonte" field. It's identical to FonteEQ.
func Fonte(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFonte), v))
	})
}

// LinkFonte applies equality check predicate on the "link_fonte" field. It's identical to LinkFonteEQ.
func LinkFonte(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLinkFonte), v))
	})
}

// Texto applies equality check predicate on the "texto" field. It's identical to TextoEQ.
func Texto(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTexto), v))
	})
}

// ImagemURL applies equality check predicate on the "imagem_url" field. It's identical to ImagemURLEQ.
func ImagemURL(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImagemURL), v))
	})
}

// Patrocinado applies equality check predicate on the "patrocinado" field. It's identical to PatrocinadoEQ.
func Patrocinado(v bool) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatrocinado), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// UnidadeIDEQ applies the EQ predicate on the "unidade_id" field.
func UnidadeIDEQ(v int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnidadeID), v))
	})
}

// UnidadeIDNEQ applies the NEQ predicate on the "unidade_id" field.
func UnidadeIDNEQ(v int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnidadeID), v))
	})
}

// UnidadeIDIn applies the In predicate on the "unidade_id" field.
func UnidadeIDIn(vs ...int) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUnidadeID), v...))
	})
}

// UnidadeIDNotIn applies the NotIn predicate on the "unidade_id" field.
func UnidadeIDNotIn(vs ...int) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUnidadeID), v...))
	})
}

// UnidadeIDGT applies the GT predicate on the "unidade_id" field.
func UnidadeIDGT(v int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnidadeID), v))
	})
}

// UnidadeIDGTE applies the GTE predicate on the "unidade_id" field.
func UnidadeIDGTE(v int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnidadeID), v))
	})
}

// UnidadeIDLT applies the LT predicate on the "unidade_id" field.
func UnidadeIDLT(v int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnidadeID), v))
	})
}

// UnidadeIDLTE applies the LTE predicate on the "unidade_id" field.
func UnidadeIDLTE(v int) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnidadeID), v))
	})
}

// TituloEQ applies the EQ predicate on the "titulo" field.
func TituloEQ(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitulo), v))
	})
}

// TituloNEQ applies the NEQ predicate on the "titulo" field.
func TituloNEQ(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitulo), v))
	})
}

// TituloIn applies the In predicate on the "titulo" field.
func TituloIn(vs ...string) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitulo), v...))
	})
}

// TituloNotIn applies the NotIn predicate on the "titulo" field.
func TituloNotIn(vs ...string) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitulo), v...))
	})
}

// TituloGT applies the GT predicate on the "titulo" field.
func TituloGT(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitulo), v))
	})
}

// TituloGTE applies the GTE predicate on the "titulo" field.
func TituloGTE(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitulo), v))
	})
}

// TituloLT applies the LT predicate on the "titulo" field.
func TituloLT(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitulo), v))
	})
}

// TituloLTE applies the LTE predicate on the "titulo" field.
func TituloLTE(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitulo), v))
	})
}

// TituloContains applies the Contains predicate on the "titulo" field.
func TituloContains(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitulo), v))
	})
}

// TituloHasPrefix applies the HasPrefix predicate on the "titulo" field.
func TituloHasPrefix(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitulo), v))
	})
}

// TituloHasSuffix applies the HasSuffix predicate on the "titulo" field.
func TituloHasSuffix(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitulo), v))
	})
}

// TituloEqualFold applies the EqualFold predicate on the "titulo" field.
func TituloEqualFold(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitulo), v))
	})
}

// TituloContainsFold applies the ContainsFold predicate on the "titulo" field.
func TituloContainsFold(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitulo), v))
	})
}

// URLAmigavelEQ applies the EQ predicate on the "url_amigavel" field.
func URLAmigavelEQ(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURLAmigavel), v))
	})
}

// URLAmigavelNEQ applies the NEQ predicate on the "url_amigavel" field.
func URLAmigavelNEQ(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURLAmigavel), v))
	})
}

// URLAmigavelIn applies the In predicate on the "url_amigavel" field.
func URLAmigavelIn(vs ...string) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURLAmigavel), v...))
	})
}

// URLAmigavelNotIn applies the NotIn predicate on the "url_amigavel" field.
func URLAmigavelNotIn(vs ...string) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURLAmigavel), v...))
	})
}

// URLAmigavelGT applies the GT predicate on the "url_amigavel" field.
func URLAmigavelGT(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURLAmigavel), v))
	})
}

// URLAmigavelGTE applies the GTE predicate on the "url_amigavel" field.
func URLAmigavelGTE(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURLAmigavel), v))
	})
}

// URLAmigavelLT applies the LT predicate on the "url_amigavel" field.
func URLAmigavelLT(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURLAmigavel), v))
	})
}

// URLAmigavelLTE applies the LTE predicate on the "url_amigavel" field.
func URLAmigavelLTE(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURLAmigavel), v))
	})
}

// URLAmigavelContains applies the Contains predicate on the "url_amigavel" field.
func URLAmigavelContains(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURLAmigavel), v))
	})
}

// URLAmigavelHasPrefix applies the HasPrefix predicate on the "url_amigavel" field.
func URLAmigavelHasPrefix(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURLAmigavel), v))
	})
}

// URLAmigavelHasSuffix applies the HasSuffix predicate on the "url_amigavel" field.
func URLAmigavelHasSuffix(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURLAmigavel), v))
	})
}

// URLAmigavelEqualFold applies the EqualFold predicate on the "url_amigavel" field.
func URLAmigavelEqualFold(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURLAmigavel), v))
	})
}

// URLAmigavelContainsFold applies the ContainsFold predicate on the "url_amigavel" field.
func URLAmigavelContainsFold(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURLAmigavel), v))
	})
}

// DataAgendamentoEQ applies the EQ predicate on the "data_agendamento" field.
func DataAgendamentoEQ(v time.Time) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDataAgendamento), v))
	})
}

// DataAgendamentoNEQ applies the NEQ predicate on the "data_agendamento" field.
func DataAgendamentoNEQ(v time.Time) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDataAgendamento), v))
	})
}

// DataAgendamentoIn applies the In predicate on the "data_agendamento" field.
func DataAgendamentoIn(vs ...time.Time) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDataAgendamento), v...))
	})
}

// DataAgendamentoNotIn applies the NotIn predicate on the "data_agendamento" field.
func DataAgendamentoNotIn(vs ...time.Time) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDataAgendamento), v...))
	})
}

// DataAgendamentoGT applies the GT predicate on the "data_agendamento" field.
func DataAgendamentoGT(v time.Time) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDataAgendamento), v))
	})
}

// DataAgendamentoGTE applies the GTE predicate on the "data_agendamento" field.
func DataAgendamentoGTE(v time.Time) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDataAgendamento), v))
	})
}

// DataAgendamentoLT applies the LT predicate on the "data_agendamento" field.
func DataAgendamentoLT(v time.Time) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDataAgendamento), v))
	})
}

// DataAgendamentoLTE applies the LTE predicate on the "data_agendamento" field.
func DataAgendamentoLTE(v time.Time) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDataAgendamento), v))
	})
}

// DataAgendamentoIsNil applies the IsNil predicate on the "data_agendamento" field.
func DataAgendamentoIsNil() predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDataAgendamento)))
	})
}

// DataAgendamentoNotNil applies the NotNil predicate on the "data_agendamento" field.
func DataAgendamentoNotNil() predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDataAgendamento)))
	})
}

// FonteEQ applies the EQ predicate on the "fonte" field.
func FonteEQ(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFonte), v))
	})
}

// FonteNEQ applies the NEQ predicate on the "fonte" field.
func FonteNEQ(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFonte), v))
	})
}

// FonteIn applies the In predicate on the "fonte" field.
func FonteIn(vs ...string) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFonte), v...))
	})
}

// FonteNotIn applies the NotIn predicate on the "fonte" field.
func FonteNotIn(vs ...string) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFonte), v...))
	})
}

// FonteGT applies the GT predicate on the "fonte" field.
func FonteGT(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFonte), v))
	})
}

// FonteGTE applies the GTE predicate on the "fonte" field.
func FonteGTE(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFonte), v))
	})
}

// FonteLT applies the LT predicate on the "fonte" field.
func FonteLT(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFonte), v))
	})
}

// FonteLTE applies the LTE predicate on the "fonte" field.
func FonteLTE(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFonte), v))
	})
}

// FonteContains applies the Contains predicate on the "fonte" field.
func FonteContains(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFonte), v))
	})
}

// FonteHasPrefix applies the HasPrefix predicate on the "fonte" field.
func FonteHasPrefix(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFonte), v))
	})
}

// FonteHasSuffix applies the HasSuffix predicate on the "fonte" field.
func FonteHasSuffix(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFonte), v))
	})
}

// FonteIsNil applies the IsNil predicate on the "fonte" field.
func FonteIsNil() predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFonte)))
	})
}

// FonteNotNil applies the NotNil predicate on the "fonte" field.
func FonteNotNil() predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFonte)))
	})
}

// FonteEqualFold applies the EqualFold predicate on the "fonte" field.
func FonteEqualFold(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFonte), v))
	})
}

// FonteContainsFold applies the ContainsFold predicate on the "fonte" field.
func FonteContainsFold(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFonte), v))
	})
}

// LinkFonteEQ applies the EQ predicate on the "link_fonte" field.
func LinkFonteEQ(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLinkFonte), v))
	})
}

// LinkFonteNEQ applies the NEQ predicate on the "link_fonte" field.
func LinkFonteNEQ(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLinkFonte), v))
	})
}

// LinkFonteIn applies the In predicate on the "link_fonte" field.
func LinkFonteIn(vs ...string) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLinkFonte), v...))
	})
}

// LinkFonteNotIn applies the NotIn predicate on the "link_fonte" field.
func LinkFonteNotIn(vs ...string) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLinkFonte), v...))
	})
}

// LinkFonteGT applies the GT predicate on the "link_fonte" field.
func LinkFonteGT(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLinkFonte), v))
	})
}

// LinkFonteGTE applies the GTE predicate on the "link_fonte" field.
func LinkFonteGTE(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLinkFonte), v))
	})
}

// LinkFonteLT applies the LT predicate on the "link_fonte" field.
func LinkFonteLT(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLinkFonte), v))
	})
}

// LinkFonteLTE applies the LTE predicate on the "link_fonte" field.
func LinkFonteLTE(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLinkFonte), v))
	})
}

// LinkFonteContains applies the Contains predicate on the "link_fonte" field.
func LinkFonteContains(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLinkFonte), v))
	})
}

// LinkFonteHasPrefix applies the HasPrefix predicate on the "link_fonte" field.
func LinkFonteHasPrefix(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLinkFonte), v))
	})
}

// LinkFonteHasSuffix applies the HasSuffix predicate on the "link_fonte" field.
func LinkFonteHasSuffix(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLinkFonte), v))
	})
}

// LinkFonteIsNil applies the IsNil predicate on the "link_fonte" field.
func LinkFonteIsNil() predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLinkFonte)))
	})
}

// LinkFonteNotNil applies the NotNil predicate on the "link_fonte" field.
func LinkFonteNotNil() predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLinkFonte)))
	})
}

// LinkFonteEqualFold applies the EqualFold predicate on the "link_fonte" field.
func LinkFonteEqualFold(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLinkFonte), v))
	})
}

// LinkFonteContainsFold applies the ContainsFold predicate on the "link_fonte" field.
func LinkFonteContainsFold(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLinkFonte), v))
	})
}

// TextoEQ applies the EQ predicate on the "texto" field.
func TextoEQ(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTexto), v))
	})
}

// TextoNEQ applies the NEQ predicate on the "texto" field.
func TextoNEQ(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTexto), v))
	})
}

// TextoIn applies the In predicate on the "texto" field.
func TextoIn(vs ...string) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTexto), v...))
	})
}

// TextoNotIn applies the NotIn predicate on the "texto" field.
func TextoNotIn(vs ...string) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTexto), v...))
	})
}

// TextoGT applies the GT predicate on the "texto" field.
func TextoGT(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTexto), v))
	})
}

// TextoGTE applies the GTE predicate on the "texto" field.
func TextoGTE(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTexto), v))
	})
}

// TextoLT applies the LT predicate on the "texto" field.
func TextoLT(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTexto), v))
	})
}

// TextoLTE applies the LTE predicate on the "texto" field.
func TextoLTE(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTexto), v))
	})
}

// TextoContains applies the Contains predicate on the "texto" field.
func TextoContains(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTexto), v))
	})
}

// TextoHasPrefix applies the HasPrefix predicate on the "texto" field.
func TextoHasPrefix(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTexto), v))
	})
}

// TextoHasSuffix applies the HasSuffix predicate on the "texto" field.
func TextoHasSuffix(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTexto), v))
	})
}

// TextoEqualFold applies the EqualFold predicate on the "texto" field.
func TextoEqualFold(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTexto), v))
	})
}

// TextoContainsFold applies the ContainsFold predicate on the "texto" field.
func TextoContainsFold(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTexto), v))
	})
}

// ImagemURLEQ applies the EQ predicate on the "imagem_url" field.
func ImagemURLEQ(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImagemURL), v))
	})
}

// ImagemURLNEQ applies the NEQ predicate on the "imagem_url" field.
func ImagemURLNEQ(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImagemURL), v))
	})
}

// ImagemURLIn applies the In predicate on the "imagem_url" field.
func ImagemURLIn(vs ...string) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImagemURL), v...))
	})
}

// ImagemURLNotIn applies the NotIn predicate on the "imagem_url" field.
func ImagemURLNotIn(vs ...string) predicate.Materia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Materia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImagemURL), v...))
	})
}

// ImagemURLGT applies the GT predicate on the "imagem_url" field.
func ImagemURLGT(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImagemURL), v))
	})
}

// ImagemURLGTE applies the GTE predicate on the "imagem_url" field.
func ImagemURLGTE(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImagemURL), v))
	})
}

// ImagemURLLT applies the LT predicate on the "imagem_url" field.
func ImagemURLLT(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImagemURL), v))
	})
}

// ImagemURLLTE applies the LTE predicate on the "imagem_url" field.
func ImagemURLLTE(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImagemURL), v))
	})
}

// ImagemURLContains applies the Contains predicate on the "imagem_url" field.
func ImagemURLContains(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImagemURL), v))
	})
}

// ImagemURLHasPrefix applies the HasPrefix predicate on the "imagem_url" field.
func ImagemURLHasPrefix(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImagemURL), v))
	})
}

// ImagemURLHasSuffix applies the HasSuffix predicate on the "imagem_url" field.
func ImagemURLHasSuffix(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImagemURL), v))
	})
}

// ImagemURLIsNil applies the IsNil predicate on the "imagem_url" field.
func ImagemURLIsNil() predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldImagemURL)))
	})
}

// ImagemURLNotNil applies the NotNil predicate on the "imagem_url" field.
func ImagemURLNotNil() predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldImagemURL)))
	})
}

// ImagemURLEqualFold applies the EqualFold predicate on the "imagem_url" field.
func ImagemURLEqualFold(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImagemURL), v))
	})
}

// ImagemURLContainsFold applies the ContainsFold predicate on the "imagem_url" field.
func ImagemURLContainsFold(v string) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImagemURL), v))
	})
}

// PatrocinadoEQ applies the EQ predicate on the "patrocinado" field.
func PatrocinadoEQ(v bool) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatrocinado), v))
	})
}

// PatrocinadoNEQ applies the NEQ predicate on the "patrocinado" field.
func PatrocinadoNEQ(v bool) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPatrocinado), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// HasCategoria applies the HasEdge predicate on the "categoria" edge.
func HasCategoria() predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoriaTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoriaTable, CategoriaColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoriaWith applies the HasEdge predicate on the "categoria" edge with a given conditions (other predicates).
func HasCategoriaWith(preds ...predicate.Categoria) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoriaInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoriaTable, CategoriaColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProfissionalMaterias applies the HasEdge predicate on the "profissional_materias" edge.
func HasProfissionalMaterias() predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfissionalMateriasTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProfissionalMateriasTable, ProfissionalMateriasColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfissionalMateriasWith applies the HasEdge predicate on the "profissional_materias" edge with a given conditions (other predicates).
func HasProfissionalMateriasWith(preds ...predicate.ProfissionalMaterias) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfissionalMateriasInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProfissionalMateriasTable, ProfissionalMateriasColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Materia) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Materia) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Materia) predicate.Materia {
	return predicate.Materia(func(s *sql.Selector) {
		p(s.Not())
	})
}
