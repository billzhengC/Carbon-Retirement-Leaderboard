// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"toucan-leaderboard/ent/predicate"
	"toucan-leaderboard/ent/tgocache"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TGoCacheUpdate is the builder for updating TGoCache entities.
type TGoCacheUpdate struct {
	config
	hooks    []Hook
	mutation *TGoCacheMutation
}

// Where appends a list predicates to the TGoCacheUpdate builder.
func (tcu *TGoCacheUpdate) Where(ps ...predicate.TGoCache) *TGoCacheUpdate {
	tcu.mutation.Where(ps...)
	return tcu
}

// SetCacheKey sets the "cache_key" field.
func (tcu *TGoCacheUpdate) SetCacheKey(s string) *TGoCacheUpdate {
	tcu.mutation.SetCacheKey(s)
	return tcu
}

// SetCacheValue sets the "cache_value" field.
func (tcu *TGoCacheUpdate) SetCacheValue(s string) *TGoCacheUpdate {
	tcu.mutation.SetCacheValue(s)
	return tcu
}

// SetNillableCacheValue sets the "cache_value" field if the given value is not nil.
func (tcu *TGoCacheUpdate) SetNillableCacheValue(s *string) *TGoCacheUpdate {
	if s != nil {
		tcu.SetCacheValue(*s)
	}
	return tcu
}

// SetMtime sets the "mtime" field.
func (tcu *TGoCacheUpdate) SetMtime(t time.Time) *TGoCacheUpdate {
	tcu.mutation.SetMtime(t)
	return tcu
}

// SetCtime sets the "ctime" field.
func (tcu *TGoCacheUpdate) SetCtime(t time.Time) *TGoCacheUpdate {
	tcu.mutation.SetCtime(t)
	return tcu
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (tcu *TGoCacheUpdate) SetNillableCtime(t *time.Time) *TGoCacheUpdate {
	if t != nil {
		tcu.SetCtime(*t)
	}
	return tcu
}

// Mutation returns the TGoCacheMutation object of the builder.
func (tcu *TGoCacheUpdate) Mutation() *TGoCacheMutation {
	return tcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcu *TGoCacheUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tcu.defaults()
	if len(tcu.hooks) == 0 {
		affected, err = tcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TGoCacheMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tcu.mutation = mutation
			affected, err = tcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tcu.hooks) - 1; i >= 0; i-- {
			if tcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tcu *TGoCacheUpdate) SaveX(ctx context.Context) int {
	affected, err := tcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcu *TGoCacheUpdate) Exec(ctx context.Context) error {
	_, err := tcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcu *TGoCacheUpdate) ExecX(ctx context.Context) {
	if err := tcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcu *TGoCacheUpdate) defaults() {
	if _, ok := tcu.mutation.Mtime(); !ok {
		v := tgocache.UpdateDefaultMtime()
		tcu.mutation.SetMtime(v)
	}
}

func (tcu *TGoCacheUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tgocache.Table,
			Columns: tgocache.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: tgocache.FieldID,
			},
		},
	}
	if ps := tcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcu.mutation.CacheKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgocache.FieldCacheKey,
		})
	}
	if value, ok := tcu.mutation.CacheValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgocache.FieldCacheValue,
		})
	}
	if value, ok := tcu.mutation.Mtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgocache.FieldMtime,
		})
	}
	if value, ok := tcu.mutation.Ctime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgocache.FieldCtime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tgocache.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TGoCacheUpdateOne is the builder for updating a single TGoCache entity.
type TGoCacheUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TGoCacheMutation
}

// SetCacheKey sets the "cache_key" field.
func (tcuo *TGoCacheUpdateOne) SetCacheKey(s string) *TGoCacheUpdateOne {
	tcuo.mutation.SetCacheKey(s)
	return tcuo
}

// SetCacheValue sets the "cache_value" field.
func (tcuo *TGoCacheUpdateOne) SetCacheValue(s string) *TGoCacheUpdateOne {
	tcuo.mutation.SetCacheValue(s)
	return tcuo
}

// SetNillableCacheValue sets the "cache_value" field if the given value is not nil.
func (tcuo *TGoCacheUpdateOne) SetNillableCacheValue(s *string) *TGoCacheUpdateOne {
	if s != nil {
		tcuo.SetCacheValue(*s)
	}
	return tcuo
}

// SetMtime sets the "mtime" field.
func (tcuo *TGoCacheUpdateOne) SetMtime(t time.Time) *TGoCacheUpdateOne {
	tcuo.mutation.SetMtime(t)
	return tcuo
}

// SetCtime sets the "ctime" field.
func (tcuo *TGoCacheUpdateOne) SetCtime(t time.Time) *TGoCacheUpdateOne {
	tcuo.mutation.SetCtime(t)
	return tcuo
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (tcuo *TGoCacheUpdateOne) SetNillableCtime(t *time.Time) *TGoCacheUpdateOne {
	if t != nil {
		tcuo.SetCtime(*t)
	}
	return tcuo
}

// Mutation returns the TGoCacheMutation object of the builder.
func (tcuo *TGoCacheUpdateOne) Mutation() *TGoCacheMutation {
	return tcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcuo *TGoCacheUpdateOne) Select(field string, fields ...string) *TGoCacheUpdateOne {
	tcuo.fields = append([]string{field}, fields...)
	return tcuo
}

// Save executes the query and returns the updated TGoCache entity.
func (tcuo *TGoCacheUpdateOne) Save(ctx context.Context) (*TGoCache, error) {
	var (
		err  error
		node *TGoCache
	)
	tcuo.defaults()
	if len(tcuo.hooks) == 0 {
		node, err = tcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TGoCacheMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tcuo.mutation = mutation
			node, err = tcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tcuo.hooks) - 1; i >= 0; i-- {
			if tcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TGoCache)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TGoCacheMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tcuo *TGoCacheUpdateOne) SaveX(ctx context.Context) *TGoCache {
	node, err := tcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcuo *TGoCacheUpdateOne) Exec(ctx context.Context) error {
	_, err := tcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcuo *TGoCacheUpdateOne) ExecX(ctx context.Context) {
	if err := tcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcuo *TGoCacheUpdateOne) defaults() {
	if _, ok := tcuo.mutation.Mtime(); !ok {
		v := tgocache.UpdateDefaultMtime()
		tcuo.mutation.SetMtime(v)
	}
}

func (tcuo *TGoCacheUpdateOne) sqlSave(ctx context.Context) (_node *TGoCache, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tgocache.Table,
			Columns: tgocache.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: tgocache.FieldID,
			},
		},
	}
	id, ok := tcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TGoCache.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tgocache.FieldID)
		for _, f := range fields {
			if !tgocache.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tgocache.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcuo.mutation.CacheKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgocache.FieldCacheKey,
		})
	}
	if value, ok := tcuo.mutation.CacheValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgocache.FieldCacheValue,
		})
	}
	if value, ok := tcuo.mutation.Mtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgocache.FieldMtime,
		})
	}
	if value, ok := tcuo.mutation.Ctime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgocache.FieldCtime,
		})
	}
	_node = &TGoCache{config: tcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tgocache.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}