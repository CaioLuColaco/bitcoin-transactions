package routes

import (
	"github.com/CaioLuColaco/bitcoin-transactions/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/transactions", controllers.ShowAllTransactions)
	r.GET("/transaction/:id", controllers.ShowOneTransactionID)
	r.POST("/transaction/:hash", controllers.CreateTransaction)
	r.PATCH("/transaction/:id", controllers.UpdateTransaction)
	r.DELETE("/transaction/:id", controllers.DeleteTransaction)
	
	r.GET("/transactions/blockNumber/:blockNumber", controllers.FindTransactionsByBlockNumber)
	r.GET("/transactions/from/:from", controllers.FindTransactionsByFrom)

	r.Run()
}
