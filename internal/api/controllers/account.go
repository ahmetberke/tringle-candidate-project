package controllers

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AccountController struct {
	service accountService
}

type accountService interface {
	FindByAccountNumber(accountNumber types.AccountNumber) (*models.Account, error)
	Create(account *models.Account) (*models.Account, error)
	Delete(accountNumber types.AccountNumber)
}

func NewAccountController(s accountService) *AccountController {
	return &AccountController{
		service: s,
	}
}

func (ac *AccountController) Get(c *gin.Context) {
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

	account, err := ac.service.FindByAccountNumber(types.AccountNumber(accountNumberI))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "account not found",
		})
		return
	}

	accountDTO := account.DTO()
	if accountDTO == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "something is wrong",
		})
		return
	}

	c.JSON(http.StatusOK, accountDTO)
	return
}

func (ac *AccountController) Create(c *gin.Context) {
	var accountDTO *models.AccountDTO
	err := c.BindJSON(&accountDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind json",
		})
		return
	}

	account := accountDTO.Normal()

	account, err = ac.service.Create(account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	respAccountDTO := account.DTO()
	if respAccountDTO == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "something is wrong",
		})
	}

	c.JSON(http.StatusCreated, respAccountDTO)

	return
}
