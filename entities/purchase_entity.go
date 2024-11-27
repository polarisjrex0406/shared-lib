package entities

import (
	"time"
)

type Purchase struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CustomerID uint `json:"customer_id"`
	ProductID  uint `json:"product_id"`

	// Password stores the hashed password for proxy credentials to identify this purchase.
	Password string `json:"pswd" gorm:"unique;column:pswd"`

	// nil means unlimited for numeric settings
	Bandwidth   *int     `json:"bandwidth"`
	TrafficLeft *int     `json:"traffic_left"`
	Duration    *int     `json:"duration"`
	IPCount     *int     `json:"ip_count,omitempty"`
	IPs         []string `json:"ips,omitempty" gorm:"serializer:json"`
	Threads     *int     `json:"threads,omitempty"`

	// Non-numeric settings
	Region Region `json:"region"`

	StartAt  time.Time `json:"start_at"`
	ExpireAt time.Time `json:"expire_at"`
}

// TableName overrides the default table name
func (Purchase) TableName() string {
	return "tbl_purchases"
}
