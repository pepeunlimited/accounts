package validator

import (
	"github.com/pepeunlimited/accounts/internal/app/app1/accountsrepo"
	"github.com/pepeunlimited/accounts/rpcaccount"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
)

type AccountServerValidator struct {}

func (v AccountServerValidator) CreateAccount(params *rpcaccount.CreateAccountParams) (*accountsrepo.AccountType, error) {
	if validator2.IsEmpty(params.AccountType) {
		return nil, twirp.RequiredArgumentError("account_type")
	}
	ac, err := v.accountType(params.AccountType)
	if err != nil {
		return nil, err
	}
	if params.UserId == 0 {
		return nil, twirp.RequiredArgumentError("user_id")
	}

	return ac, nil
}

func (v AccountServerValidator) GetAccount(params *rpcaccount.GetAccountParams) error {
	if params.AccountId == 0 {
		return twirp.RequiredArgumentError("account_id")
	}
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (v AccountServerValidator) GetAccounts(params *rpcaccount.GetAccountsParams) (*accountsrepo.AccountType, error) {
	if params.UserId == 0 {
		return nil, twirp.RequiredArgumentError("user_id")
	}
	if params.AccountType == nil || validator2.IsEmpty(params.AccountType.Value) {
		return nil, nil
	}
	ac, err := v.accountType(params.AccountType.Value)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

func (v AccountServerValidator) CreateDeposit(params *rpcaccount.CreateDepositParams) (*accountsrepo.AccountType, error) {
	if params.Amount < 0 {
		return nil, twirp.InvalidArgumentError("amount","amount < 0")
	}
	if params.UserId == 0 {
		return nil, twirp.RequiredArgumentError("user_id")
	}
	if params.Amount == 0 {
		return nil, twirp.RequiredArgumentError("amount")
	}
	ac, err := v.accountType(params.AccountType)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

func (v AccountServerValidator) accountType(accountType string) (*accountsrepo.AccountType, error) {
	ac := accountsrepo.AccountTypeFromString(accountType)
	if ac == accountsrepo.Unknown {
		return nil, twirp.InvalidArgumentError("account_type", accountType)
	}
	return &ac, nil
}

func (v AccountServerValidator) CreateWithdraw(params *rpcaccount.CreateWithdrawParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	if params.Amount > 0 {
		return twirp.InvalidArgumentError("amount","amount > 0")
	}
	return nil
}

func (v AccountServerValidator) CreateTransfer(params *rpcaccount.CreateTransferParams) error {
	if params.FromAmount > 0 {
		return twirp.InvalidArgumentError("from_amount","amount > 0")
	}
	if params.FromUserId == 0 {
		return twirp.RequiredArgumentError("from_user_id")
	}
	if params.ToUserId == 0 {
		return twirp.RequiredArgumentError("to_user_id")
	}
	if params.ToAmount < 0 {
		return twirp.InvalidArgumentError("to_amount","amount < 0")
	}
	return nil
}

func NewAccountServerValidator() AccountServerValidator {
	return AccountServerValidator{}
}

