package main

import (

	"github.com/CaioLuColaco/bitcoin-transactions/config"
	"github.com/CaioLuColaco/bitcoin-transactions/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return 
	}

	router.Initialize()
}
