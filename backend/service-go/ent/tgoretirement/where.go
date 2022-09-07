// Code generated by ent, DO NOT EDIT.

package tgoretirement

import (
	"time"
	"toucan-leaderboard/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreationTx applies equality check predicate on the "creation_tx" field. It's identical to CreationTxEQ.
func CreationTx(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreationTx), v))
	})
}

// CreatorAddress applies equality check predicate on the "creator_address" field. It's identical to CreatorAddressEQ.
func CreatorAddress(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatorAddress), v))
	})
}

// BeneficiaryAddress applies equality check predicate on the "beneficiary_address" field. It's identical to BeneficiaryAddressEQ.
func BeneficiaryAddress(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBeneficiaryAddress), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// TokenAddress applies equality check predicate on the "token_address" field. It's identical to TokenAddressEQ.
func TokenAddress(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenAddress), v))
	})
}

// TokenName applies equality check predicate on the "token_name" field. It's identical to TokenNameEQ.
func TokenName(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenName), v))
	})
}

// TokenType applies equality check predicate on the "token_type" field. It's identical to TokenTypeEQ.
func TokenType(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenType), v))
	})
}

// RetirementTime applies equality check predicate on the "retirement_time" field. It's identical to RetirementTimeEQ.
func RetirementTime(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRetirementTime), v))
	})
}

// Mtime applies equality check predicate on the "mtime" field. It's identical to MtimeEQ.
func Mtime(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMtime), v))
	})
}

// Ctime applies equality check predicate on the "ctime" field. It's identical to CtimeEQ.
func Ctime(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCtime), v))
	})
}

// CreationTxEQ applies the EQ predicate on the "creation_tx" field.
func CreationTxEQ(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreationTx), v))
	})
}

// CreationTxNEQ applies the NEQ predicate on the "creation_tx" field.
func CreationTxNEQ(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreationTx), v))
	})
}

// CreationTxIn applies the In predicate on the "creation_tx" field.
func CreationTxIn(vs ...string) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreationTx), v...))
	})
}

// CreationTxNotIn applies the NotIn predicate on the "creation_tx" field.
func CreationTxNotIn(vs ...string) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreationTx), v...))
	})
}

// CreationTxGT applies the GT predicate on the "creation_tx" field.
func CreationTxGT(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreationTx), v))
	})
}

// CreationTxGTE applies the GTE predicate on the "creation_tx" field.
func CreationTxGTE(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreationTx), v))
	})
}

// CreationTxLT applies the LT predicate on the "creation_tx" field.
func CreationTxLT(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreationTx), v))
	})
}

// CreationTxLTE applies the LTE predicate on the "creation_tx" field.
func CreationTxLTE(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreationTx), v))
	})
}

// CreationTxContains applies the Contains predicate on the "creation_tx" field.
func CreationTxContains(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreationTx), v))
	})
}

// CreationTxHasPrefix applies the HasPrefix predicate on the "creation_tx" field.
func CreationTxHasPrefix(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreationTx), v))
	})
}

// CreationTxHasSuffix applies the HasSuffix predicate on the "creation_tx" field.
func CreationTxHasSuffix(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreationTx), v))
	})
}

// CreationTxEqualFold applies the EqualFold predicate on the "creation_tx" field.
func CreationTxEqualFold(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreationTx), v))
	})
}

// CreationTxContainsFold applies the ContainsFold predicate on the "creation_tx" field.
func CreationTxContainsFold(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreationTx), v))
	})
}

// CreatorAddressEQ applies the EQ predicate on the "creator_address" field.
func CreatorAddressEQ(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatorAddress), v))
	})
}

// CreatorAddressNEQ applies the NEQ predicate on the "creator_address" field.
func CreatorAddressNEQ(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatorAddress), v))
	})
}

// CreatorAddressIn applies the In predicate on the "creator_address" field.
func CreatorAddressIn(vs ...string) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatorAddress), v...))
	})
}

// CreatorAddressNotIn applies the NotIn predicate on the "creator_address" field.
func CreatorAddressNotIn(vs ...string) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatorAddress), v...))
	})
}

// CreatorAddressGT applies the GT predicate on the "creator_address" field.
func CreatorAddressGT(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatorAddress), v))
	})
}

