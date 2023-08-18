// structure model is created here
package models

type Transactions struct {
	Id              string `json:"id" bson:"id" binding:"required"`
	TransactionDate string `json:"transaction_date" bson:"transaction_date" binding:"required"`
	Amount          int    `json:"amount" bson:"amount" binding:"required"`
	TransactionType string `json:"transaction_type" bson:"transaction_type" binding:"required"`
	CustomerId      string `json:"customer_id" bson:"customer_id" binding:"required"`
}
