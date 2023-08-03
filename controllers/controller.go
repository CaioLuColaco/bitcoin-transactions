package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/CaioLuColaco/bitcoin-transactions/database"
	"github.com/CaioLuColaco/bitcoin-transactions/models"
	"github.com/gin-gonic/gin"
)

func ShowAllTransactions(c *gin.Context) {
	var transactions []models.Transaction
	database.DB.Find(&transactions)
	c.JSON(200, transactions)
}

func ShowOneTransaction(c *gin.Context) {

}

func CreateTransaction(c *gin.Context) {
	hash := c.Params.ByName("hash")

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
	transaction.BlockHash = responseData["result"].(map[string]interface{})["blockHash"].(string)
	transaction.BlockNumber = responseData["result"].(map[string]interface{})["blockNumber"].(string)
	transaction.ChainId = responseData["result"].(map[string]interface{})["chainId"].(string)
	transaction.From = responseData["result"].(map[string]interface{})["from"].(string)
	transaction.Gas = responseData["result"].(map[string]interface{})["gas"].(string)
	transaction.GasPrice = responseData["result"].(map[string]interface{})["gasPrice"].(string)
	transaction.Hash = responseData["result"].(map[string]interface{})["hash"].(string)
	transaction.MaxFeePerGas = responseData["result"].(map[string]interface{})["maxFeePerGas"].(string)
	transaction.MaxPriorityFeePerGas = responseData["result"].(map[string]interface{})["maxPriorityFeePerGas"].(string)
	transaction.Nonce = responseData["result"].(map[string]interface{})["nonce"].(string)
	transaction.R = responseData["result"].(map[string]interface{})["r"].(string)
	transaction.To = responseData["result"].(map[string]interface{})["to"].(string)
	transaction.TransactionIndex = responseData["result"].(map[string]interface{})["transactionIndex"].(string)
	transaction.Tp = responseData["result"].(map[string]interface{})["type"].(string)
	transaction.V = responseData["result"].(map[string]interface{})["v"].(string)
	transaction.Value = responseData["result"].(map[string]interface{})["value"].(string)

	database.DB.Create(&transaction)

	c.JSON(http.StatusOK, transaction)
}

func UpdateTransaction(c *gin.Context) {

}

func DeleteTransaction(c *gin.Context) {

}
