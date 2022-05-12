package api

import "github.com/ahmetberke/tringle-candidate-project/internal/api/controllers"

func (a *api) AccountRoutesInitialize(c *controllers.AccountController) {
	ag := a.Router.Group("/account")
	{
		ag.POST("/", c.Create)
		ag.GET("/:accountNumber", c.Get)
	}
}
