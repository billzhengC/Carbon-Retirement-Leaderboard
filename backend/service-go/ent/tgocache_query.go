// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"toucan-leaderboard/ent/predicate"
	"toucan-leaderboard/ent/tgocache"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TGoCacheQuery is the builder for querying TGoCache entities.
type TGoCacheQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TGoCache
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TGoCacheQuery builder.
func (tcq *TGoCacheQuery) Where(ps ...predicate.TGoCache) *TGoCacheQuery {
	tcq.predicates = append(tcq.predicates, ps...)
	return tcq
}

// Limit adds a limit step to the query.
func (tcq *TGoCacheQuery) Limit(limit int) *TGoCacheQuery {
	tcq.limit = &limit
	return tcq
}

// Offset adds an offset step to the query.
func (tcq *TGoCacheQuery) Offset(offset int) *TGoCacheQuery {
	tcq.offset = &offset
	return tcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tcq *TGoCacheQuery) Unique(unique bool) *TGoCacheQuery {
	tcq.unique = &unique
	return tcq
}

// Order adds an order step to the query.
func (tcq *TGoCacheQuery) Order(o ...OrderFunc) *TGoCacheQuery {
	tcq.order = append(tcq.order, o...)
	return tcq
}

// First returns the first TGoCache entity from the query.
// Returns a *NotFoundError when no TGoCache was found.
func (tcq *TGoCacheQuery) First(ctx context.Context) (*TGoCache, error) {
	nodes, err := tcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tgocache.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tcq *TGoCacheQuery) FirstX(ctx context.Context) *TGoCache {
	node, err := tcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TGoCache ID from the query.
// Returns a *NotFoundError when no TGoCache ID was found.
func (tcq *TGoCacheQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = tcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tgocache.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tcq *TGoCacheQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := tcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TGoCache entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TGoCache entity is found.
// Returns a *NotFoundError when no TGoCache entities are found.
func (tcq *TGoCacheQuery) Only(ctx context.Context) (*TGoCache, error) {
	nodes, err := tcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tgocache.Label}
	default:
		return nil, &NotSingularError{tgocache.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tcq *TGoCacheQuery) OnlyX(ctx context.Context) *TGoCache {
	node, err := tcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TGoCache ID in the query.
// Returns a *NotSingularError when more than one TGoCache ID is found.
// Returns a *NotFoundError when no entities are found.
func (tcq *TGoCacheQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = tcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tgocache.Label}
	default:
		err = &NotSingularError{tgocache.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tcq *TGoCacheQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := tcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TGoCaches.
func (tcq *TGoCacheQuery) All(ctx context.Context) ([]*TGoCache, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tcq *TGoCacheQuery) AllX(ctx context.Context) []*TGoCache {
	nodes, err := tcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TGoCache IDs.
func (tcq *TGoCacheQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := tcq.Select(tgocache.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tcq *TGoCacheQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := tcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tcq *TGoCacheQuery) Count(ctx context.Context) (int, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tcq *TGoCacheQuery) CountX(ctx context.Context) int {
	count, err := tcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tcq *TGoCacheQuery) Exist(ctx context.Context) (bool, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tcq *TGoCacheQuery) ExistX(ctx context.Context) bool {
	exist, err := tcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TGoCacheQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tcq *TGoCacheQuery) Clone() *TGoCacheQuery {
	if tcq == nil {
		return nil
	}
	return &TGoCacheQuery{
		config:     tcq.config,
		limit:      tcq.limit,
		offset:     tcq.offset,
		order:      append([]OrderFunc{}, tcq.order...),
		predicates: append([]predicate.TGoCache{}, tcq.predicates...),
		// clone intermediate query.
		sql:    tcq.sql.Clone(),
		path:   tcq.path,
		unique: tcq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CacheKey string `json:"cache_key,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TGoCache.Query().
//		GroupBy(tgocache.FieldCacheKey).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tcq *TGoCacheQuery) GroupBy(field string, fields ...string) *TGoCacheGroupBy {
	grbuild := &TGoCacheGroupBy{config: tcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tcq.sqlQuery(ctx), nil
	}
	grbuild.label = tgocache.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CacheKey string `json:"cache_key,omitempty"`
//	}
//
//	client.TGoCache.Query().
//		Select(tgocache.FieldCacheKey).
//		Scan(ctx, &v)
//
func (tcq *TGoCacheQuery) Select(fields ...string) *TGoCacheSelect {
	tcq.fields = append(tcq.fields, fields...)
	selbuild := &TGoCacheSelect{TGoCacheQuery: tcq}
	selbuild.label = tgocache.Label
	selbuild.flds, selbuild.scan = &tcq.fields, selbuild.Scan
	return selbuild
}

func (tcq *TGoCacheQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tcq.fields {
		if !tgocache.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tcq.path != nil {
		prev, err := tcq.path(ctx)
		if err != nil {
			return err
		}
		tcq.sql = prev
	}
	return nil
}

func (tcq *TGoCacheQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TGoCache, error) {
	var (
		nodes = []*TGoCache{}
		_spec = tcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*TGoCache).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &TGoCache{config: tcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tcq *TGoCacheQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tcq.querySpec()
	_spec.Node.Columns = tcq.fields
	if len(tcq.fields) > 0 {
		_spec.Unique = tcq.unique != nil && *tcq.unique
	}
	return sqlgraph.CountNodes(ctx, tcq.driver, _spec)
}

func (tcq *TGoCacheQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tcq *TGoCacheQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tgocache.Table,
			Columns: tgocache.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: tgocache.FieldID,
			},
		},
		From:   tcq.sql,
		Unique: true,
	}
	if unique := tcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tgocache.FieldID)
		for i := range fields {
			if fields[i] != tgocache.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tcq *TGoCacheQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tcq.driver.Dialect())
	t1 := builder.Table(tgocache.Table)
	columns := tcq.fields
	if len(columns) == 0 {
		columns = tgocache.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tcq.sql != nil {
		selector = tcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tcq.unique != nil && *tcq.unique {
		selector.Distinct()
	}
	for _, p := range tcq.predicates {
		p(selector)
	}
	for _, p := range tcq.order {
		p(selector)
	}
	if offset := tcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TGoCacheGroupBy is the group-by builder for TGoCache entities.
type TGoCacheGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tcgb *TGoCacheGroupBy) Aggregate(fns ...AggregateFunc) *TGoCacheGroupBy {
	tcgb.fns = append(tcgb.fns, fns...)
	return tcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tcgb *TGoCacheGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tcgb.path(ctx)
	if err != nil {
		return err
	}
	tcgb.sql = query
	return tcgb.sqlScan(ctx, v)
}

func (tcgb *TGoCacheGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tcgb.fields {
		if !tgocache.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tcgb *TGoCacheGroupBy) sqlQuery() *sql.Selector {
	selector := tcgb.sql.Select()
	aggregation := make([]string, 0, len(tcgb.fns))
	for _, fn := range tcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tcgb.fields)+len(tcgb.fns))
		for _, f := range tcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tcgb.fields...)...)
}

// TGoCacheSelect is the builder for selecting fields of TGoCache entities.
type TGoCacheSelect struct {
	*TGoCacheQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tcs *TGoCacheSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tcs.prepareQuery(ctx); err != nil {
		return err
	}
	tcs.sql = tcs.TGoCacheQuery.sqlQuery(ctx)
	return tcs.sqlScan(ctx, v)
}

func (tcs *TGoCacheSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tcs.sql.Query()
	if err := tcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
