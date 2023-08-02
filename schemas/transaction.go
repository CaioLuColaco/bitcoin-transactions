package schemas

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	exOrdId int64
	clOrdId string
	ordType string
	ordStatus string
	side string
	price float64
	text string
	symbol string
	lastShares float64
	lastPx float64
	leavesQty int64
	cumQty float64
	avgPx float64
	timestamp int64
}