package accountsrpc

import (
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
)

const (
	AccountNotFound        	  	= "account_not_found"
	AccountExist        	  	= "account_exist"
	AccountTXsRollback   		= "account_txs_rollback"
	AccountTXsCommit     		= "account_txs_commit"
	AccountInvalidAmount     	= "account_invalid_amount"
	LowAccountBalance			= "low_account_balance"
	WithdrawIsDisabled			= "withdraw_is_disabled"
)

func IsReason(error twirp.Error, key string) bool {
	reason := error.Meta(rpcz.Reason)
	if validator.IsEmpty(reason) {
		return false
	}
	return reason == key
}