package controllers

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/services"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TransactionController struct {
	service *services.TransactionService
}

func NewTransactionController(service *services.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func (tc *TransactionController) GetTransactionHistory(c *gin.Context) {
	accountNumber, ok := c.Params.Get("accountNumber")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid argument",
		})
		return
	}

	// account number's type must be integer so string value converting to integer type
	accountNumberI, err := strconv.ParseInt(accountNumber, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid argument",
		})
		return
	}

	transactionHistory, err := tc.service.GetTransactionHistory(types.AccountNumber(accountNumberI))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactionHistory)
	return

}

func (tc *TransactionController) Payment(c *gin.Context) {
	var payment *models.Payment
	err := c.BindJSON(&payment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind json",
		})
		return
	}

	transaction, err := tc.service.NewPayment(payment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transaction)
	return
}

func (tc *TransactionController) Deposit(c *gin.Context) {
	var deposit *models.Deposit
	err := c.BindJSON(&deposit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind json",
		})
		return
	}

	transaction, err := tc.service.NewDeposit(deposit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transaction)
	return
}

func (tc *TransactionController) Withdraw(c *gin.Context) {
	var withdraw *models.Withdraw
	err := c.BindJSON(&withdraw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind json",
		})
		return
	}

	transaction, err := tc.service.NewWithdraw(withdraw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transaction)
	return
}
