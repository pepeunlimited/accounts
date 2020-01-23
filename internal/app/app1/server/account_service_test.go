package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pepeunlimited/accounts/accountsrpc"
	"github.com/pepeunlimited/accounts/internal/app/app1/accountsrepo"
	"github.com/pepeunlimited/accounts/internal/app/app1/mysql"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/twitchtv/twirp"
	"testing"
)

func TestAccountServer_CreateAccountAndGet(t *testing.T) {
	ctx := context.TODO()
	userId := int64(1)

	server := NewAccountServer(mysql.NewEntClient())
	server.accounts.DeleteAll(ctx)
	coin, err := server.CreateAccount(ctx, &accountsrpc.CreateAccountParams{
		AccountType: "coin",
		UserId:userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if coin == nil {
		t.FailNow()
	}
	cash, err := server.CreateAccount(ctx, &accountsrpc.CreateAccountParams{
		AccountType: "cash",
		UserId:userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if cash == nil {
		t.FailNow()
	}
	if _, err := server.accounts.GetAccountByUserIDAndType(ctx, userId, accountsrepo.Cash); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := server.accounts.GetAccountByUserIDAndType(ctx, userId, accountsrepo.Coin); err != nil {
		t.Error(err)
		t.FailNow()
	}

	account, err := server.GetAccount(ctx, &accountsrpc.GetAccountParams{
		AccountId: coin.Id,
		UserId:userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if account.Id != coin.Id {
		t.FailNow()
	}
}

func TestAccountServer_NotFound(t *testing.T) {
	ctx := context.TODO()
	userId := int64(1)
	ctx = rpcz.AddUserId(userId)
	server := NewAccountServer(mysql.NewEntClient())
	server.accounts.DeleteAll(ctx)

	account, err := server.GetAccount(ctx, &accountsrpc.GetAccountParams{
		AccountId: 333333333333,
		UserId:userId,
	})
	if err == nil {
		t.FailNow()
	}
	if account != nil {
		t.FailNow()
	}
	if !accountsrpc.IsReason(err.(twirp.Error), accountsrpc.AccountNotFound) {
		t.Error(err.(twirp.Error).Meta(accountsrpc.AccountNotFound))
		t.FailNow()
	}
}

func TestAccountServer_GetAccounts(t *testing.T) {
	ctx := context.TODO()
	userId := int64(1)

	server := NewAccountServer(mysql.NewEntClient())
	server.accounts.DeleteAll(ctx)
	resp0, err := server.GetAccounts(ctx, &accountsrpc.GetAccountsParams{UserId: userId})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(resp0.Accounts) != 0 {
		t.FailNow()
	}
	server.CreateAccount(ctx, &accountsrpc.CreateAccountParams{
		AccountType: "coin",
		UserId:userId,
	})
	server.CreateAccount(ctx, &accountsrpc.CreateAccountParams{
		AccountType: "cash",
		UserId:userId,
	})
	resp1, err := server.GetAccounts(ctx,&accountsrpc.GetAccountsParams{UserId: userId,})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(resp1.Accounts) != 2 {
		t.FailNow()
	}
	resp2, err := server.GetAccounts(ctx, &accountsrpc.GetAccountsParams{
		AccountType: &wrappers.StringValue{
			Value: "coin",
		},
		UserId:userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(resp2.Accounts) != 1 {
		t.FailNow()
	}
	if resp2.Accounts[0].Type != "COIN" {
		t.Log(resp2)
		t.FailNow()
	}
	resp3, err := server.GetAccounts(ctx, &accountsrpc.GetAccountsParams{
		AccountType: &wrappers.StringValue{
			Value: "cash",
		},
		UserId:userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(resp3.Accounts) != 1 {
		t.FailNow()
	}
	if resp3.Accounts[0].Type != "CASH" {
		t.Log(resp3)
		t.FailNow()
	}
}

func TestAccountServer_CreateDeposit(t *testing.T) {
	ctx := context.TODO()
	userId := int64(1)

	server := NewAccountServer(mysql.NewEntClient())
	server.accounts.DeleteAll(ctx)

	coin, err := server.CreateAccount(ctx, &accountsrpc.CreateAccountParams{
		AccountType: "Coin",
		UserId:      userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	server.accounts.DoDeposit(ctx, 10, int(coin.Id), userId)
	account, err := server.CreateDeposit(ctx, &accountsrpc.CreateDepositParams{
		UserId:      userId,
		Amount:      10,
		AccountType: "Coin",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if account.Balance != 20 {
		t.FailNow()
	}
}

func TestAccountServer_CreateWithdraw(t *testing.T) {
	ctx := context.TODO()
	userId := int64(1)

	server := NewAccountServer(mysql.NewEntClient())
	server.accounts.DeleteAll(ctx)

	_, err := server.CreateAccount(ctx, &accountsrpc.CreateAccountParams{
		AccountType: "Cash",
		UserId:      userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	server.CreateDeposit(ctx, &accountsrpc.CreateDepositParams{
		UserId:      userId,
		Amount:      20,
		AccountType: "Cash",
	})

	withdrawed, err := server.CreateWithdraw(ctx, &accountsrpc.CreateWithdrawParams{
		UserId: userId,
		Amount: -20,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if withdrawed.Balance != 0 {
		t.FailNow()
	}
}

func TestAccountServer_CreateTransfer(t *testing.T) {
	ctx := context.TODO()
	toUserID := int64(1)
	fromUserID := int64(2)
	server := NewAccountServer(mysql.NewEntClient())
	server.accounts.DeleteAll(ctx)

	server.CreateAccount(ctx, &accountsrpc.CreateAccountParams{
		AccountType: "cash",
		UserId:      toUserID,
	})
	server.CreateAccount(ctx, &accountsrpc.CreateAccountParams{
		AccountType: "coin",
		UserId:      fromUserID,
	})
	server.CreateDeposit(ctx, &accountsrpc.CreateDepositParams{
		UserId: 	 fromUserID,
		Amount:      200,
		AccountType: "coin",
	})
	referenceNumber := "reference-number"
	transfer, err := server.CreateTransfer(ctx, &accountsrpc.CreateTransferParams{
		FromUserId:      fromUserID,
		FromAmount:      -200,
		ToUserId:        toUserID,
		ToAmount:        100,
		ReferenceNumber: &wrappers.StringValue{Value: referenceNumber},
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if transfer == nil {
		t.FailNow()
	}
}