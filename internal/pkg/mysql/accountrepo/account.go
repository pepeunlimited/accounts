package accountrepo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/pepeunlimited/accounts/internal/pkg/ent"
	"github.com/pepeunlimited/accounts/internal/pkg/ent/account"
	"github.com/pepeunlimited/microservice-kit/misc"
	"log"
	"strconv"
	"time"
)

const (
	updateBalanceSQL 	= "UPDATE accounts SET balance = ?, version = ? WHERE id = ? AND version = ?"
	createTXsSQL 		= "INSERT INTO txs (tx_type, created_at, amount, account_txs, reference_number) VALUES (?, ?, ?, ?, ?)"
)

var (
	ErrAccountNotExist 			= errors.New("accounts: account not exist")
	ErrUserAccountExist 		= errors.New("accounts: user account exist")
	ErrLowAccountBalance    	= errors.New("accounts: unable to process payment; low account balance")
	ErrOptimisticConcurrency 	= errors.New("accounts: unable to process payment; optimistic concurrency exception")
	ErrInvalidAmount 			= errors.New("accounts: invalid amount")
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, userId int64) (*ent.Account, error)

	GetAccountByUserID(ctx context.Context, userID int64) (*ent.Account, error)

	GetAccountByUserAndAccountID(ctx context.Context, userID int64, accountID int) (*ent.Account, error)

	Deposit(ctx context.Context, amount int64, toAccountID int, toUserID int64, referenceNumber *string) (*sql.Tx, error)
	Withdraw(ctx context.Context, withdrawAmount int64, fromAccountID int, fromUserID int64, referenceNumber *string) (*sql.Tx, error)

	DoWithdraw(ctx context.Context, withdrawAmount int64, fromAccountID int, fromUserID int64, referenceNumber *string) error
	DoDeposit(ctx context.Context, amount int64, toAccountID int, toUserID int64, referenceNumber *string) error

	UpdateAccountVerified(ctx context.Context, userID int64) (*ent.Account, error)

	DeleteAll(ctx context.Context)
}

type accountsMySQL struct {
	client *ent.Client
	isDebug bool
}

func (mysql accountsMySQL) UpdateAccountVerified(ctx context.Context, userID int64) (*ent.Account, error) {
	a, err := mysql.GetAccountByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	save, err := a.Update().SetIsVerified(true).Save(ctx)
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (mysql accountsMySQL) GetAccountByUserID(ctx context.Context, userID int64) (*ent.Account, error) {
	account, err := mysql.client.Account.Query().Where(account.UserID(userID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrAccountNotExist
		}
		return nil, err
	}
	return account, nil
}

func (mysql accountsMySQL) CreateAccount(ctx context.Context, userId int64) (*ent.Account, error) {
	return mysql.create(ctx, userId)
}

func (mysql accountsMySQL) DoWithdraw(ctx context.Context, withdrawAmount int64, fromCashAccountID int, fromUserID int64, referenceNumber *string) error {
	deposit, err := mysql.Withdraw(ctx, withdrawAmount, fromCashAccountID, fromUserID, referenceNumber)
	if err != nil {
		return err
	}
	return deposit.Commit()
}

func (mysql accountsMySQL) DoDeposit(ctx context.Context, amount int64, toAccountID int, toUserID int64, referenceNumber *string) error {
	deposit, err := mysql.Deposit(ctx, amount, toAccountID, toUserID, referenceNumber)
	if err != nil {
		return err
	}
	return deposit.Commit()
}

func (mysql accountsMySQL) Deposit(ctx context.Context, amount int64, toAccountID int, toUserID int64, referenceNumber *string) (*sql.Tx, error) {
	if amount < 0 {
		return nil, ErrInvalidAmount
	}
	toAccount, err := mysql.GetAccountByUserAndAccountID(ctx, toUserID, toAccountID)
	if err != nil {
		return nil, err
	}
	tx, err := mysql.client.DB().Begin()
	if err != nil {
		return nil, err
	}
	if err := mysql.updateBalance(ctx, tx, toUserID, toAccountID, amount, toAccount.Version, toAccount.Balance); err != nil {
		return nil, err
	}
	//write tx history
	if err = mysql.createTX(ctx, toAccountID, amount, Deposit, tx, referenceNumber); err != nil {
		return nil, err
	}
	return tx, nil
}

func (mysql accountsMySQL) Withdraw(ctx context.Context, withdrawAmount int64, fromAccountID int, fromUserID int64, referenceNumber *string) (*sql.Tx, error) {
	if withdrawAmount > 0 {
		return nil, ErrInvalidAmount
	}
	fromAccount, err := mysql.GetAccountByUserAndAccountID(ctx, fromUserID, fromAccountID)
	if err != nil {
		return nil, err
	}
	tx, err := mysql.client.DB().Begin()
	if err != nil {
		return nil, err
	}
	if err := mysql.updateBalance(ctx, tx, fromUserID, fromAccountID, withdrawAmount, fromAccount.Version, fromAccount.Balance); err != nil {
		return nil, err
	}
	//write tx history
	if err = mysql.createTX(ctx, fromAccountID, withdrawAmount, Withdraw, tx, referenceNumber); err != nil {
		return nil, err
	}
	return tx, nil
}

func (mysql accountsMySQL) DeleteAll(ctx context.Context) {
	mysql.client.Txs.Delete().ExecX(ctx)
	mysql.client.Account.Delete().ExecX(ctx)
}

func (mysql accountsMySQL) create(ctx context.Context, userId int64) (*ent.Account, error) {
	account, err := mysql.client.Account.Create().SetBalance(0).SetIsVerified(false).SetUserID(userId).SetVersion(0).Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, ErrUserAccountExist
		}
		return nil, err
	}
	return account, nil
}

