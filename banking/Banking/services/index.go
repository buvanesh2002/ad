package services

import (
	"Banking/interfaces"
	"Banking/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)
type NewService struct{
	DataCollection *mongo.Collection
	Ctx context.Context
}
func StartNewService(collection *mongo.Collection,ctx context.Context)interfaces.CustomerFunc{
	return &NewService{collection,ctx}

}

func (n*NewService)CreateCustomer(customer *models.Customer){

	_,err:=n.DataCollection.InsertOne(n.Ctx,&customer)
	if err!=nil {
       log.Fatal(err.Error())
	}

}
func(n*NewService)UpdateCustomer(customer *models.Customer){
	

}