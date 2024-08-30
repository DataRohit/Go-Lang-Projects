package schemas

type Stock struct {
	Issuer            string `json:"issuer"`
	Symbol            string `json:"symbol"`
	Currency          string `json:"currency"`
	Volume            int64  `json:"volume"`
	OutstandingShares int64  `json:"outstanding_shares"`
	MarketCap         int64  `json:"market_cap"`
	DividendRate      int64  `json:"dividend_rate"`
}
