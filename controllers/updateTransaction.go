package controllers

import (
	"net/http"

	"github.com/CaioLuColaco/etherum-transactions/database"
	"github.com/CaioLuColaco/etherum-transactions/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// UpdateTransaction godoc
// @Summary      Update one transaction
// @Description  Route used to update a one transaction by ID
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Transaction ID"
// @Param        transaction   body    models.Transaction  true  "Model of Transaction"
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /transaction/{id} [patch]
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
