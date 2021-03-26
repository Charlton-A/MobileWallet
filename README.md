# MobileWallet

A simple Go REST API for a  mobile wallet allowing creation of users, transfer of  funds, checking of balances  and transactions.


## Structure

```bash
├── cli
│   ├── go.mod
│   ├── main.go
│   └── main_test.go
├── db.sql
├── forms
│   ├── go.mod
│   ├── transactions.go
│   └── transfer.go
├── handlers
│   ├── app.go
│   ├── go.mod
│   ├── ping.go
│   ├── transactions.go
│   ├── users.go
│   └── wallets.go
├── models
│   ├── go.mod
│   ├── transactions.go
│   ├── user.go
│   └── wallet.go
├── README.md
└── utils
    ├── go.mod

```

## Installation

Set up postgres  using the db.sql file.\
Configure database dsn and basic auth setting on an .env file.\
You can create one similar to the sample(.env.sample).\
Load the .env file.

```bash
source .env
cd cli/ && go mod tidy && go get
go run main.go

```

## Testing

Ensure you have followed the installation set up first.
```bash
source .env
cd cli/ && go test
```
## API Endpoints

| METHOD  |      ENDPOINT      |  DESCRIPTION |
|----------|:-------------:|------:|
| GET |  /api/v1/users/ping  | Ping to check if API is up |
| POST |    /api/v1/users/create    |   Create  a user with a wallet and a balance of 0  |
| GET |  /api/v1/users/balance/:user_id |   Get user balance on user wallets |
| POST | /api/v1/users/wallet/create  |  Create a wallet for existing user |
| POST | /api/v1/users/wallet/update  |  Update user wallet balance(credit or debit) |
| POST | /api/v1/users/wallet/transfer  |  Transfer funds from one wallet to another |
| GET | /api/v1/users/transactions/:user_id  |Get user transaction details |

To acces the ednpoints use credentials APP_KEY and app APP_PASS via basic auth.


###### POST endpoints and their  respective  json data formats.
##### POST    /api/v1/users/create
```json
{
	"first_name":"Rick101",
	"last_name":"Sanchez",
	"phone":"+1844-666-3645",
	"email":"rick101@c132.earth"
}
```

##### POST    /api/v1/users/wallet/create
```json
{
	"user_id":1,
	"balance":20000
}
```

##### POST    /api/v1/users/wallet/create
```json
{
	"user_id":1,
	"balance":20000
}
```

##### POST    /api/v1/users/wallet/update
Action codes meaning
   1 - Deposit to wallet
   4 - Withdraw from  wallet
```json
{
	"user_id":1,
	"wallet_id":1,
	"amount":500,
	"action":1
}

```

##### POST    /api/v1/users/wallet/transfer
```json
{
	"source_user_id":1,
	"source_wallet_id":1,
	"dest_user_id":2,
	"dest_wallet_id":2,
	"amount":50
}

```
## Monitoring

Simple monitoring can be done by checking the log file .
A Prometheus module exitst for gin https://github.com/zsais/go-gin-prometheus .It can be added as a middleware.




## Assumptions
*API is not meant for public release but for internal communication between applications and the database .\
*All wallets use the same cuurency
*Password managed is handled by another party and the API only offers authentication for resoure acccess\
*The API is still  at prototyping level and has yet to be optimized for scalability

