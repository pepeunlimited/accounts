// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
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

// AccountQuery is the builder for querying Account entities.
type AccountQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Account
	// eager-loading edges.
	withTxs *TxsQuery
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (aq *AccountQuery) Where(ps ...predicate.Account) *AccountQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AccountQuery) Limit(limit int) *AccountQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AccountQuery) Offset(offset int) *AccountQuery {
	aq.offset = &offset
	return aq
}

// Order adds an order step to the query.
func (aq *AccountQuery) Order(o ...Order) *AccountQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryTxs chains the current query on the txs edge.
func (aq *AccountQuery) QueryTxs() *TxsQuery {
	query := &TxsQuery{config: aq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(account.Table, account.FieldID, aq.sqlQuery()),
		sqlgraph.To(txs.Table, txs.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, account.TxsTable, account.TxsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
	return query
}

// First returns the first Account entity in the query. Returns *NotFoundError when no account was found.
func (aq *AccountQuery) First(ctx context.Context) (*Account, error) {
	as, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(as) == 0 {
		return nil, &NotFoundError{account.Label}
	}
	return as[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AccountQuery) FirstX(ctx context.Context) *Account {
	a, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return a
}

// FirstID returns the first Account id in the query. Returns *NotFoundError when no id was found.
func (aq *AccountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{account.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (aq *AccountQuery) FirstXID(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Account entity in the query, returns an error if not exactly one entity was returned.
func (aq *AccountQuery) Only(ctx context.Context) (*Account, error) {
	as, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(as) {
	case 1:
		return as[0], nil
	case 0:
		return nil, &NotFoundError{account.Label}
	default:
		return nil, &NotSingularError{account.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AccountQuery) OnlyX(ctx context.Context) *Account {
	a, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// OnlyID returns the only Account id in the query, returns an error if not exactly one id was returned.
func (aq *AccountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{account.Label}
	default:
		err = &NotSingularError{account.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (aq *AccountQuery) OnlyXID(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Accounts.
func (aq *AccountQuery) All(ctx context.Context) ([]*Account, error) {
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AccountQuery) AllX(ctx context.Context) []*Account {
	as, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return as
}

// IDs executes the query and returns a list of Account ids.
func (aq *AccountQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aq.Select(account.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AccountQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AccountQuery) Count(ctx context.Context) (int, error) {
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AccountQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AccountQuery) Exist(ctx context.Context) (bool, error) {
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AccountQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AccountQuery) Clone() *AccountQuery {
	return &AccountQuery{
		config:     aq.config,
		limit:      aq.limit,
		offset:     aq.offset,
		order:      append([]Order{}, aq.order...),
		unique:     append([]string{}, aq.unique...),
		predicates: append([]predicate.Account{}, aq.predicates...),
		// clone intermediate query.
		sql: aq.sql.Clone(),
	}
}

//  WithTxs tells the query-builder to eager-loads the nodes that are connected to
// the "txs" edge. The optional arguments used to configure the query builder of the edge.
func (aq *AccountQuery) WithTxs(opts ...func(*TxsQuery)) *AccountQuery {
	query := &TxsQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withTxs = query
	return aq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Balance int64 `json:"balance,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Account.Query().
//		GroupBy(account.FieldBalance).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aq *AccountQuery) GroupBy(field string, fields ...string) *AccountGroupBy {
	group := &AccountGroupBy{config: aq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = aq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Balance int64 `json:"balance,omitempty"`
//	}
//
//	client.Account.Query().
//		Select(account.FieldBalance).
//		Scan(ctx, &v)
//
func (aq *AccountQuery) Select(field string, fields ...string) *AccountSelect {
	selector := &AccountSelect{config: aq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = aq.sqlQuery()
	return selector
}

func (aq *AccountQuery) sqlAll(ctx context.Context) ([]*Account, error) {
	var (
		nodes       = []*Account{}
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withTxs != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Account{config: aq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
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
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aq.withTxs; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Account)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Txs(func(s *sql.Selector) {
			s.Where(sql.InValues(account.TxsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.account_txs
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "account_txs" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "account_txs" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Txs = append(node.Edges.Txs, n)
		}
	}

	return nodes, nil
}

func (aq *AccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AccountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (aq *AccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AccountQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(account.Table)
	selector := builder.Select(t1.Columns(account.Columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(account.Columns...)...)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AccountGroupBy is the builder for group-by Account entities.
type AccountGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AccountGroupBy) Aggregate(fns ...Aggregate) *AccountGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scan the result into the given value.
func (agb *AccountGroupBy) Scan(ctx context.Context, v interface{}) error {
	return agb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (agb *AccountGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := agb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (agb *AccountGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AccountGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (agb *AccountGroupBy) StringsX(ctx context.Context) []string {
	v, err := agb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (agb *AccountGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AccountGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (agb *AccountGroupBy) IntsX(ctx context.Context) []int {
	v, err := agb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (agb *AccountGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AccountGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (agb *AccountGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := agb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (agb *AccountGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AccountGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (agb *AccountGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := agb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (agb *AccountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := agb.sqlQuery().Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AccountGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql
	columns := make([]string, 0, len(agb.fields)+len(agb.fns))
	columns = append(columns, agb.fields...)
	for _, fn := range agb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(agb.fields...)
}

// AccountSelect is the builder for select fields of Account entities.
type AccountSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (as *AccountSelect) Scan(ctx context.Context, v interface{}) error {
	return as.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (as *AccountSelect) ScanX(ctx context.Context, v interface{}) {
	if err := as.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (as *AccountSelect) Strings(ctx context.Context) ([]string, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AccountSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (as *AccountSelect) StringsX(ctx context.Context) []string {
	v, err := as.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (as *AccountSelect) Ints(ctx context.Context) ([]int, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AccountSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (as *AccountSelect) IntsX(ctx context.Context) []int {
	v, err := as.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (as *AccountSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AccountSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (as *AccountSelect) Float64sX(ctx context.Context) []float64 {
	v, err := as.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (as *AccountSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AccountSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (as *AccountSelect) BoolsX(ctx context.Context) []bool {
	v, err := as.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (as *AccountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sqlQuery().Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (as *AccountSelect) sqlQuery() sql.Querier {
	selector := as.sql
	selector.Select(selector.Columns(as.fields...)...)
	return selector
}
