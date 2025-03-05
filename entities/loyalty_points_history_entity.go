package entities

import (
	"time"

	"gorm.io/gorm"
)

// LoyaltyPointsHistory is a struct that represents a history of spending balance.
type LoyaltyPointsHistory struct {
	gorm.Model

	// BalanceSpent is the amount of balance that the customer spends for that time.
	BalanceSpent float64 `json:"balance_spent"`
	// PointsEarned is number of points that the customer earns for that time.
	PointsEarned int `json:"points_earned"`
	// TotalPoints is number of total points that the customer has earned.
	TotalPoints int `json:"total_points"`
	// SpentDate indicates the date when the customer spends balance.
	SpentDate time.Time `json:"spent_date"`
	// CustomerID is a unique identifier of a customer who earns points by spending balance.
	CustomerID uint `json:"customer_id"`
}

// TableName overrides the default table name
func (LoyaltyPointsHistory) TableName() string {
	return "tbl_loyalty_points_histories"
}
