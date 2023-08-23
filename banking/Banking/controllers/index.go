package controllers

import (
	"Banking/interfaces"
	"Banking/models"
	"net/http"
	//"strings"

	"github.com/gin-gonic/gin"
)

type NewController struct {
	NewService interfaces.CustomerFunc
}

func StartNewController(NewService interfaces.CustomerFunc) NewController {
	return NewController{NewService}
}
func (nc *NewController) CreateCustomer(ctx *gin.Context) {
	var customer *models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	 nc.NewService.CreateCustomer(customer)
	
}
