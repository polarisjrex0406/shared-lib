package entities

import "time"

type ReferralEarning struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

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
