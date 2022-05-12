package api

import (
	"github.com/ahmetberke/tringle-candidate-project/configs"
	"github.com/ahmetberke/tringle-candidate-project/internal/api/controllers"
	"github.com/ahmetberke/tringle-candidate-project/internal/cache"
	"github.com/ahmetberke/tringle-candidate-project/internal/services"
	"github.com/gin-gonic/gin"
)

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

	// Creating services
	accountService := services.NewAccountService(accountCache)

	// Creating controllers
	accountController := controllers.NewAccountController(accountService)

	// Initializing routes
	a.AccountRoutesInitialize(accountController)

	return a
}

func (a *api) Run() error {
	err := a.Router.Run(a.PORT)
	if err != nil {
		return err
	}
	return nil
}
