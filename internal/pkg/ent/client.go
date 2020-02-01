// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"github.com/pepeunlimited/accounts/internal/pkg/ent/accounts"
	"github.com/pepeunlimited/accounts/internal/pkg/ent/migrate"
	"github.com/pepeunlimited/accounts/internal/pkg/ent/txs"
	"log"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Accounts is the client for interacting with the Accounts builders.
	Accounts *AccountsClient
	// Txs is the client for interacting with the Txs builders.
	Txs *TxsClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config:   c,
		Schema:   migrate.NewSchema(c.driver),
		Accounts: NewAccountsClient(c),
		Txs:      NewTxsClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config:   cfg,
		Accounts: NewAccountsClient(cfg),
		Txs:      NewTxsClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Accounts.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config:   cfg,
		Schema:   migrate.NewSchema(cfg.driver),
		Accounts: NewAccountsClient(cfg),
		Txs:      NewTxsClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// AccountsClient is a client for the Accounts schema.
type AccountsClient struct {
	config
}

// NewAccountsClient returns a client for the Accounts from the given config.
func NewAccountsClient(c config) *AccountsClient {
	return &AccountsClient{config: c}
}

// Create returns a create builder for Accounts.
func (c *AccountsClient) Create() *AccountsCreate {
	return &AccountsCreate{config: c.config}
}

// Update returns an update builder for Accounts.
func (c *AccountsClient) Update() *AccountsUpdate {
	return &AccountsUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountsClient) UpdateOne(a *Accounts) *AccountsUpdateOne {
	return c.UpdateOneID(a.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountsClient) UpdateOneID(id int) *AccountsUpdateOne {
	return &AccountsUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Accounts.
func (c *AccountsClient) Delete() *AccountsDelete {
	return &AccountsDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AccountsClient) DeleteOne(a *Accounts) *AccountsDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AccountsClient) DeleteOneID(id int) *AccountsDeleteOne {
	return &AccountsDeleteOne{c.Delete().Where(accounts.ID(id))}
}

// Create returns a query builder for Accounts.
func (c *AccountsClient) Query() *AccountsQuery {
	return &AccountsQuery{config: c.config}
}

// Get returns a Accounts entity by its id.
func (c *AccountsClient) Get(ctx context.Context, id int) (*Accounts, error) {
	return c.Query().Where(accounts.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountsClient) GetX(ctx context.Context, id int) *Accounts {
	a, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return a
}

// QueryTxs queries the txs edge of a Accounts.
func (c *AccountsClient) QueryTxs(a *Accounts) *TxsQuery {
	query := &TxsQuery{config: c.config}
	id := a.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(accounts.Table, accounts.FieldID, id),
		sqlgraph.To(txs.Table, txs.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, accounts.TxsTable, accounts.TxsColumn),
	)
	query.sql = sqlgraph.Neighbors(a.driver.Dialect(), step)

	return query
}

// TxsClient is a client for the Txs schema.
type TxsClient struct {
	config
}

// NewTxsClient returns a client for the Txs from the given config.
func NewTxsClient(c config) *TxsClient {
	return &TxsClient{config: c}
}

// Create returns a create builder for Txs.
func (c *TxsClient) Create() *TxsCreate {
	return &TxsCreate{config: c.config}
}

// Update returns an update builder for Txs.
func (c *TxsClient) Update() *TxsUpdate {
	return &TxsUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *TxsClient) UpdateOne(t *Txs) *TxsUpdateOne {
	return c.UpdateOneID(t.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *TxsClient) UpdateOneID(id int) *TxsUpdateOne {
	return &TxsUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Txs.
func (c *TxsClient) Delete() *TxsDelete {
	return &TxsDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TxsClient) DeleteOne(t *Txs) *TxsDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TxsClient) DeleteOneID(id int) *TxsDeleteOne {
	return &TxsDeleteOne{c.Delete().Where(txs.ID(id))}
}

// Create returns a query builder for Txs.
func (c *TxsClient) Query() *TxsQuery {
	return &TxsQuery{config: c.config}
}

// Get returns a Txs entity by its id.
func (c *TxsClient) Get(ctx context.Context, id int) (*Txs, error) {
	return c.Query().Where(txs.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TxsClient) GetX(ctx context.Context, id int) *Txs {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// QueryAccounts queries the accounts edge of a Txs.
func (c *TxsClient) QueryAccounts(t *Txs) *AccountsQuery {
	query := &AccountsQuery{config: c.config}
	id := t.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(txs.Table, txs.FieldID, id),
		sqlgraph.To(accounts.Table, accounts.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, txs.AccountsTable, txs.AccountsColumn),
	)
	query.sql = sqlgraph.Neighbors(t.driver.Dialect(), step)

	return query
}