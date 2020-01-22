package server

import (
	"github.com/pepeunlimited/accounts/internal/app/app1/ent"
	"github.com/pepeunlimited/accounts/accountrpc"
)

func toAccountRPC(account *ent.Accounts) *accountrpc.Account {
	return &accountrpc.Account{
		Balance: account.Balance,
		Type:    account.Type,
		UserId:  account.UserID,
		Id:   int64(account.ID),
	}
}

func toAccountsRPC(accounts []*ent.Accounts) *accountrpc.GetAccountsResponse {
	toRPCs := make([]*accountrpc.Account, 0)
	for _, account := range accounts {
		toRPCs = append(toRPCs, toAccountRPC(account))
	}
	return &accountrpc.GetAccountsResponse{
		Accounts: toRPCs,
	}
}
