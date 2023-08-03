package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializePostgresql()

	if err != nil {
		return fmt.Errorf("error initializing postgresql: %v", err)
	}

	return nil
}

func GetPostgresql() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
