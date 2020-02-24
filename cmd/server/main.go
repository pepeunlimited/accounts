package main

import (
	"github.com/pepeunlimited/accounts/internal/pkg/ent"
	"github.com/pepeunlimited/accounts/internal/server/twirp"
	"github.com/pepeunlimited/accounts/pkg/rpc/account"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"log"
	"net/http"
)

const (
	Version = "0.1.11"
)

func main() {
	log.Printf("Starting the AccountServer... version=[%v]", Version)

	client := ent.NewEntClient()
	ts := account.NewAccountServiceServer(twirp.NewAccountServer(client), nil)

	mux := http.NewServeMux()
	mux.Handle(ts.PathPrefix(), middleware.Adapt(ts, headers.UserId()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}