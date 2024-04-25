// package main

// import (
// 	"encoding/json"

// 	"github.com/hyperledger/fabric-contract-api-go/contractapi"
// )

// // getting all the ratings by user id
// /*
// This chaincode is used to find all the ratings of the particular user
// with the help of user id
// startdate and enddate will be empty as of now , if we want from particular dtae to a specific date
// we will set both values else we pass empty.

// */
// func (s *SmartContract) GetUser(ctx contractapi.TransactionContextInterface, startDate string, endDate string, email string) ([]*Users, error) {
// 	resultsIterator, err := ctx.GetStub().GetStateByRange(startDate, endDate)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resultsIterator.Close()

// 	var users []*Users
// 	// sum := 0.0
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return nil, err
// 		}

// 		var u Users
// 		err = json.Unmarshal(queryResponse.Value, &u)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if u.Email == email {
// 			users = append(users, &u)
// 		}
// 	}

// 	return users, nil
// }

// // add rating for a transaction
// // this chaincode will be used to add rating to a particular user from a particular buyer
// func (s *SmartContract) AddUser(ctx contractapi.TransactionContextInterface, id string, name string, email string, password string) error {

// 	user := Users{
// 		ID:       id,
// 		Name:     name,
// 		Email:    email,
// 		Password: password,
// 	}
// 	r, err := json.Marshal(user)
// 	if err != nil {
// 		return err
// 	}

// 	return ctx.GetStub().PutState(id, r)
// }

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type UserStruct struct {
	Id                  string   `json:"id"`
	MobileNumber        string   `json:"mobileNumber"`
	Orders              []string `json:"orders"`
	OrderCount          int      `json:"orderCount"`
	ServicesOffered     []string `json:"servicesOffered"`
	ServiceOfferedCount int      `json:"serviceOfferedCount"`
	ServicesDeleted     []string `json:"servicesDeleted"`
	ServiceDeletedCount int      `json:"serviceDeletedCount"`
	ServicesSold        []string `json:"servicesSold"`
	ServiceSoldCount    int      `json:"serviceSoldCount"`
}

type UserReturnStruct struct {
	Id           string `json:"id"`
	MobileNumber string `json:"mobileNumber"`
}

// 	Purpose: This function is used to create a new user profile by storing user-related data in the ledger. It initializes the user's profile with empty arrays for orders, offered services, and deleted services.
// 	Parameters:
// 		`ctx contractapi.TransactionContextInterface`: Provides access to the blockchain context.
// 		`userId`: The unique identifier of the user to be created.
// 	Logic:
// 		It checks if a user with the given `userId` already exists by attempting to retrieve their data from the ledger.
// 		If the user doesn't exist (`exists == nil`), it creates a new `UserStruct`, sets the `userId`, and initializes the user's arrays and counts.
// 		It marshals the user data into JSON.
// 		It stores the user's data on the ledger.
// 	Returns: A boolean indicating the success of the user creation operation and an error message if there's an issue.

func (s *SmartContract) CreateUserKey(ctx contractapi.TransactionContextInterface, userId string, mobileNumber string) (bool, error) {
	exists, err := ctx.GetStub().GetState(userId)
	if err != nil {
		return false, fmt.Errorf("some error in reading state %v", err)
	}
	if exists != nil {
		return false, fmt.Errorf("user already exist")
	}

	var user UserStruct
	user.Id = userId
	user.Orders = []string{}
	user.ServicesOffered = []string{}
	user.ServicesDeleted = []string{}
	user.OrderCount = 0
	user.ServiceOfferedCount = 0
	user.ServiceDeletedCount = 0
	user.MobileNumber = mobileNumber
	userJSON, err := json.Marshal(user)

	if err != nil {
		return false, fmt.Errorf("some error in JSON Marshal %v", err)
	}

	if err = ctx.GetStub().PutState(userId, userJSON); err != nil {
		return false, fmt.Errorf("some error in putting state %v", err)
	}

	return true, nil
}

// 	Purpose: This function is a utility function to check if a user's profile (by `userId`) already exists in the ledger.
// 	Parameters:
// 		`ctx contractapi.TransactionContextInterface`: Provides access to the blockchain context.
// 		`userId`: The unique identifier of the user to be checked.
// 	Logic:
// 		It attempts to retrieve the user's data from the ledger based on the `userId`.
// 		If the user's data exists in the ledger (i.e., `exists != nil`), it returns `true`; otherwise, it returns `false`.
// 	Returns: A boolean indicating whether the user's profile exists in the ledger (true if it exists, false if not).

func (s *SmartContract) IsUserKeyExists(ctx contractapi.TransactionContextInterface, userId string) bool {
	exists, err := ctx.GetStub().GetState(userId)
	if err != nil {
		return false
	}

	return exists != nil
}

func (s *SmartContract) ReadUser(ctx contractapi.TransactionContextInterface, userId string) (UserReturnStruct, error) {
	userBytes, err := ctx.GetStub().GetState(userId)
	returnUser := UserReturnStruct{}
	if err != nil {
		return returnUser, fmt.Errorf("some error in reading state %v", err)
	}
	user := UserStruct{}
	if err = json.Unmarshal(userBytes, &user); err != nil {
		return returnUser, fmt.Errorf("some error in json unmarshal %v", err)
	}

	returnUser.Id = user.Id
	returnUser.MobileNumber = user.MobileNumber

	return returnUser, nil
}
