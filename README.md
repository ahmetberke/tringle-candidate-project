# TRINGLE PAYMENT REST API

A RESTful API for payment systems with GO.

## Allowed Endpoints and Methods

- click on endpoint to go to details 

| Endpoint                                                    | Method |
|-------------------------------------------------------------|--------|
| [/account](#account-endpoint)                               | POST   |
| [/account/:accountNumber](#account-endpoint)                | GET    |
| [/payment](#payment-endpoint)                               | POST   |
| [/deposit](#deposit-endpoint)                               | POST   |
| [/withdraw](#withdraw-endpoint)                             | POST   |
| [/accounting/:accountNumber](#transaction-history-endpoint) | GET    |


## Installation & Run
### Download
```
    $ git clone https://github.com/ahmetberke/tringle-candidate-project
```

### Build & Run With Docker
```
    $ docker build --tag tringle-candidate-project .
    $ docker run --publish 5000:5000 tringle-candidate-project
```
### Build & Run With Docker-Compose
```
    $ docker compose up -d
```
How to whatch logs in docker?
```
    $ docker ps
    $ docker logs <container_name>
```

### Build & Run With GO
```
    $ go mod download
```
```
    $ go build -o /tringle-candidate-project
    $ ./tringle-candidate-project
```
or
```
    $ go run main.go
```

## Production

![heroku](https://www.vectorlogo.zone/logos/heroku/heroku-ar21.png)

This api already published on heroku

click [here](https://tringle-payment-rest-api.herokuapp.com/) to go

## API Structure

![api structure](https://github.com/ahmetberke/tringle-candidate-project/blob/main/images/arc.png?raw=true)

## Folder Structure
```
.
├── configs
│   └── manager.go
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── api.go
│   │   ├── controllers
│   │   │   ├── account.go
│   │   │   ├── account_test.go
│   │   │   ├── transaction.go
│   │   │   └── transaction_test.go
│   │   └── routes.go
│   ├── cache
│   │   ├── account.go
│   │   ├── account_test.go
│   │   ├── transaction_history.go
│   │   └── transaction_history_test.go
│   ├── models
│   │   ├── account.go
│   │   ├── deposit.go
│   │   ├── payment.go
│   │   ├── transaction.go
│   │   └── withdraw.go
│   ├── services
│   │   ├── account.go
│   │   ├── account_test.go
│   │   ├── transaction.go
│   │   └── transaction_test.go
│   └── types
│       └── types.go
└── main.go

```


# Account Endpoint

*Request body*

```
{
  "ownerName": string,
  "currencyCode": {enum: ["TRY", "USD", "EUR"]},
  "accountType": {enum: ["individual", "corporate"]}
}
```

*Response*

```
{
  "accountNumber" : number,
  "ownerName" : string,
  "currencyCode" : {enum : ["TRY", "USD", "EUR"]},
  "accountType" : {enum : ["individual", "corporate"]},
  "balance" : number
}
```


# Payment Endpoint

*Request body*

```
{
  "senderAccount" : number,
  "receiverAccount" : number,
  "amount" : number
}
```

*Response*

```
{
  "accountNumber" : number,
  "amount" :  number,
  "transactionType" : "payment",
  "createdAt" : date
}
```


# Deposit Endpoint

*Request body*

```
{
  "accountNumber": number,
  "amount": number
}
```

*Response*

```
{
  "accountNumber" : number,
  "amount" :  number,
  "transactionType" : "deposit",
  "createdAt" : date
}
```


# Withdraw Endpoint

*Request body*

```
{
  "accountNumber": number,
  "amount": number
}
```

*Response*

```
{
  "accountNumber" : number,
  "amount" :  number,
  "transactionType" : "withdraw",
  "createdAt" : date
}
```


# Transaction History Endpoint

*Response*

```
{
  "accountNumber" : number,
  "amount" :  number,
  "transactionType" : { enum: ["payment", "deposit", "withdraw"] },
  "createdAt" : date
}
```
