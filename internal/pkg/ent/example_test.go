// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleAccount() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the account's edges.
	t0 := client.Txs.
		Create().
		SetTxType("string").
		SetCreatedAt(time.Now()).
		SetAmount(1).
		SetReferenceNumber("string").
		SaveX(ctx)
	log.Println("txs created:", t0)

	// create account vertex with its edges.
	a := client.Account.
		Create().
		SetBalance(1).
		SetVersion(1).
		SetIsVerified(true).
		SetUserID(1).
		AddTxs(t0).
		SaveX(ctx)
	log.Println("account created:", a)

	// query edges.
	t0, err = a.QueryTxs().First(ctx)
	if err != nil {
		log.Fatalf("failed querying txs: %v", err)
	}
	log.Println("txs found:", t0)

	// Output:
}
func ExampleTxs() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the txs's edges.

	// create txs vertex with its edges.
	t := client.Txs.
		Create().
		SetTxType("string").
		SetCreatedAt(time.Now()).
		SetAmount(1).
		SetReferenceNumber("string").
		SaveX(ctx)
	log.Println("txs created:", t)

	// query edges.

	// Output:
}
