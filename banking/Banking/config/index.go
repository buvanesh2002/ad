package config

import (
	"Banking/constants"
	"context"
	"fmt"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDatabase() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	mongoConnection := options.Client().ApplyURI(constants.ConnectionString)
	mongoClient, err := mongo.Connect(ctx, mongoConnection) //mongo connection established
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return mongoClient, nil
}
func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	client, err := ConnectDatabase()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	collection := client.Database(dbName).Collection(collectionName) // getting info about the database and collection
	return collection
}
