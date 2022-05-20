package controllers

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TransactionController struct {
	service transactionService
}

type transactionService interface {
	NewPayment(payment *models.Payment) (*models.Transaction, error)
	NewDeposit(deposit *models.Deposit) (*models.Transaction, error)
	NewWithdraw(withdraw *models.Withdraw) (*models.Transaction, error)
	GetTransactionHistory(accountNumber types.AccountNumber) ([]*models.Transaction, error)
}

func NewTransactionController(s transactionService) *TransactionController {
	return &TransactionController{service: s}
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

	var transactionHistoryDTO []*models.TransactionDTO
	for _, t := range transactionHistory {
		transactionHistoryDTO = append(transactionHistoryDTO, t.DTO())
	}

	c.JSON(http.StatusOK, transactionHistoryDTO)
	return

}

func (tc *TransactionController) Payment(c *gin.Context) {
	var paymentDTO *models.PaymentDTO
	err := c.BindJSON(&paymentDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind json",
		})
		return
	}

	payment := paymentDTO.Normal()

	transaction, err := tc.service.NewPayment(payment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	transactionDTO := transaction.DTO()

	c.JSON(http.StatusOK, transactionDTO)
	return
}

func (tc *TransactionController) Deposit(c *gin.Context) {
	var depositDTO *models.DepositDTO
	err := c.BindJSON(&depositDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind json",
		})
		return
	}

	deposit := depositDTO.Normal()

	transaction, err := tc.service.NewDeposit(deposit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	transactionDTO := transaction.DTO()
	if transactionDTO == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "something is wrong",
		})
		return
	}

	c.JSON(http.StatusOK, transactionDTO)
	return
}

func (tc *TransactionController) Withdraw(c *gin.Context) {
	var withdrawDTO *models.WithdrawDTO
	err := c.BindJSON(&withdrawDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind json",
		})
		return
	}

	withdraw := withdrawDTO.Normal()

	transaction, err := tc.service.NewWithdraw(withdraw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	transactionDTO := transaction.DTO()

	c.JSON(http.StatusOK, transactionDTO)
	return
}
