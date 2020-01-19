// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pepeunlimited/accounts/internal/app/app1/ent/txs"
)

// Txs is the model entity for the Txs schema.
type Txs struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TxType holds the value of the "tx_type" field.
	TxType string `json:"tx_type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int64 `json:"amount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TxsQuery when eager-loading is set.
	Edges struct {
		// Accounts holds the value of the accounts edge.
		Accounts *Accounts
	}
	accounts_id *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Txs) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // tx_type
		&sql.NullTime{},   // created_at
		&sql.NullInt64{},  // amount
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Txs) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // accounts_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Txs fields.
func (t *Txs) assignValues(values ...interface{}) error {
	if m, n := len(values), len(txs.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field tx_type", values[0])
	} else if value.Valid {
		t.TxType = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[1])
	} else if value.Valid {
		t.CreatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field amount", values[2])
	} else if value.Valid {
		t.Amount = value.Int64
	}
	values = values[3:]
	if len(values) == len(txs.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field accounts_id", value)
		} else if value.Valid {
			t.accounts_id = new(int)
			*t.accounts_id = int(value.Int64)
		}
	}
	return nil
}

// QueryAccounts queries the accounts edge of the Txs.
func (t *Txs) QueryAccounts() *AccountsQuery {
	return (&TxsClient{t.config}).QueryAccounts(t)
}

// Update returns a builder for updating this Txs.
// Note that, you need to call Txs.Unwrap() before calling this method, if this Txs
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Txs) Update() *TxsUpdateOne {
	return (&TxsClient{t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Txs) Unwrap() *Txs {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Txs is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Txs) String() string {
	var builder strings.Builder
	builder.WriteString("Txs(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", tx_type=")
	builder.WriteString(t.TxType)
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", t.Amount))
	builder.WriteByte(')')
	return builder.String()
}

// TxsSlice is a parsable slice of Txs.
type TxsSlice []*Txs

func (t TxsSlice) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