// CreatorAddressGTE applies the GTE predicate on the "creator_address" field.
func CreatorAddressGTE(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatorAddress), v))
	})
}

// CreatorAddressLT applies the LT predicate on the "creator_address" field.
func CreatorAddressLT(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatorAddress), v))
	})
}

// CreatorAddressLTE applies the LTE predicate on the "creator_address" field.
func CreatorAddressLTE(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatorAddress), v))
	})
}

// CreatorAddressContains applies the Contains predicate on the "creator_address" field.
func CreatorAddressContains(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreatorAddress), v))
	})
}

// CreatorAddressHasPrefix applies the HasPrefix predicate on the "creator_address" field.
func CreatorAddressHasPrefix(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreatorAddress), v))
	})
}

// CreatorAddressHasSuffix applies the HasSuffix predicate on the "creator_address" field.
func CreatorAddressHasSuffix(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreatorAddress), v))
	})
}

// CreatorAddressEqualFold applies the EqualFold predicate on the "creator_address" field.
func CreatorAddressEqualFold(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreatorAddress), v))
	})
}

// CreatorAddressContainsFold applies the ContainsFold predicate on the "creator_address" field.
func CreatorAddressContainsFold(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreatorAddress), v))
	})
}

// BeneficiaryAddressEQ applies the EQ predicate on the "beneficiary_address" field.
func BeneficiaryAddressEQ(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBeneficiaryAddress), v))
	})
}

// BeneficiaryAddressNEQ applies the NEQ predicate on the "beneficiary_address" field.
func BeneficiaryAddressNEQ(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBeneficiaryAddress), v))
	})
}

// BeneficiaryAddressIn applies the In predicate on the "beneficiary_address" field.
func BeneficiaryAddressIn(vs ...string) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBeneficiaryAddress), v...))
	})
}

// BeneficiaryAddressNotIn applies the NotIn predicate on the "beneficiary_address" field.
func BeneficiaryAddressNotIn(vs ...string) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBeneficiaryAddress), v...))
	})
}

// BeneficiaryAddressGT applies the GT predicate on the "beneficiary_address" field.
func BeneficiaryAddressGT(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBeneficiaryAddress), v))
	})
}

// BeneficiaryAddressGTE applies the GTE predicate on the "beneficiary_address" field.
func BeneficiaryAddressGTE(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBeneficiaryAddress), v))
	})
}

// BeneficiaryAddressLT applies the LT predicate on the "beneficiary_address" field.
func BeneficiaryAddressLT(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBeneficiaryAddress), v))
	})
}

// BeneficiaryAddressLTE applies the LTE predicate on the "beneficiary_address" field.
func BeneficiaryAddressLTE(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBeneficiaryAddress), v))
	})
}

// BeneficiaryAddressContains applies the Contains predicate on the "beneficiary_address" field.
func BeneficiaryAddressContains(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBeneficiaryAddress), v))
	})
}

// BeneficiaryAddressHasPrefix applies the HasPrefix predicate on the "beneficiary_address" field.
func BeneficiaryAddressHasPrefix(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBeneficiaryAddress), v))
	})
}

// BeneficiaryAddressHasSuffix applies the HasSuffix predicate on the "beneficiary_address" field.
func BeneficiaryAddressHasSuffix(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBeneficiaryAddress), v))
	})
}

// BeneficiaryAddressEqualFold applies the EqualFold predicate on the "beneficiary_address" field.
func BeneficiaryAddressEqualFold(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBeneficiaryAddress), v))
	})
}

// BeneficiaryAddressContainsFold applies the ContainsFold predicate on the "beneficiary_address" field.
func BeneficiaryAddressContainsFold(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBeneficiaryAddress), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// TokenAddressEQ applies the EQ predicate on the "token_address" field.
func TokenAddressEQ(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenAddress), v))
	})
}

// TokenAddressNEQ applies the NEQ predicate on the "token_address" field.
func TokenAddressNEQ(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTokenAddress), v))
	})
}

// TokenAddressIn applies the In predicate on the "token_address" field.
func TokenAddressIn(vs ...string) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTokenAddress), v...))
	})
}

// TokenAddressNotIn applies the NotIn predicate on the "token_address" field.
func TokenAddressNotIn(vs ...string) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTokenAddress), v...))
	})
}

