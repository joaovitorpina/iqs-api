// Code generated by entc, DO NOT EDIT.

package profissionalmaterias

import (
	"materia/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ProfissionalID applies equality check predicate on the "profissional_id" field. It's identical to ProfissionalIDEQ.
func ProfissionalID(v int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProfissionalID), v))
	})
}

// ProfissionalIDEQ applies the EQ predicate on the "profissional_id" field.
func ProfissionalIDEQ(v int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProfissionalID), v))
	})
}

// ProfissionalIDNEQ applies the NEQ predicate on the "profissional_id" field.
func ProfissionalIDNEQ(v int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProfissionalID), v))
	})
}

// ProfissionalIDIn applies the In predicate on the "profissional_id" field.
func ProfissionalIDIn(vs ...int) predicate.ProfissionalMaterias {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProfissionalID), v...))
	})
}

// ProfissionalIDNotIn applies the NotIn predicate on the "profissional_id" field.
func ProfissionalIDNotIn(vs ...int) predicate.ProfissionalMaterias {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProfissionalID), v...))
	})
}

// ProfissionalIDGT applies the GT predicate on the "profissional_id" field.
func ProfissionalIDGT(v int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProfissionalID), v))
	})
}

// ProfissionalIDGTE applies the GTE predicate on the "profissional_id" field.
func ProfissionalIDGTE(v int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProfissionalID), v))
	})
}

// ProfissionalIDLT applies the LT predicate on the "profissional_id" field.
func ProfissionalIDLT(v int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProfissionalID), v))
	})
}

// ProfissionalIDLTE applies the LTE predicate on the "profissional_id" field.
func ProfissionalIDLTE(v int) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProfissionalID), v))
	})
}

// HasMateria applies the HasEdge predicate on the "materia" edge.
func HasMateria() predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MateriaTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MateriaTable, MateriaColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMateriaWith applies the HasEdge predicate on the "materia" edge with a given conditions (other predicates).
func HasMateriaWith(preds ...predicate.Materia) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MateriaInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MateriaTable, MateriaColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProfissionalMaterias) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProfissionalMaterias) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
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
func Not(p predicate.ProfissionalMaterias) predicate.ProfissionalMaterias {
	return predicate.ProfissionalMaterias(func(s *sql.Selector) {
		p(s.Not())
	})
}
