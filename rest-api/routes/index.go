package routes

import (
	"rest-api/controllers"

	"github.com/gin-gonic/gin"
	
)

func TransactionRoute(router *gin.Engine, controller controllers.TransactionController) {
	router.POST("/api/profile/createNew", controller.CreateTransaction)
	//router.POST("/api/profile/edit", controller.EditProfile)
	//router.DELETE("/api/profile/delete/:id", controller.DeleteProfile)
	router.GET("/api/profile/customer", controller.GetTransaction)
	//router.POST("/api/profile/search", controller.GetProfileBySearch)
}
