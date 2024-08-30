package utils

import (
	"fmt"

	"github.com/datarohit/go-stock-data-api/pkg/schemas"
)

func ValidateStock(stock schemas.Stock) error {
	if stock.Issuer == "" {
		return fmt.Errorf("issuer is required")
	}

	if stock.Symbol == "" {
		return fmt.Errorf("symbol is required")
	}

	if stock.Currency == "" {
		return fmt.Errorf("currency is required")
	}

	if stock.Volume <= 0 {
		return fmt.Errorf("volume must be greater than zero")
	}

	if stock.OutstandingShares <= 0 {
		return fmt.Errorf("outstanding shares must be greater than zero")
	}

	if stock.MarketCap <= 0 {
		return fmt.Errorf("market cap must be greater than zero")
	}

	if stock.DividendRate <= 0 {
		return fmt.Errorf("dividend rate must be greater than zero")
	}

	return nil
}
