package routes

import (
	"github.com/CaioLuColaco/bitcoin-transactions/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/transactions", controllers.ShowAllTransactions)
	r.GET("/transaction/:id", controllers.ShowOneTransaction)
	r.POST("/transaction/:hash", controllers.CreateTransaction)
	r.PUT("/transaction/:id", controllers.UpdateTransaction)
	r.DELETE("/transaction/:id", controllers.DeleteTransaction)

	r.Run()
}
