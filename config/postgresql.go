package config

import (
	"os"

	"github.com/CaioLuColaco/bitcoin-transactions/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgresql() (*gorm.DB, error) {
	logger = GetLogger("postgresql")
	logger.Infof(os.Getenv("DATABASE_URL"))
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgresql opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Transaction{})
	if err != nil {
		logger.Errorf("postgresql automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
