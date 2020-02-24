package twirp

import (
	"github.com/pepeunlimited/accounts/internal/pkg/ent"
	"github.com/pepeunlimited/accounts/pkg/accounts"
)

func toAccountRPC(account *ent.Account) *accounts.Account {
	return &accounts.Account{
		Balance: account.Balance,
		UserId:  account.UserID,
		Id:   int64(account.ID),
	}
}
