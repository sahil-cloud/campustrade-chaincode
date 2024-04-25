package main

import "time"

// for rating
// id, producr id,seller id,buyer id
/*
This will store the ratings of the users that have been listed on the platform so as to create
an image of correctness or genuiness among the buyers
based on user rating
*/
type Ratings struct {
	ID          string `json:"ID"`
	SellerEmail string `json:"sellerEmail"`
	BuyerEmail  string `json:"buyerEmail"`
	Rating      string `json:"Rating"`
}

//for transaction
//ID , amount , buyerid, products id[] , payment mode, date
/*
This will store the transactions that have happend during the buy/sell of the products
it will stroe all the below things in order to identify it correctly
and we can check for that anytime we want
*/

type Transactions struct {
	ID            string    `json:"ID"`
	Amount        string    `json:"amount"`
	BuyerEmail    string    `json:"buyerEmail"`
	SellerEmail   string    `json:"sellerEmail"`
	ProductID     string    `json:"productID"`
	PaymentMode   string    `json:"paymentMode"`
	PhoneNumber   string    `json:"phoneNumber"`
	ProductName   string    `json:"productName"`
	ProductImg    string    `json:"productImg"`
	Date          time.Time `json:"date"`
	TransactionID string    `json:"transactionID"`
	Delivered     string    `json:"delivered"`
	Otp           string    `json:"otp"`
}

// seller history
// id , prodct id,buyer id,
/*
This will store the buyer and seller history
what have the particular buyer hav bought till now
what are the products the seller has sold till now
*/
type DeliveredProducts struct {
	ID          string    `json:"ID"`
	ProductID   string    `json:"productID"`
	BuyerEmail  string    `json:"buyerEmail"`
	SellerEmail string    `json:"sellerEmail"`
	Delivered   string    `json:"delivered"`
	Mobile      string    `json:"mobile"`
	Image       string    `json:"image"`
	Amount      string    `json:"amount"`
	Date        time.Time `json:"date"`
}

type OrderedProducts struct {
	ID          string    `json:"ID"`
	ProductID   string    `json:"productID"`
	BuyerEmail  string    `json:"buyerEmail"`
	SellerEmail string    `json:"sellerEmail"`
	Delivered   string    `json:"delivered"`
	Mobile      string    `json:"mobile"`
	Image       string    `json:"image"`
	Amount      string    `json:"amount"`
	Date        time.Time `json:"date"`
}

// reviews
// id, seller id, product id, buyer id
/*
This will store the reviews of the users that have been listed on the platform so as to create
an image of correctness or genuiness among the buyers
based on user rating
*/
type Reviews struct {
	ID          string `json:"ID"`
	SellerEmail string `json:"sellerEmail"`
	BuyerEmail  string `json:"buyerEmail"`
	Review      string `json:"Review"`
	Rating      string `json:"Rating"`
}

// products for every seller
// id, seller id, product id, buyer id
/*
It will stroe the product details for a particular seller , we set the flag if its avsilable
and reset if not
*/
type Products struct {
	ID                string `json:"ID"`
	ProductName       string `json:"productName"`
	ProductPrice      string `json:"ProductPrice"`
	ProductSubName    string `json:"ProductSubName"`
	ProductDesc       string `json:"ProductDesc"`
	ProductPrimaryImg string `json:"ProductPrimaryImg"`
	ProductSecImg1    string `json:"ProductSecImg1"`
	ProductSecImg2    string `json:"ProductSecImg2"`
	SellerEmail       string `json:"sellerEmail"`
	SellerUPI         string `json:"sellerUPI"`
	Flag              string `json:"flag"`
}

type Users struct {
	ID       string `json:"ID"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
