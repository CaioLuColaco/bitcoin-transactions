package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	BlockHash            string `json:"blockHash"`
	BlockNumber          string `json:"blockNumber"`
	ChainId              string `json:"chainId"`
	From                 string `json:"from"`
	Gas                  string `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	Hash                 string `json:"hash" gorm:"unique"`
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
