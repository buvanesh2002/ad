// we will implement the controlling operators here
package controllers

import (
	"net/http"
	"rest-api/interfaces"
	"rest-api/models"
	"strings"

	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService interfaces.ITransaction
}

func InitTransactionController(TransactionService interfaces.ITransaction) TransactionController {
	return TransactionController{TransactionService}
}
func (tc *TransactionController) CreateTransaction(ctx *gin.Context) {
	var transaction *models.Transactions
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	 newTransaction,err:=tc.TransactionService.CreateTransaction(transaction)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newTransaction})


}
func (tc *TransactionController) GetTransaction(context *gin.Context) {

}
