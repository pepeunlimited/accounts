package twirp

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pepeunlimited/accounts/internal/pkg/ent"
	"github.com/pepeunlimited/accounts/pkg/accountsrpc"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/twitchtv/twirp"
	"testing"
)

func TestAccountServer_CreateAccountAndGet(t *testing.T) {
	ctx := context.TODO()
	userId := int64(1)

	server := NewAccountServer(ent.NewEntClient())
	server.accounts.DeleteAll(ctx)
	coin, err := server.CreateAccount(ctx, &accountsrpc.CreateAccountParams{
		UserId:userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if coin == nil {
		t.FailNow()
	}
	account, err := server.GetAccount(ctx, &accountsrpc.GetAccountParams{
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
	server := NewAccountServer(ent.NewEntClient())
	server.accounts.DeleteAll(ctx)

	account, err := server.GetAccount(ctx, &accountsrpc.GetAccountParams{
		UserId:userId,
	})
	if err == nil {
		t.FailNow()
	}
	if account != nil {
		t.FailNow()
	}
	if err.(twirp.Error).Msg() != accountsrpc.AccountNotFound {
		t.Error(err.(twirp.Error).Meta(accountsrpc.AccountNotFound))
		t.FailNow()
	}
}

func TestAccountServer_CreateDeposit(t *testing.T) {
	ctx := context.TODO()
	userId := int64(1)

	server := NewAccountServer(ent.NewEntClient())
	server.accounts.DeleteAll(ctx)

	coin, err := server.CreateAccount(ctx, &accountsrpc.CreateAccountParams{
		UserId:      userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	server.accounts.DoDeposit(ctx, 10, int(coin.Id), userId, nil)
	account, err := server.CreateDeposit(ctx, &accountsrpc.CreateDepositParams{
		UserId:      userId,
		Amount:      10,
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

	server := NewAccountServer(ent.NewEntClient())
	server.accounts.DeleteAll(ctx)

	_, err := server.CreateAccount(ctx, &accountsrpc.CreateAccountParams{
		UserId:      userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	server.CreateDeposit(ctx, &accountsrpc.CreateDepositParams{
		UserId:          userId,
		Amount:          20,
		ReferenceNumber: &wrappers.StringValue{Value: "reference-number-0"},
	})
	_, err = server.UpdateAccountVerified(ctx, &accountsrpc.UpdateAccountVerifiedParams{
		UserId: userId,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

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