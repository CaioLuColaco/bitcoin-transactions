package controllers

import (
	"net/http"

	"github.com/CaioLuColaco/bitcoin-transactions/database"
	"github.com/CaioLuColaco/bitcoin-transactions/models"
	"github.com/gin-gonic/gin"
)

func FindTransactionsByBlockNumber(c *gin.Context) {
	blockNumber := c.Params.ByName("blockNumber")

	var transactions []models.Transaction
	if err := database.DB.Find(&transactions, "block_number = ?", blockNumber).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch transactions",
		})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
