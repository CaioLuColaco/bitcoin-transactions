package controllers

import (
	"net/http"

	"github.com/CaioLuColaco/etherum-transactions/database"
	"github.com/CaioLuColaco/etherum-transactions/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// ShowOneTransactionID godoc
// @Summary      Show one transaction registred
// @Description  Route used to get a one transaction registred by ID
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Transaction ID"
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /transaction/{id} [get]
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
