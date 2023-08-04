package controllers

import (
	"github.com/CaioLuColaco/etherum-transactions/database"
	"github.com/CaioLuColaco/etherum-transactions/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
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
