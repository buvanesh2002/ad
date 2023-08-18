package interfaces

import "rest-api/models"
type ITransaction interface{
	CreateTransaction(transaction *models.Transactions)(string,error)
	GetTransaction(transaction*models.Transactions)(string,error)
}
