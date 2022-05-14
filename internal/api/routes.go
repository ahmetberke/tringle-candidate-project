package api

import "github.com/ahmetberke/tringle-candidate-project/internal/api/controllers"

func (a *api) AccountRoutesInitialize(c *controllers.AccountController) {
	// implementing account routes
	ag := a.Router.Group("/account")
	{
		ag.POST("/", c.Create)
		ag.GET("/:accountNumber", c.Get)
	}
}

func (a *api) TransactionRoutesInitialize(c *controllers.TransactionController) {
	// implementing account routes
	a.Router.POST("/payment", c.Payment)
	a.Router.POST("/deposit", c.Deposit)
	a.Router.POST("/withdraw", c.Withdraw)
	a.Router.GET("/accounting/:accountNumber", c.GetTransactionHistory)
}
