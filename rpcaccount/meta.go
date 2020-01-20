package rpcaccount

import (
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
)

const (
	AccountNotFound        	  = "account_not_found"
	AccountExist        	  = "account_exist"
)

func IsReason(error twirp.Error, key string) bool {
	reason := error.Meta(rpcz.Reason)
	if validator.IsEmpty(reason) {
		return false
	}
	return reason == key
}