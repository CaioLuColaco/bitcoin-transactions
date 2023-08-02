package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindOneTransactionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get transaction",
	})
}

func CreateTransactionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Post transaction",
	})
}

func DeleteTransactionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete transaction",
	})
}

func UpdateTransactionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Put transaction",
	})
}

func FindAllTransactionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get transactions",
	})
}


