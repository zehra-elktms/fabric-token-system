# Hyperledger Fabric Token System

A simple blockchain-based token and point management system built with Hyperledger Fabric and Go chaincode.

## Features

- Create users
- Add token/points
- Transfer points between users
- Query user information from ledger

## Technologies

- Hyperledger Fabric
- Go
- Docker
- WSL
- VS Code

## Smart Contract Functions

### CreateUser
Creates a new user on blockchain ledger.

### AddPoints
Adds points to a user account.

### TransferPoints
Transfers points between users.

### GetUser
Returns user information from blockchain ledger.

---

## Deploy Chaincode

```bash
./network.sh deployCC \
-ccn token \
-ccp ../token-chaincode \
-ccl go
```

## Query Example

```bash
peer chaincode query \
-C mychannel \
-n token \
-c '{"function":"GetUser","Args":["user1"]}'
```

