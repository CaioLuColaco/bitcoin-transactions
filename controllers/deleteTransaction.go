package controllers

import (
	"net/http"

	"github.com/CaioLuColaco/bitcoin-transactions/database"
	"github.com/CaioLuColaco/bitcoin-transactions/models"
	"github.com/gin-gonic/gin"
)

func DeleteTransaction(c *gin.Context) {
	var transaction models.Transaction
	id := c.Params.ByName("id")
	database.DB.Delete(&transaction, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}
