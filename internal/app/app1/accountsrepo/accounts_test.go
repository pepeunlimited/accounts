package accountsrepo

import (
	"context"
	"github.com/pepeunlimited/accounts/internal/app/app1/mysql"
	"testing"
	"time"
)

func TestAccountsMySQL_Create(t *testing.T) {
	ctx := context.TODO()
	client := mysql.NewEntClient()
	accounts := NewAccountsRepository(client)
	accounts.DeleteAll(ctx)
	userId := int64(1)
	account, err := accounts.CreateCoinAccount(ctx, userId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if account.Balance != 0 {
		t.FailNow()
	}
	//_, err = accounts.Create(ctx, "name2", userId)
	//if err == nil {
	//	t.FailNow()
	//}
	//if err != ErrUserAccountExist {
	//	t.FailNow()
	//}
}

func TestAccountsMySQL_Transfer(t *testing.T) {
	ctx := context.TODO()
	client := mysql.NewEntClient()
	accounts := NewAccountsRepository(client)
	accounts.DeleteAll(ctx)

	fromUserId := int64(1)
	fromAmount := int64(50)

	toUserId := int64(2)
	toAmount := int64(40)

	fromAccount, err := accounts.CreateCoinAccount(ctx, fromUserId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	accounts.DoDeposit(ctx, fromAmount, fromAccount.ID, fromUserId)
	toAccount, err := accounts.CreateCashAccount(ctx, toUserId)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	tx, err := accounts.Transfer(ctx, -fromAmount, fromAccount.ID, fromUserId, toAccount.ID, toUserId, toAmount)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	tx.Commit()
}



func TestAccountsMySQL_TransferLowBalance(t *testing.T) {
	ctx := context.TODO()
	client := mysql.NewEntClient()
	accounts := NewAccountsRepository(client)
	accounts.DeleteAll(ctx)

	deposit := int64(40)
	fromUserId := int64(1)
	fromAmount := int64(50)

	toUserId := int64(2)
	toAmount := int64(40)

	fromAccount, err := accounts.CreateCoinAccount(ctx, fromUserId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	accounts.DoDeposit(ctx, deposit, fromAccount.ID, fromUserId)
	toAccount, err := accounts.CreateCashAccount(ctx, toUserId)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	tx, err := accounts.Transfer(ctx, -fromAmount, fromAccount.ID, fromUserId, toAccount.ID, toUserId, toAmount)
	if err == nil {
		t.FailNow()
	}
	if tx != nil {
		t.FailNow()
	}
	if err != ErrLowAccountBalance {
		t.FailNow()
	}
}


func TestAccountsMySQL_DoTransfer(t *testing.T) {
	ctx := context.TODO()
	client := mysql.NewEntClient()
	accounts := NewAccountsRepository(client)
	accounts.DeleteAll(ctx)

	fromUserId := int64(1)
	fromAmount := int64(50)

	toUserId := int64(2)
	toAmount := int64(40)
	fromAccount, err := accounts.CreateCoinAccount(ctx, fromUserId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	accounts.DoDeposit(ctx, int64(50000), fromAccount.ID, fromUserId)
	toAccount, err := accounts.CreateCashAccount(ctx, toUserId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for i := 0; i < 250; i++ {
		tx, err := accounts.Transfer(ctx, -fromAmount, fromAccount.ID, fromUserId, toAccount.ID, toUserId, toAmount)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		tx.Commit()
	}
}


func TestAccountsMySQL_DoTransferOCC(t *testing.T) {
	ctx := context.TODO()
	client := mysql.NewEntClient()
	accounts := NewAccountsRepository(client)
	accounts.DeleteAll(ctx)

	fromUserId := int64(1)
	fromAmount := int64(50)

	toUserId := int64(2)
	toAmount := int64(40)
	fromAccount, err := accounts.CreateCoinAccount(ctx, fromUserId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	accounts.DoDeposit(ctx, int64(50000), fromAccount.ID, fromUserId)
	toAccount, err := accounts.CreateCashAccount(ctx, toUserId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for i := 0; i < 100; i++ {
		go accounts.DoTransfer(ctx, -fromAmount, fromAccount.ID, fromUserId, toAccount.ID, toUserId, toAmount)
	}
	time.Sleep(3 * time.Second)
}

func TestAccountsMySQL_WithdrawOcc(t *testing.T) {
	ctx := context.TODO()
	client := mysql.NewEntClient()
	accounts := NewAccountsRepository(client)
	accounts.DeleteAll(ctx)
	fromUserId := int64(1)
	fromAccount, err := accounts.CreateCashAccount(ctx, fromUserId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	accounts.DoDeposit(ctx, int64(200000), fromAccount.ID, fromUserId)
	withdrawAmount := int64(100)
	for i := 0; i < 100; i++ {
		go accounts.DoWithdraw(ctx, -withdrawAmount, fromAccount.ID, fromUserId)
	}
	time.Sleep(3 * time.Second)
}


func TestAccountsMySQL_Withdraw(t *testing.T) {
	ctx := context.TODO()
	client := mysql.NewEntClient()
	accounts := NewAccountsRepository(client)
	accounts.DeleteAll(ctx)

	fromUserId := int64(1)
	fromAccount, err := accounts.CreateCashAccount(ctx, fromUserId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	accounts.DoDeposit(ctx, int64(200), fromAccount.ID, fromUserId)
	withdrawAmount := int64(100)
	err = accounts.DoWithdraw(ctx, -withdrawAmount, fromAccount.ID, fromUserId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestAccountsMySQL_UpdateBalanceLowBalance(t *testing.T) {
	ctx := context.TODO()
	client := mysql.NewEntClient()
	accounts := NewAccountsRepository(client)
	accounts.DeleteAll(ctx)


	userId := int64(1)
	account, err := accounts.CreateCashAccount(ctx, userId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	tx, err := accounts.Deposit(ctx, 10, account.ID, userId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	tx.Commit()

	tx, err = accounts.Deposit(ctx, 10, account.ID, userId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	tx.Commit()

	tx, err = accounts.Withdraw(ctx, -30, account.ID, userId)
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
	client := mysql.NewEntClient()
	accounts := NewAccountsRepository(client)
	accounts.DeleteAll(ctx)

	userId := int64(1)
	account, err := accounts.CreateCoinAccount(ctx, userId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for i := 0; i < 256; i++ {
		err := accounts.DoDeposit(ctx, 10,  account.ID, userId)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
}

func TestAccountsMySQL_UpdateBalanceVersionOccTest(t *testing.T) {
	ctx := context.TODO()
	client := mysql.NewEntClient()
	accounts := NewAccountsRepository(client)
	accounts.DeleteAll(ctx)

	userId := int64(1)
	account, err := accounts.CreateCoinAccount(ctx, userId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for i := 0; i < 3000; i++ {
		go accounts.DoDeposit(ctx, 10, account.ID, userId)
	}
	time.Sleep(8 * time.Second)
}

func TestAccountsMySQL_GetAccountByUserAndAccountID(t *testing.T) {
	ctx := context.TODO()
	client := mysql.NewEntClient()
	accounts := NewAccountsRepository(client)
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
	client := mysql.NewEntClient()
	accounts := NewAccountsRepository(client)
	accounts.DeleteAll(ctx)

	userId := int64(2)
	created, err := accounts.CreateCoinAccount(ctx, userId)
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
