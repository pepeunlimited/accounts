package server

import (
	"github.com/pepeunlimited/accounts/internal/app/app1/ent"
	"github.com/pepeunlimited/accounts/accountsrpc"
)

func toAccountRPC(account *ent.Accounts) *accountsrpc.Account {
	return &accountsrpc.Account{
		Balance: account.Balance,
		Type:    account.Type,
		UserId:  account.UserID,
		Id:   int64(account.ID),
	}
}

func toAccountsRPC(accounts []*ent.Accounts) *accountsrpc.GetAccountsResponse {
	toRPCs := make([]*accountsrpc.Account, 0)
	for _, account := range accounts {
		toRPCs = append(toRPCs, toAccountRPC(account))
	}
	return &accountsrpc.GetAccountsResponse{
		Accounts: toRPCs,
	}
}
