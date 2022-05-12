package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/ahmetberke/tringle-candidate-project/internal/cache"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

func StartServer() {
	accountCache := cache.NewAccountCache()
	accountService := services.NewAccountService(accountCache)
	accountController := NewAccountController(accountService)

	router := gin.Default()
	router.POST("/account", accountController.Create)

	go router.Run(":8080")

	time.Sleep(time.Second)
}

type testCase struct {
	account     *models.Account
	isIncorrect bool
}

func TestAccountController_Create(t *testing.T) {

	StartServer()

	testCases := []*testCase{
		{
			account: &models.Account{
				CurrencyCode: "TRY",
				OwnerName:    "Mustafa Sarigül",
				AccountType:  "individual",
			},
			isIncorrect: false,
		},
		{
			account: &models.Account{
				CurrencyCode: "USDD",
				OwnerName:    "Orkun Aydoğdu",
				AccountType:  "individual",
			},
			isIncorrect: true,
		},
		{
			account: &models.Account{
				CurrencyCode: "USDD",
				OwnerName:    "Selim İçerler",
				AccountType:  "individ",
			},
			isIncorrect: true,
		},
	}

	for i, tCase := range testCases {
		accountJSON, _ := json.Marshal(tCase.account)
		resp, err := http.Post("http://localhost:8080/account", "json", bytes.NewBuffer(accountJSON))
		if err != nil {
			t.Errorf("%s", err.Error())
		}

		errSt := struct {
			Error string `json:"error"`
		}{
			Error: "",
		}

		json.NewDecoder(resp.Body).Decode(&errSt)

		if (errSt.Error != "") != tCase.isIncorrect {
			t.Errorf("error on %d. index, response error message : %s", i, errSt.Error)
		}

	}

}
