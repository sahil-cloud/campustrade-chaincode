package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	return nil
}

func main() {

	//invoking chaincode
	rateChainCode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating rateChainCode : %v", err)
	}

	if err := rateChainCode.Start(); err != nil {
		log.Panicf("Error starting rateChainCode: %v", err)
	}

}
