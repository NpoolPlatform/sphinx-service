// Code generated by entc, DO NOT EDIT.

package review

import (
	"sphinx/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IsApproved applies equality check predicate on the "is_approved" field. It's identical to IsApprovedEQ.
func IsApproved(v bool) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsApproved), v))
	})
}

// OperatorNote applies equality check predicate on the "operator_note" field. It's identical to OperatorNoteEQ.
func OperatorNote(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperatorNote), v))
	})
}

// CreatetimeUtc applies equality check predicate on the "createtime_utc" field. It's identical to CreatetimeUtcEQ.
func CreatetimeUtc(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatetimeUtc), v))
	})
}

// UpdatetimeUtc applies equality check predicate on the "updatetime_utc" field. It's identical to UpdatetimeUtcEQ.
func UpdatetimeUtc(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatetimeUtc), v))
	})
}

// IsApprovedEQ applies the EQ predicate on the "is_approved" field.
func IsApprovedEQ(v bool) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsApproved), v))
	})
}

// IsApprovedNEQ applies the NEQ predicate on the "is_approved" field.
func IsApprovedNEQ(v bool) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsApproved), v))
	})
}

// OperatorNoteEQ applies the EQ predicate on the "operator_note" field.
func OperatorNoteEQ(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperatorNote), v))
	})
}

// OperatorNoteNEQ applies the NEQ predicate on the "operator_note" field.
func OperatorNoteNEQ(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOperatorNote), v))
	})
}

// OperatorNoteIn applies the In predicate on the "operator_note" field.
func OperatorNoteIn(vs ...string) predicate.Review {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOperatorNote), v...))
	})
}

// OperatorNoteNotIn applies the NotIn predicate on the "operator_note" field.
func OperatorNoteNotIn(vs ...string) predicate.Review {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOperatorNote), v...))
	})
}

// OperatorNoteGT applies the GT predicate on the "operator_note" field.
func OperatorNoteGT(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOperatorNote), v))
	})
}

// OperatorNoteGTE applies the GTE predicate on the "operator_note" field.
func OperatorNoteGTE(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOperatorNote), v))
	})
}

// OperatorNoteLT applies the LT predicate on the "operator_note" field.
func OperatorNoteLT(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOperatorNote), v))
	})
}

// OperatorNoteLTE applies the LTE predicate on the "operator_note" field.
func OperatorNoteLTE(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOperatorNote), v))
	})
}

// OperatorNoteContains applies the Contains predicate on the "operator_note" field.
func OperatorNoteContains(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOperatorNote), v))
	})
}

// OperatorNoteHasPrefix applies the HasPrefix predicate on the "operator_note" field.
func OperatorNoteHasPrefix(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOperatorNote), v))
	})
}

// OperatorNoteHasSuffix applies the HasSuffix predicate on the "operator_note" field.
func OperatorNoteHasSuffix(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOperatorNote), v))
	})
}

// OperatorNoteEqualFold applies the EqualFold predicate on the "operator_note" field.
func OperatorNoteEqualFold(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOperatorNote), v))
	})
}

// OperatorNoteContainsFold applies the ContainsFold predicate on the "operator_note" field.
func OperatorNoteContainsFold(v string) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOperatorNote), v))
	})
}

// CreatetimeUtcEQ applies the EQ predicate on the "createtime_utc" field.
func CreatetimeUtcEQ(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatetimeUtc), v))
	})
}

// CreatetimeUtcNEQ applies the NEQ predicate on the "createtime_utc" field.
func CreatetimeUtcNEQ(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatetimeUtc), v))
	})
}

// CreatetimeUtcIn applies the In predicate on the "createtime_utc" field.
func CreatetimeUtcIn(vs ...int) predicate.Review {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatetimeUtc), v...))
	})
}

// CreatetimeUtcNotIn applies the NotIn predicate on the "createtime_utc" field.
func CreatetimeUtcNotIn(vs ...int) predicate.Review {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatetimeUtc), v...))
	})
}

// CreatetimeUtcGT applies the GT predicate on the "createtime_utc" field.
func CreatetimeUtcGT(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatetimeUtc), v))
	})
}

// CreatetimeUtcGTE applies the GTE predicate on the "createtime_utc" field.
func CreatetimeUtcGTE(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatetimeUtc), v))
	})
}

// CreatetimeUtcLT applies the LT predicate on the "createtime_utc" field.
func CreatetimeUtcLT(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatetimeUtc), v))
	})
}

// CreatetimeUtcLTE applies the LTE predicate on the "createtime_utc" field.
func CreatetimeUtcLTE(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatetimeUtc), v))
	})
}

// UpdatetimeUtcEQ applies the EQ predicate on the "updatetime_utc" field.
func UpdatetimeUtcEQ(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatetimeUtc), v))
	})
}

// UpdatetimeUtcNEQ applies the NEQ predicate on the "updatetime_utc" field.
func UpdatetimeUtcNEQ(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatetimeUtc), v))
	})
}

// UpdatetimeUtcIn applies the In predicate on the "updatetime_utc" field.
func UpdatetimeUtcIn(vs ...int) predicate.Review {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatetimeUtc), v...))
	})
}

// UpdatetimeUtcNotIn applies the NotIn predicate on the "updatetime_utc" field.
func UpdatetimeUtcNotIn(vs ...int) predicate.Review {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Review(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatetimeUtc), v...))
	})
}

// UpdatetimeUtcGT applies the GT predicate on the "updatetime_utc" field.
func UpdatetimeUtcGT(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatetimeUtc), v))
	})
}

// UpdatetimeUtcGTE applies the GTE predicate on the "updatetime_utc" field.
func UpdatetimeUtcGTE(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatetimeUtc), v))
	})
}

// UpdatetimeUtcLT applies the LT predicate on the "updatetime_utc" field.
func UpdatetimeUtcLT(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatetimeUtc), v))
	})
}

// UpdatetimeUtcLTE applies the LTE predicate on the "updatetime_utc" field.
func UpdatetimeUtcLTE(v int) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatetimeUtc), v))
	})
}

// HasTransaction applies the HasEdge predicate on the "transaction" edge.
func HasTransaction() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TransactionTable, TransactionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionWith applies the HasEdge predicate on the "transaction" edge with a given conditions (other predicates).
func HasTransactionWith(preds ...predicate.Transaction) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TransactionTable, TransactionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
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
func Not(p predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		p(s.Not())
	})
}
