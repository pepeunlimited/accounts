package accountsrpc

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pepeunlimited/microservice-kit/errorz"
)

type AccountsMock struct {
	Errors 			errorz.Stack
	Account     	*Account
	ReferenceNumber *wrappers.StringValue
	IsVerified		bool
}

func (a *AccountsMock) CreateDeposit(ctx context.Context, params *CreateDepositParams) (*Account, error) {
	a.ReferenceNumber = params.ReferenceNumber
	if a.Errors.IsEmpty() {
		a.Account.Balance += params.Amount
	}
	return nil, a.Errors.Pop()
}

func (a *AccountsMock) CreateWithdraw(ctx context.Context, params *CreateWithdrawParams) (*Account, error) {
	if a.Errors.IsEmpty() {
		if !a.IsVerified {
			return nil, fmt.Errorf("not verified")
		}
		a.Account.Balance += params.Amount
		return a.Account, nil
	}
	return nil, a.Errors.Pop()
}

func (a *AccountsMock) CreateAccount(context.Context, *CreateAccountParams) (*Account, error) {
	panic("implement me")
}

func (a *AccountsMock) GetAccount(context.Context, *GetAccountParams) (*Account, error) {
	panic("implement me")
}

func NewAccountsMock(errors []error, account *Account) *AccountsMock {
	mock := &AccountsMock{
		Errors:  errorz.NewErrorStack(errors),
	}
	if account == nil {
		mock.Account = mock.account(1)
		mock.IsVerified = false
	}
	return mock
}

func (a AccountsMock) account(id int64) *Account {
	return &Account{
		Balance: 100,
		UserId:  1,
		Id:      id,
	}
}