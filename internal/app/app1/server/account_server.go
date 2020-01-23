package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pepeunlimited/accounts/internal/app/app1/accountsrepo"
	"github.com/pepeunlimited/accounts/internal/app/app1/ent"
	"github.com/pepeunlimited/accounts/internal/app/app1/validator"
	"github.com/pepeunlimited/accounts/accountsrpc"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/twitchtv/twirp"
	"log"
)

type AccountServer struct {
	accounts accountsrepo.AccountsRepository
	validator validator.AccountServerValidator
}

func (server AccountServer) CreateDeposit(ctx context.Context, params *accountsrpc.CreateDepositParams) (*accountsrpc.Account, error) {
	accountType, err := server.validator.CreateDeposit(params)
	if err != nil {
		return nil, err
	}
	account, err := server.accounts.GetAccountByUserIDAndType(ctx, params.UserId, *accountType)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	tx, err := server.accounts.Deposit(ctx, params.Amount, account.ID, params.UserId)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Print("accounts-service: deposit commit failure: "+err.Error())
		return nil, twirp.NewError(twirp.Aborted, accountsrpc.AccountTXsCommit)
	}
	deposited, err := server.accounts.GetAccountByUserIDAndType(ctx, params.UserId, *accountType)
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
	account, err := server.accounts.GetAccountByUserIDAndType(ctx, params.UserId, accountsrepo.Cash)
	if err != nil {
		return nil, server.isAccountError(err)
	}

	if !account.IsWithdrawable {
		return nil, twirp.NewError(twirp.Aborted, accountsrpc.WithdrawIsDisabled)
	}

	tx, err := server.accounts.Withdraw(ctx, params.Amount, account.ID, params.UserId)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Print("accounts-service: withdraw commit failure: "+err.Error())
		return nil, twirp.NewError(twirp.Aborted, accountsrpc.AccountTXsCommit)
	}
	deposited, err := server.accounts.GetAccountByUserIDAndType(ctx, params.UserId, accountsrepo.Cash)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return toAccountRPC(deposited), nil
}

func (server AccountServer) CreateTransfer(ctx context.Context, params *accountsrpc.CreateTransferParams) (*accountsrpc.CreateTransferResponse, error) {
	err := server.validator.CreateTransfer(params)
	if err != nil {
		return nil, err
	}
	fromAccount, err := server.accounts.GetAccountByUserIDAndType(ctx, params.FromUserId, accountsrepo.Coin)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	toAccount, err := server.accounts.GetAccountByUserIDAndType(ctx, params.ToUserId, accountsrepo.Cash)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	referenceNumber := server.referenceNumber(params.ReferenceNumber)
	tx, err := server.accounts.Transfer(ctx, params.FromAmount, fromAccount.ID, params.FromUserId, toAccount.ID, params.ToUserId, params.ToAmount, referenceNumber)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Print("accounts-service: transfer commit failure: "+err.Error())
		return nil, twirp.NewError(twirp.Aborted, accountsrpc.AccountTXsCommit)
	}
	from, err := server.accounts.GetAccountByUserIDAndType(ctx, params.FromUserId, accountsrepo.Coin)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	to, err := server.accounts.GetAccountByUserIDAndType(ctx, params.ToUserId, accountsrepo.Cash)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return &accountsrpc.CreateTransferResponse{
		From: toAccountRPC(from),
		To:   toAccountRPC(to),
	}, nil
}

func (server AccountServer) referenceNumber(referenceNumber *wrappers.StringValue) *string {
	if referenceNumber == nil {
		return nil
	}
	return  &referenceNumber.Value
}

func (server AccountServer) CreateAccount(ctx context.Context, params *accountsrpc.CreateAccountParams) (*accountsrpc.Account, error) {
	ac, err := server.validator.CreateAccount(params)
	if err != nil {
		return nil, err
	}
	var account *ent.Accounts
	switch *ac {
	case accountsrepo.Coin:
		account, err = server.accounts.CreateCoinAccount(ctx, params.UserId)
	case accountsrepo.Cash:
		account, err = server.accounts.CreateCashAccount(ctx, params.UserId)
	}
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return toAccountRPC(account), nil
}

func (server AccountServer) GetAccounts(ctx context.Context, params *accountsrpc.GetAccountsParams) (*accountsrpc.GetAccountsResponse, error) {
	ac, err := server.validator.GetAccounts(params)
	if err != nil {
		return nil, err
	}
	accounts, err := server.accounts.GetAccountsByUserID(ctx, params.UserId, ac)
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return toAccountsRPC(accounts), nil
}

func (server AccountServer) GetAccount(ctx context.Context, params *accountsrpc.GetAccountParams) (*accountsrpc.Account, error) {
	err := server.validator.GetAccount(params)
	if err != nil {
		return nil, err
	}
	account, err := server.accounts.GetAccountByUserAndAccountID(ctx, params.UserId, int(params.AccountId))
	if err != nil {
		return nil, server.isAccountError(err)
	}
	return toAccountRPC(account), nil
}

func (server AccountServer) isAccountError(err error) error {
	switch err {
	case accountsrepo.ErrAccountNotExist:
		return twirp.NotFoundError(accountsrpc.AccountNotFound).WithMeta(rpcz.Reason, accountsrpc.AccountNotFound)
	case accountsrepo.ErrUserAccountExist:
		return twirp.NewError(twirp.AlreadyExists, err.Error()).WithMeta(rpcz.Reason, accountsrpc.AccountExist)
	case accountsrepo.ErrInvalidAmount:
		return twirp.NewError(twirp.Aborted, err.Error()).WithMeta(rpcz.Reason, accountsrpc.AccountInvalidAmount)
	case accountsrepo.ErrLowAccountBalance:
		return twirp.NewError(twirp.Aborted, err.Error()).WithMeta(rpcz.Reason, accountsrpc.LowAccountBalance)
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