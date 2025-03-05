package entities

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model

	CustomerID    uint              `json:"customer_id"`
	Status        TransactionStatus `json:"status"`
	PaymentMethod PaymentMethod     `json:"payment_method"`
}

// TableName overrides the default table name
func (Transaction) TableName() string {
	return "tbl_transactions"
}
