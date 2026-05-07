package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}

// Kullanıcı oluştur
func (s *SmartContract) CreateUser(ctx contractapi.TransactionContextInterface, id string, name string) error {
	exists, _ := ctx.GetStub().GetState(id)
	if exists != nil {
		return fmt.Errorf("user already exists")
	}

	user := User{
		ID:     id,
		Name:   name,
		Points: 0,
	}

	userJSON, _ := json.Marshal(user)
	return ctx.GetStub().PutState(id, userJSON)
}

// Puan ekle
func (s *SmartContract) AddPoints(ctx contractapi.TransactionContextInterface, id string, points string) error {
	userJSON, _ := ctx.GetStub().GetState(id)
	if userJSON == nil {
		return fmt.Errorf("user not found")
	}

	var user User
	json.Unmarshal(userJSON, &user)

	p, _ := strconv.Atoi(points)
	user.Points += p

	updatedJSON, _ := json.Marshal(user)
	return ctx.GetStub().PutState(id, updatedJSON)
}

// Transfer
func (s *SmartContract) TransferPoints(ctx contractapi.TransactionContextInterface, from string, to string, points string) error {
	fromJSON, _ := ctx.GetStub().GetState(from)
	toJSON, _ := ctx.GetStub().GetState(to)

	if fromJSON == nil || toJSON == nil {
		return fmt.Errorf("user not found")
	}

	var fromUser, toUser User
	json.Unmarshal(fromJSON, &fromUser)
	json.Unmarshal(toJSON, &toUser)

	p, _ := strconv.Atoi(points)

	if fromUser.Points < p {
		return fmt.Errorf("not enough points")
	}

	fromUser.Points -= p
	toUser.Points += p

	fJSON, _ := json.Marshal(fromUser)
	tJSON, _ := json.Marshal(toUser)

	ctx.GetStub().PutState(from, fJSON)
	ctx.GetStub().PutState(to, tJSON)

	return nil
}

// Kullanıcı getir
func (s *SmartContract) GetUser(ctx contractapi.TransactionContextInterface, id string) (*User, error) {
	userJSON, _ := ctx.GetStub().GetState(id)
	if userJSON == nil {
		return nil, fmt.Errorf("user not found")
	}

	var user User
	json.Unmarshal(userJSON, &user)
	return &user, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		panic(err)
	}

	if err := chaincode.Start(); err != nil {
		panic(err)
	}
}