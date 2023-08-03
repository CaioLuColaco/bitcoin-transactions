package controllers

import (
	"net/http"

	"github.com/CaioLuColaco/bitcoin-transactions/database"
	"github.com/CaioLuColaco/bitcoin-transactions/models"
	"github.com/gin-gonic/gin"
)

func ShowOneTransactionID(c *gin.Context) {
	var transaction models.Transaction
	id := c.Params.ByName("id")
	database.DB.First(&transaction, id)

	if transaction.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Transaction not found",
		})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
