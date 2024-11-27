package entities

import (
	"time"
)

// Balance is a struct that represents balance status for a customer.
type Balance struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// CustomerID is a unique identifier for the customer
	// associated with this balance.
	CustomerID uint `json:"customer_id" gorm:"unique"`
	// Currency is the type of currency for this balance,
	// e.g. "USD".
	Currency Currency `json:"currency"`
	// Current is the amount of money currently available to the customer.
	Current float64 `json:"current"`
	// Pending is the amount of money that is not yet available for use, typically due to pending transactions or holds.
	Pending float64 `json:"pending"`
	// Total is the sum of the current and pending.
	Total float64 `json:"total"`
}

// TableName overrides the default table name
func (Balance) TableName() string {
	return "tbl_balances"
}
