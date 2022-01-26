// Code generated by entc, DO NOT EDIT.

package especializacao

import (
	"profissional/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Descricao applies equality check predicate on the "descricao" field. It's identical to DescricaoEQ.
func Descricao(v string) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescricao), v))
	})
}

// DescricaoEQ applies the EQ predicate on the "descricao" field.
func DescricaoEQ(v string) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescricao), v))
	})
}

// DescricaoNEQ applies the NEQ predicate on the "descricao" field.
func DescricaoNEQ(v string) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescricao), v))
	})
}

// DescricaoIn applies the In predicate on the "descricao" field.
func DescricaoIn(vs ...string) predicate.Especializacao {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Especializacao(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescricao), v...))
	})
}

// DescricaoNotIn applies the NotIn predicate on the "descricao" field.
func DescricaoNotIn(vs ...string) predicate.Especializacao {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Especializacao(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescricao), v...))
	})
}

// DescricaoGT applies the GT predicate on the "descricao" field.
func DescricaoGT(v string) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescricao), v))
	})
}

// DescricaoGTE applies the GTE predicate on the "descricao" field.
func DescricaoGTE(v string) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescricao), v))
	})
}

// DescricaoLT applies the LT predicate on the "descricao" field.
func DescricaoLT(v string) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescricao), v))
	})
}

// DescricaoLTE applies the LTE predicate on the "descricao" field.
func DescricaoLTE(v string) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescricao), v))
	})
}

// DescricaoContains applies the Contains predicate on the "descricao" field.
func DescricaoContains(v string) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescricao), v))
	})
}

// DescricaoHasPrefix applies the HasPrefix predicate on the "descricao" field.
func DescricaoHasPrefix(v string) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescricao), v))
	})
}

// DescricaoHasSuffix applies the HasSuffix predicate on the "descricao" field.
func DescricaoHasSuffix(v string) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescricao), v))
	})
}

// DescricaoEqualFold applies the EqualFold predicate on the "descricao" field.
func DescricaoEqualFold(v string) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescricao), v))
	})
}

// DescricaoContainsFold applies the ContainsFold predicate on the "descricao" field.
func DescricaoContainsFold(v string) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescricao), v))
	})
}

// HasAreasaude applies the HasEdge predicate on the "areasaude" edge.
func HasAreasaude() predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AreasaudeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AreasaudeTable, AreasaudeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAreasaudeWith applies the HasEdge predicate on the "areasaude" edge with a given conditions (other predicates).
func HasAreasaudeWith(preds ...predicate.AreaSaude) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AreasaudeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AreasaudeTable, AreasaudeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProfissionais applies the HasEdge predicate on the "profissionais" edge.
func HasProfissionais() predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfissionaisTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ProfissionaisTable, ProfissionaisPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfissionaisWith applies the HasEdge predicate on the "profissionais" edge with a given conditions (other predicates).
func HasProfissionaisWith(preds ...predicate.Profissional) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfissionaisInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ProfissionaisTable, ProfissionaisPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Especializacao) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Especializacao) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
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
func Not(p predicate.Especializacao) predicate.Especializacao {
	return predicate.Especializacao(func(s *sql.Selector) {
		p(s.Not())
	})
}
