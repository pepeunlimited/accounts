package errorz

import (
	accountrepo "github.com/pepeunlimited/accounts/internal/pkg/mysql/account"
	"github.com/pepeunlimited/accounts/pkg/rpc/account"
	"github.com/twitchtv/twirp"
	"log"
)

func Account(err error) error {
	switch err {
	case accountrepo.ErrAccountNotExist:
		return twirp.NotFoundError(account.AccountNotFound)
	case accountrepo.ErrUserAccountExist:
		return twirp.NewError(twirp.AlreadyExists, account.AccountExist)
	case accountrepo.ErrInvalidAmount:
		return twirp.NewError(twirp.Aborted, account.AccountInvalidAmount)
	case accountrepo.ErrLowAccountBalance:
		return twirp.NewError(twirp.Aborted, account.LowAccountBalance)
	}
	log.Print("accounts-service: unknown error: "+err.Error())
	return twirp.InternalErrorWith(err)
}