func (mysql accountsMySQL) GetAccountByUserAndAccountID(ctx context.Context, userID int64, accountID int) (*ent.Account, error) {
	account, err := mysql.client.Account.Query().Where(account.UserID(userID), account.ID(accountID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrAccountNotExist
		}
		return nil, err
	}
	return account, nil
}

func (mysql accountsMySQL) updateBalance(ctx context.Context, tx *sql.Tx, userId int64, accountId int, amount int64, version uint8, balance int64) error {
	// NOTICE: required to use raw sql because .ent doesn't support occ..
	balance += amount
	if balance < 0 {
		rollback(tx)
		return ErrLowAccountBalance
	}
	result, err := tx.ExecContext(ctx, updateBalanceSQL, balance, version+1, accountId, version)
	if err != nil {
		rollback(tx)
		return err
	}
	isOccChanged, err := result.RowsAffected()
	if err != nil {
		rollback(tx)
		return err
	}
	// validate does the version has changed during tx..
	if isOccChanged != 1 {
		rollback(tx)
		if mysql.isDebug {
			log.Printf("accounts: occ issue for userID=%v accountID=%v amount=%v .. should do rollback..", userId, accountId, amount)
		}
		return ErrOptimisticConcurrency
	}
	if mysql.isDebug {
		log.Printf("accounts: balance change for userID=%v accountID=%v amount=%v is ok..", userId, accountId, amount)
	}
	return nil
}

func isDebug() bool {
	isDebug, err := strconv.ParseBool(misc.GetEnv("ACCOUNTS_DEBUG", "true"))
	if err != nil {
		log.Panic("accounts: env for ACCOUNTS_DEBUG is not boolean")
	}
	return isDebug
}

func (mysql accountsMySQL) createTX(ctx context.Context, accountID int, amount int64, types TxType, tx *sql.Tx, referenceNumber *string) error {
	rn := sql.NullString{Valid: false}
	if referenceNumber != nil {
		rn.String = *referenceNumber
		rn.Valid = true
	}
	_, err := tx.ExecContext(ctx, createTXsSQL, types.String(), time.Now().UTC(), amount, accountID, rn)
	if err != nil {
		rollback(tx)
		return err
	}
	return nil
}

func rollback(tx *sql.Tx) {
	if err := tx.Rollback(); err != nil {
		log.Print("accounts-repository: rollback failed: "+err.Error())
	}
}

func NewAccountRepository(client *ent.Client) AccountRepository {
	return accountsMySQL{client:client, isDebug:isDebug()}
}