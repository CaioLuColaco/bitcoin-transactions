package main

import (
	"github.com/CaioLuColaco/bitcoin-transactions/database"
	"github.com/CaioLuColaco/bitcoin-transactions/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
