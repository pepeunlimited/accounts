package server

import (
	"github.com/pepeunlimited/accounts/internal/app/app1/ent"
	"github.com/pepeunlimited/accounts/rpcaccount"
)

func toAccountRPC(account *ent.Accounts) *rpcaccount.Account {
	return &rpcaccount.Account{
		Balance: account.Balance,
		Type:    account.Type,
		UserId:  account.UserID,
		Id:   int64(account.ID),
	}
}
