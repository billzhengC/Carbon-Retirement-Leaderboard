// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"toucan-leaderboard/ent/tgoens"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TGoEnsCreate is the builder for creating a TGoEns entity.
type TGoEnsCreate struct {
	config
	mutation *TGoEnsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetWalletPub sets the "wallet_pub" field.
func (tec *TGoEnsCreate) SetWalletPub(s string) *TGoEnsCreate {
	tec.mutation.SetWalletPub(s)
	return tec
}

// SetEns sets the "ens" field.
func (tec *TGoEnsCreate) SetEns(s string) *TGoEnsCreate {
	tec.mutation.SetEns(s)
	return tec
}

// SetNillableEns sets the "ens" field if the given value is not nil.
func (tec *TGoEnsCreate) SetNillableEns(s *string) *TGoEnsCreate {
	if s != nil {
		tec.SetEns(*s)
	}
	return tec
}

// SetMtime sets the "mtime" field.
func (tec *TGoEnsCreate) SetMtime(t time.Time) *TGoEnsCreate {
	tec.mutation.SetMtime(t)
	return tec
}

// SetNillableMtime sets the "mtime" field if the given value is not nil.
func (tec *TGoEnsCreate) SetNillableMtime(t *time.Time) *TGoEnsCreate {
	if t != nil {
		tec.SetMtime(*t)
	}
	return tec
}

// SetCtime sets the "ctime" field.
func (tec *TGoEnsCreate) SetCtime(t time.Time) *TGoEnsCreate {
	tec.mutation.SetCtime(t)
	return tec
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (tec *TGoEnsCreate) SetNillableCtime(t *time.Time) *TGoEnsCreate {
	if t != nil {
		tec.SetCtime(*t)
	}
	return tec
}

// SetID sets the "id" field.
func (tec *TGoEnsCreate) SetID(u uint64) *TGoEnsCreate {
	tec.mutation.SetID(u)
	return tec
}

// Mutation returns the TGoEnsMutation object of the builder.
func (tec *TGoEnsCreate) Mutation() *TGoEnsMutation {
	return tec.mutation
}

// Save creates the TGoEns in the database.
func (tec *TGoEnsCreate) Save(ctx context.Context) (*TGoEns, error) {
	var (
		err  error
		node *TGoEns
	)
	tec.defaults()
	if len(tec.hooks) == 0 {
		if err = tec.check(); err != nil {
			return nil, err
		}
		node, err = tec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TGoEnsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tec.check(); err != nil {
				return nil, err
			}
			tec.mutation = mutation
			if node, err = tec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tec.hooks) - 1; i >= 0; i-- {
			if tec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TGoEns)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TGoEnsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tec *TGoEnsCreate) SaveX(ctx context.Context) *TGoEns {
	v, err := tec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tec *TGoEnsCreate) Exec(ctx context.Context) error {
	_, err := tec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tec *TGoEnsCreate) ExecX(ctx context.Context) {
	if err := tec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tec *TGoEnsCreate) defaults() {
	if _, ok := tec.mutation.Ens(); !ok {
		v := tgoens.DefaultEns
		tec.mutation.SetEns(v)
	}
	if _, ok := tec.mutation.Mtime(); !ok {
		v := tgoens.DefaultMtime
		tec.mutation.SetMtime(v)
	}
	if _, ok := tec.mutation.Ctime(); !ok {
		v := tgoens.DefaultCtime
		tec.mutation.SetCtime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tec *TGoEnsCreate) check() error {
	if _, ok := tec.mutation.WalletPub(); !ok {
		return &ValidationError{Name: "wallet_pub", err: errors.New(`ent: missing required field "TGoEns.wallet_pub"`)}
	}
	if _, ok := tec.mutation.Ens(); !ok {
		return &ValidationError{Name: "ens", err: errors.New(`ent: missing required field "TGoEns.ens"`)}
	}
	if _, ok := tec.mutation.Mtime(); !ok {
		return &ValidationError{Name: "mtime", err: errors.New(`ent: missing required field "TGoEns.mtime"`)}
	}
	if _, ok := tec.mutation.Ctime(); !ok {
		return &ValidationError{Name: "ctime", err: errors.New(`ent: missing required field "TGoEns.ctime"`)}
	}
	return nil
}

func (tec *TGoEnsCreate) sqlSave(ctx context.Context) (*TGoEns, error) {
	_node, _spec := tec.createSpec()
	if err := sqlgraph.CreateNode(ctx, tec.driver, _spec); err != nil {
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

func (tec *TGoEnsCreate) createSpec() (*TGoEns, *sqlgraph.CreateSpec) {
	var (
		_node = &TGoEns{config: tec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tgoens.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: tgoens.FieldID,
			},
		}
	)
	_spec.OnConflict = tec.conflict
	if id, ok := tec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tec.mutation.WalletPub(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgoens.FieldWalletPub,
		})
		_node.WalletPub = value
	}
	if value, ok := tec.mutation.Ens(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgoens.FieldEns,
		})
		_node.Ens = value
	}
	if value, ok := tec.mutation.Mtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgoens.FieldMtime,
		})
		_node.Mtime = value
	}
	if value, ok := tec.mutation.Ctime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgoens.FieldCtime,
		})
		_node.Ctime = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TGoEns.Create().
