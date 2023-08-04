package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/CaioLuColaco/bitcoin-transactions/database"
	"github.com/CaioLuColaco/bitcoin-transactions/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// CreateTransaction godoc
// @Summary      Create one transaction
// @Description  Route used to create a one transaction
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        hash   path      string  true  "Transaction Hash"
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /transaction/{hash} [post]
func CreateTransaction(c *gin.Context) {
	hash := c.Params.ByName("hash")

	var findTransaction models.Transaction
	if err := database.DB.Where("hash = ?", hash).First(&findTransaction).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Transaction already exists",
		})
		return
	}

	destURL := "https://mainnet.infura.io/v3/4355c5add5574242b01ce888c231f7bf"
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionByHash",
		"params":  []interface{}{hash},
		"id":      1,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao serializar JSON"})
		return
	}

	resp, err := http.Post(destURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao fazer a requisição"})
		return
	}
	defer resp.Body.Close()

	var responseData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao decodificar a resposta JSON"})
		return
	}

	var transaction models.Transaction
	responseResult, _ := responseData["result"].(map[string]interface{})

	transaction.BlockHash = getStringFromResponse(responseResult, "blockHash")
	transaction.BlockNumber = getStringFromResponse(responseResult, "blockNumber")
	transaction.ChainId = getStringFromResponse(responseResult, "chainId")
	transaction.From = getStringFromResponse(responseResult, "from")
	transaction.Gas = getStringFromResponse(responseResult, "gas")
	transaction.GasPrice = getStringFromResponse(responseResult, "gasPrice")
	transaction.Hash = getStringFromResponse(responseResult, "hash")
	transaction.MaxFeePerGas = getStringFromResponse(responseResult, "maxFeePerGas")
	transaction.MaxPriorityFeePerGas = getStringFromResponse(responseResult, "maxPriorityFeePerGas")
	transaction.Nonce = getStringFromResponse(responseResult, "nonce")
	transaction.R = getStringFromResponse(responseResult, "r")
	transaction.To = getStringFromResponse(responseResult, "to")
	transaction.TransactionIndex = getStringFromResponse(responseResult, "transactionIndex")
	transaction.Tp = getStringFromResponse(responseResult, "type")
	transaction.V = getStringFromResponse(responseResult, "v")
	transaction.Value = getStringFromResponse(responseResult, "value")

	database.DB.Create(&transaction)

	c.JSON(http.StatusOK, transaction)
}

func getStringFromResponse(data map[string]interface{}, key string) string {
	value, ok := data[key]
	if !ok || value == nil {
		return ""
	}
	strValue, _ := value.(string)
	return strValue
}
