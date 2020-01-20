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
	ac := accountsrepo.AccountTypeFromString(params.AccountType)
	if ac == accountsrepo.Unknown {
		return nil, twirp.InvalidArgumentError("account_type", params.AccountType)
	}
	return &ac, nil
}

func (v AccountServerValidator) GetAccount(params *rpcaccount.GetAccountParams) error {
	if params.AccountId == 0 {
		return twirp.RequiredArgumentError("account_id")
	}
	return nil
}

func (v AccountServerValidator) GetAccounts(params *rpcaccount.GetAccountsParams) (*accountsrepo.AccountType, error) {
	if params.AccountType == nil || validator2.IsEmpty(params.AccountType.Value) {
		return nil, nil
	}
	ac := accountsrepo.AccountTypeFromString(params.AccountType.Value)
	if ac == accountsrepo.Unknown {
		return nil, twirp.InvalidArgumentError("account_type", params.AccountType.Value)
	}
	return &ac, nil
}


func NewAccountServerValidator() AccountServerValidator {
	return AccountServerValidator{}
}

