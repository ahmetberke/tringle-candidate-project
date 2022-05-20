package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockAccountService struct {
	FindByAccountNumberMock func(accountNumber types.AccountNumber) (*models.Account, error)
	CreateMock              func(account *models.Account) (*models.Account, error)
	DeleteMock              func(accountNumber types.AccountNumber)
}

func (m mockAccountService) FindByAccountNumber(accountNumber types.AccountNumber) (*models.Account, error) {
	return m.FindByAccountNumberMock(accountNumber)
}

func (m mockAccountService) Create(account *models.Account) (*models.Account, error) {
	return m.CreateMock(account)
}

func (m *mockAccountService) Delete(accountNumber types.AccountNumber) {
	m.DeleteMock(accountNumber)
}

func TestAccountController_Create(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockAccountServ := mockAccountService{
			CreateMock: func(account *models.Account) (*models.Account, error) {
				return &models.Account{
					AccountNumber: 1,
					CurrencyCode:  account.CurrencyCode,
					OwnerName:     account.OwnerName,
					AccountType:   account.AccountType,
					Balance:       decimal.NewFromFloat(0),
				}, nil
			},
		}

		mockAccountController := NewAccountController(&mockAccountServ)

		mockAccountResp := &models.AccountDTO{
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

		var incomingAccount *models.AccountDTO
		err = json.NewDecoder(rr.Body).Decode(&incomingAccount)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, mockAccountResp.OwnerName, incomingAccount.OwnerName)

	})

	t.Run("NoContextAccount", func(t *testing.T) {
		mockAccountServ := mockAccountService{}

		mockAccountController := NewAccountController(&mockAccountServ)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/account", mockAccountController.Create)

		req, err := http.NewRequest(http.MethodPost, "/account", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var incomingAccount *models.AccountDTO
		err = json.NewDecoder(rr.Body).Decode(&incomingAccount)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidOwnerNameForIndividualAccount", func(t *testing.T) {
		mockAccountServ := mockAccountService{
			CreateMock: func(account *models.Account) (*models.Account, error) {
				return nil, errors.New("invalid owner name")
			},
		}

		mockAccountController := NewAccountController(&mockAccountServ)

		mockAccountResp := &models.AccountDTO{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe",
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

		var incomingAccount *models.AccountDTO
		err = json.NewDecoder(rr.Body).Decode(&incomingAccount)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("InvalidCurrencyCode", func(t *testing.T) {
		mockAccountServ := mockAccountService{
			CreateMock: func(account *models.Account) (*models.Account, error) {
				return nil, errors.New("invalid currency code")
			},
		}

		mockAccountController := NewAccountController(&mockAccountServ)

		mockAccountResp := &models.AccountDTO{
			CurrencyCode: "X",
			OwnerName:    "Ayşe",
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

		var incomingAccount *models.AccountDTO
		err = json.NewDecoder(rr.Body).Decode(&incomingAccount)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("InvalidAccountType", func(t *testing.T) {
		mockAccountServ := mockAccountService{
			CreateMock: func(account *models.Account) (*models.Account, error) {
				return nil, errors.New("invalid account type")
			},
		}

		mockAccountController := NewAccountController(&mockAccountServ)

		mockAccountResp := &models.AccountDTO{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe",
			AccountType:  "x",
		}
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/account", mockAccountController.Create)

		accountJSON, err := json.Marshal(mockAccountResp)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/account", bytes.NewBuffer(accountJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var incomingAccount *models.AccountDTO
		err = json.NewDecoder(rr.Body).Decode(&incomingAccount)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestAccountController_Get(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		mockAccountServ := mockAccountService{
			FindByAccountNumberMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				return &models.Account{
					AccountNumber: accountNumber,
					CurrencyCode:  types.TRY,
					OwnerName:     "test",
					AccountType:   types.Individual,
					Balance:       decimal.NewFromFloat(123),
				}, nil
			},
		}

		mockAccountController := NewAccountController(&mockAccountServ)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/account/:accountNumber", mockAccountController.Get)

		req, err := http.NewRequest(http.MethodGet, "/account/1", nil)
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

		assert.Equal(t, types.AccountNumber(1), incomingAccount.AccountNumber)
		assert.Equal(t, http.StatusOK, rr.Code)

	})
	t.Run("AccountNotFound", func(t *testing.T) {
		mockAccountServ := mockAccountService{
			FindByAccountNumberMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				return nil, errors.New("account not found")
			},
		}

		mockAccountController := NewAccountController(&mockAccountServ)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/account/:accountNumber", mockAccountController.Get)

		req, err := http.NewRequest(http.MethodGet, "/account/1", nil)
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
