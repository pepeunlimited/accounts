package rpcaccount

import (
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
)

const (
	AccountNotFound        	  = "account_not_found"
	AccountExist        	  = "account_exist"
	UserIsBanned        	  = "user_is_banned"
	UserIsLocked        	  = "user_is_locked"
	InvalidCredentials  	  = "invalid_credentials"
	UsernameExist       	  = "username_exist"
	EmailExist          	  = "email_exist"
	TicketExist    			  = "ticket_exist"
	TicketNotFound 			  = "ticket_not_found"
	TicketExpired       	  = "ticket_expired"
	ProfilePictureAccessDenied = "profile_picture_access_denied"
)

func IsReason(error twirp.Error, key string) bool {
	reason := error.Meta(rpcz.Reason)
	if validator.IsEmpty(reason) {
		return false
	}
	return reason == key
}