//		SetWalletPub(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TGoEnsUpsert) {
//			SetWalletPub(v+v).
//		}).
//		Exec(ctx)
//
func (tec *TGoEnsCreate) OnConflict(opts ...sql.ConflictOption) *TGoEnsUpsertOne {
	tec.conflict = opts
	return &TGoEnsUpsertOne{
		create: tec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TGoEns.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tec *TGoEnsCreate) OnConflictColumns(columns ...string) *TGoEnsUpsertOne {
	tec.conflict = append(tec.conflict, sql.ConflictColumns(columns...))
	return &TGoEnsUpsertOne{
		create: tec,
	}
}

type (
	// TGoEnsUpsertOne is the builder for "upsert"-ing
	//  one TGoEns node.
	TGoEnsUpsertOne struct {
		create *TGoEnsCreate
	}

	// TGoEnsUpsert is the "OnConflict" setter.
	TGoEnsUpsert struct {
		*sql.UpdateSet
	}
)

// SetWalletPub sets the "wallet_pub" field.
func (u *TGoEnsUpsert) SetWalletPub(v string) *TGoEnsUpsert {
	u.Set(tgoens.FieldWalletPub, v)
	return u
}

// UpdateWalletPub sets the "wallet_pub" field to the value that was provided on create.
func (u *TGoEnsUpsert) UpdateWalletPub() *TGoEnsUpsert {
	u.SetExcluded(tgoens.FieldWalletPub)
	return u
}

// SetEns sets the "ens" field.
func (u *TGoEnsUpsert) SetEns(v string) *TGoEnsUpsert {
	u.Set(tgoens.FieldEns, v)
	return u
}

// UpdateEns sets the "ens" field to the value that was provided on create.
func (u *TGoEnsUpsert) UpdateEns() *TGoEnsUpsert {
	u.SetExcluded(tgoens.FieldEns)
	return u
}

// SetMtime sets the "mtime" field.
func (u *TGoEnsUpsert) SetMtime(v time.Time) *TGoEnsUpsert {
	u.Set(tgoens.FieldMtime, v)
	return u
}

// UpdateMtime sets the "mtime" field to the value that was provided on create.
func (u *TGoEnsUpsert) UpdateMtime() *TGoEnsUpsert {
	u.SetExcluded(tgoens.FieldMtime)
	return u
}

// SetCtime sets the "ctime" field.
func (u *TGoEnsUpsert) SetCtime(v time.Time) *TGoEnsUpsert {
	u.Set(tgoens.FieldCtime, v)
	return u
}

// UpdateCtime sets the "ctime" field to the value that was provided on create.
func (u *TGoEnsUpsert) UpdateCtime() *TGoEnsUpsert {
	u.SetExcluded(tgoens.FieldCtime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TGoEns.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tgoens.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TGoEnsUpsertOne) UpdateNewValues() *TGoEnsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(tgoens.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TGoEns.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TGoEnsUpsertOne) Ignore() *TGoEnsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TGoEnsUpsertOne) DoNothing() *TGoEnsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TGoEnsCreate.OnConflict
// documentation for more info.
func (u *TGoEnsUpsertOne) Update(set func(*TGoEnsUpsert)) *TGoEnsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TGoEnsUpsert{UpdateSet: update})
	}))
	return u
}

