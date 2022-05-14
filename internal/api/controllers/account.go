package controllers

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/services"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AccountController struct {
	Service *services.AccountService
}

func NewAccountController(service *services.AccountService) *AccountController {
	return &AccountController{
		Service: service,
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

	account, err := ac.Service.FindByAccountNumber(types.AccountNumber(accountNumberI))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "account not found",
		})
		return
	}

	c.JSON(http.StatusOK, account)
	return
}

func (ac *AccountController) Create(c *gin.Context) {
	var account *models.Account
	err := c.BindJSON(&account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind json",
		})
		return
	}

	account, err = ac.Service.Create(account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, account)

	return
}
