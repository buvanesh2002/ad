package routes

import (
	"Banking/controllers"

	"github.com/gin-gonic/gin"
)
func BankRoutes(router *gin.Engine,controller controllers.NewController){
	router.POST("/api/createNew", controller.CreateCustomer)
}