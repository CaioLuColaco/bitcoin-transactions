package controllers

import (
	"net/http"

	"github.com/CaioLuColaco/bitcoin-transactions/database"
	"github.com/CaioLuColaco/bitcoin-transactions/models"
	"github.com/gin-gonic/gin"
)

func UpdateTransaction(c *gin.Context) {
	id := c.Params.ByName("id")

	var transaction models.Transaction

	database.DB.First(&transaction, id)
	if transaction.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Transaction not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateTransaction(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	database.DB.Model(&transaction).UpdateColumns(transaction)

	c.JSON(http.StatusOK, transaction)

}
