// Code generated by entc, DO NOT EDIT.

package estado

import (
	"endereco/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Nome applies equality check predicate on the "nome" field. It's identical to NomeEQ.
func Nome(v string) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNome), v))
	})
}

// NomeEQ applies the EQ predicate on the "nome" field.
func NomeEQ(v string) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNome), v))
	})
}

// NomeNEQ applies the NEQ predicate on the "nome" field.
func NomeNEQ(v string) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNome), v))
	})
}

// NomeIn applies the In predicate on the "nome" field.
func NomeIn(vs ...string) predicate.Estado {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Estado(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNome), v...))
	})
}

// NomeNotIn applies the NotIn predicate on the "nome" field.
func NomeNotIn(vs ...string) predicate.Estado {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Estado(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNome), v...))
	})
}

// NomeGT applies the GT predicate on the "nome" field.
func NomeGT(v string) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNome), v))
	})
}

// NomeGTE applies the GTE predicate on the "nome" field.
func NomeGTE(v string) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNome), v))
	})
}

// NomeLT applies the LT predicate on the "nome" field.
func NomeLT(v string) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNome), v))
	})
}

// NomeLTE applies the LTE predicate on the "nome" field.
func NomeLTE(v string) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNome), v))
	})
}

// NomeContains applies the Contains predicate on the "nome" field.
func NomeContains(v string) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNome), v))
	})
}

// NomeHasPrefix applies the HasPrefix predicate on the "nome" field.
func NomeHasPrefix(v string) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNome), v))
	})
}

// NomeHasSuffix applies the HasSuffix predicate on the "nome" field.
func NomeHasSuffix(v string) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNome), v))
	})
}

// NomeEqualFold applies the EqualFold predicate on the "nome" field.
func NomeEqualFold(v string) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNome), v))
	})
}

// NomeContainsFold applies the ContainsFold predicate on the "nome" field.
func NomeContainsFold(v string) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNome), v))
	})
}

// HasCidades applies the HasEdge predicate on the "cidades" edge.
func HasCidades() predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CidadesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CidadesTable, CidadesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCidadesWith applies the HasEdge predicate on the "cidades" edge with a given conditions (other predicates).
func HasCidadesWith(preds ...predicate.Cidade) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CidadesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CidadesTable, CidadesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Estado) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Estado) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
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
func Not(p predicate.Estado) predicate.Estado {
	return predicate.Estado(func(s *sql.Selector) {
		p(s.Not())
	})
}
