package models

//import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	Password  string `json:"password" bson:"password"`
	BankId    string `json:"bank_id" bson:"bank_id"`
	AccountId string `json:"account_id" bson:"account_id"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"mail" bson:"mail"`
}
type TransactionLoan struct {
	LoanAmount float64 `json:"loan_amount" bson:"loan_amount"`
	LoanType   string  `json:"type_of_loan" bson:"type_of_loan"`
}
type Transactions struct {
	CustomerId string            `json:"customer_id" bson:"customer_id"`
	Amount     float64           `json:"transaction_amount" bson:"transaction_amount"`
	loan       []TransactionLoan `json:"loans" bson:"loans"`
}
type Accounts struct{
	AccountId string `json:"account_id" bson:"account_id"`
	AccountType string `json:"account_type" bson:"account_type"`
	Branch string `json:"account_branch" bson:"account_branch"`
	Address string `json:"address" bson:"address"`
}
type Bank struct{
	BankId string `json:"bank_id" bson:"bank_id"`
	BankAddress string`json:"address" bson:"address"`
	IFSC string`json:"IFSC_code" bson:"IFSC_code"`
}