// TokenAddressGT applies the GT predicate on the "token_address" field.
func TokenAddressGT(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTokenAddress), v))
	})
}

// TokenAddressGTE applies the GTE predicate on the "token_address" field.
func TokenAddressGTE(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTokenAddress), v))
	})
}

// TokenAddressLT applies the LT predicate on the "token_address" field.
func TokenAddressLT(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTokenAddress), v))
	})
}

// TokenAddressLTE applies the LTE predicate on the "token_address" field.
func TokenAddressLTE(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTokenAddress), v))
	})
}

// TokenAddressContains applies the Contains predicate on the "token_address" field.
func TokenAddressContains(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTokenAddress), v))
	})
}

// TokenAddressHasPrefix applies the HasPrefix predicate on the "token_address" field.
func TokenAddressHasPrefix(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTokenAddress), v))
	})
}

// TokenAddressHasSuffix applies the HasSuffix predicate on the "token_address" field.
func TokenAddressHasSuffix(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTokenAddress), v))
	})
}

// TokenAddressEqualFold applies the EqualFold predicate on the "token_address" field.
func TokenAddressEqualFold(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTokenAddress), v))
	})
}

// TokenAddressContainsFold applies the ContainsFold predicate on the "token_address" field.
func TokenAddressContainsFold(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTokenAddress), v))
	})
}

// TokenNameEQ applies the EQ predicate on the "token_name" field.
func TokenNameEQ(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenName), v))
	})
}

// TokenNameNEQ applies the NEQ predicate on the "token_name" field.
func TokenNameNEQ(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTokenName), v))
	})
}

// TokenNameIn applies the In predicate on the "token_name" field.
func TokenNameIn(vs ...string) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTokenName), v...))
	})
}

// TokenNameNotIn applies the NotIn predicate on the "token_name" field.
func TokenNameNotIn(vs ...string) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTokenName), v...))
	})
}

// TokenNameGT applies the GT predicate on the "token_name" field.
func TokenNameGT(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTokenName), v))
	})
}

// TokenNameGTE applies the GTE predicate on the "token_name" field.
func TokenNameGTE(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTokenName), v))
	})
}

// TokenNameLT applies the LT predicate on the "token_name" field.
func TokenNameLT(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTokenName), v))
	})
}

// TokenNameLTE applies the LTE predicate on the "token_name" field.
func TokenNameLTE(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTokenName), v))
	})
}

// TokenNameContains applies the Contains predicate on the "token_name" field.
func TokenNameContains(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTokenName), v))
	})
}

// TokenNameHasPrefix applies the HasPrefix predicate on the "token_name" field.
func TokenNameHasPrefix(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTokenName), v))
	})
}

// TokenNameHasSuffix applies the HasSuffix predicate on the "token_name" field.
func TokenNameHasSuffix(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTokenName), v))
	})
}

// TokenNameEqualFold applies the EqualFold predicate on the "token_name" field.
func TokenNameEqualFold(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTokenName), v))
	})
}

// TokenNameContainsFold applies the ContainsFold predicate on the "token_name" field.
func TokenNameContainsFold(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTokenName), v))
	})
}

// TokenTypeEQ applies the EQ predicate on the "token_type" field.
func TokenTypeEQ(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenType), v))
	})
}

// TokenTypeNEQ applies the NEQ predicate on the "token_type" field.
func TokenTypeNEQ(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTokenType), v))
	})
}

// TokenTypeIn applies the In predicate on the "token_type" field.
func TokenTypeIn(vs ...string) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTokenType), v...))
	})
}

// TokenTypeNotIn applies the NotIn predicate on the "token_type" field.
func TokenTypeNotIn(vs ...string) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTokenType), v...))
	})
}

// TokenTypeGT applies the GT predicate on the "token_type" field.
func TokenTypeGT(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTokenType), v))
	})
}

// TokenTypeGTE applies the GTE predicate on the "token_type" field.
func TokenTypeGTE(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTokenType), v))
	})
}

// TokenTypeLT applies the LT predicate on the "token_type" field.
func TokenTypeLT(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTokenType), v))
	})
}

// TokenTypeLTE applies the LTE predicate on the "token_type" field.
func TokenTypeLTE(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTokenType), v))
	})
}

// TokenTypeContains applies the Contains predicate on the "token_type" field.
func TokenTypeContains(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTokenType), v))
	})
}

