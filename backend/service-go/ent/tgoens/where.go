// Code generated by ent, DO NOT EDIT.

package tgoens

import (
	"time"
	"toucan-leaderboard/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// WalletPub applies equality check predicate on the "wallet_pub" field. It's identical to WalletPubEQ.
func WalletPub(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWalletPub), v))
	})
}

// Ens applies equality check predicate on the "ens" field. It's identical to EnsEQ.
func Ens(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEns), v))
	})
}

// Mtime applies equality check predicate on the "mtime" field. It's identical to MtimeEQ.
func Mtime(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMtime), v))
	})
}

// Ctime applies equality check predicate on the "ctime" field. It's identical to CtimeEQ.
func Ctime(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCtime), v))
	})
}

// WalletPubEQ applies the EQ predicate on the "wallet_pub" field.
func WalletPubEQ(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWalletPub), v))
	})
}

// WalletPubNEQ applies the NEQ predicate on the "wallet_pub" field.
func WalletPubNEQ(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWalletPub), v))
	})
}

// WalletPubIn applies the In predicate on the "wallet_pub" field.
func WalletPubIn(vs ...string) predicate.TGoEns {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldWalletPub), v...))
	})
}

// WalletPubNotIn applies the NotIn predicate on the "wallet_pub" field.
func WalletPubNotIn(vs ...string) predicate.TGoEns {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldWalletPub), v...))
	})
}

// WalletPubGT applies the GT predicate on the "wallet_pub" field.
func WalletPubGT(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWalletPub), v))
	})
}

// WalletPubGTE applies the GTE predicate on the "wallet_pub" field.
func WalletPubGTE(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWalletPub), v))
	})
}

// WalletPubLT applies the LT predicate on the "wallet_pub" field.
func WalletPubLT(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWalletPub), v))
	})
}

// WalletPubLTE applies the LTE predicate on the "wallet_pub" field.
func WalletPubLTE(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWalletPub), v))
	})
}

// WalletPubContains applies the Contains predicate on the "wallet_pub" field.
func WalletPubContains(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWalletPub), v))
	})
}

// WalletPubHasPrefix applies the HasPrefix predicate on the "wallet_pub" field.
func WalletPubHasPrefix(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWalletPub), v))
	})
}

// WalletPubHasSuffix applies the HasSuffix predicate on the "wallet_pub" field.
func WalletPubHasSuffix(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWalletPub), v))
	})
}

// WalletPubEqualFold applies the EqualFold predicate on the "wallet_pub" field.
func WalletPubEqualFold(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWalletPub), v))
	})
}

// WalletPubContainsFold applies the ContainsFold predicate on the "wallet_pub" field.
func WalletPubContainsFold(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWalletPub), v))
	})
}

// EnsEQ applies the EQ predicate on the "ens" field.
func EnsEQ(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEns), v))
	})
}

// EnsNEQ applies the NEQ predicate on the "ens" field.
func EnsNEQ(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEns), v))
	})
}

// EnsIn applies the In predicate on the "ens" field.
func EnsIn(vs ...string) predicate.TGoEns {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEns), v...))
	})
}

// EnsNotIn applies the NotIn predicate on the "ens" field.
func EnsNotIn(vs ...string) predicate.TGoEns {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEns), v...))
	})
}

// EnsGT applies the GT predicate on the "ens" field.
func EnsGT(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEns), v))
	})
}

// EnsGTE applies the GTE predicate on the "ens" field.
func EnsGTE(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEns), v))
	})
}

// EnsLT applies the LT predicate on the "ens" field.
func EnsLT(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEns), v))
	})
}

// EnsLTE applies the LTE predicate on the "ens" field.
func EnsLTE(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEns), v))
	})
}

// EnsContains applies the Contains predicate on the "ens" field.
func EnsContains(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEns), v))
	})
}

// EnsHasPrefix applies the HasPrefix predicate on the "ens" field.
func EnsHasPrefix(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEns), v))
	})
}

// EnsHasSuffix applies the HasSuffix predicate on the "ens" field.
func EnsHasSuffix(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEns), v))
	})
}

// EnsEqualFold applies the EqualFold predicate on the "ens" field.
func EnsEqualFold(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEns), v))
	})
}

// EnsContainsFold applies the ContainsFold predicate on the "ens" field.
func EnsContainsFold(v string) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEns), v))
	})
}

// MtimeEQ applies the EQ predicate on the "mtime" field.
func MtimeEQ(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMtime), v))
	})
}

// MtimeNEQ applies the NEQ predicate on the "mtime" field.
func MtimeNEQ(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMtime), v))
	})
}

// MtimeIn applies the In predicate on the "mtime" field.
func MtimeIn(vs ...time.Time) predicate.TGoEns {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMtime), v...))
	})
}

// MtimeNotIn applies the NotIn predicate on the "mtime" field.
func MtimeNotIn(vs ...time.Time) predicate.TGoEns {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMtime), v...))
	})
}

// MtimeGT applies the GT predicate on the "mtime" field.
func MtimeGT(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMtime), v))
	})
}

// MtimeGTE applies the GTE predicate on the "mtime" field.
func MtimeGTE(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMtime), v))
	})
}

// MtimeLT applies the LT predicate on the "mtime" field.
func MtimeLT(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMtime), v))
	})
}

// MtimeLTE applies the LTE predicate on the "mtime" field.
func MtimeLTE(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMtime), v))
	})
}

// CtimeEQ applies the EQ predicate on the "ctime" field.
func CtimeEQ(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCtime), v))
	})
}

// CtimeNEQ applies the NEQ predicate on the "ctime" field.
func CtimeNEQ(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCtime), v))
	})
}

// CtimeIn applies the In predicate on the "ctime" field.
func CtimeIn(vs ...time.Time) predicate.TGoEns {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCtime), v...))
	})
}

// CtimeNotIn applies the NotIn predicate on the "ctime" field.
func CtimeNotIn(vs ...time.Time) predicate.TGoEns {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCtime), v...))
	})
}

// CtimeGT applies the GT predicate on the "ctime" field.
func CtimeGT(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCtime), v))
	})
}

// CtimeGTE applies the GTE predicate on the "ctime" field.
func CtimeGTE(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCtime), v))
	})
}

// CtimeLT applies the LT predicate on the "ctime" field.
func CtimeLT(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCtime), v))
	})
}

// CtimeLTE applies the LTE predicate on the "ctime" field.
func CtimeLTE(v time.Time) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCtime), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TGoEns) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TGoEns) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
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
func Not(p predicate.TGoEns) predicate.TGoEns {
	return predicate.TGoEns(func(s *sql.Selector) {
		p(s.Not())
	})
}