// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/accounts/internal/app/app1/ent/accounts"
	"github.com/pepeunlimited/accounts/internal/app/app1/ent/txs"
)

// AccountsCreate is the builder for creating a Accounts entity.
type AccountsCreate struct {
	config
	balance         *int64
	version         *uint8
	_type           *string
	is_withdrawable *bool
	user_id         *int64
	txs             map[int]struct{}
}

// SetBalance sets the balance field.
func (ac *AccountsCreate) SetBalance(i int64) *AccountsCreate {
	ac.balance = &i
	return ac
}

// SetVersion sets the version field.
func (ac *AccountsCreate) SetVersion(u uint8) *AccountsCreate {
	ac.version = &u
	return ac
}

// SetType sets the type field.
func (ac *AccountsCreate) SetType(s string) *AccountsCreate {
	ac._type = &s
	return ac
}

// SetIsWithdrawable sets the is_withdrawable field.
func (ac *AccountsCreate) SetIsWithdrawable(b bool) *AccountsCreate {
	ac.is_withdrawable = &b
	return ac
}

// SetUserID sets the user_id field.
func (ac *AccountsCreate) SetUserID(i int64) *AccountsCreate {
	ac.user_id = &i
	return ac
}

// AddTxIDs adds the txs edge to Txs by ids.
func (ac *AccountsCreate) AddTxIDs(ids ...int) *AccountsCreate {
	if ac.txs == nil {
		ac.txs = make(map[int]struct{})
	}
	for i := range ids {
		ac.txs[ids[i]] = struct{}{}
	}
	return ac
}

// AddTxs adds the txs edges to Txs.
func (ac *AccountsCreate) AddTxs(t ...*Txs) *AccountsCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTxIDs(ids...)
}

// Save creates the Accounts in the database.
func (ac *AccountsCreate) Save(ctx context.Context) (*Accounts, error) {
	if ac.balance == nil {
		return nil, errors.New("ent: missing required field \"balance\"")
	}
	if ac.version == nil {
		return nil, errors.New("ent: missing required field \"version\"")
	}
	if ac._type == nil {
		return nil, errors.New("ent: missing required field \"type\"")
	}
	if err := accounts.TypeValidator(*ac._type); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"type\": %v", err)
	}
	if ac.is_withdrawable == nil {
		return nil, errors.New("ent: missing required field \"is_withdrawable\"")
	}
	if ac.user_id == nil {
		return nil, errors.New("ent: missing required field \"user_id\"")
	}
	return ac.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountsCreate) SaveX(ctx context.Context) *Accounts {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *AccountsCreate) sqlSave(ctx context.Context) (*Accounts, error) {
	var (
		a     = &Accounts{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: accounts.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: accounts.FieldID,
			},
		}
	)
	if value := ac.balance; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: accounts.FieldBalance,
		})
		a.Balance = *value
	}
	if value := ac.version; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  *value,
			Column: accounts.FieldVersion,
		})
		a.Version = *value
	}
	if value := ac._type; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: accounts.FieldType,
		})
		a.Type = *value
	}
	if value := ac.is_withdrawable; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: accounts.FieldIsWithdrawable,
		})
		a.IsWithdrawable = *value
	}
	if value := ac.user_id; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: accounts.FieldUserID,
		})
		a.UserID = *value
	}
	if nodes := ac.txs; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   accounts.TxsTable,
			Columns: []string{accounts.TxsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: txs.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	a.ID = int(id)
	return a, nil
}
