package entities

import "time"

// Coupon is a struct that represents a promo with code.
type Coupon struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// TargetCustomerIDs stores a list of customer IDs this coupon
	// applies to and is empty when open to all customers.
	TargetCustomerIDs []uint `json:"target_customer_ids,omitempty" gorm:"serializer:json"`
	// RedeemingCustomerIDs stores a list of customer IDs who used
	// this coupon.
	RedeemingCustomerIDs []uint `json:"redeeming_customer_ids,omitempty" gorm:"serializer:json"`
	// ProductIDs stores a list of product IDs this coupon applies
	// to and is empty when open to all products.
	ProductIDs []uint `json:"product_ids,omitempty" gorm:"serializer:json"`
	// Description is explanation about this coupon e.g. "2025 -
	// New Year's Promotion (15%)".
	Description string `json:"description"`
	// DiscountRate indicates discount rate of this coupon.
	DiscountRate float64 `json:"discount_rate"`
	// Code indicates coupon code.
	Code string `json:"code"`
	// StartAt indicates the beginning date when this coupon applies.
	StartAt time.Time `json:"start_at"`
	// ExpireAt indicates the date when this coupon will be expired.
	ExpireAt time.Time `json:"expire_at"`
	// Active indicates whether this coupon can be used or not.
	Active bool `gorm:"default:true" json:"active"`
}

// TableName overrides the default table name
func (Coupon) TableName() string {
	return "tbl_coupons"
}
