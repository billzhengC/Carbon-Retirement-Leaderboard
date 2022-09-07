// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"toucan-leaderboard/ent/predicate"
	"toucan-leaderboard/ent/tgoens"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TGoEnsQuery is the builder for querying TGoEns entities.
type TGoEnsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TGoEns
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TGoEnsQuery builder.
func (teq *TGoEnsQuery) Where(ps ...predicate.TGoEns) *TGoEnsQuery {
	teq.predicates = append(teq.predicates, ps...)
	return teq
}

// Limit adds a limit step to the query.
func (teq *TGoEnsQuery) Limit(limit int) *TGoEnsQuery {
	teq.limit = &limit
	return teq
}

// Offset adds an offset step to the query.
func (teq *TGoEnsQuery) Offset(offset int) *TGoEnsQuery {
	teq.offset = &offset
	return teq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (teq *TGoEnsQuery) Unique(unique bool) *TGoEnsQuery {
	teq.unique = &unique
	return teq
}

// Order adds an order step to the query.
func (teq *TGoEnsQuery) Order(o ...OrderFunc) *TGoEnsQuery {
	teq.order = append(teq.order, o...)
	return teq
}

// First returns the first TGoEns entity from the query.
// Returns a *NotFoundError when no TGoEns was found.
func (teq *TGoEnsQuery) First(ctx context.Context) (*TGoEns, error) {
	nodes, err := teq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tgoens.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (teq *TGoEnsQuery) FirstX(ctx context.Context) *TGoEns {
	node, err := teq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TGoEns ID from the query.
// Returns a *NotFoundError when no TGoEns ID was found.
func (teq *TGoEnsQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = teq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tgoens.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (teq *TGoEnsQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := teq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TGoEns entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TGoEns entity is found.
// Returns a *NotFoundError when no TGoEns entities are found.
func (teq *TGoEnsQuery) Only(ctx context.Context) (*TGoEns, error) {
	nodes, err := teq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tgoens.Label}
	default:
		return nil, &NotSingularError{tgoens.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (teq *TGoEnsQuery) OnlyX(ctx context.Context) *TGoEns {
	node, err := teq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TGoEns ID in the query.
// Returns a *NotSingularError when more than one TGoEns ID is found.
// Returns a *NotFoundError when no entities are found.
func (teq *TGoEnsQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = teq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tgoens.Label}
	default:
		err = &NotSingularError{tgoens.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (teq *TGoEnsQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := teq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TGoEnsSlice.
func (teq *TGoEnsQuery) All(ctx context.Context) ([]*TGoEns, error) {
	if err := teq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return teq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (teq *TGoEnsQuery) AllX(ctx context.Context) []*TGoEns {
	nodes, err := teq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TGoEns IDs.
func (teq *TGoEnsQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := teq.Select(tgoens.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (teq *TGoEnsQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := teq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (teq *TGoEnsQuery) Count(ctx context.Context) (int, error) {
	if err := teq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return teq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (teq *TGoEnsQuery) CountX(ctx context.Context) int {
	count, err := teq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (teq *TGoEnsQuery) Exist(ctx context.Context) (bool, error) {
	if err := teq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return teq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (teq *TGoEnsQuery) ExistX(ctx context.Context) bool {
	exist, err := teq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TGoEnsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (teq *TGoEnsQuery) Clone() *TGoEnsQuery {
	if teq == nil {
		return nil
	}
	return &TGoEnsQuery{
		config:     teq.config,
		limit:      teq.limit,
		offset:     teq.offset,
		order:      append([]OrderFunc{}, teq.order...),
		predicates: append([]predicate.TGoEns{}, teq.predicates...),
		// clone intermediate query.
		sql:    teq.sql.Clone(),
		path:   teq.path,
		unique: teq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		WalletPub string `json:"wallet_pub,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TGoEns.Query().
//		GroupBy(tgoens.FieldWalletPub).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (teq *TGoEnsQuery) GroupBy(field string, fields ...string) *TGoEnsGroupBy {
	grbuild := &TGoEnsGroupBy{config: teq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := teq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return teq.sqlQuery(ctx), nil
	}
	grbuild.label = tgoens.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		WalletPub string `json:"wallet_pub,omitempty"`
//	}
//
//	client.TGoEns.Query().
//		Select(tgoens.FieldWalletPub).
//		Scan(ctx, &v)
//
func (teq *TGoEnsQuery) Select(fields ...string) *TGoEnsSelect {
	teq.fields = append(teq.fields, fields...)
	selbuild := &TGoEnsSelect{TGoEnsQuery: teq}
	selbuild.label = tgoens.Label
	selbuild.flds, selbuild.scan = &teq.fields, selbuild.Scan
	return selbuild
}

func (teq *TGoEnsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range teq.fields {
		if !tgoens.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if teq.path != nil {
		prev, err := teq.path(ctx)
		if err != nil {
			return err
		}
		teq.sql = prev
	}
	return nil
}

func (teq *TGoEnsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TGoEns, error) {
	var (
		nodes = []*TGoEns{}
		_spec = teq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*TGoEns).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &TGoEns{config: teq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, teq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (teq *TGoEnsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := teq.querySpec()
	_spec.Node.Columns = teq.fields
	if len(teq.fields) > 0 {
		_spec.Unique = teq.unique != nil && *teq.unique
	}
	return sqlgraph.CountNodes(ctx, teq.driver, _spec)
}

func (teq *TGoEnsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := teq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (teq *TGoEnsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tgoens.Table,
			Columns: tgoens.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: tgoens.FieldID,
			},
		},
		From:   teq.sql,
		Unique: true,
	}
	if unique := teq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := teq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tgoens.FieldID)
		for i := range fields {
			if fields[i] != tgoens.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := teq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := teq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := teq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := teq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (teq *TGoEnsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(teq.driver.Dialect())
	t1 := builder.Table(tgoens.Table)
	columns := teq.fields
	if len(columns) == 0 {
		columns = tgoens.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if teq.sql != nil {
		selector = teq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if teq.unique != nil && *teq.unique {
		selector.Distinct()
	}
	for _, p := range teq.predicates {
		p(selector)
	}
	for _, p := range teq.order {
		p(selector)
	}
	if offset := teq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := teq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TGoEnsGroupBy is the group-by builder for TGoEns entities.
type TGoEnsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tegb *TGoEnsGroupBy) Aggregate(fns ...AggregateFunc) *TGoEnsGroupBy {
	tegb.fns = append(tegb.fns, fns...)
	return tegb
}

// Scan applies the group-by query and scans the result into the given value.
func (tegb *TGoEnsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tegb.path(ctx)
	if err != nil {
		return err
	}
	tegb.sql = query
	return tegb.sqlScan(ctx, v)
}

func (tegb *TGoEnsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tegb.fields {
		if !tgoens.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tegb *TGoEnsGroupBy) sqlQuery() *sql.Selector {
	selector := tegb.sql.Select()
	aggregation := make([]string, 0, len(tegb.fns))
	for _, fn := range tegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tegb.fields)+len(tegb.fns))
		for _, f := range tegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tegb.fields...)...)
}

// TGoEnsSelect is the builder for selecting fields of TGoEns entities.
type TGoEnsSelect struct {
	*TGoEnsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tes *TGoEnsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tes.prepareQuery(ctx); err != nil {
		return err
	}
	tes.sql = tes.TGoEnsQuery.sqlQuery(ctx)
	return tes.sqlScan(ctx, v)
}

func (tes *TGoEnsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tes.sql.Query()
	if err := tes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
