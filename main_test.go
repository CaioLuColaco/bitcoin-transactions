package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/CaioLuColaco/etherum-transactions/controllers"
	"github.com/CaioLuColaco/etherum-transactions/database"
	"github.com/CaioLuColaco/etherum-transactions/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int
var HASH string
var BLOCKNUMBER string

func CreateTransaction() {
	transaction := models.Transaction{Hash: "1234567890", BlockNumber: "0987654321"}
	database.DB.Create(&transaction)
	ID = int(transaction.ID)
	HASH = transaction.Hash
	BLOCKNUMBER = transaction.BlockNumber
}

func DeleteTransaction() {
	var transaction models.Transaction
	database.DB.Delete(&transaction, ID)
}

func SetUpRoutesTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func TestShowAllTransactionHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateTransaction()
	defer DeleteTransaction()
	r := SetUpRoutesTest()
	r.GET("/transactions", controllers.ShowAllTransactions)
	req, _ := http.NewRequest("GET", "/transactions", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "The status must be the same")
}

func TestShowOneTransactionIDHandler(t *testing.T) {

	database.ConnectToDatabase()

	CreateTransaction()
	defer DeleteTransaction()

	r := SetUpRoutesTest()
	r.GET("/transaction/:id", controllers.ShowOneTransactionID)
	routeTest := `/transaction/` + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", routeTest, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var transactionTest models.Transaction
	json.Unmarshal(res.Body.Bytes(), &transactionTest)

	assert.Equal(t, "1234567890", transactionTest.Hash, "The values must be the same")
	assert.Equal(t, "0987654321", transactionTest.BlockNumber, "The values must be the same")
}

func TestFindTransactionsByBlockNumber(t *testing.T) {

	database.ConnectToDatabase()

	CreateTransaction()
	defer DeleteTransaction()

	r := SetUpRoutesTest()
	r.GET("/transactions/blockNumber/:blockNumber", controllers.FindTransactionsByBlockNumber)
	routeTest := `/transactions/blockNumber/` + BLOCKNUMBER
	req, _ := http.NewRequest("GET", routeTest, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var transactionTest models.Transaction
	json.Unmarshal(res.Body.Bytes(), &transactionTest)

	assert.Equal(t, http.StatusOK, res.Code, "The status must be the same")
}

func TestDeleteTransactionHandler(t *testing.T) {

	database.ConnectToDatabase()

	CreateTransaction()

	r := SetUpRoutesTest()
	r.DELETE("/transaction/:id", controllers.DeleteTransaction)
	routeTest := `/transaction/` + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", routeTest, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "The status must be the same")

}

func TestUpdateTransactionHandler(t *testing.T) {

	database.ConnectToDatabase()

	CreateTransaction()
	defer DeleteTransaction()

	r := SetUpRoutesTest()
	r.PATCH("/transaction/:id", controllers.UpdateTransaction)

	transaction := models.Transaction{Hash: "0987654321", BlockNumber: "1234567890"}
	transactionJson, _ := json.Marshal(transaction)
	routeTest := `/transaction/` + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", routeTest, bytes.NewBuffer(transactionJson))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var transactionUpdated models.Transaction
	json.Unmarshal(res.Body.Bytes(), &transactionUpdated)

	assert.Equal(t, "0987654321", transactionUpdated.Hash, "The values must be the same")
	assert.Equal(t, "1234567890", transactionUpdated.BlockNumber, "The values must be the same")
}
