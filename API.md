# cURL

### Install
```$ brew install jq > curl ... | jq```

##### UpdateAccountVerified

```
$  
```

##### CreateAccount

```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/CreateAccount" \
-d '{"user_id": 1}'
```

##### GetAccount

```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/GetAccount" \
 -d '{"user_id": 1}'
```

##### CreateDeposit

```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/CreateDeposit" \
-d '{"user_id": 1, "account_type": "cash", "amount": 200, "reference_number": "$UUID"}'
```
##### CreateWithdraw

```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.accounts.AccountService/CreateWithdraw" \
-d '{"user_id": 1, "amount": -200}'
```