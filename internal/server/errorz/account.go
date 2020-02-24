package errorz

import (
	"github.com/pepeunlimited/accounts/internal/pkg/mysql/account"
	"github.com/pepeunlimited/accounts/pkg/accounts"
	"github.com/twitchtv/twirp"
	"log"
)

func Account(err error) error {
	switch err {
	case account.ErrAccountNotExist:
		return twirp.NotFoundError(accounts.AccountNotFound)
	case account.ErrUserAccountExist:
		return twirp.NewError(twirp.AlreadyExists, accounts.AccountExist)
	case account.ErrInvalidAmount:
		return twirp.NewError(twirp.Aborted, accounts.AccountInvalidAmount)
	case account.ErrLowAccountBalance:
		return twirp.NewError(twirp.Aborted, accounts.LowAccountBalance)
	}
	log.Print("accounts-service: unknown error: "+err.Error())
	return twirp.InternalErrorWith(err)
}