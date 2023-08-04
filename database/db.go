package database

import (
	"log"

	"github.com/CaioLuColaco/etherum-transactions/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Error to connect with database")
	}
	DB.AutoMigrate(&models.Transaction{})
}
