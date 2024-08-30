package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Stock struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Issuer            string    `json:"issuer"`
	Symbol            string    `gorm:"unique;not null" json:"symbol"`
	Currency          string    `json:"currency"`
	Volume            int64     `json:"volume"`
	OutstandingShares int64     `json:"outstanding_shares"`
	MarketCap         int64     `json:"market_cap"`
	DividendRate      int64     `json:"dividend_rate"`
}

func (s *Stock) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}
