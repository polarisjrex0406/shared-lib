package entities

import (
	"time"

	"gorm.io/gorm"
)

type ReferralEarning struct {
	gorm.Model

	CommissionRate float64   `json:"commission_rate"`
	Earned         float64   `json:"earned"`
	OrderDate      time.Time `json:"order_date"`
	Payment        float64   `json:"payment"`
	RefereeID      uint      `json:"referee_id"`
	CustomerID     uint      `json:"customer_id"`
}

// TableName overrides the default table name
func (ReferralEarning) TableName() string {
	return "tbl_referral_earnings"
}
