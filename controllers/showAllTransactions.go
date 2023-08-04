package controllers

import (
	"github.com/CaioLuColaco/bitcoin-transactions/database"
	"github.com/CaioLuColaco/bitcoin-transactions/models"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"github.com/gin-gonic/gin"
)

// ShowAllTransactions godoc
// @Summary      Show all transactions registred
// @Description  Route used to list all transactions registred
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /transactions [get]
func ShowAllTransactions(c *gin.Context) {
	var transactions []models.Transaction
	database.DB.Find(&transactions)
	c.JSON(200, transactions)
}
