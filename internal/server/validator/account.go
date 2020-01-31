package validator

import (
	"github.com/pepeunlimited/accounts/pkg/accountsrpc"
	"github.com/twitchtv/twirp"
)

type AccountServerValidator struct {}

func (v AccountServerValidator) CreateAccount(params *accountsrpc.CreateAccountParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (v AccountServerValidator) GetAccount(params *accountsrpc.GetAccountParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (v AccountServerValidator) CreateDeposit(params *accountsrpc.CreateDepositParams) error {
	if params.Amount < 0 {
		return twirp.InvalidArgumentError("amount","amount < 0")
	}
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (v AccountServerValidator) CreateWithdraw(params *accountsrpc.CreateWithdrawParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	if params.Amount > 0 {
		return twirp.InvalidArgumentError("amount","amount > 0")
	}
	return nil
}

func (v AccountServerValidator) UpdateAccountVerified(params *accountsrpc.UpdateAccountVerifiedParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func NewAccountServerValidator() AccountServerValidator {
	return AccountServerValidator{}
}