// SetWalletPub sets the "wallet_pub" field.
func (u *TGoEnsUpsertOne) SetWalletPub(v string) *TGoEnsUpsertOne {
	return u.Update(func(s *TGoEnsUpsert) {
		s.SetWalletPub(v)
	})
}

// UpdateWalletPub sets the "wallet_pub" field to the value that was provided on create.
func (u *TGoEnsUpsertOne) UpdateWalletPub() *TGoEnsUpsertOne {
	return u.Update(func(s *TGoEnsUpsert) {
		s.UpdateWalletPub()
	})
}

// SetEns sets the "ens" field.
func (u *TGoEnsUpsertOne) SetEns(v string) *TGoEnsUpsertOne {
	return u.Update(func(s *TGoEnsUpsert) {
		s.SetEns(v)
	})
}

// UpdateEns sets the "ens" field to the value that was provided on create.
func (u *TGoEnsUpsertOne) UpdateEns() *TGoEnsUpsertOne {
	return u.Update(func(s *TGoEnsUpsert) {
		s.UpdateEns()
	})
}

// SetMtime sets the "mtime" field.
func (u *TGoEnsUpsertOne) SetMtime(v time.Time) *TGoEnsUpsertOne {
	return u.Update(func(s *TGoEnsUpsert) {
		s.SetMtime(v)
	})
}

// UpdateMtime sets the "mtime" field to the value that was provided on create.
func (u *TGoEnsUpsertOne) UpdateMtime() *TGoEnsUpsertOne {
	return u.Update(func(s *TGoEnsUpsert) {
		s.UpdateMtime()
	})
}

// SetCtime sets the "ctime" field.
func (u *TGoEnsUpsertOne) SetCtime(v time.Time) *TGoEnsUpsertOne {
	return u.Update(func(s *TGoEnsUpsert) {
		s.SetCtime(v)
	})
}

// UpdateCtime sets the "ctime" field to the value that was provided on create.
func (u *TGoEnsUpsertOne) UpdateCtime() *TGoEnsUpsertOne {
	return u.Update(func(s *TGoEnsUpsert) {
		s.UpdateCtime()
	})
}

