// Code generated by entc, DO NOT EDIT.

package accounts

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/pepeunlimited/accounts/internal/app/app1/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Accounts {
	return predicate.Accounts(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
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
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
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
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// Balance applies equality check predicate on the "balance" field. It's identical to BalanceEQ.
func Balance(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBalance), v))
	},
	)
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v uint8) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	},
	)
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	},
	)
}

// IsWithdrawable applies equality check predicate on the "is_withdrawable" field. It's identical to IsWithdrawableEQ.
func IsWithdrawable(v bool) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsWithdrawable), v))
	},
	)
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	},
	)
}

// BalanceEQ applies the EQ predicate on the "balance" field.
func BalanceEQ(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBalance), v))
	},
	)
}

// BalanceNEQ applies the NEQ predicate on the "balance" field.
func BalanceNEQ(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBalance), v))
	},
	)
}

// BalanceIn applies the In predicate on the "balance" field.
func BalanceIn(vs ...int64) predicate.Accounts {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Accounts(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBalance), v...))
	},
	)
}

// BalanceNotIn applies the NotIn predicate on the "balance" field.
func BalanceNotIn(vs ...int64) predicate.Accounts {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Accounts(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBalance), v...))
	},
	)
}

// BalanceGT applies the GT predicate on the "balance" field.
func BalanceGT(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBalance), v))
	},
	)
}

// BalanceGTE applies the GTE predicate on the "balance" field.
func BalanceGTE(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBalance), v))
	},
	)
}

// BalanceLT applies the LT predicate on the "balance" field.
func BalanceLT(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBalance), v))
	},
	)
}

// BalanceLTE applies the LTE predicate on the "balance" field.
func BalanceLTE(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBalance), v))
	},
	)
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v uint8) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	},
	)
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v uint8) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVersion), v))
	},
	)
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...uint8) predicate.Accounts {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Accounts(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVersion), v...))
	},
	)
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...uint8) predicate.Accounts {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Accounts(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVersion), v...))
	},
	)
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v uint8) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVersion), v))
	},
	)
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v uint8) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVersion), v))
	},
	)
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v uint8) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVersion), v))
	},
	)
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v uint8) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVersion), v))
	},
	)
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	},
	)
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	},
	)
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Accounts {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Accounts(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	},
	)
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Accounts {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Accounts(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	},
	)
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	},
	)
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	},
	)
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	},
	)
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	},
	)
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	},
	)
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	},
	)
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	},
	)
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	},
	)
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	},
	)
}

// IsWithdrawableEQ applies the EQ predicate on the "is_withdrawable" field.
func IsWithdrawableEQ(v bool) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsWithdrawable), v))
	},
	)
}

// IsWithdrawableNEQ applies the NEQ predicate on the "is_withdrawable" field.
func IsWithdrawableNEQ(v bool) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsWithdrawable), v))
	},
	)
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	},
	)
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	},
	)
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Accounts {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Accounts(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	},
	)
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Accounts {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Accounts(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	},
	)
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	},
	)
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	},
	)
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	},
	)
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	},
	)
}

// HasTxs applies the HasEdge predicate on the "txs" edge.
func HasTxs() predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TxsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TxsTable, TxsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasTxsWith applies the HasEdge predicate on the "txs" edge with a given conditions (other predicates).
func HasTxsWith(preds ...predicate.Txs) predicate.Accounts {
	return predicate.Accounts(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TxsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TxsTable, TxsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Accounts) predicate.Accounts {
	return predicate.Accounts(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Accounts) predicate.Accounts {
	return predicate.Accounts(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Accounts) predicate.Accounts {
	return predicate.Accounts(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
