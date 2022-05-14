# TRINGLE PAYMENT REST API

A RESTful API for payment systems with GO.

## Installation & Run
### Download
```
    $ git clone https://github.com/ahmetberke/tringle-candidate-project
```

### Run With Docker
```
    $ docker build --tag tringle-candidate-project .
    $ docker run --publish 5000:5000 tringle-candidate-project
```
### Run With Docker-Compose
```
    $ docker compose up -d
```

### Run With GO
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
│       ├── account-type.go
│       ├── currency.go
│       └── transaction-type.go
└── main.go
```