package entities

import "time"

type ClaimedPrize struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CustomerID uint `json:"customer_id"`

	PrizeID  uint      `json:"prize_id"`
	ClaimAt  time.Time `json:"claim_at"`
	ExpireAt time.Time `json:"expire_at"`
	RedeemAt time.Time `json:"redeem_at"`
}

// TableName overrides the default table name
func (ClaimedPrize) TableName() string {
	return "tbl_claimed_prizes"
}
