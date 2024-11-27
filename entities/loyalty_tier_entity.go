package entities

import "time"

type LoyaltyTier struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CommissionRate float64  `json:"commission_rate" gorm:"default:0.0"`
	DailySpins     int      `json:"daily_spins" gorm:"default:0"`
	DiscountRate   float64  `json:"discount_rate" gorm:"default:0.0"`
	DiscountCap    *float64 `json:"discount_cap,omitempty"`
	Points         int      `json:"points"`
	Rank           string   `json:"rank"`
}

// TableName overrides the default table name
func (LoyaltyTier) TableName() string {
	return "tbl_loyalty_tiers"
}
