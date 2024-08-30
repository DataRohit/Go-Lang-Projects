package utils

import (
	"fmt"

	"github.com/datarohit/go-stock-data-api/pkg/schemas"
)

func ValidateStock(stock *schemas.Stock, fieldsToValidate ...string) error {
	validateMap := make(map[string]bool)
	for _, field := range fieldsToValidate {
		validateMap[field] = true
	}

	if len(fieldsToValidate) == 0 {
		validateMap["Issuer"] = true
		validateMap["Symbol"] = true
		validateMap["Currency"] = true
		validateMap["Volume"] = true
		validateMap["OutstandingShares"] = true
		validateMap["MarketCap"] = true
		validateMap["DividendRate"] = true
	}

	if validateMap["Issuer"] && stock.Issuer == "" {
		return fmt.Errorf("issuer is required")
	}
	if validateMap["Symbol"] && stock.Symbol == "" {
		return fmt.Errorf("symbol is required")
	}
	if validateMap["Currency"] && stock.Currency == "" {
		return fmt.Errorf("currency is required")
	}
	if validateMap["Volume"] && stock.Volume <= 0 {
		return fmt.Errorf("volume must be greater than zero")
	}
	if validateMap["OutstandingShares"] && stock.OutstandingShares <= 0 {
		return fmt.Errorf("outstanding shares must be greater than zero")
	}
	if validateMap["MarketCap"] && stock.MarketCap <= 0 {
		return fmt.Errorf("market cap must be greater than zero")
	}
	if validateMap["DividendRate"] && stock.DividendRate <= 0 {
		return fmt.Errorf("dividend rate must be greater than zero")
	}

	return nil
}
