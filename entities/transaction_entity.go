package entities

import (
	"time"
)

type Transaction struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CustomerID    uint              `json:"customer_id"`
	Status        TransactionStatus `json:"status"`
	PaymentMethod PaymentMethod     `json:"payment_method"`
}

// TableName overrides the default table name
func (Transaction) TableName() string {
	return "tbl_transactions"
}
