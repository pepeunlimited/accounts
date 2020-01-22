# cURL

### Install
```$ brew install jq > curl ... | jq```

##### CreateAccount

###### COIN
```
$ curl -H "Content-Type: application/json" -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/CreateAccount"  -d '{"account_type": "coin", "user_id": 1}'
```
###### CASH
```
$ curl -H "Content-Type: application/json" -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/CreateAccount"  -d '{"account_type": "cash", "user_id": 1}'
```

##### GetAccounts
```
$ curl -H "Content-Type: application/json" -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/GetAccounts"  -d '{"user_id": 1, "accoun_type": "coin"}'
```

##### GetAccount
```
$ curl -H "Content-Type: application/json" -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/GetAccount"  -d '{"user_id": 1, "account_type": "coin"}'
```

##### CreateDeposit
```
$ curl -H "Content-Type: application/json" -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.Acco©©untService/CreateDeposit"  -d '{"user_id": 1, "account_type": "cash", "amount": 200}'
```
##### CreateTransfer
```
$ curl -H "Content-Type: application/json" -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/CreateTransfer"  -d '{"from_user_id": 2, "from_amount": -200,"to_user_id": 2, "to_amount": 200}'
```

##### CreateWithdraw
```
$ curl -H "Content-Type: application/json" -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/CreateWithdraw"  -d '{"user_id": 1, "amount": -200}'
```