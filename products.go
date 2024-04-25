package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// getting all the products by user id
/*
This chaincode is used to get all the products of a particular user using its userid
startdate and enddate will be empty as of now , if we want from particular dtae to a specific date
we will set both values else we pass empty.

*/
func (s *SmartContract) GetAllProducts(ctx contractapi.TransactionContextInterface) ([]*Products, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("product_", "q")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// var ratings []*Ratings
	var products []*Products
	// sum := 0.0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		// var r Ratings
		var p Products
		err = json.Unmarshal(queryResponse.Value, &p)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)

	}

	return products, nil
}

func (s *SmartContract) GetAllProductsUser(ctx contractapi.TransactionContextInterface, userid string) ([]*Products, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("product_", "q")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// var ratings []*Ratings
	var products []*Products
	// sum := 0.0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		// var r Ratings
		var p Products
		err = json.Unmarshal(queryResponse.Value, &p)
		if err != nil {
			return nil, err
		}
		if p.SellerEmail != userid {
			// sum = sum + float64(r.Rating)
			products = append(products, &p)
		}
	}

	return products, nil
}

// getting all the unsold products by user id
func (s *SmartContract) GetAllUnsoldProducts(ctx contractapi.TransactionContextInterface, userid string) ([]*Products, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("product_", "q")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// var ratings []*Ratings
	var products []*Products
	// sum := 0.0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		// var r Ratings
		var p Products
		err = json.Unmarshal(queryResponse.Value, &p)
		if err != nil {
			return nil, err
		}
		if p.SellerEmail == userid && p.Flag == "true" {
			// sum = sum + float64(r.Rating)
			products = append(products, &p)
		}
	}

	return products, nil
}

// getting all the sold products by user id
func (s *SmartContract) GetAllSoldProducts(ctx contractapi.TransactionContextInterface, userid string) ([]*Products, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("product_", "q")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// var ratings []*Ratings
	var products []*Products
	// sum := 0.0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		// var r Ratings
		var p Products
		err = json.Unmarshal(queryResponse.Value, &p)
		if err != nil {
			return nil, err
		}
		if p.SellerEmail == userid && p.Flag == "false" {
			// sum = sum + float64(r.Rating)
			products = append(products, &p)
		}
	}

	return products, nil
}

// add a prouct for a a particular user
func (s *SmartContract) AddProduct(ctx contractapi.TransactionContextInterface, id string, productname string, productprice string, productsubname string, productdesc string, productprimaryimg string, productsecimg1 string, productsecimg2 string, selleremail string, sellerupi string, flag string) error {

	pro := Products{
		ID:                id,
		ProductName:       productname,
		ProductPrice:      productprice,
		ProductSubName:    productsubname,
		ProductDesc:       productdesc,
		ProductPrimaryImg: productprimaryimg,
		ProductSecImg1:    productsecimg1,
		ProductSecImg2:    productsecimg2,
		SellerEmail:       selleremail,
		SellerUPI:         sellerupi,
		Flag:              flag,
	}
	p, err := json.Marshal(pro)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, p)
}

func (s *SmartContract) DeleteProduct(ctx contractapi.TransactionContextInterface, id string) error {
	product, err := ctx.GetStub().GetState(id)
	if err != nil {
		return err
	}
	if product == nil {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)

}
