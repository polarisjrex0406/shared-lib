package entities

import "gorm.io/gorm"

// Balance is a struct that represents balance status for a customer.
type Balance struct {
	gorm.Model

	// CustomerID is a unique identifier for the customer
	// associated with this balance.
	CustomerID uint `gorm:"uniqueIndex:idx_customer_balance;not null"`
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
