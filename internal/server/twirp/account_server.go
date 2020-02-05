package twirp

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"github.com/pepeunlimited/accounts/internal/pkg/ent"
	"github.com/pepeunlimited/accounts/internal/pkg/mysql/accountrepo"
	"github.com/pepeunlimited/accounts/internal/server/validator"
	"github.com/pepeunlimited/accounts/pkg/accountsrpc"
	"github.com/twitchtv/twirp"
	"log"
)

type AccountServer struct {
	accounts  accountrepo.AccountRepository
	validator validator.AccountServerValidator
}

func (server AccountServer) UpdateAccountVerified(ctx context.Context, params *accountsrpc.UpdateAccountVerifiedParams) (*accountsrpc.Account, error) {
	err := server.validator.UpdateAccountVerified(params)
	if err != nil {
		return nil, err
	}
	verified, err := server.accounts.UpdateAccountVerified(ctx, params.UserId)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return toAccountRPC(verified), nil
}

func (server AccountServer) CreateDeposit(ctx context.Context, params *accountsrpc.CreateDepositParams) (*accountsrpc.Account, error) {
	err := server.validator.CreateDeposit(params)
	if err != nil {
		return nil, err
	}
	account, err := server.accounts.GetAccountByUserID(ctx, params.UserId)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	tx, err := server.accounts.Deposit(ctx, params.Amount, account.ID, params.UserId, server.referenceNumber(params.ReferenceNumber))
	if err != nil {
		return nil, server.isAccountError(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Print("accounts-service: deposit commit failure: "+err.Error())
		return nil, twirp.NewError(twirp.Aborted, accountsrpc.AccountTXsCommit)
	}
	deposited, err := server.accounts.GetAccountByUserID(ctx, params.UserId)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return toAccountRPC(deposited), nil
}

func (server AccountServer) CreateWithdraw(ctx context.Context, params *accountsrpc.CreateWithdrawParams) (*accountsrpc.Account, error) {
	err := server.validator.CreateWithdraw(params)
	if err != nil {
		return nil, err
	}
	account, err := server.accounts.GetAccountByUserID(ctx, params.UserId)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	if !account.IsVerified {
		return nil, twirp.NewError(twirp.Aborted, accountsrpc.AccountIsNotVerified)
	}
	referenceNumber := uuid.New().String()
	tx, err := server.accounts.Withdraw(ctx, params.Amount, account.ID, params.UserId, &referenceNumber)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Print("accounts-service: withdraw commit failure: "+err.Error())
		return nil, twirp.NewError(twirp.Aborted, accountsrpc.AccountTXsCommit)
	}
	withdrawn, err := server.accounts.GetAccountByUserID(ctx, params.UserId)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return toAccountRPC(withdrawn), nil
}

func (server AccountServer) referenceNumber(referenceNumber *wrappers.StringValue) *string {
	if referenceNumber == nil {
		return nil
	}
	return  &referenceNumber.Value
}

func (server AccountServer) CreateAccount(ctx context.Context, params *accountsrpc.CreateAccountParams) (*accountsrpc.Account, error) {
	err := server.validator.CreateAccount(params)
	if err != nil {
		return nil, err
	}
	account, err := server.accounts.CreateAccount(ctx, params.UserId)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return toAccountRPC(account), nil
}

func (server AccountServer) GetAccount(ctx context.Context, params *accountsrpc.GetAccountParams) (*accountsrpc.Account, error) {
	err := server.validator.GetAccount(params)
	if err != nil {
		return nil, err
	}
	account, err := server.accounts.GetAccountByUserID(ctx, params.UserId)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return toAccountRPC(account), nil
}

func (server AccountServer) isAccountError(err error) error {
	switch err {
	case accountrepo.ErrAccountNotExist:
		return twirp.NotFoundError(accountsrpc.AccountNotFound)
	case accountrepo.ErrUserAccountExist:
		return twirp.NewError(twirp.AlreadyExists, accountsrpc.AccountExist)
	case accountrepo.ErrInvalidAmount:
		return twirp.NewError(twirp.Aborted, accountsrpc.AccountInvalidAmount)
	case accountrepo.ErrLowAccountBalance:
		return twirp.NewError(twirp.Aborted, accountsrpc.LowAccountBalance)
	}
	log.Print("accounts-service: unknown error: "+err.Error())
	return twirp.InternalErrorWith(err)
}

func NewAccountServer(client *ent.Client) AccountServer {
	return AccountServer{
		accounts:  accountrepo.NewAccountRepository(client),
		validator: validator.NewAccountServerValidator(),
	}
}