package controllers

import (
	"net/http"

	"github.com/CaioLuColaco/etherum-transactions/database"
	"github.com/CaioLuColaco/etherum-transactions/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// DeleteTransaction godoc
// @Summary      Delete one transaction
// @Description  Route used to delete a one transaction by ID
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Transaction ID"
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /transaction/{id} [delete]
func DeleteTransaction(c *gin.Context) {
	var transaction models.Transaction
	id := c.Params.ByName("id")
	database.DB.Delete(&transaction, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}
