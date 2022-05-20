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
	"strconv"
	"testing"
	"time"
)

type mockTransactionService struct {
	NewPaymentMock               func(payment *models.Payment) (*models.Transaction, error)
	NewDepositMock               func(deposit *models.Deposit) (*models.Transaction, error)
	NewWithdrawMock              func(withdraw *models.Withdraw) (*models.Transaction, error)
	NewGetTransactionHistoryMock func(accountNumber types.AccountNumber) ([]*models.Transaction, error)
}

func (m mockTransactionService) NewPayment(payment *models.Payment) (*models.Transaction, error) {
	return m.NewPaymentMock(payment)
}

func (m mockTransactionService) NewDeposit(deposit *models.Deposit) (*models.Transaction, error) {
	return m.NewDepositMock(deposit)
}

func (m mockTransactionService) NewWithdraw(withdraw *models.Withdraw) (*models.Transaction, error) {
	return m.NewWithdrawMock(withdraw)
}

func (m mockTransactionService) GetTransactionHistory(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
	return m.NewGetTransactionHistoryMock(accountNumber)
}

func TestTransactionController_GetTransactionHistory(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewGetTransactionHistoryMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				testData := []*models.Transaction{
					{
						AccountNumber:   1,
						Amount:          decimal.NewFromFloat(121.1),
						TransactionType: types.Payment,
						CreatedAt:       time.Now(),
					},
					{
						AccountNumber:   1,
						Amount:          decimal.NewFromFloat(312),
						TransactionType: types.Deposit,
						CreatedAt:       time.Now(),
					},
				}

				return testData, nil

			},
		}

		mockTransactionController := NewTransactionController(mockTransactionServ)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/accounting/:accountNumber", mockTransactionController.GetTransactionHistory)

		req, err := http.NewRequest(http.MethodGet, "/accounting/"+strconv.FormatInt(1, 10), nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualTH []*models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualTH)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, 2, len(actualTH))

	})

	t.Run("invalid account number", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewGetTransactionHistoryMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				return nil, errors.New("invalid account number")
			},
		}

		mockTransactionController := NewTransactionController(mockTransactionServ)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.GET("/accounting/:accountNumber", mockTransactionController.GetTransactionHistory)

		req, err := http.NewRequest(http.MethodGet, "/accounting/"+strconv.FormatInt(1, 10), nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}