// TokenTypeHasPrefix applies the HasPrefix predicate on the "token_type" field.
func TokenTypeHasPrefix(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTokenType), v))
	})
}

// TokenTypeHasSuffix applies the HasSuffix predicate on the "token_type" field.
func TokenTypeHasSuffix(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTokenType), v))
	})
}

// TokenTypeEqualFold applies the EqualFold predicate on the "token_type" field.
func TokenTypeEqualFold(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTokenType), v))
	})
}

// TokenTypeContainsFold applies the ContainsFold predicate on the "token_type" field.
func TokenTypeContainsFold(v string) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTokenType), v))
	})
}

// RetirementTimeEQ applies the EQ predicate on the "retirement_time" field.
func RetirementTimeEQ(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRetirementTime), v))
	})
}

// RetirementTimeNEQ applies the NEQ predicate on the "retirement_time" field.
func RetirementTimeNEQ(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRetirementTime), v))
	})
}

// RetirementTimeIn applies the In predicate on the "retirement_time" field.
func RetirementTimeIn(vs ...time.Time) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRetirementTime), v...))
	})
}

// RetirementTimeNotIn applies the NotIn predicate on the "retirement_time" field.
func RetirementTimeNotIn(vs ...time.Time) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRetirementTime), v...))
	})
}

// RetirementTimeGT applies the GT predicate on the "retirement_time" field.
func RetirementTimeGT(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRetirementTime), v))
	})
}

// RetirementTimeGTE applies the GTE predicate on the "retirement_time" field.
func RetirementTimeGTE(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRetirementTime), v))
	})
}

// RetirementTimeLT applies the LT predicate on the "retirement_time" field.
func RetirementTimeLT(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRetirementTime), v))
	})
}

// RetirementTimeLTE applies the LTE predicate on the "retirement_time" field.
func RetirementTimeLTE(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRetirementTime), v))
	})
}

// MtimeEQ applies the EQ predicate on the "mtime" field.
func MtimeEQ(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMtime), v))
	})
}

// MtimeNEQ applies the NEQ predicate on the "mtime" field.
func MtimeNEQ(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMtime), v))
	})
}

// MtimeIn applies the In predicate on the "mtime" field.
func MtimeIn(vs ...time.Time) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMtime), v...))
	})
}

// MtimeNotIn applies the NotIn predicate on the "mtime" field.
func MtimeNotIn(vs ...time.Time) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMtime), v...))
	})
}

// MtimeGT applies the GT predicate on the "mtime" field.
func MtimeGT(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMtime), v))
	})
}

// MtimeGTE applies the GTE predicate on the "mtime" field.
func MtimeGTE(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMtime), v))
	})
}

// MtimeLT applies the LT predicate on the "mtime" field.
func MtimeLT(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMtime), v))
	})
}

// MtimeLTE applies the LTE predicate on the "mtime" field.
func MtimeLTE(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMtime), v))
	})
}

// CtimeEQ applies the EQ predicate on the "ctime" field.
func CtimeEQ(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCtime), v))
	})
}

// CtimeNEQ applies the NEQ predicate on the "ctime" field.
func CtimeNEQ(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCtime), v))
	})
}

// CtimeIn applies the In predicate on the "ctime" field.
func CtimeIn(vs ...time.Time) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCtime), v...))
	})
}

// CtimeNotIn applies the NotIn predicate on the "ctime" field.
func CtimeNotIn(vs ...time.Time) predicate.TGoRetirement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCtime), v...))
	})
}

// CtimeGT applies the GT predicate on the "ctime" field.
func CtimeGT(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCtime), v))
	})
}

// CtimeGTE applies the GTE predicate on the "ctime" field.
func CtimeGTE(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCtime), v))
	})
}

// CtimeLT applies the LT predicate on the "ctime" field.
func CtimeLT(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCtime), v))
	})
}

// CtimeLTE applies the LTE predicate on the "ctime" field.
func CtimeLTE(v time.Time) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCtime), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TGoRetirement) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TGoRetirement) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
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
func Not(p predicate.TGoRetirement) predicate.TGoRetirement {
	return predicate.TGoRetirement(func(s *sql.Selector) {
		p(s.Not())
	})
}
