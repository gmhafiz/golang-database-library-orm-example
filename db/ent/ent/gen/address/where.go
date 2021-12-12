// Code generated by entc, DO NOT EDIT.

package address

import (
	"godb/db/ent/ent/gen/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
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
func IDNotIn(ids ...uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
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
func IDGT(id uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Line1 applies equality check predicate on the "line_1" field. It's identical to Line1EQ.
func Line1(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLine1), v))
	})
}

// Line2 applies equality check predicate on the "line_2" field. It's identical to Line2EQ.
func Line2(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLine2), v))
	})
}

// Postcode applies equality check predicate on the "postcode" field. It's identical to PostcodeEQ.
func Postcode(v uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPostcode), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// Line1EQ applies the EQ predicate on the "line_1" field.
func Line1EQ(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLine1), v))
	})
}

// Line1NEQ applies the NEQ predicate on the "line_1" field.
func Line1NEQ(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLine1), v))
	})
}

// Line1In applies the In predicate on the "line_1" field.
func Line1In(vs ...string) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLine1), v...))
	})
}

// Line1NotIn applies the NotIn predicate on the "line_1" field.
func Line1NotIn(vs ...string) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLine1), v...))
	})
}

// Line1GT applies the GT predicate on the "line_1" field.
func Line1GT(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLine1), v))
	})
}

// Line1GTE applies the GTE predicate on the "line_1" field.
func Line1GTE(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLine1), v))
	})
}

// Line1LT applies the LT predicate on the "line_1" field.
func Line1LT(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLine1), v))
	})
}

// Line1LTE applies the LTE predicate on the "line_1" field.
func Line1LTE(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLine1), v))
	})
}

// Line1Contains applies the Contains predicate on the "line_1" field.
func Line1Contains(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLine1), v))
	})
}

// Line1HasPrefix applies the HasPrefix predicate on the "line_1" field.
func Line1HasPrefix(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLine1), v))
	})
}

// Line1HasSuffix applies the HasSuffix predicate on the "line_1" field.
func Line1HasSuffix(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLine1), v))
	})
}

// Line1EqualFold applies the EqualFold predicate on the "line_1" field.
func Line1EqualFold(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLine1), v))
	})
}

// Line1ContainsFold applies the ContainsFold predicate on the "line_1" field.
func Line1ContainsFold(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLine1), v))
	})
}

// Line2EQ applies the EQ predicate on the "line_2" field.
func Line2EQ(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLine2), v))
	})
}

// Line2NEQ applies the NEQ predicate on the "line_2" field.
func Line2NEQ(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLine2), v))
	})
}

// Line2In applies the In predicate on the "line_2" field.
func Line2In(vs ...string) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLine2), v...))
	})
}

// Line2NotIn applies the NotIn predicate on the "line_2" field.
func Line2NotIn(vs ...string) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLine2), v...))
	})
}

// Line2GT applies the GT predicate on the "line_2" field.
func Line2GT(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLine2), v))
	})
}

// Line2GTE applies the GTE predicate on the "line_2" field.
func Line2GTE(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLine2), v))
	})
}

// Line2LT applies the LT predicate on the "line_2" field.
func Line2LT(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLine2), v))
	})
}

// Line2LTE applies the LTE predicate on the "line_2" field.
func Line2LTE(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLine2), v))
	})
}

// Line2Contains applies the Contains predicate on the "line_2" field.
func Line2Contains(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLine2), v))
	})
}

// Line2HasPrefix applies the HasPrefix predicate on the "line_2" field.
func Line2HasPrefix(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLine2), v))
	})
}

// Line2HasSuffix applies the HasSuffix predicate on the "line_2" field.
func Line2HasSuffix(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLine2), v))
	})
}

// Line2IsNil applies the IsNil predicate on the "line_2" field.
func Line2IsNil() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLine2)))
	})
}

// Line2NotNil applies the NotNil predicate on the "line_2" field.
func Line2NotNil() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLine2)))
	})
}

// Line2EqualFold applies the EqualFold predicate on the "line_2" field.
func Line2EqualFold(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLine2), v))
	})
}

// Line2ContainsFold applies the ContainsFold predicate on the "line_2" field.
func Line2ContainsFold(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLine2), v))
	})
}

// PostcodeEQ applies the EQ predicate on the "postcode" field.
func PostcodeEQ(v uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPostcode), v))
	})
}

// PostcodeNEQ applies the NEQ predicate on the "postcode" field.
func PostcodeNEQ(v uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPostcode), v))
	})
}

// PostcodeIn applies the In predicate on the "postcode" field.
func PostcodeIn(vs ...uint) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPostcode), v...))
	})
}

// PostcodeNotIn applies the NotIn predicate on the "postcode" field.
func PostcodeNotIn(vs ...uint) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPostcode), v...))
	})
}

// PostcodeGT applies the GT predicate on the "postcode" field.
func PostcodeGT(v uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPostcode), v))
	})
}

// PostcodeGTE applies the GTE predicate on the "postcode" field.
func PostcodeGTE(v uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPostcode), v))
	})
}

// PostcodeLT applies the LT predicate on the "postcode" field.
func PostcodeLT(v uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPostcode), v))
	})
}

// PostcodeLTE applies the LTE predicate on the "postcode" field.
func PostcodeLTE(v uint) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPostcode), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// HasCountry applies the HasEdge predicate on the "country" edge.
func HasCountry() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CountryTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CountryTable, CountryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCountryWith applies the HasEdge predicate on the "country" edge with a given conditions (other predicates).
func HasCountryWith(preds ...predicate.Country) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CountryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CountryTable, CountryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
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
func Not(p predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		p(s.Not())
	})
}
