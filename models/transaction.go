package models

import (
	"gorm.io/gorm"
	"gopkg.in/validator.v2"
)

type Transaction struct {
	gorm.Model
	BlockHash            string `json:"blockHash"`
	BlockNumber          string `json:"blockNumber"`
	ChainId              string `json:"chainId"`
	From                 string `json:"from"`
	Gas                  string `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	Hash                 string `json:"hash" gorm:"unique" validate:"nonzero"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	Nonce                string `json:"nonce"`
	R                    string `json:"r"`
	To                   string `json:"to"`
	TransactionIndex     string `json:"transactionIndex"`
	Tp                   string `json:"type"`
	V                    string `json:"v"`
	Value                string `json:"value"`
}

func ValidateTransaction(transaction *Transaction) error {
	if err := validator.Validate(transaction); err != nil {
		return err
	}
	return nil
}