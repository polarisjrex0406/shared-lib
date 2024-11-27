package entities

import "time"

type Prize struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Kind    PrizeKind `json:"kind"`
	GroupID uint      `json:"group_id"`

	ProductID *uint   `json:"product_id,omitempty"`
	Bandwidth *int    `json:"bandwidth,omitempty"`
	Duration  *int    `json:"duration,omitempty"`
	IPCount   *int    `json:"ip_count,omitempty"`
	Threads   *int    `json:"threads,omitempty"`
	Region    *Region `json:"region,omitempty"`

	LoyaltyPoints *int     `json:"loyalty_points,omitempty"`
	Credit        *float64 `json:"credit,omitempty"`
	DiscountRate  *float64 `json:"discount_rate,omitempty"`
}

// TableName overrides the default table name
func (Prize) TableName() string {
	return "tbl_prizes"
}
