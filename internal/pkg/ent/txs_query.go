// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/accounts/internal/pkg/ent/account"
	"github.com/pepeunlimited/accounts/internal/pkg/ent/predicate"
	"github.com/pepeunlimited/accounts/internal/pkg/ent/txs"
)

// TxsQuery is the builder for querying Txs entities.
type TxsQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Txs
	// eager-loading edges.
	withAccounts *AccountQuery
	withFKs      bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (tq *TxsQuery) Where(ps ...predicate.Txs) *TxsQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TxsQuery) Limit(limit int) *TxsQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TxsQuery) Offset(offset int) *TxsQuery {
	tq.offset = &offset
	return tq
}

// Order adds an order step to the query.
func (tq *TxsQuery) Order(o ...Order) *TxsQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryAccounts chains the current query on the accounts edge.
func (tq *TxsQuery) QueryAccounts() *AccountQuery {
	query := &AccountQuery{config: tq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(txs.Table, txs.FieldID, tq.sqlQuery()),
		sqlgraph.To(account.Table, account.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, txs.AccountsTable, txs.AccountsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
	return query
}

// First returns the first Txs entity in the query. Returns *NotFoundError when no txs was found.
func (tq *TxsQuery) First(ctx context.Context) (*Txs, error) {
	ts, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ts) == 0 {
		return nil, &NotFoundError{txs.Label}
	}
	return ts[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TxsQuery) FirstX(ctx context.Context) *Txs {
	t, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return t
}

// FirstID returns the first Txs id in the query. Returns *NotFoundError when no id was found.
func (tq *TxsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{txs.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (tq *TxsQuery) FirstXID(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Txs entity in the query, returns an error if not exactly one entity was returned.
func (tq *TxsQuery) Only(ctx context.Context) (*Txs, error) {
	ts, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ts) {
	case 1:
		return ts[0], nil
	case 0:
		return nil, &NotFoundError{txs.Label}
	default:
		return nil, &NotSingularError{txs.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TxsQuery) OnlyX(ctx context.Context) *Txs {
	t, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// OnlyID returns the only Txs id in the query, returns an error if not exactly one id was returned.
func (tq *TxsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{txs.Label}
	default:
		err = &NotSingularError{txs.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (tq *TxsQuery) OnlyXID(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TxsSlice.
func (tq *TxsQuery) All(ctx context.Context) ([]*Txs, error) {
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TxsQuery) AllX(ctx context.Context) []*Txs {
	ts, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ts
}

// IDs executes the query and returns a list of Txs ids.
func (tq *TxsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tq.Select(txs.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TxsQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TxsQuery) Count(ctx context.Context) (int, error) {
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TxsQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TxsQuery) Exist(ctx context.Context) (bool, error) {
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TxsQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TxsQuery) Clone() *TxsQuery {
	return &TxsQuery{
		config:     tq.config,
		limit:      tq.limit,
		offset:     tq.offset,
		order:      append([]Order{}, tq.order...),
		unique:     append([]string{}, tq.unique...),
		predicates: append([]predicate.Txs{}, tq.predicates...),
		// clone intermediate query.
		sql: tq.sql.Clone(),
	}
}

//  WithAccounts tells the query-builder to eager-loads the nodes that are connected to
// the "accounts" edge. The optional arguments used to configure the query builder of the edge.
func (tq *TxsQuery) WithAccounts(opts ...func(*AccountQuery)) *TxsQuery {
	query := &AccountQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withAccounts = query
	return tq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TxType string `json:"tx_type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Txs.Query().
//		GroupBy(txs.FieldTxType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tq *TxsQuery) GroupBy(field string, fields ...string) *TxsGroupBy {
	group := &TxsGroupBy{config: tq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = tq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		TxType string `json:"tx_type,omitempty"`
//	}
//
//	client.Txs.Query().
//		Select(txs.FieldTxType).
//		Scan(ctx, &v)
//
func (tq *TxsQuery) Select(field string, fields ...string) *TxsSelect {
	selector := &TxsSelect{config: tq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = tq.sqlQuery()
	return selector
}

func (tq *TxsQuery) sqlAll(ctx context.Context) ([]*Txs, error) {
	var (
		nodes       = []*Txs{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withAccounts != nil,
		}
	)
	if tq.withAccounts != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, txs.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Txs{config: tq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tq.withAccounts; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Txs)
		for i := range nodes {
			if fk := nodes[i].account_txs; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(account.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "account_txs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Accounts = n
			}
		}
	}

	return nodes, nil
}

func (tq *TxsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TxsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (tq *TxsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   txs.Table,
			Columns: txs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: txs.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TxsQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(txs.Table)
	selector := builder.Select(t1.Columns(txs.Columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(txs.Columns...)...)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TxsGroupBy is the builder for group-by Txs entities.
type TxsGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TxsGroupBy) Aggregate(fns ...Aggregate) *TxsGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scan the result into the given value.
func (tgb *TxsGroupBy) Scan(ctx context.Context, v interface{}) error {
	return tgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tgb *TxsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (tgb *TxsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TxsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tgb *TxsGroupBy) StringsX(ctx context.Context) []string {
	v, err := tgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (tgb *TxsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TxsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tgb *TxsGroupBy) IntsX(ctx context.Context) []int {
	v, err := tgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (tgb *TxsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TxsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tgb *TxsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (tgb *TxsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TxsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tgb *TxsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tgb *TxsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tgb.sqlQuery().Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TxsGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql
	columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
	columns = append(columns, tgb.fields...)
	for _, fn := range tgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(tgb.fields...)
}

// TxsSelect is the builder for select fields of Txs entities.
type TxsSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ts *TxsSelect) Scan(ctx context.Context, v interface{}) error {
	return ts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ts *TxsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ts *TxsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TxsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ts *TxsSelect) StringsX(ctx context.Context) []string {
	v, err := ts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ts *TxsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TxsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ts *TxsSelect) IntsX(ctx context.Context) []int {
	v, err := ts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ts *TxsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TxsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ts *TxsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ts *TxsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TxsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ts *TxsSelect) BoolsX(ctx context.Context) []bool {
	v, err := ts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ts *TxsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sqlQuery().Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ts *TxsSelect) sqlQuery() sql.Querier {
	selector := ts.sql
	selector.Select(selector.Columns(ts.fields...)...)
	return selector
}
