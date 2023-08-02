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
