// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"toucan-leaderboard/ent/predicate"
	"toucan-leaderboard/ent/tgoretirement"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TGoRetirementDelete is the builder for deleting a TGoRetirement entity.
type TGoRetirementDelete struct {
	config
	hooks    []Hook
	mutation *TGoRetirementMutation
}

// Where appends a list predicates to the TGoRetirementDelete builder.
func (trd *TGoRetirementDelete) Where(ps ...predicate.TGoRetirement) *TGoRetirementDelete {
	trd.mutation.Where(ps...)
	return trd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (trd *TGoRetirementDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(trd.hooks) == 0 {
		affected, err = trd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TGoRetirementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			trd.mutation = mutation
			affected, err = trd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(trd.hooks) - 1; i >= 0; i-- {
			if trd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = trd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, trd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (trd *TGoRetirementDelete) ExecX(ctx context.Context) int {
	n, err := trd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (trd *TGoRetirementDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tgoretirement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: tgoretirement.FieldID,
			},
		},
	}
	if ps := trd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, trd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TGoRetirementDeleteOne is the builder for deleting a single TGoRetirement entity.
type TGoRetirementDeleteOne struct {
	trd *TGoRetirementDelete
}

// Exec executes the deletion query.
func (trdo *TGoRetirementDeleteOne) Exec(ctx context.Context) error {
	n, err := trdo.trd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tgoretirement.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (trdo *TGoRetirementDeleteOne) ExecX(ctx context.Context) {
	trdo.trd.ExecX(ctx)
}
