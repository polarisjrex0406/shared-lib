package entities

import "time"

// BasePrice is a struct that represents base price matrix for
// a product.
type BasePrice struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// ProductID is a unique identifier for the product
	// associated with basic price matrix.
	ProductID uint `json:"product_id"`
	// RowIndex stores a string value for a row of a basic price
	// matrix, e.g. "mixed" or "2500".
	RowIndex string `json:"row_index"`
	// ColIndex stores a string value for a column of a basic price
	// matrix and it can be empty string when there aren't any
	// columns for matrix, e.g. Premium residential product.
	ColIndex string `json:"col_index"`
	// PriceValue is an original price of a product with certain
	// features, e.g. $12.00 when IP count is 2500 and duration is
	// 7 days for Standard datacenter product.
	PriceValue float64 `json:"price_value"`
}

// TableName overrides the default table name
func (BasePrice) TableName() string {
	return "tbl_baseprices"
}
