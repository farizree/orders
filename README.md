# orders
orders service handles wallets, transaction and merchant(soon)

## Folder Structure
```text
├── README.md
├── main.go
├── go.mod
├── go.sum
├── config
│   └── config.go
├── handler
│   └── httphandler
│       ├── transaction_handler.go
│       └── wallet_handler.go
├── model
├── pkg // Store external dependencies package

```

## Prerequisite
1. Postgre
3. Go

## Run The Service
1. Set Up `config.go`
2. Run API Server
```
go run main.go
```