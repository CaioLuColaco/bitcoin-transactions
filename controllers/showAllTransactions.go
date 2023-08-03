package controllers

import (
	"github.com/CaioLuColaco/bitcoin-transactions/database"
	"github.com/CaioLuColaco/bitcoin-transactions/models"
	"github.com/gin-gonic/gin"
)

func ShowAllTransactions(c *gin.Context) {
	var transactions []models.Transaction
	database.DB.Find(&transactions)
	c.JSON(200, transactions)
}
