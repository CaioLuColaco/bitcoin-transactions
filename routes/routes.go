package routes

import (
	"github.com/CaioLuColaco/bitcoin-transactions/controllers"
	"github.com/gin-gonic/gin"
	docs "github.com/CaioLuColaco/bitcoin-transactions/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	r.GET("/transactions", controllers.ShowAllTransactions)
	r.GET("/transaction/:id", controllers.ShowOneTransactionID)
	r.POST("/transaction/:hash", controllers.CreateTransaction)
	r.PATCH("/transaction/:id", controllers.UpdateTransaction)
	r.DELETE("/transaction/:id", controllers.DeleteTransaction)
	r.GET("/transactions/blockNumber/:blockNumber", controllers.FindTransactionsByBlockNumber)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
