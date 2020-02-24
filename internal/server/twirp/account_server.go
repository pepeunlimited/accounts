package twirp

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"github.com/pepeunlimited/accounts/internal/pkg/ent"
	accountrepo "github.com/pepeunlimited/accounts/internal/pkg/mysql/account"
	"github.com/pepeunlimited/accounts/internal/server/errorz"
	"github.com/pepeunlimited/accounts/internal/server/validator"
	"github.com/pepeunlimited/accounts/pkg/account"
	"github.com/twitchtv/twirp"
	"log"
)

type AccountServer struct {
	accounts  accountrepo.AccountRepository
	validator validator.AccountServerValidator
}

func (server AccountServer) UpdateAccountVerified(ctx context.Context, params *account.UpdateAccountVerifiedParams) (*account.Account, error) {
	err := server.validator.UpdateAccountVerified(params)
	if err != nil {
		return nil, err
	}
	verified, err := server.accounts.UpdateAccountVerified(ctx, params.UserId)
	if err != nil {
		return nil, errorz.Account(err)
	}
	return toAccountRPC(verified), nil
}

func (server AccountServer) CreateDeposit(ctx context.Context, params *account.CreateDepositParams) (*account.Account, error) {
	err := server.validator.CreateDeposit(params)
	if err != nil {
		return nil, err
	}
	fromDB, err := server.accounts.GetAccountByUserID(ctx, params.UserId)
	if err != nil {
		return nil, errorz.Account(err)
	}
	tx, err := server.accounts.Deposit(ctx, params.Amount, fromDB.ID, params.UserId, server.referenceNumber(params.ReferenceNumber))
	if err != nil {
		return nil, errorz.Account(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Print("accounts-service: deposit commit failure: "+err.Error())
		return nil, twirp.NewError(twirp.Aborted, account.AccountTXsCommit)
	}
	deposited, err := server.accounts.GetAccountByUserID(ctx, params.UserId)
	if err != nil {
		return nil, errorz.Account(err)
	}
	return toAccountRPC(deposited), nil
}

func (server AccountServer) CreateWithdraw(ctx context.Context, params *account.CreateWithdrawParams) (*account.Account, error) {
	err := server.validator.CreateWithdraw(params)
	if err != nil {
		return nil, err
	}
	fromDB, err := server.accounts.GetAccountByUserID(ctx, params.UserId)
	if err != nil {
		return nil, errorz.Account(err)
	}
	if !fromDB.IsVerified {
		return nil, twirp.NewError(twirp.Aborted, account.AccountIsNotVerified)
	}
	referenceNumber := uuid.New().String()
	tx, err := server.accounts.Withdraw(ctx, params.Amount, fromDB.ID, params.UserId, &referenceNumber)
	if err != nil {
		return nil, errorz.Account(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Print("accounts-service: withdraw commit failure: "+err.Error())
		return nil, twirp.NewError(twirp.Aborted, account.AccountTXsCommit)
	}
	withdrawn, err := server.accounts.GetAccountByUserID(ctx, params.UserId)
	if err != nil {
		return nil, errorz.Account(err)
	}
	return toAccountRPC(withdrawn), nil
}

func (server AccountServer) referenceNumber(referenceNumber *wrappers.StringValue) *string {
	if referenceNumber == nil {
		return nil
	}
	return  &referenceNumber.Value
}

func (server AccountServer) CreateAccount(ctx context.Context, params *account.CreateAccountParams) (*account.Account, error) {
	err := server.validator.CreateAccount(params)
	if err != nil {
		return nil, err
	}
	account, err := server.accounts.CreateAccount(ctx, params.UserId)
	if err != nil {
		return nil, errorz.Account(err)
	}
	return toAccountRPC(account), nil
}

func (server AccountServer) GetAccount(ctx context.Context, params *account.GetAccountParams) (*account.Account, error) {
	err := server.validator.GetAccount(params)
	if err != nil {
		return nil, err
	}
	account, err := server.accounts.GetAccountByUserID(ctx, params.UserId)
	if err != nil {
		return nil, errorz.Account(err)
	}
	return toAccountRPC(account), nil
}

func NewAccountServer(client *ent.Client) AccountServer {
	return AccountServer{
		accounts:  accountrepo.New(client),
		validator: validator.NewAccountServerValidator(),
	}
}