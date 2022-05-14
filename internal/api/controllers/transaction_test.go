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
	"testing"
)

func TestTransactionController_GetTransactionHistory(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		account := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		})

		_, _ = mockTransactionService.NewDeposit(&models.Deposit{
			AccountNumber: account.AccountNumber,
			Amount:        100,
		})

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/accounting/:accountNumber", mockTransactionController.GetTransactionHistory)

		req, err := http.NewRequest(http.MethodGet, "/accounting/1", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualTH []*models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualTH)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, 1, len(actualTH))

	})

	t.Run("invalid account number", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/accounting/:accountNumber", mockTransactionController.GetTransactionHistory)

		req, err := http.NewRequest(http.MethodGet, "/accounting/12", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}

func TestTransactionController_Deposit(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		account := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		})

		depositJSON, err := json.Marshal(&models.Deposit{
			AccountNumber: account.AccountNumber,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/deposit", mockTransactionController.Deposit)

		req, err := http.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(depositJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, types.Deposit, actualT.TransactionType)
		assert.Equal(t, http.StatusOK, rr.Code)

	})

	t.Run("InvalidAccountNumber", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		account := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		})

		depositJSON, err := json.Marshal(&models.Deposit{
			AccountNumber: account.AccountNumber + 1,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/deposit", mockTransactionController.Deposit)

		req, err := http.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(depositJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidAccountType", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		account := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "corporate",
		})

		depositJSON, err := json.Marshal(&models.Deposit{
			AccountNumber: account.AccountNumber,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/deposit", mockTransactionController.Deposit)

		req, err := http.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(depositJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}

func TestTransactionController_Withdraw(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		account := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		})

		err := mockAccountCache.UpdateBalance(account.AccountNumber, 500)
		assert.NoError(t, err)

		withdrawJSON, err := json.Marshal(&models.Withdraw{
			AccountNumber: account.AccountNumber,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/withdraw", mockTransactionController.Withdraw)

		req, err := http.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(withdrawJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, types.Withdraw, actualT.TransactionType)
		assert.Equal(t, http.StatusOK, rr.Code)

	})

	t.Run("InvalidAccountNumber", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		account := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		})

		err := mockAccountCache.UpdateBalance(account.AccountNumber, 500)
		assert.NoError(t, err)

		withdrawJSON, err := json.Marshal(&models.Withdraw{
			AccountNumber: account.AccountNumber + 1,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/withdraw", mockTransactionController.Withdraw)

		req, err := http.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(withdrawJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidAccountType", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		account := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "corporate",
		})

		err := mockAccountCache.UpdateBalance(account.AccountNumber, 500)
		assert.NoError(t, err)

		withdrawJSON, err := json.Marshal(&models.Withdraw{
			AccountNumber: account.AccountNumber,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/withdraw", mockTransactionController.Withdraw)

		req, err := http.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(withdrawJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InsufficientBalance", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		account := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		})

		err := mockAccountCache.UpdateBalance(account.AccountNumber, 500)
		assert.NoError(t, err)

		withdrawJSON, err := json.Marshal(&models.Withdraw{
			AccountNumber: account.AccountNumber,
			Amount:        750,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/withdraw", mockTransactionController.Withdraw)

		req, err := http.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(withdrawJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}

func TestTransactionController_Payment(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		sender := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		})

		receiver := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Tringle",
			AccountType:  "corporate",
		})

		err := mockAccountCache.UpdateBalance(sender.AccountNumber, 500)
		assert.NoError(t, err)

		paymentJSON, err := json.Marshal(&models.Payment{
			SenderAccount:   sender.AccountNumber,
			ReceiverAccount: receiver.AccountNumber,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, types.Payment, actualT.TransactionType)
		assert.Equal(t, 400, sender.Balance)
		assert.Equal(t, 100, receiver.Balance)
		assert.Equal(t, http.StatusOK, rr.Code)

	})

	t.Run("InvalidAccountNumberForSender", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		sender := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		})

		receiver := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Tringle",
			AccountType:  "corporate",
		})

		err := mockAccountCache.UpdateBalance(sender.AccountNumber, 500)
		assert.NoError(t, err)

		paymentJSON, err := json.Marshal(&models.Payment{
			SenderAccount:   sender.AccountNumber + 1,
			ReceiverAccount: receiver.AccountNumber,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidAccountNumberForReceiver", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		sender := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		})

		receiver := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Tringle",
			AccountType:  "corporate",
		})

		err := mockAccountCache.UpdateBalance(sender.AccountNumber, 500)
		assert.NoError(t, err)

		paymentJSON, err := json.Marshal(&models.Payment{
			SenderAccount:   sender.AccountNumber,
			ReceiverAccount: receiver.AccountNumber + 1,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidAccountNumberForBoth", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		sender := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		})

		receiver := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Tringle",
			AccountType:  "corporate",
		})

		err := mockAccountCache.UpdateBalance(sender.AccountNumber, 500)
		assert.NoError(t, err)

		paymentJSON, err := json.Marshal(&models.Payment{
			SenderAccount:   sender.AccountNumber + 1,
			ReceiverAccount: receiver.AccountNumber + 1,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidAccountTypeForSender", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		sender := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "corporate",
		})

		receiver := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Tringle",
			AccountType:  "corporate",
		})

		err := mockAccountCache.UpdateBalance(sender.AccountNumber, 500)
		assert.NoError(t, err)

		paymentJSON, err := json.Marshal(&models.Payment{
			SenderAccount:   sender.AccountNumber,
			ReceiverAccount: receiver.AccountNumber,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidAccountTypeForReceiver", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		sender := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		})

		receiver := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Tringle",
			AccountType:  "individual",
		})

		err := mockAccountCache.UpdateBalance(sender.AccountNumber, 500)
		assert.NoError(t, err)

		paymentJSON, err := json.Marshal(&models.Payment{
			SenderAccount:   sender.AccountNumber,
			ReceiverAccount: receiver.AccountNumber,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidAccountTypeForBoth", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		sender := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "corporate",
		})

		receiver := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Tringle",
			AccountType:  "individual",
		})

		err := mockAccountCache.UpdateBalance(sender.AccountNumber, 500)
		assert.NoError(t, err)

		paymentJSON, err := json.Marshal(&models.Payment{
			SenderAccount:   sender.AccountNumber,
			ReceiverAccount: receiver.AccountNumber,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InsufficientBalance", func(t *testing.T) {

		mockTransactionCache := cache.NewTransactionCache()
		mockAccountCache := cache.NewAccountCache()
		mockTransactionService := services.NewTransactionService(mockAccountCache, mockTransactionCache)
		mockTransactionController := NewTransactionController(mockTransactionService)

		sender := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Ayşe Durmaz",
			AccountType:  "individual",
		})

		receiver := mockAccountCache.Create(&models.Account{
			CurrencyCode: "TRY",
			OwnerName:    "Tringle",
			AccountType:  "corporate",
		})

		err := mockAccountCache.UpdateBalance(sender.AccountNumber, 500)
		assert.NoError(t, err)

		paymentJSON, err := json.Marshal(&models.Payment{
			SenderAccount:   sender.AccountNumber,
			ReceiverAccount: receiver.AccountNumber,
			Amount:          501,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.Transaction
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}
