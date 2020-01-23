package accountsrpc

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/errorz"
	"strings"
)

type AccountsMock struct {
	Errors 			errorz.Stack
	Coin     		*Account
	Cash     		*Account
}

func (a *AccountsMock) CreateDeposit(ctx context.Context, params *CreateDepositParams) (*Account, error) {
	if a.Errors.IsEmpty() {
		if strings.ToLower(params.AccountType) == "coin" {
			a.Coin.Balance += params.Amount
			return a.Coin, nil
		}
		if strings.ToLower(params.AccountType) == "cash" {
			a.Cash.Balance += params.Amount
			return a.Cash, nil
		}
	}
	return nil, a.Errors.Pop()
}

func (a *AccountsMock) CreateWithdraw(context.Context, *CreateWithdrawParams) (*Account, error) {
	panic("implement me")
}

func (a *AccountsMock) CreateTransfer(context.Context, *CreateTransferParams) (*CreateTransferResponse, error) {
	panic("implement me")
}

func (a *AccountsMock) CreateAccount(context.Context, *CreateAccountParams) (*Account, error) {
	panic("implement me")
}

func (a *AccountsMock) GetAccounts(context.Context, *GetAccountsParams) (*GetAccountsResponse, error) {
	panic("implement me")
}

func (a *AccountsMock) GetAccount(context.Context, *GetAccountParams) (*Account, error) {
	panic("implement me")
}

func NewAccountsMock(errors []error, coin *Account, cash *Account) *AccountsMock {
	mock := &AccountsMock{
		Errors:  errorz.NewErrorStack(errors),
	}
	if coin == nil {
		mock.Coin = mock.account("COIN", 1)
	}
	if cash == nil {
		mock.Coin = mock.account("CASH", 2)
	}
	return mock
}

func (a AccountsMock) account(types string, id int64) *Account {
	return &Account{
		Balance: 100,
		Type:    types,
		UserId:  1,
		Id:      id,
	}
}