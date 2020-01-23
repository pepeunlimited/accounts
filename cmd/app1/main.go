package main

import (
	"github.com/pepeunlimited/accounts/internal/app/app1/mysql"
	"github.com/pepeunlimited/accounts/internal/app/app1/server"
	"github.com/pepeunlimited/accounts/accountsrpc"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"log"
	"net/http"
)

const (
	Version = "0.1.5"
)

func main() {
	log.Printf("Starting the AccountServer... version=[%v]", Version)

	client := mysql.NewEntClient()
	ts := accountsrpc.NewAccountServiceServer(server.NewAccountServer(client), nil)

	mux := http.NewServeMux()
	mux.Handle(ts.PathPrefix(), middleware.Adapt(ts, headers.UserId()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}