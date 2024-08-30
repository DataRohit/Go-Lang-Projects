package models

import (
	"log"

	"github.com/datarohit/go-stock-data-api/pkg/config/dbConfig"
	"github.com/datarohit/go-stock-data-api/pkg/schemas"
)

func GetAllStocks() ([]schemas.Stock, error) {
	var stocks []schemas.Stock

	result := dbConfig.DatabaseConnection.Find(&stocks)
	if result.Error != nil {
		log.Printf("Error retrieving stocks: %v", result.Error)
		return nil, result.Error
	}

	return stocks, nil
}

func GetStockBySymbol(symbol string) (*schemas.Stock, error) {
	var stock schemas.Stock

	result := dbConfig.DatabaseConnection.Where("symbol = ?", symbol).First(&stock)
	if result.Error != nil {
		log.Printf("Error retrieving stock with symbol %s: %v", symbol, result.Error)
		return nil, result.Error
	}

	return &stock, nil
}
