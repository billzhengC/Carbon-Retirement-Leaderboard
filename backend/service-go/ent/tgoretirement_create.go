// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"toucan-leaderboard/ent/tgoretirement"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TGoRetirementCreate is the builder for creating a TGoRetirement entity.
type TGoRetirementCreate struct {
	config
	mutation *TGoRetirementMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreationTx sets the "creation_tx" field.
func (trc *TGoRetirementCreate) SetCreationTx(s string) *TGoRetirementCreate {
	trc.mutation.SetCreationTx(s)
	return trc
}

// SetNillableCreationTx sets the "creation_tx" field if the given value is not nil.
func (trc *TGoRetirementCreate) SetNillableCreationTx(s *string) *TGoRetirementCreate {
	if s != nil {
		trc.SetCreationTx(*s)
	}
	return trc
}

// SetCreatorAddress sets the "creator_address" field.
func (trc *TGoRetirementCreate) SetCreatorAddress(s string) *TGoRetirementCreate {
	trc.mutation.SetCreatorAddress(s)
	return trc
}

// SetNillableCreatorAddress sets the "creator_address" field if the given value is not nil.
func (trc *TGoRetirementCreate) SetNillableCreatorAddress(s *string) *TGoRetirementCreate {
	if s != nil {
		trc.SetCreatorAddress(*s)
	}
	return trc
}

// SetBeneficiaryAddress sets the "beneficiary_address" field.
func (trc *TGoRetirementCreate) SetBeneficiaryAddress(s string) *TGoRetirementCreate {
	trc.mutation.SetBeneficiaryAddress(s)
	return trc
}

// SetNillableBeneficiaryAddress sets the "beneficiary_address" field if the given value is not nil.
func (trc *TGoRetirementCreate) SetNillableBeneficiaryAddress(s *string) *TGoRetirementCreate {
	if s != nil {
		trc.SetBeneficiaryAddress(*s)
	}
	return trc
}

// SetAmount sets the "amount" field.
func (trc *TGoRetirementCreate) SetAmount(f float64) *TGoRetirementCreate {
	trc.mutation.SetAmount(f)
	return trc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (trc *TGoRetirementCreate) SetNillableAmount(f *float64) *TGoRetirementCreate {
	if f != nil {
		trc.SetAmount(*f)
	}
	return trc
}

// SetTokenAddress sets the "token_address" field.
func (trc *TGoRetirementCreate) SetTokenAddress(s string) *TGoRetirementCreate {
	trc.mutation.SetTokenAddress(s)
	return trc
}

// SetNillableTokenAddress sets the "token_address" field if the given value is not nil.
func (trc *TGoRetirementCreate) SetNillableTokenAddress(s *string) *TGoRetirementCreate {
	if s != nil {
		trc.SetTokenAddress(*s)
	}
	return trc
}

// SetTokenName sets the "token_name" field.
func (trc *TGoRetirementCreate) SetTokenName(s string) *TGoRetirementCreate {
	trc.mutation.SetTokenName(s)
	return trc
}

// SetNillableTokenName sets the "token_name" field if the given value is not nil.
func (trc *TGoRetirementCreate) SetNillableTokenName(s *string) *TGoRetirementCreate {
	if s != nil {
		trc.SetTokenName(*s)
	}
	return trc
}

// SetTokenType sets the "token_type" field.
func (trc *TGoRetirementCreate) SetTokenType(s string) *TGoRetirementCreate {
	trc.mutation.SetTokenType(s)
	return trc
}

// SetNillableTokenType sets the "token_type" field if the given value is not nil.
func (trc *TGoRetirementCreate) SetNillableTokenType(s *string) *TGoRetirementCreate {
	if s != nil {
		trc.SetTokenType(*s)
	}
	return trc
}

// SetRetirementTime sets the "retirement_time" field.
func (trc *TGoRetirementCreate) SetRetirementTime(t time.Time) *TGoRetirementCreate {
	trc.mutation.SetRetirementTime(t)
	return trc
}

// SetMtime sets the "mtime" field.
func (trc *TGoRetirementCreate) SetMtime(t time.Time) *TGoRetirementCreate {
	trc.mutation.SetMtime(t)
	return trc
}

// SetNillableMtime sets the "mtime" field if the given value is not nil.
func (trc *TGoRetirementCreate) SetNillableMtime(t *time.Time) *TGoRetirementCreate {
	if t != nil {
		trc.SetMtime(*t)
	}
	return trc
}

// SetCtime sets the "ctime" field.
func (trc *TGoRetirementCreate) SetCtime(t time.Time) *TGoRetirementCreate {
	trc.mutation.SetCtime(t)
	return trc
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (trc *TGoRetirementCreate) SetNillableCtime(t *time.Time) *TGoRetirementCreate {
	if t != nil {
		trc.SetCtime(*t)
	}
	return trc
}

// SetID sets the "id" field.
func (trc *TGoRetirementCreate) SetID(u uint64) *TGoRetirementCreate {
	trc.mutation.SetID(u)
	return trc
}

// Mutation returns the TGoRetirementMutation object of the builder.
func (trc *TGoRetirementCreate) Mutation() *TGoRetirementMutation {
	return trc.mutation
}

// Save creates the TGoRetirement in the database.
func (trc *TGoRetirementCreate) Save(ctx context.Context) (*TGoRetirement, error) {
	var (
		err  error
		node *TGoRetirement
	)
	trc.defaults()
	if len(trc.hooks) == 0 {
		if err = trc.check(); err != nil {
			return nil, err
		}
		node, err = trc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TGoRetirementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = trc.check(); err != nil {
				return nil, err
			}
			trc.mutation = mutation
			if node, err = trc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(trc.hooks) - 1; i >= 0; i-- {
			if trc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = trc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, trc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TGoRetirement)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TGoRetirementMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (trc *TGoRetirementCreate) SaveX(ctx context.Context) *TGoRetirement {
	v, err := trc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (trc *TGoRetirementCreate) Exec(ctx context.Context) error {
	_, err := trc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trc *TGoRetirementCreate) ExecX(ctx context.Context) {
	if err := trc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (trc *TGoRetirementCreate) defaults() {
	if _, ok := trc.mutation.CreationTx(); !ok {
		v := tgoretirement.DefaultCreationTx
		trc.mutation.SetCreationTx(v)
	}
	if _, ok := trc.mutation.CreatorAddress(); !ok {
		v := tgoretirement.DefaultCreatorAddress
		trc.mutation.SetCreatorAddress(v)
	}
	if _, ok := trc.mutation.BeneficiaryAddress(); !ok {
		v := tgoretirement.DefaultBeneficiaryAddress
		trc.mutation.SetBeneficiaryAddress(v)
	}
	if _, ok := trc.mutation.Amount(); !ok {
		v := tgoretirement.DefaultAmount
		trc.mutation.SetAmount(v)
	}
	if _, ok := trc.mutation.TokenAddress(); !ok {
		v := tgoretirement.DefaultTokenAddress
		trc.mutation.SetTokenAddress(v)
	}
	if _, ok := trc.mutation.TokenName(); !ok {
		v := tgoretirement.DefaultTokenName
		trc.mutation.SetTokenName(v)
	}
	if _, ok := trc.mutation.TokenType(); !ok {
		v := tgoretirement.DefaultTokenType
		trc.mutation.SetTokenType(v)
	}
	if _, ok := trc.mutation.Mtime(); !ok {
		v := tgoretirement.DefaultMtime
		trc.mutation.SetMtime(v)
	}
	if _, ok := trc.mutation.Ctime(); !ok {
		v := tgoretirement.DefaultCtime
		trc.mutation.SetCtime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (trc *TGoRetirementCreate) check() error {
	if _, ok := trc.mutation.CreationTx(); !ok {
		return &ValidationError{Name: "creation_tx", err: errors.New(`ent: missing required field "TGoRetirement.creation_tx"`)}
	}
	if _, ok := trc.mutation.CreatorAddress(); !ok {
		return &ValidationError{Name: "creator_address", err: errors.New(`ent: missing required field "TGoRetirement.creator_address"`)}
	}
	if _, ok := trc.mutation.BeneficiaryAddress(); !ok {
		return &ValidationError{Name: "beneficiary_address", err: errors.New(`ent: missing required field "TGoRetirement.beneficiary_address"`)}
	}
	if _, ok := trc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "TGoRetirement.amount"`)}
	}
	if _, ok := trc.mutation.TokenAddress(); !ok {
		return &ValidationError{Name: "token_address", err: errors.New(`ent: missing required field "TGoRetirement.token_address"`)}
	}
	if _, ok := trc.mutation.TokenName(); !ok {
		return &ValidationError{Name: "token_name", err: errors.New(`ent: missing required field "TGoRetirement.token_name"`)}
	}
	if _, ok := trc.mutation.TokenType(); !ok {
		return &ValidationError{Name: "token_type", err: errors.New(`ent: missing required field "TGoRetirement.token_type"`)}
	}
	if _, ok := trc.mutation.RetirementTime(); !ok {
		return &ValidationError{Name: "retirement_time", err: errors.New(`ent: missing required field "TGoRetirement.retirement_time"`)}
	}
	if _, ok := trc.mutation.Mtime(); !ok {
		return &ValidationError{Name: "mtime", err: errors.New(`ent: missing required field "TGoRetirement.mtime"`)}
	}
	if _, ok := trc.mutation.Ctime(); !ok {
		return &ValidationError{Name: "ctime", err: errors.New(`ent: missing required field "TGoRetirement.ctime"`)}
	}
	return nil
}

func (trc *TGoRetirementCreate) sqlSave(ctx context.Context) (*TGoRetirement, error) {
	_node, _spec := trc.createSpec()
	if err := sqlgraph.CreateNode(ctx, trc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (trc *TGoRetirementCreate) createSpec() (*TGoRetirement, *sqlgraph.CreateSpec) {
	var (
		_node = &TGoRetirement{config: trc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tgoretirement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: tgoretirement.FieldID,
			},
		}
	)
	_spec.OnConflict = trc.conflict
	if id, ok := trc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := trc.mutation.CreationTx(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgoretirement.FieldCreationTx,
		})
		_node.CreationTx = value
	}
	if value, ok := trc.mutation.CreatorAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgoretirement.FieldCreatorAddress,
		})
		_node.CreatorAddress = value
	}
	if value, ok := trc.mutation.BeneficiaryAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgoretirement.FieldBeneficiaryAddress,
		})
		_node.BeneficiaryAddress = value
	}
	if value, ok := trc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: tgoretirement.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := trc.mutation.TokenAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgoretirement.FieldTokenAddress,
		})
		_node.TokenAddress = value
	}
	if value, ok := trc.mutation.TokenName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgoretirement.FieldTokenName,
		})
		_node.TokenName = value
	}
	if value, ok := trc.mutation.TokenType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgoretirement.FieldTokenType,
		})
		_node.TokenType = value
	}
	if value, ok := trc.mutation.RetirementTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgoretirement.FieldRetirementTime,
		})
		_node.RetirementTime = value
	}
	if value, ok := trc.mutation.Mtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgoretirement.FieldMtime,
		})
		_node.Mtime = value
	}
	if value, ok := trc.mutation.Ctime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgoretirement.FieldCtime,
		})
		_node.Ctime = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TGoRetirement.Create().
//		SetCreationTx(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TGoRetirementUpsert) {
//			SetCreationTx(v+v).
//		}).
//		Exec(ctx)
//
func (trc *TGoRetirementCreate) OnConflict(opts ...sql.ConflictOption) *TGoRetirementUpsertOne {
	trc.conflict = opts
	return &TGoRetirementUpsertOne{
		create: trc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TGoRetirement.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (trc *TGoRetirementCreate) OnConflictColumns(columns ...string) *TGoRetirementUpsertOne {
	trc.conflict = append(trc.conflict, sql.ConflictColumns(columns...))
	return &TGoRetirementUpsertOne{
		create: trc,
	}
}

type (
	// TGoRetirementUpsertOne is the builder for "upsert"-ing
	//  one TGoRetirement node.
	TGoRetirementUpsertOne struct {
		create *TGoRetirementCreate
	}

	// TGoRetirementUpsert is the "OnConflict" setter.
	TGoRetirementUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreationTx sets the "creation_tx" field.
func (u *TGoRetirementUpsert) SetCreationTx(v string) *TGoRetirementUpsert {
	u.Set(tgoretirement.FieldCreationTx, v)
	return u
}

// UpdateCreationTx sets the "creation_tx" field to the value that was provided on create.
func (u *TGoRetirementUpsert) UpdateCreationTx() *TGoRetirementUpsert {
	u.SetExcluded(tgoretirement.FieldCreationTx)
	return u
}

// SetCreatorAddress sets the "creator_address" field.
func (u *TGoRetirementUpsert) SetCreatorAddress(v string) *TGoRetirementUpsert {
	u.Set(tgoretirement.FieldCreatorAddress, v)
	return u
}

// UpdateCreatorAddress sets the "creator_address" field to the value that was provided on create.
func (u *TGoRetirementUpsert) UpdateCreatorAddress() *TGoRetirementUpsert {
	u.SetExcluded(tgoretirement.FieldCreatorAddress)
	return u
}

// SetBeneficiaryAddress sets the "beneficiary_address" field.
func (u *TGoRetirementUpsert) SetBeneficiaryAddress(v string) *TGoRetirementUpsert {
	u.Set(tgoretirement.FieldBeneficiaryAddress, v)
	return u
}

// UpdateBeneficiaryAddress sets the "beneficiary_address" field to the value that was provided on create.
func (u *TGoRetirementUpsert) UpdateBeneficiaryAddress() *TGoRetirementUpsert {
	u.SetExcluded(tgoretirement.FieldBeneficiaryAddress)
	return u
}

// SetAmount sets the "amount" field.
func (u *TGoRetirementUpsert) SetAmount(v float64) *TGoRetirementUpsert {
	u.Set(tgoretirement.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *TGoRetirementUpsert) UpdateAmount() *TGoRetirementUpsert {
	u.SetExcluded(tgoretirement.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *TGoRetirementUpsert) AddAmount(v float64) *TGoRetirementUpsert {
	u.Add(tgoretirement.FieldAmount, v)
	return u
}

// SetTokenAddress sets the "token_address" field.
func (u *TGoRetirementUpsert) SetTokenAddress(v string) *TGoRetirementUpsert {
	u.Set(tgoretirement.FieldTokenAddress, v)
	return u
}

// UpdateTokenAddress sets the "token_address" field to the value that was provided on create.
func (u *TGoRetirementUpsert) UpdateTokenAddress() *TGoRetirementUpsert {
	u.SetExcluded(tgoretirement.FieldTokenAddress)
	return u
}

// SetTokenName sets the "token_name" field.
func (u *TGoRetirementUpsert) SetTokenName(v string) *TGoRetirementUpsert {
	u.Set(tgoretirement.FieldTokenName, v)
	return u
}

// UpdateTokenName sets the "token_name" field to the value that was provided on create.
func (u *TGoRetirementUpsert) UpdateTokenName() *TGoRetirementUpsert {
	u.SetExcluded(tgoretirement.FieldTokenName)
	return u
}

// SetTokenType sets the "token_type" field.
func (u *TGoRetirementUpsert) SetTokenType(v string) *TGoRetirementUpsert {
	u.Set(tgoretirement.FieldTokenType, v)
	return u
}

// UpdateTokenType sets the "token_type" field to the value that was provided on create.
func (u *TGoRetirementUpsert) UpdateTokenType() *TGoRetirementUpsert {
	u.SetExcluded(tgoretirement.FieldTokenType)
	return u
}

// SetRetirementTime sets the "retirement_time" field.
func (u *TGoRetirementUpsert) SetRetirementTime(v time.Time) *TGoRetirementUpsert {
	u.Set(tgoretirement.FieldRetirementTime, v)
	return u
}

// UpdateRetirementTime sets the "retirement_time" field to the value that was provided on create.
func (u *TGoRetirementUpsert) UpdateRetirementTime() *TGoRetirementUpsert {
	u.SetExcluded(tgoretirement.FieldRetirementTime)
	return u
}

// SetMtime sets the "mtime" field.
func (u *TGoRetirementUpsert) SetMtime(v time.Time) *TGoRetirementUpsert {
	u.Set(tgoretirement.FieldMtime, v)
	return u
}

// UpdateMtime sets the "mtime" field to the value that was provided on create.
func (u *TGoRetirementUpsert) UpdateMtime() *TGoRetirementUpsert {
	u.SetExcluded(tgoretirement.FieldMtime)
	return u
}

// SetCtime sets the "ctime" field.
func (u *TGoRetirementUpsert) SetCtime(v time.Time) *TGoRetirementUpsert {
	u.Set(tgoretirement.FieldCtime, v)
	return u
}

// UpdateCtime sets the "ctime" field to the value that was provided on create.
func (u *TGoRetirementUpsert) UpdateCtime() *TGoRetirementUpsert {
	u.SetExcluded(tgoretirement.FieldCtime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TGoRetirement.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tgoretirement.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TGoRetirementUpsertOne) UpdateNewValues() *TGoRetirementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(tgoretirement.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TGoRetirement.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TGoRetirementUpsertOne) Ignore() *TGoRetirementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TGoRetirementUpsertOne) DoNothing() *TGoRetirementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TGoRetirementCreate.OnConflict
// documentation for more info.
func (u *TGoRetirementUpsertOne) Update(set func(*TGoRetirementUpsert)) *TGoRetirementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TGoRetirementUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreationTx sets the "creation_tx" field.
func (u *TGoRetirementUpsertOne) SetCreationTx(v string) *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetCreationTx(v)
	})
}

// UpdateCreationTx sets the "creation_tx" field to the value that was provided on create.
func (u *TGoRetirementUpsertOne) UpdateCreationTx() *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateCreationTx()
	})
}

// SetCreatorAddress sets the "creator_address" field.
func (u *TGoRetirementUpsertOne) SetCreatorAddress(v string) *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetCreatorAddress(v)
	})
}

// UpdateCreatorAddress sets the "creator_address" field to the value that was provided on create.
func (u *TGoRetirementUpsertOne) UpdateCreatorAddress() *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateCreatorAddress()
	})
}

// SetBeneficiaryAddress sets the "beneficiary_address" field.
func (u *TGoRetirementUpsertOne) SetBeneficiaryAddress(v string) *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetBeneficiaryAddress(v)
	})
}

// UpdateBeneficiaryAddress sets the "beneficiary_address" field to the value that was provided on create.
func (u *TGoRetirementUpsertOne) UpdateBeneficiaryAddress() *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateBeneficiaryAddress()
	})
}

// SetAmount sets the "amount" field.
func (u *TGoRetirementUpsertOne) SetAmount(v float64) *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *TGoRetirementUpsertOne) AddAmount(v float64) *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *TGoRetirementUpsertOne) UpdateAmount() *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateAmount()
	})
}

// SetTokenAddress sets the "token_address" field.
func (u *TGoRetirementUpsertOne) SetTokenAddress(v string) *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetTokenAddress(v)
	})
}

// UpdateTokenAddress sets the "token_address" field to the value that was provided on create.
func (u *TGoRetirementUpsertOne) UpdateTokenAddress() *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateTokenAddress()
	})
}

// SetTokenName sets the "token_name" field.
func (u *TGoRetirementUpsertOne) SetTokenName(v string) *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetTokenName(v)
	})
}

// UpdateTokenName sets the "token_name" field to the value that was provided on create.
func (u *TGoRetirementUpsertOne) UpdateTokenName() *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateTokenName()
	})
}

// SetTokenType sets the "token_type" field.
func (u *TGoRetirementUpsertOne) SetTokenType(v string) *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetTokenType(v)
	})
}

// UpdateTokenType sets the "token_type" field to the value that was provided on create.
func (u *TGoRetirementUpsertOne) UpdateTokenType() *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateTokenType()
	})
}

// SetRetirementTime sets the "retirement_time" field.
func (u *TGoRetirementUpsertOne) SetRetirementTime(v time.Time) *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetRetirementTime(v)
	})
}

// UpdateRetirementTime sets the "retirement_time" field to the value that was provided on create.
func (u *TGoRetirementUpsertOne) UpdateRetirementTime() *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateRetirementTime()
	})
}

// SetMtime sets the "mtime" field.
func (u *TGoRetirementUpsertOne) SetMtime(v time.Time) *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetMtime(v)
	})
}

// UpdateMtime sets the "mtime" field to the value that was provided on create.
func (u *TGoRetirementUpsertOne) UpdateMtime() *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateMtime()
	})
}

// SetCtime sets the "ctime" field.
func (u *TGoRetirementUpsertOne) SetCtime(v time.Time) *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetCtime(v)
	})
}

// UpdateCtime sets the "ctime" field to the value that was provided on create.
func (u *TGoRetirementUpsertOne) UpdateCtime() *TGoRetirementUpsertOne {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateCtime()
	})
}

// Exec executes the query.
func (u *TGoRetirementUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TGoRetirementCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TGoRetirementUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TGoRetirementUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TGoRetirementUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TGoRetirementCreateBulk is the builder for creating many TGoRetirement entities in bulk.
type TGoRetirementCreateBulk struct {
	config
	builders []*TGoRetirementCreate
	conflict []sql.ConflictOption
}

// Save creates the TGoRetirement entities in the database.
func (trcb *TGoRetirementCreateBulk) Save(ctx context.Context) ([]*TGoRetirement, error) {
	specs := make([]*sqlgraph.CreateSpec, len(trcb.builders))
	nodes := make([]*TGoRetirement, len(trcb.builders))
	mutators := make([]Mutator, len(trcb.builders))
	for i := range trcb.builders {
		func(i int, root context.Context) {
			builder := trcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TGoRetirementMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, trcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = trcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, trcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, trcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (trcb *TGoRetirementCreateBulk) SaveX(ctx context.Context) []*TGoRetirement {
	v, err := trcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (trcb *TGoRetirementCreateBulk) Exec(ctx context.Context) error {
	_, err := trcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trcb *TGoRetirementCreateBulk) ExecX(ctx context.Context) {
	if err := trcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TGoRetirement.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TGoRetirementUpsert) {
//			SetCreationTx(v+v).
//		}).
//		Exec(ctx)
//
func (trcb *TGoRetirementCreateBulk) OnConflict(opts ...sql.ConflictOption) *TGoRetirementUpsertBulk {
	trcb.conflict = opts
	return &TGoRetirementUpsertBulk{
		create: trcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TGoRetirement.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (trcb *TGoRetirementCreateBulk) OnConflictColumns(columns ...string) *TGoRetirementUpsertBulk {
	trcb.conflict = append(trcb.conflict, sql.ConflictColumns(columns...))
	return &TGoRetirementUpsertBulk{
		create: trcb,
	}
}

// TGoRetirementUpsertBulk is the builder for "upsert"-ing
// a bulk of TGoRetirement nodes.
type TGoRetirementUpsertBulk struct {
	create *TGoRetirementCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TGoRetirement.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tgoretirement.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TGoRetirementUpsertBulk) UpdateNewValues() *TGoRetirementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(tgoretirement.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TGoRetirement.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TGoRetirementUpsertBulk) Ignore() *TGoRetirementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TGoRetirementUpsertBulk) DoNothing() *TGoRetirementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TGoRetirementCreateBulk.OnConflict
// documentation for more info.
func (u *TGoRetirementUpsertBulk) Update(set func(*TGoRetirementUpsert)) *TGoRetirementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TGoRetirementUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreationTx sets the "creation_tx" field.
func (u *TGoRetirementUpsertBulk) SetCreationTx(v string) *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetCreationTx(v)
	})
}

// UpdateCreationTx sets the "creation_tx" field to the value that was provided on create.
func (u *TGoRetirementUpsertBulk) UpdateCreationTx() *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateCreationTx()
	})
}

// SetCreatorAddress sets the "creator_address" field.
func (u *TGoRetirementUpsertBulk) SetCreatorAddress(v string) *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetCreatorAddress(v)
	})
}

// UpdateCreatorAddress sets the "creator_address" field to the value that was provided on create.
func (u *TGoRetirementUpsertBulk) UpdateCreatorAddress() *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateCreatorAddress()
	})
}

// SetBeneficiaryAddress sets the "beneficiary_address" field.
func (u *TGoRetirementUpsertBulk) SetBeneficiaryAddress(v string) *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetBeneficiaryAddress(v)
	})
}

// UpdateBeneficiaryAddress sets the "beneficiary_address" field to the value that was provided on create.
func (u *TGoRetirementUpsertBulk) UpdateBeneficiaryAddress() *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateBeneficiaryAddress()
	})
}

// SetAmount sets the "amount" field.
func (u *TGoRetirementUpsertBulk) SetAmount(v float64) *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *TGoRetirementUpsertBulk) AddAmount(v float64) *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *TGoRetirementUpsertBulk) UpdateAmount() *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateAmount()
	})
}

// SetTokenAddress sets the "token_address" field.
func (u *TGoRetirementUpsertBulk) SetTokenAddress(v string) *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetTokenAddress(v)
	})
}

// UpdateTokenAddress sets the "token_address" field to the value that was provided on create.
func (u *TGoRetirementUpsertBulk) UpdateTokenAddress() *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateTokenAddress()
	})
}

// SetTokenName sets the "token_name" field.
func (u *TGoRetirementUpsertBulk) SetTokenName(v string) *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetTokenName(v)
	})
}

// UpdateTokenName sets the "token_name" field to the value that was provided on create.
func (u *TGoRetirementUpsertBulk) UpdateTokenName() *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateTokenName()
	})
}

// SetTokenType sets the "token_type" field.
func (u *TGoRetirementUpsertBulk) SetTokenType(v string) *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetTokenType(v)
	})
}

// UpdateTokenType sets the "token_type" field to the value that was provided on create.
func (u *TGoRetirementUpsertBulk) UpdateTokenType() *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateTokenType()
	})
}

// SetRetirementTime sets the "retirement_time" field.
func (u *TGoRetirementUpsertBulk) SetRetirementTime(v time.Time) *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetRetirementTime(v)
	})
}

// UpdateRetirementTime sets the "retirement_time" field to the value that was provided on create.
func (u *TGoRetirementUpsertBulk) UpdateRetirementTime() *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateRetirementTime()
	})
}

// SetMtime sets the "mtime" field.
func (u *TGoRetirementUpsertBulk) SetMtime(v time.Time) *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetMtime(v)
	})
}

// UpdateMtime sets the "mtime" field to the value that was provided on create.
func (u *TGoRetirementUpsertBulk) UpdateMtime() *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateMtime()
	})
}

// SetCtime sets the "ctime" field.
func (u *TGoRetirementUpsertBulk) SetCtime(v time.Time) *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.SetCtime(v)
	})
}

// UpdateCtime sets the "ctime" field to the value that was provided on create.
func (u *TGoRetirementUpsertBulk) UpdateCtime() *TGoRetirementUpsertBulk {
	return u.Update(func(s *TGoRetirementUpsert) {
		s.UpdateCtime()
	})
}

// Exec executes the query.
func (u *TGoRetirementUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TGoRetirementCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TGoRetirementCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TGoRetirementUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
