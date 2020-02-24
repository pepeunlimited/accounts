package account

import (
	"context"
	"github.com/pepeunlimited/accounts/internal/pkg/ent"
	"testing"
	"time"
)

func TestAccountsMySQL_Create(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	accounts := New(client)
	accounts.DeleteAll(ctx)
	userId := int64(1)
	coin, err := accounts.CreateAccount(ctx, userId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if coin.Balance != 0 {
		t.FailNow()
	}
	_, err = accounts.CreateAccount(ctx, userId)
	if err == nil {
		t.FailNow()
	}
	if err != ErrUserAccountExist {
		t.FailNow()
	}
}

func TestAccountsMySQL_WithdrawOcc(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	accounts := New(client)
	accounts.DeleteAll(ctx)
	fromUserId := int64(1)
	fromAccount, err := accounts.CreateAccount(ctx, fromUserId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	accounts.DoDeposit(ctx, int64(200000), fromAccount.ID, fromUserId, nil)
	withdrawAmount := int64(100)
	for i := 0; i < 100; i++ {
		go accounts.DoWithdraw(ctx, -withdrawAmount, fromAccount.ID, fromUserId, nil)
	}
	time.Sleep(3 * time.Second)
}


func TestAccountsMySQL_Withdraw(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	accounts := New(client)
	accounts.DeleteAll(ctx)

	fromUserId := int64(1)
	fromAccount, err := accounts.CreateAccount(ctx, fromUserId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	accounts.DoDeposit(ctx, int64(200), fromAccount.ID, fromUserId, nil)
	withdrawAmount := int64(100)
	err = accounts.DoWithdraw(ctx, -withdrawAmount, fromAccount.ID, fromUserId, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestAccountsMySQL_UpdateBalanceLowBalance(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	accounts := New(client)
	accounts.DeleteAll(ctx)


	userId := int64(1)
	account, err := accounts.CreateAccount(ctx, userId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	tx, err := accounts.Deposit(ctx, 10, account.ID, userId, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	tx.Commit()

	tx, err = accounts.Deposit(ctx, 10, account.ID, userId, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	tx.Commit()

	tx, err = accounts.Withdraw(ctx, -30, account.ID, userId,nil)
	if err == nil {
		t.FailNow()
	}
	if err != ErrLowAccountBalance {
		t.FailNow()
	}
	if tx != nil {
		t.FailNow()
	}
}

func TestAccountsMySQL_UpdateBalanceVersionOver255(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	accounts := New(client)
	accounts.DeleteAll(ctx)

	userId := int64(1)
	account, err := accounts.CreateAccount(ctx, userId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for i := 0; i < 256; i++ {
		err := accounts.DoDeposit(ctx, 10,  account.ID, userId, nil)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
}

func TestAccountsMySQL_UpdateBalanceVersionOccTest(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	accounts := New(client)
	accounts.DeleteAll(ctx)

	userId := int64(1)
	account, err := accounts.CreateAccount(ctx, userId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for i := 0; i < 10; i++ {
		go accounts.DoDeposit(ctx, 10, account.ID, userId, nil)
	}
	time.Sleep(8 * time.Second)
}

func TestAccountsMySQL_GetAccountByUserAndAccountID(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	accounts := New(client)
	accounts.DeleteAll(ctx)

	_, err := accounts.GetAccountByUserAndAccountID(ctx, 1, 1)
	if err == nil {
		t.FailNow()
	}
	if err != ErrAccountNotExist {
		t.FailNow()
	}
}

func TestAccountsMySQL_GetAccountByUserAndAccountID2(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	accounts := New(client)
	accounts.DeleteAll(ctx)

	userId := int64(2)
	created, err := accounts.CreateAccount(ctx, userId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	query, err := accounts.GetAccountByUserAndAccountID(ctx, userId, created.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if created.ID != query.ID {
		t.FailNow()
	}
}