package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/ahmetberke/tringle-candidate-project/internal/cache"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/services"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestAccountController_Create(t *testing.T) {

	gin.SetMode(gin.TestMode)

	mockAccountCache := cache.NewAccountCache()
	mockAccountService := services.NewAccountService(mockAccountCache)
	mockAccountController := NewAccountController(mockAccountService)

	t.Run("Success", func(t *testing.T) {
		mockAccountResp := &models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		}
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/account", mockAccountController.Create)

		accountJSON, err := json.Marshal(mockAccountResp)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/account", bytes.NewBuffer(accountJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var incomingAccount *models.Account
		err = json.NewDecoder(rr.Body).Decode(&incomingAccount)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, mockAccountResp.OwnerName, incomingAccount.OwnerName)

	})

	t.Run("NoContextAccount", func(t *testing.T) {
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/account", mockAccountController.Create)

		req, err := http.NewRequest(http.MethodPost, "/account", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidOwnerNameForIndividualAccount", func(t *testing.T) {
		mockAccountResp := &models.Account{
			CurrencyCode: "X",
			OwnerName:    "Ayşe",
			AccountType:  types.Individual,
		}
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/account", mockAccountController.Create)

		accountJSON, err := json.Marshal(mockAccountResp)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/account", bytes.NewBuffer(accountJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("InvalidCurrencyCode", func(t *testing.T) {
		mockAccountResp := &models.Account{
			CurrencyCode: "X",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		}
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/account", mockAccountController.Create)

		accountJSON, err := json.Marshal(mockAccountResp)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/account", bytes.NewBuffer(accountJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("InvalidAccountType", func(t *testing.T) {
		mockAccountResp := &models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "X",
		}
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/account", mockAccountController.Create)

		accountJSON, err := json.Marshal(mockAccountResp)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/account", bytes.NewBuffer(accountJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestAccountController_Get(t *testing.T) {

	gin.SetMode(gin.TestMode)

	mockAccountCache := cache.NewAccountCache()
	mockAccountService := services.NewAccountService(mockAccountCache)
	mockAccountController := NewAccountController(mockAccountService)

	t.Run("Success", func(t *testing.T) {
		mockAccountResp := &models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		}
		mockAccountResp, err := mockAccountService.Create(mockAccountResp)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/account/:accountNumber", mockAccountController.Get)

		req, err := http.NewRequest(http.MethodGet, "/account/"+strconv.FormatInt(int64(mockAccountResp.AccountNumber), 10), nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		/*
			b, err := io.ReadAll(rr.Body)
			assert.NoError(t, err)

			fmt.Printf("BODY : %s", string(b))
		*/

		var incomingAccount *models.Account

		err = json.NewDecoder(rr.Body).Decode(&incomingAccount)
		assert.NoError(t, err)

		assert.Equal(t, mockAccountResp.AccountNumber, incomingAccount.AccountNumber)
		assert.Equal(t, http.StatusOK, rr.Code)

	})
	t.Run("AccountNotFound", func(t *testing.T) {
		mockAccountResp := &models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		}
		mockAccountResp, err := mockAccountService.Create(mockAccountResp)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/account/:accountNumber", mockAccountController.Get)

		req, err := http.NewRequest(http.MethodGet, "/account/"+strconv.FormatInt(int64(mockAccountResp.AccountNumber+1), 10), nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		/*
			b, err := io.ReadAll(rr.Body)
			assert.NoError(t, err)

			fmt.Printf("BODY : %s", string(b))
		*/

		var incomingAccount *models.Account

		err = json.NewDecoder(rr.Body).Decode(&incomingAccount)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}
