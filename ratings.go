package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// getting all the ratings by user id
/*
This chaincode is used to find all the ratings of the particular user
with the help of user id
startdate and enddate will be empty as of now , if we want from particular dtae to a specific date
we will set both values else we pass empty.

*/
func (s *SmartContract) GetAllRatings(ctx contractapi.TransactionContextInterface, userid string) ([]*Ratings, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var ratings []*Ratings
	// sum := 0.0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var r Ratings
		err = json.Unmarshal(queryResponse.Value, &r)
		if err != nil {
			return nil, err
		}
		if r.SellerEmail == userid {
			// sum = sum + float64(r.Rating)
			ratings = append(ratings, &r)
		}
	}

	return ratings, nil
}

// add rating for a transaction
// this chaincode will be used to add rating to a particular user from a particular buyer
func (s *SmartContract) AddRating(ctx contractapi.TransactionContextInterface, id string, sellerid string, buyerid string, rating string) error {

	rate := Ratings{
		ID:          id,
		SellerEmail: sellerid,
		BuyerEmail:  buyerid,
		Rating:      rating,
	}
	r, err := json.Marshal(rate)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, r)
}
