package entities

import "time"

// Invoice is a struct that represents an internal invoice to show purchasing history.
type Invoice struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Number indicates the invoice number unique for a customer.
	Number uint `json:"number"`

	// PurchaseID is a unique identifier of a purchase which is paid by this invoice.
	PurchaseID uint `json:"purchase_id"`

	// IsNewBuying indicates whether this invoice is for buying or adding bandwidth.
	IsNewBuying bool `json:"is_new_buying"`

	// AdditionalBandwidth is the amount of additional bandwidth in unit of GB.
	AdditionalBandwidth *int `json:"additional_bandwidth,omitempty"`

	// Status indicates whether this invoice is paid or unpaid.
	Status InvoiceStatus `json:"status"`

	// Email is the same as the one of Customer.
	Email string `json:"email"`

	// Firstname is the same as the one of BillingAddress.
	Firstname string `json:"firstname"`

	// Lastname is the same as the one of BillingAddress.
	Lastname string `json:"lastname"`

	// Country is the same as the one of BillingAddress.
	Country string `json:"country"`

	// StreetAddress is the same as the one of BillingAddress.
	StreetAddress string `json:"street_address"`

	// StateAbbr is the same as the one of BillingAddress.
	StateAbbr string `json:"state_abbr"`

	// City is the same as the one of BillingAddress.
	City string `json:"city"`

	// Zipcode is the same as the one of BillingAddress.
	Zipcode string `json:"zipcode"`

	// Subtotal is the price before applying discount, credit or coupon.
	Subtotal float64 `json:"subtotal"`

	// Total is the price after applying discount, credit and coupon.
	Total float64 `json:"total"`

	// LoyaltyRank is the same as the one of LoyaltyTier.
	LoyaltyRank string `json:"loyalty_rank"`

	// LoyaltyDiscount is discount rate by the loyalty tier.
	LoyaltyDiscount float64 `json:"loyalty_discount"`

	// DiscountCap is cap of discount by the loyalty tier.
	DiscountCap *float64 `json:"discount_cap,omitempty"`

	// CouponID indicates a unique identifier of a coupon applied to this invoice. nil means no coupon is applied.
	CouponID *uint `json:"coupon_id,omitempty"`

	// ClaimedPrizeID indicates a unique identifier of a claimed prize redeemd to this invoice. nil means no prize is redeemed.
	ClaimedPrizeID *uint `json:"claimed_prize_id,omitempty"`
}

// TableName overrides the default table name
func (Invoice) TableName() string {
	return "tbl_invoices"
}
