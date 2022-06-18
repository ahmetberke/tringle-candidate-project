package api

import (
	"github.com/ahmetberke/tringle-candidate-project/configs"
	"github.com/ahmetberke/tringle-candidate-project/internal/api/controllers"
	"github.com/ahmetberke/tringle-candidate-project/internal/cache"
	"github.com/ahmetberke/tringle-candidate-project/internal/services"
	"github.com/gin-gonic/gin"
)

// this struct has a PORT for custom host setting
// and it has a router that distributes the routes to the relevant controllers
// the server is run through this router
// TODO: Add balance and amount is negative or zero check
type api struct {
	PORT   string
	Router *gin.Engine
}

func NewAPI() *api {
	a := &api{
		PORT:   configs.Manager.HostCredentials.PORT,
		Router: gin.Default(),
	}

	// Creating cache layers
	accountCache := cache.NewAccountCache()
	transactionCache := cache.NewTransactionCache()

	// Creating services
	accountService := services.NewAccountService(accountCache)
	transactionService := services.NewTransactionService(accountCache, transactionCache)

	// Creating controllers
	accountController := controllers.NewAccountController(accountService)
	transactionController := controllers.NewTransactionController(transactionService)

	// Initializing routes
	a.AccountRoutesInitialize(accountController)
	a.TransactionRoutesInitialize(transactionController)

	return a
}

// Run starts the server with the port set in the api
func (a *api) Run() error {
	err := a.Router.Run(a.PORT)
	if err != nil {
		return err
	}
	return nil
}
