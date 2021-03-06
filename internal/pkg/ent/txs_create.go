// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/accounts/internal/pkg/ent/account"
	"github.com/pepeunlimited/accounts/internal/pkg/ent/txs"
)

// TxsCreate is the builder for creating a Txs entity.
type TxsCreate struct {
	config
	tx_type          *string
	created_at       *time.Time
	amount           *int64
	reference_number *string
	accounts         map[int]struct{}
}

// SetTxType sets the tx_type field.
func (tc *TxsCreate) SetTxType(s string) *TxsCreate {
	tc.tx_type = &s
	return tc
}

// SetCreatedAt sets the created_at field.
func (tc *TxsCreate) SetCreatedAt(t time.Time) *TxsCreate {
	tc.created_at = &t
	return tc
}

// SetAmount sets the amount field.
func (tc *TxsCreate) SetAmount(i int64) *TxsCreate {
	tc.amount = &i
	return tc
}

// SetReferenceNumber sets the reference_number field.
func (tc *TxsCreate) SetReferenceNumber(s string) *TxsCreate {
	tc.reference_number = &s
	return tc
}

// SetNillableReferenceNumber sets the reference_number field if the given value is not nil.
func (tc *TxsCreate) SetNillableReferenceNumber(s *string) *TxsCreate {
	if s != nil {
		tc.SetReferenceNumber(*s)
	}
	return tc
}

// SetAccountsID sets the accounts edge to Account by id.
func (tc *TxsCreate) SetAccountsID(id int) *TxsCreate {
	if tc.accounts == nil {
		tc.accounts = make(map[int]struct{})
	}
	tc.accounts[id] = struct{}{}
	return tc
}

// SetNillableAccountsID sets the accounts edge to Account by id if the given value is not nil.
func (tc *TxsCreate) SetNillableAccountsID(id *int) *TxsCreate {
	if id != nil {
		tc = tc.SetAccountsID(*id)
	}
	return tc
}

// SetAccounts sets the accounts edge to Account.
func (tc *TxsCreate) SetAccounts(a *Account) *TxsCreate {
	return tc.SetAccountsID(a.ID)
}

// Save creates the Txs in the database.
func (tc *TxsCreate) Save(ctx context.Context) (*Txs, error) {
	if tc.tx_type == nil {
		return nil, errors.New("ent: missing required field \"tx_type\"")
	}
	if err := txs.TxTypeValidator(*tc.tx_type); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"tx_type\": %v", err)
	}
	if tc.created_at == nil {
		return nil, errors.New("ent: missing required field \"created_at\"")
	}
	if tc.amount == nil {
		return nil, errors.New("ent: missing required field \"amount\"")
	}
	if tc.reference_number != nil {
		if err := txs.ReferenceNumberValidator(*tc.reference_number); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"reference_number\": %v", err)
		}
	}
	if len(tc.accounts) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"accounts\"")
	}
	return tc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TxsCreate) SaveX(ctx context.Context) *Txs {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TxsCreate) sqlSave(ctx context.Context) (*Txs, error) {
	var (
		t     = &Txs{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: txs.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: txs.FieldID,
			},
		}
	)
	if value := tc.tx_type; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: txs.FieldTxType,
		})
		t.TxType = *value
	}
	if value := tc.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: txs.FieldCreatedAt,
		})
		t.CreatedAt = *value
	}
	if value := tc.amount; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: txs.FieldAmount,
		})
		t.Amount = *value
	}
	if value := tc.reference_number; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: txs.FieldReferenceNumber,
		})
		t.ReferenceNumber = value
	}
	if nodes := tc.accounts; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   txs.AccountsTable,
			Columns: []string{txs.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}
