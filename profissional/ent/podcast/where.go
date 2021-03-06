// Code generated by entc, DO NOT EDIT.

package podcast

import (
	"profissional/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Titulo applies equality check predicate on the "titulo" field. It's identical to TituloEQ.
func Titulo(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitulo), v))
	})
}

// Codigo applies equality check predicate on the "codigo" field. It's identical to CodigoEQ.
func Codigo(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCodigo), v))
	})
}

// TituloEQ applies the EQ predicate on the "titulo" field.
func TituloEQ(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitulo), v))
	})
}

// TituloNEQ applies the NEQ predicate on the "titulo" field.
func TituloNEQ(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitulo), v))
	})
}

// TituloIn applies the In predicate on the "titulo" field.
func TituloIn(vs ...string) predicate.Podcast {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Podcast(func(s *sql.Selector) {
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
func TituloNotIn(vs ...string) predicate.Podcast {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Podcast(func(s *sql.Selector) {
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
func TituloGT(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitulo), v))
	})
}

// TituloGTE applies the GTE predicate on the "titulo" field.
func TituloGTE(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitulo), v))
	})
}

// TituloLT applies the LT predicate on the "titulo" field.
func TituloLT(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitulo), v))
	})
}

// TituloLTE applies the LTE predicate on the "titulo" field.
func TituloLTE(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitulo), v))
	})
}

// TituloContains applies the Contains predicate on the "titulo" field.
func TituloContains(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitulo), v))
	})
}

// TituloHasPrefix applies the HasPrefix predicate on the "titulo" field.
func TituloHasPrefix(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitulo), v))
	})
}

// TituloHasSuffix applies the HasSuffix predicate on the "titulo" field.
func TituloHasSuffix(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitulo), v))
	})
}

// TituloEqualFold applies the EqualFold predicate on the "titulo" field.
func TituloEqualFold(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitulo), v))
	})
}

// TituloContainsFold applies the ContainsFold predicate on the "titulo" field.
func TituloContainsFold(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitulo), v))
	})
}

// CodigoEQ applies the EQ predicate on the "codigo" field.
func CodigoEQ(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCodigo), v))
	})
}

// CodigoNEQ applies the NEQ predicate on the "codigo" field.
func CodigoNEQ(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCodigo), v))
	})
}

// CodigoIn applies the In predicate on the "codigo" field.
func CodigoIn(vs ...string) predicate.Podcast {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Podcast(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCodigo), v...))
	})
}

// CodigoNotIn applies the NotIn predicate on the "codigo" field.
func CodigoNotIn(vs ...string) predicate.Podcast {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Podcast(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCodigo), v...))
	})
}

// CodigoGT applies the GT predicate on the "codigo" field.
func CodigoGT(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCodigo), v))
	})
}

// CodigoGTE applies the GTE predicate on the "codigo" field.
func CodigoGTE(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCodigo), v))
	})
}

// CodigoLT applies the LT predicate on the "codigo" field.
func CodigoLT(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCodigo), v))
	})
}

// CodigoLTE applies the LTE predicate on the "codigo" field.
func CodigoLTE(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCodigo), v))
	})
}

// CodigoContains applies the Contains predicate on the "codigo" field.
func CodigoContains(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCodigo), v))
	})
}

// CodigoHasPrefix applies the HasPrefix predicate on the "codigo" field.
func CodigoHasPrefix(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCodigo), v))
	})
}

// CodigoHasSuffix applies the HasSuffix predicate on the "codigo" field.
func CodigoHasSuffix(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCodigo), v))
	})
}

// CodigoEqualFold applies the EqualFold predicate on the "codigo" field.
func CodigoEqualFold(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCodigo), v))
	})
}

// CodigoContainsFold applies the ContainsFold predicate on the "codigo" field.
func CodigoContainsFold(v string) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCodigo), v))
	})
}

// HasProfissional applies the HasEdge predicate on the "profissional" edge.
func HasProfissional() predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfissionalTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProfissionalTable, ProfissionalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfissionalWith applies the HasEdge predicate on the "profissional" edge with a given conditions (other predicates).
func HasProfissionalWith(preds ...predicate.Profissional) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfissionalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProfissionalTable, ProfissionalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Podcast) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Podcast) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
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
func Not(p predicate.Podcast) predicate.Podcast {
	return predicate.Podcast(func(s *sql.Selector) {
		p(s.Not())
	})
}
