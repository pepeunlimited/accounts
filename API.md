# cURL

### Install
```$ brew install jq > curl ... | jq```

##### CreateAccount
```
$ curl -H "Content-Type: application/json" -H "X-JWT-UserId: 1"  -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/CreateAccount"  -d '{"account_type": "coin"}'
```

##### GetAccounts
```
$ curl -H "Content-Type: application/json" -H "X-JWT-UserId: 1"  -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/GetAccounts"  -d '{}'
```

##### GetAccount
```
$ curl -H "Content-Type: application/json" -H "X-JWT-UserId: 1"  -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/GetAccount"  -d '{"account_id": 1}'
```