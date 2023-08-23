package main

import (
	"Banking/config"
	"Banking/constants"
	"Banking/controllers"
	"Banking/routes"
	"Banking/services"
	"context"
	"fmt"
	"log"

	//	"rest-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	//mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initRoutes() {
	routes.Default(server)
}

func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	profileCollection := mongoClient.Database(constants.DatabaseName).Collection(constants.Collection)
	newService:= services.StartNewService(profileCollection, ctx)
	newController := controllers.StartNewController(newService)
	routes.BankRoutes(server, newController)
}

func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDatabase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
