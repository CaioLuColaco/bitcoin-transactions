package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTransactionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Put transaction",
	})
}
