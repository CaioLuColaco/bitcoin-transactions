package controllers

import (
	"net/http"

	"github.com/CaioLuColaco/etherum-transactions/database"
	"github.com/CaioLuColaco/etherum-transactions/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// FindTransactionsByBlockNumber godoc
// @Summary      Show one transaction registred
// @Description  Route used to get a one transaction registred by BlockNumber
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        BlockNumber   path      int  true  "Transaction BlockNumber"
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /transactions/blockNumber/{BlockNumber} [get]
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