// Exec executes the query.
func (u *TGoEnsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TGoEnsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TGoEnsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TGoEnsUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TGoEnsUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TGoEnsCreateBulk is the builder for creating many TGoEns entities in bulk.
type TGoEnsCreateBulk struct {
	config
	builders []*TGoEnsCreate
	conflict []sql.ConflictOption
}

// Save creates the TGoEns entities in the database.
func (tecb *TGoEnsCreateBulk) Save(ctx context.Context) ([]*TGoEns, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tecb.builders))
	nodes := make([]*TGoEns, len(tecb.builders))
	mutators := make([]Mutator, len(tecb.builders))
	for i := range tecb.builders {
		func(i int, root context.Context) {
			builder := tecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TGoEnsMutation)
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
					_, err = mutators[i+1].Mutate(root, tecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tecb *TGoEnsCreateBulk) SaveX(ctx context.Context) []*TGoEns {
	v, err := tecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tecb *TGoEnsCreateBulk) Exec(ctx context.Context) error {
	_, err := tecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tecb *TGoEnsCreateBulk) ExecX(ctx context.Context) {
	if err := tecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TGoEns.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TGoEnsUpsert) {
//			SetWalletPub(v+v).
//		}).
//		Exec(ctx)
//
func (tecb *TGoEnsCreateBulk) OnConflict(opts ...sql.ConflictOption) *TGoEnsUpsertBulk {
	tecb.conflict = opts
	return &TGoEnsUpsertBulk{
		create: tecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TGoEns.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tecb *TGoEnsCreateBulk) OnConflictColumns(columns ...string) *TGoEnsUpsertBulk {
	tecb.conflict = append(tecb.conflict, sql.ConflictColumns(columns...))
	return &TGoEnsUpsertBulk{
		create: tecb,
	}
}

// TGoEnsUpsertBulk is the builder for "upsert"-ing
// a bulk of TGoEns nodes.
type TGoEnsUpsertBulk struct {
	create *TGoEnsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TGoEns.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tgoens.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TGoEnsUpsertBulk) UpdateNewValues() *TGoEnsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(tgoens.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TGoEns.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TGoEnsUpsertBulk) Ignore() *TGoEnsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TGoEnsUpsertBulk) DoNothing() *TGoEnsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TGoEnsCreateBulk.OnConflict
// documentation for more info.
func (u *TGoEnsUpsertBulk) Update(set func(*TGoEnsUpsert)) *TGoEnsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TGoEnsUpsert{UpdateSet: update})
	}))
	return u
}

// SetWalletPub sets the "wallet_pub" field.
func (u *TGoEnsUpsertBulk) SetWalletPub(v string) *TGoEnsUpsertBulk {
	return u.Update(func(s *TGoEnsUpsert) {
		s.SetWalletPub(v)
	})
}

// UpdateWalletPub sets the "wallet_pub" field to the value that was provided on create.
func (u *TGoEnsUpsertBulk) UpdateWalletPub() *TGoEnsUpsertBulk {
	return u.Update(func(s *TGoEnsUpsert) {
		s.UpdateWalletPub()
	})
}

// SetEns sets the "ens" field.
func (u *TGoEnsUpsertBulk) SetEns(v string) *TGoEnsUpsertBulk {
	return u.Update(func(s *TGoEnsUpsert) {
		s.SetEns(v)
	})
}

// UpdateEns sets the "ens" field to the value that was provided on create.
func (u *TGoEnsUpsertBulk) UpdateEns() *TGoEnsUpsertBulk {
	return u.Update(func(s *TGoEnsUpsert) {
		s.UpdateEns()
	})
}

// SetMtime sets the "mtime" field.
func (u *TGoEnsUpsertBulk) SetMtime(v time.Time) *TGoEnsUpsertBulk {
	return u.Update(func(s *TGoEnsUpsert) {
		s.SetMtime(v)
	})
}

// UpdateMtime sets the "mtime" field to the value that was provided on create.
func (u *TGoEnsUpsertBulk) UpdateMtime() *TGoEnsUpsertBulk {
	return u.Update(func(s *TGoEnsUpsert) {
		s.UpdateMtime()
	})
}

// SetCtime sets the "ctime" field.
func (u *TGoEnsUpsertBulk) SetCtime(v time.Time) *TGoEnsUpsertBulk {
	return u.Update(func(s *TGoEnsUpsert) {
		s.SetCtime(v)
	})
}

// UpdateCtime sets the "ctime" field to the value that was provided on create.
func (u *TGoEnsUpsertBulk) UpdateCtime() *TGoEnsUpsertBulk {
	return u.Update(func(s *TGoEnsUpsert) {
		s.UpdateCtime()
	})
}

// Exec executes the query.
func (u *TGoEnsUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TGoEnsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TGoEnsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TGoEnsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}