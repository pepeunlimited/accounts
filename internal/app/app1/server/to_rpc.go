package server

import (
	"github.com/pepeunlimited/accounts/internal/app/app1/ent"
	"github.com/pepeunlimited/accounts/accountsrpc"
)

func toAccountRPC(account *ent.Accounts) *accountsrpc.Account {
	return &accountsrpc.Account{
		Balance: account.Balance,
		UserId:  account.UserID,
		Id:   int64(account.ID),
	}
}
