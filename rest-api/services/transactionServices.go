// function implementation is done here.
package services

import (
	"context"
	"rest-api/interfaces"
	"rest-api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionService struct {
	TransactionCollection *mongo.Collection
	ctx                   context.Context
}

// GetTransaction implements interfaces.ITransaction.


func InitTransactionService(collection *mongo.Collection, ctx context.Context) interfaces.ITransaction {
	return &TransactionService{collection, ctx}
}
func (t *TransactionService) CreateTransaction(transaction *models.Transactions) (string,error){
	// transaction.Amount = 15000
	// transaction.TransactionType = "deposit"
	// transaction.Id="001"
	// transaction.TransactionDate="21.04.2023"
	// transaction.CustomerId="51555"
	// return []models.Transactions{},nil
	_, err := t.TransactionCollection.InsertOne(t.ctx, &transaction)
	if err!=nil{
		return "error",nil
	}
	return "success",nil

	
}
func (t*TransactionService) GetTransaction(transaction *models.Transactions) (string,error) {
	panic("unimplemented")
}

