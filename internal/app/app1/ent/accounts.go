// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pepeunlimited/accounts/internal/app/app1/ent/accounts"
)

// Accounts is the model entity for the Accounts schema.
type Accounts struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance int64 `json:"balance,omitempty"`
	// Version holds the value of the "version" field.
	Version uint8 `json:"version,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// IsWithdrawable holds the value of the "is_withdrawable" field.
	IsWithdrawable bool `json:"is_withdrawable,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountsQuery when eager-loading is set.
	Edges struct {
		// Txs holds the value of the txs edge.
		Txs []*Txs
	} `json:"edges"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Accounts) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // balance
		&sql.NullInt64{},  // version
		&sql.NullString{}, // type
		&sql.NullBool{},   // is_withdrawable
		&sql.NullInt64{},  // user_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Accounts fields.
func (a *Accounts) assignValues(values ...interface{}) error {
	if m, n := len(values), len(accounts.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field balance", values[0])
	} else if value.Valid {
		a.Balance = value.Int64
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field version", values[1])
	} else if value.Valid {
		a.Version = uint8(value.Int64)
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[2])
	} else if value.Valid {
		a.Type = value.String
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field is_withdrawable", values[3])
	} else if value.Valid {
		a.IsWithdrawable = value.Bool
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field user_id", values[4])
	} else if value.Valid {
		a.UserID = value.Int64
	}
	return nil
}

// QueryTxs queries the txs edge of the Accounts.
func (a *Accounts) QueryTxs() *TxsQuery {
	return (&AccountsClient{a.config}).QueryTxs(a)
}

// Update returns a builder for updating this Accounts.
// Note that, you need to call Accounts.Unwrap() before calling this method, if this Accounts
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Accounts) Update() *AccountsUpdateOne {
	return (&AccountsClient{a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Accounts) Unwrap() *Accounts {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Accounts is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Accounts) String() string {
	var builder strings.Builder
	builder.WriteString("Accounts(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", balance=")
	builder.WriteString(fmt.Sprintf("%v", a.Balance))
	builder.WriteString(", version=")
	builder.WriteString(fmt.Sprintf("%v", a.Version))
	builder.WriteString(", type=")
	builder.WriteString(a.Type)
	builder.WriteString(", is_withdrawable=")
	builder.WriteString(fmt.Sprintf("%v", a.IsWithdrawable))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", a.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// AccountsSlice is a parsable slice of Accounts.
type AccountsSlice []*Accounts

func (a AccountsSlice) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
