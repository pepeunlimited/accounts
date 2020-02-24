package twirp

import (
	"github.com/pepeunlimited/accounts/internal/pkg/ent"
	"github.com/pepeunlimited/accounts/pkg/account"
)

func toAccountRPC(from *ent.Account) *account.Account {
	return &account.Account{
		Balance: from.Balance,
		UserId:  from.UserID,
		Id:   int64(from.ID),
	}
}
