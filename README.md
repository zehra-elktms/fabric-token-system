# Hyperledger Fabric Token System

## Türkçe

Bu proje, Hyperledger Fabric ve Go chaincode kullanılarak geliştirilmiş blockchain tabanlı bir token ve puan yönetim sistemidir.

Projede kullanıcı oluşturma, puan ekleme, kullanıcılar arasında token transferi yapma ve blockchain ledger üzerinden veri sorgulama işlemleri gerçekleştirilmektedir.

Ayrıca proje; Hyperledger Fabric ağı kurma, chaincode deploy etme, endorsement politikalarıyla çalışma ve peer'ler arası transaction doğrulama süreçlerini de içermektedir.

Bu çalışma, Hyperledger Fabric üzerinde temel seviyede çalışan bir blockchain uygulaması geliştirmek amacıyla hazırlanmıştır.

---

## English

A blockchain-based token and point management system developed using Hyperledger Fabric and Go chaincode.

This project demonstrates how to build, deploy, and interact with a custom smart contract (chaincode) on a Hyperledger Fabric network.

---

# Features

* Create blockchain users
* Add token/points to users
* Transfer points between users
* Query user information from the ledger
* Multi-peer endorsement support
* Hyperledger Fabric smart contract deployment

---

# Technologies Used

* Hyperledger Fabric 2.5
* Go (Golang)
* Docker
* WSL
* VS Code

---

# Project Structure

```text
fabric-token-system/
│
├── chaincode.go
├── go.mod
├── go.sum
├── README.md
└── .gitignore
```

---

# Smart Contract Functions

## CreateUser

Creates a new blockchain user.

### Example

```bash
peer chaincode invoke \
-o localhost:7050 \
--ordererTLSHostnameOverride orderer.example.com \
--tls \
--cafile $PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
--waitForEvent \
--peerAddresses localhost:7051 \
--tlsRootCertFiles $PWD/organizations/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem \
--peerAddresses localhost:9051 \
--tlsRootCertFiles $PWD/organizations/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem \
-C mychannel \
-n token \
-c '{"function":"CreateUser","Args":["user1","Zehra"]}'
```

---

## AddPoints

Adds token/points to a user.

### Example

```bash
peer chaincode invoke \
-o localhost:7050 \
--ordererTLSHostnameOverride orderer.example.com \
--tls \
--cafile $PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
--waitForEvent \
--peerAddresses localhost:7051 \
--tlsRootCertFiles $PWD/organizations/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem \
--peerAddresses localhost:9051 \
--tlsRootCertFiles $PWD/organizations/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem \
-C mychannel \
-n token \
-c '{"function":"AddPoints","Args":["user1","100"]}'
```

---

## TransferPoints

Transfers points between users.

### Example

```bash
peer chaincode invoke \
-o localhost:7050 \
--ordererTLSHostnameOverride orderer.example.com \
--tls \
--cafile $PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
--waitForEvent \
--peerAddresses localhost:7051 \
--tlsRootCertFiles $PWD/organizations/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem \
--peerAddresses localhost:9051 \
--tlsRootCertFiles $PWD/organizations/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem \
-C mychannel \
-n token \
-c '{"function":"TransferPoints","Args":["user1","user2","40"]}'
```

---

## GetUser

Queries user information from blockchain ledger.

### Example

```bash
peer chaincode query \
-C mychannel \
-n token \
-c '{"function":"GetUser","Args":["user1"]}'
```

---

# Example Workflow

## 1. Create User

```text
user1 -> Zehra
```

## 2. Add Points

```text
user1 += 100
```

## 3. Transfer Points

```text
user1 -> user2 : 40 points
```

## 4. Final State

```text
user1 = 60
user2 = 40
```

---

# Deploy Chaincode

```bash
./network.sh deployCC \
-ccn token \
-ccp ../token-chaincode \
-ccl go
```

---

# Author

Zehra

---

# GitHub Repository

[https://github.com/zehra-elktms/fabric-token-system](https://github.com/zehra-elktms/fabric-token-system)

