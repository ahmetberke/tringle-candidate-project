package api

import "github.com/ahmetberke/tringle-candidate-project/internal/api/controllers"

// AccountRoutesInitialize takes the AccountController as a parameter
// and implements the relevant handlers to the account routes.
func (a *api) AccountRoutesInitialize(c *controllers.AccountController) {
	ag := a.Router.Group("/account")
	{
		ag.POST("/", c.Create)
		ag.GET("/:accountNumber", c.Get)
	}
}

// TransactionRoutesInitialize takes the TransactionController as a parameter
// and implements the relevant handlers to the transaction routes.
func (a *api) TransactionRoutesInitialize(c *controllers.TransactionController) {
	a.Router.POST("/payment", c.Payment)
	a.Router.POST("/deposit", c.Deposit)
	a.Router.POST("/withdraw", c.Withdraw)
	a.Router.GET("/accounting/:accountNumber", c.GetTransactionHistory)
}
