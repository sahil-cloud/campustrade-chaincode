package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// getting all the ratings by user id
/*
This chaincode is used to find all the reviews of the particular user
with the help of user id
startdate and enddate will be empty as of now , if we want from particular dtae to a specific date
we will set both values else we pass empty.

*/
func (s *SmartContract) GetAllReviews(ctx contractapi.TransactionContextInterface, userid string) ([]*Reviews, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("review_", "s")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var reviews []*Reviews
	// sum := 0.0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var r Reviews
		err = json.Unmarshal(queryResponse.Value, &r)
		if err != nil {
			return nil, err
		}
		if r.SellerEmail == userid {
			// sum = sum + float64(r.Rating)
			reviews = append(reviews, &r)
		}
	}

	return reviews, nil
}

// add review for a particular user b the buyer
func (s *SmartContract) AddReview(ctx contractapi.TransactionContextInterface, id string, sellerid string, buyerid string, review string, rating string) error {

	rev := Reviews{
		ID:          id,
		SellerEmail: sellerid,
		BuyerEmail:  buyerid,
		Review:      review,
		Rating:      rating,
	}
	r, err := json.Marshal(rev)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, r)
}
