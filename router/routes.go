package router

import (
	"github.com/CaioLuColaco/bitcoin-transactions/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/transaction", handler.FindOneTransactionHandler)

		v1.POST("/transaction", handler.CreateTransactionHandler)

		v1.DELETE("/transaction", handler.DeleteTransactionHandler)

		v1.PUT("/transaction", handler.UpdateTransactionHandler)
		
		v1.GET("/transactions", handler.FindAllTransactionHandler)
	}
}
