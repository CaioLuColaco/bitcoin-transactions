package main

import (
	"github.com/CaioLuColaco/etherum-transactions/database"
	"github.com/CaioLuColaco/etherum-transactions/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}
