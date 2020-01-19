package server

import (
	"context"
	"github.com/pepeunlimited/accounts/internal/app/app1/accountsrepo"
	"github.com/pepeunlimited/accounts/internal/app/app1/mysql"
	"github.com/pepeunlimited/accounts/rpcaccount"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/twitchtv/twirp"
	"testing"
)

func TestAccountServer_CreateAccountAndGet(t *testing.T) {
	ctx := context.TODO()
	userId := int64(1)
	ctx = rpcz.AddUserId(userId)
	server := NewAccountServer(mysql.NewEntClient())
	server.accounts.DeleteAll(ctx)
	coin, err := server.CreateAccount(ctx, &rpcaccount.CreateAccountParams{
		AccountType: "coin",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if coin == nil {
		t.FailNow()
	}
	cash, err := server.CreateAccount(ctx, &rpcaccount.CreateAccountParams{
		AccountType: "cash",
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

	account, err := server.GetAccount(ctx, &rpcaccount.GetAccountParams{
		AccountId: coin.Id,
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

	account, err := server.GetAccount(ctx, &rpcaccount.GetAccountParams{
		AccountId: 333333333333,
	})
	if err == nil {
		t.FailNow()
	}
	if account != nil {
		t.FailNow()
	}
	if !rpcaccount.IsReason(err.(twirp.Error), rpcaccount.AccountNotFound) {
		t.Error(err.(twirp.Error).Meta(rpcaccount.AccountNotFound))
		t.FailNow()
	}
}