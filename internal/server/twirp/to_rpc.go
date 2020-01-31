package twirp

import (
	"github.com/pepeunlimited/accounts/internal/pkg/ent"
	"github.com/pepeunlimited/accounts/pkg/accountsrpc"
)

func toAccountRPC(account *ent.Accounts) *accountsrpc.Account {
	return &accountsrpc.Account{
		Balance: account.Balance,
		UserId:  account.UserID,
		Id:   int64(account.ID),
	}
}
