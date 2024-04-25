package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// getting all the transactions
/*
This chaincode is used to get all the transactions
startdate and enddate will be empty as of now , if we want from particular dtae to a specific date
we will set both values else we pass empty.
*/
func (s *SmartContract) GetAllTransactions(ctx contractapi.TransactionContextInterface) ([]*Transactions, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("transaction_", "u")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var transactions []*Transactions
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var trans Transactions
		err = json.Unmarshal(queryResponse.Value, &trans)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &trans)

	}

	return transactions, nil
}

/*
This chaincode is used to get all the transactions for a particular buyer
startdate and enddate will be empty as of now , if we want from particular dtae to a specific date
we will set both values else we pass empty.
*/

func (s *SmartContract) GetTransactionByBuyerId(ctx contractapi.TransactionContextInterface, userid string) ([]*Transactions, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("transaction_", "u")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var transactions []*Transactions
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var trans Transactions
		err = json.Unmarshal(queryResponse.Value, &trans)
		if err != nil {
			return nil, err
		}
		if trans.BuyerEmail == userid && trans.Delivered == "true" {
			transactions = append(transactions, &trans)
		}
	}

	return transactions, nil
}

/*
This chaincode is used to get all the transactions for a particular seller
startdate and enddate will be empty as of now , if we want from particular dtae to a specific date
we will set both values else we pass empty.
*/
func (s *SmartContract) GetTransactionBySellerId(ctx contractapi.TransactionContextInterface, userid string) ([]*Transactions, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("transaction_", "u")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var transactions []*Transactions
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var trans Transactions
		err = json.Unmarshal(queryResponse.Value, &trans)
		if err != nil {
			return nil, err
		}
		if trans.SellerEmail == userid && trans.Delivered == "true" {
			transactions = append(transactions, &trans)
		}
	}

	return transactions, nil
}

/*
This chaincode is used to get all the products that buyer has bought till now
startdate and enddate will be empty as of now , if we want from particular dtae to a specific date
we will set both values else we pass empty.
*/

func (s *SmartContract) GetBuyerOrderedProducts(ctx contractapi.TransactionContextInterface, buyerid string) ([]*Transactions, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("transaction_", "u")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var history []*Transactions
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var his Transactions
		err = json.Unmarshal(queryResponse.Value, &his)
		if err != nil {
			return nil, err
		}
		if his.BuyerEmail == buyerid && his.Delivered == "false" {
			history = append(history, &his)
		}
	}

	return history, nil
}

func (s *SmartContract) GetSellerOrderedProducts(ctx contractapi.TransactionContextInterface, buyerid string) ([]*Transactions, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("transaction_", "u")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var history []*Transactions
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var his Transactions
		err = json.Unmarshal(queryResponse.Value, &his)
		if err != nil {
			return nil, err
		}
		if his.SellerEmail == buyerid && his.Delivered == "false" {
			history = append(history, &his)
		}
	}

	return history, nil
}

/*
This chaincode is used to get all the products that seller has sold till now
startdate and enddate will be empty as of now , if we want from particular dtae to a specific date
we will set both values else we pass empty.
*/

func (s *SmartContract) GetSellerHistory(ctx contractapi.TransactionContextInterface, sellerid string) ([]*Transactions, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("transaction_", "u")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var history []*Transactions
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var his Transactions
		err = json.Unmarshal(queryResponse.Value, &his)
		if err != nil {
			return nil, err
		}
		if his.SellerEmail == sellerid {
			history = append(history, &his)
		}
	}

	return history, nil
}

/*
This chaincode is used to add a transaction and also to add in buyer history as well as seller history
startdate and enddate will be empty as of now , if we want from particular dtae to a specific date
we will set both values else we pass empty.
*/

func (s *SmartContract) AddTransaction(ctx contractapi.TransactionContextInterface, id string, amount string, buyerid string, sellerid string, productid string, paymentmode string, phonenumber string, productname string, productimg string, date time.Time, transactionID string, status string, otp string) error {

	trans := Transactions{
		ID:            id,
		Amount:        amount,
		BuyerEmail:    buyerid,
		SellerEmail:   sellerid,
		ProductID:     productid,
		PaymentMode:   paymentmode,
		PhoneNumber:   phonenumber,
		ProductName:   productname,
		ProductImg:    productimg,
		Date:          date,
		TransactionID: transactionID,
		Delivered:     status,
		Otp:           otp,
	}
	buyer := OrderedProducts{
		ID:          id,
		ProductID:   productid,
		SellerEmail: sellerid,
		BuyerEmail:  buyerid,
		Delivered:   "false",
		Mobile:      phonenumber,
		Image:       productimg,
		Amount:      amount,
		Date:        date,
	}
	b, err := json.Marshal(buyer)
	if err != nil {
		return err
	}

	t, err := json.Marshal(trans)
	if err != nil {
		return err
	}

	ctx.GetStub().PutState(id, b)

	return ctx.GetStub().PutState(id, t)
}

func (s *SmartContract) VerifyTransaction(ctx contractapi.TransactionContextInterface, id string, otp string) (string, error) {
	trans, err := ctx.GetStub().GetState(id)
	if err != nil {
		return "false", fmt.Errorf("some error in reading state %v", err)
	}

	var trans1 Transactions
	err = json.Unmarshal(trans, &trans1)
	if err != nil {
		return "false", fmt.Errorf("some error in  error in json Unmarshal %v", err)
	}

	// overwriting original asset with new asset
	if trans1.Otp == otp {
		trans1.Delivered = "true"
	} else {
		return "false", fmt.Errorf("wrong otp")
	}
	assetJSON, err := json.Marshal(trans1)

	if err != nil {
		return "false", fmt.Errorf("some error in json Marshal %v", err)
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "false", fmt.Errorf("some error in writing to Ledger %v", err)
	}

	return "true", nil

}
