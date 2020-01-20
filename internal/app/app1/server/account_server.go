package server

import (
	"context"
	"github.com/pepeunlimited/accounts/internal/app/app1/accountsrepo"
	"github.com/pepeunlimited/accounts/internal/app/app1/ent"
	"github.com/pepeunlimited/accounts/internal/app/app1/validator"
	"github.com/pepeunlimited/accounts/rpcaccount"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/twitchtv/twirp"
	"log"
)

type AccountServer struct {
	accounts accountsrepo.AccountsRepository
	validator validator.AccountServerValidator
}

func (server AccountServer) CreateDeposit(context.Context, *rpcaccount.CreateDepositParams) (*rpcaccount.CreateDepositResponse, error) {
	panic("implement me")
}

func (server AccountServer) CreateWithdraw(context.Context, *rpcaccount.CreateWithdrawParams) (*rpcaccount.CreateWithdrawResponse, error) {
	panic("implement me")
}

func (server AccountServer) CreateTransfer(ctx context.Context, params *rpcaccount.CreateTransferParams) (*rpcaccount.CreateTransferResponse, error) {
	panic("implement me")
}

func (server AccountServer) CreateAccount(ctx context.Context, params *rpcaccount.CreateAccountParams) (*rpcaccount.Account, error) {
	userId, err := rpcz.GetUserId(ctx)
	if err != nil {
		return nil, twirp.RequiredArgumentError("user_id")
	}
	ac, err := server.validator.CreateAccount(params)
	if err != nil {
		return nil, err
	}
	var account *ent.Accounts
	switch *ac {
	case accountsrepo.Coin:
		account, err = server.accounts.CreateCoinAccount(ctx, userId)
	case accountsrepo.Cash:
		account, err = server.accounts.CreateCashAccount(ctx, userId)
	}
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return toAccountRPC(account), nil
}

func (server AccountServer) GetAccounts(ctx context.Context, params *rpcaccount.GetAccountsParams) (*rpcaccount.GetAccountsResponse, error) {
	userId, err := rpcz.GetUserId(ctx)
	if err != nil {
		return nil, twirp.RequiredArgumentError("user_id")
	}
	ac, err := server.validator.GetAccounts(params)
	if err != nil {
		return nil, err
	}
	accounts, err := server.accounts.GetAccountsByUserID(ctx, userId, ac)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return toAccountsRPC(accounts), nil
}

func (server AccountServer) GetAccount(ctx context.Context, params *rpcaccount.GetAccountParams) (*rpcaccount.Account, error) {
	userId, err := rpcz.GetUserId(ctx)
	if err != nil {
		return nil, twirp.RequiredArgumentError("user_id")
	}
	err = server.validator.GetAccount(params)
	if err != nil {
		return nil, err
	}
	account, err := server.accounts.GetAccountByUserAndAccountID(ctx, userId, int(params.AccountId))
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return toAccountRPC(account), nil
}

func (server AccountServer) isAccountError(err error) error {
	switch err {
	case accountsrepo.ErrAccountNotExist:
		return twirp.NotFoundError(rpcaccount.AccountNotFound).WithMeta(rpcz.Reason, rpcaccount.AccountNotFound)
	case accountsrepo.ErrUserAccountExist:
		return twirp.NewError(twirp.AlreadyExists, err.Error()).WithMeta(rpcz.Reason, rpcaccount.AccountExist)
	}
	log.Print("accounts: unknown error: "+err.Error())
	return twirp.NewError(twirp.Internal ,"unknown error: "+err.Error())
}

func NewAccountServer(client *ent.Client) AccountServer {
	return AccountServer{
		accounts:accountsrepo.NewAccountsRepository(client),
		validator: validator.NewAccountServerValidator(),
	}
}