func TestTransactionController_Deposit(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockTransactionServ := mockTransactionService{
			NewDepositMock: func(deposit *models.Deposit) (*models.Transaction, error) {
				return &models.Transaction{
					AccountNumber:   deposit.AccountNumber,
					Amount:          deposit.Amount,
					TransactionType: types.Deposit,
					CreatedAt:       time.Now(),
				}, nil
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		depositJSON, err := json.Marshal(&models.DepositDTO{
			AccountNumber: 1,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/deposit", mockTransactionController.Deposit)

		req, err := http.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(depositJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, types.Deposit, actualT.TransactionType)
		assert.Equal(t, http.StatusOK, rr.Code)

	})

	t.Run("InvalidAccountNumber", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewDepositMock: func(deposit *models.Deposit) (*models.Transaction, error) {
				return nil, errors.New("invalid account number")
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		depositJSON, err := json.Marshal(&models.DepositDTO{
			AccountNumber: 1,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/deposit", mockTransactionController.Deposit)

		req, err := http.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(depositJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidAccountType", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewDepositMock: func(deposit *models.Deposit) (*models.Transaction, error) {
				return nil, errors.New("invalid account type")
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		depositJSON, err := json.Marshal(&models.DepositDTO{
			AccountNumber: 1,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/deposit", mockTransactionController.Deposit)

		req, err := http.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(depositJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}

func TestTransactionController_Withdraw(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewWithdrawMock: func(withdraw *models.Withdraw) (*models.Transaction, error) {
				return &models.Transaction{
					AccountNumber:   withdraw.AccountNumber,
					Amount:          withdraw.Amount,
					TransactionType: types.Withdraw,
					CreatedAt:       time.Now(),
				}, nil
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		withdrawJSON, err := json.Marshal(&models.WithdrawDTO{
			AccountNumber: 1,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/withdraw", mockTransactionController.Withdraw)

		req, err := http.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(withdrawJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, types.Withdraw, actualT.TransactionType)
		assert.Equal(t, http.StatusOK, rr.Code)

	})

	t.Run("InvalidAccountNumber", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewWithdrawMock: func(withdraw *models.Withdraw) (*models.Transaction, error) {
				return nil, errors.New("invalid account number")
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		withdrawJSON, err := json.Marshal(&models.WithdrawDTO{
			AccountNumber: 1,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/withdraw", mockTransactionController.Withdraw)

		req, err := http.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(withdrawJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidAccountType", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewWithdrawMock: func(withdraw *models.Withdraw) (*models.Transaction, error) {
				return nil, errors.New("invalid account type")
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		withdrawJSON, err := json.Marshal(&models.WithdrawDTO{
			AccountNumber: 1,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/withdraw", mockTransactionController.Withdraw)

		req, err := http.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(withdrawJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InsufficientBalance", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewWithdrawMock: func(withdraw *models.Withdraw) (*models.Transaction, error) {
				return nil, errors.New("invalid insufficient balance")
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		withdrawJSON, err := json.Marshal(&models.WithdrawDTO{
			AccountNumber: 1,
			Amount:        100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/withdraw", mockTransactionController.Withdraw)

		req, err := http.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(withdrawJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}

func TestTransactionController_Payment(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewPaymentMock: func(payment *models.Payment) (*models.Transaction, error) {
				return &models.Transaction{
					AccountNumber:   payment.SenderAccount,
					Amount:          payment.Amount,
					TransactionType: types.Payment,
					CreatedAt:       time.Now(),
				}, nil
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		paymentJSON, err := json.Marshal(&models.PaymentDTO{
			SenderAccount:   1,
			ReceiverAccount: 2,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, types.AccountNumber(1), actualT.AccountNumber)
		assert.Equal(t, types.Payment, actualT.TransactionType)
		assert.Equal(t, http.StatusOK, rr.Code)

	})

	t.Run("InvalidAccountNumberForSender", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewPaymentMock: func(payment *models.Payment) (*models.Transaction, error) {
				return nil, errors.New("invalid account number")
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		paymentJSON, err := json.Marshal(&models.PaymentDTO{
			SenderAccount:   1,
			ReceiverAccount: 2,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidAccountNumberForReceiver", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewPaymentMock: func(payment *models.Payment) (*models.Transaction, error) {
				return nil, errors.New("invalid account number")
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		paymentJSON, err := json.Marshal(&models.PaymentDTO{
			SenderAccount:   1,
			ReceiverAccount: 2,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidAccountNumberForBoth", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewPaymentMock: func(payment *models.Payment) (*models.Transaction, error) {
				return nil, errors.New("invalid account number")
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		paymentJSON, err := json.Marshal(&models.PaymentDTO{
			SenderAccount:   1,
			ReceiverAccount: 2,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("InvalidAccountTypeForSender", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewPaymentMock: func(payment *models.Payment) (*models.Transaction, error) {
				return nil, errors.New("sender must be individual and receiver must be corporate")
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		paymentJSON, err := json.Marshal(&models.PaymentDTO{
			SenderAccount:   1,
			ReceiverAccount: 2,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("InvalidAccountTypeForReceiver", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewPaymentMock: func(payment *models.Payment) (*models.Transaction, error) {
				return nil, errors.New("sender must be individual and receiver must be corporate")
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		paymentJSON, err := json.Marshal(&models.PaymentDTO{
			SenderAccount:   1,
			ReceiverAccount: 2,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("InvalidAccountTypeForBoth", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewPaymentMock: func(payment *models.Payment) (*models.Transaction, error) {
				return nil, errors.New("sender must be individual and receiver must be corporate")
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		paymentJSON, err := json.Marshal(&models.PaymentDTO{
			SenderAccount:   1,
			ReceiverAccount: 2,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("InsufficientBalance", func(t *testing.T) {

		mockTransactionServ := mockTransactionService{
			NewPaymentMock: func(payment *models.Payment) (*models.Transaction, error) {
				return nil, errors.New("insufficient balance")
			},
		}
		mockTransactionController := NewTransactionController(mockTransactionServ)

		paymentJSON, err := json.Marshal(&models.PaymentDTO{
			SenderAccount:   1,
			ReceiverAccount: 2,
			Amount:          100,
		})
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.POST("/payment", mockTransactionController.Payment)

		req, err := http.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentJSON))
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		var actualT *models.TransactionDTO
		err = json.NewDecoder(rr.Body).Decode(&actualT)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}
