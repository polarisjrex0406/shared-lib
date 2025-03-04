package entities

import "gorm.io/gorm"

// BillingAddress is a struct that represents the billing address of
// a customer, which is used for invoicing and payment processing.
type BillingAddress struct {
	gorm.Model

	// CustomerID is a unique identifier for the customer
	// associated with this billing address.
	CustomerID uint `gorm:"uniqueIndex:idx_customer_billing_address;not null"`
	// Firstname stores the first name of this customer, e.g. P.
	Firstname string `json:"firstname"`
	// Lastname stores the last name of this customer, e.g. Sherman.
	Lastname string `json:"lastname"`
	// Country stores the name of a country for this address,
	// e.g. Australia.
	Country string `json:"country"`
	// StreetAddress stores house number and street for this address,
	// e.g. 42 Wallaby Way.
	StreetAddress string `json:"street_address"`
	// StateAbbr stores abbreviation of a state for this address,
	// e.g. NSW.
	StateAbbr string `json:"state_abbr"`
	// City stores the name of a city for this address, e.g. Sydney.
	City string `json:"city"`
	// Zipcode stores the zip code for this address, e.g. 2000.
	Zipcode string `json:"zipcode"`
}

// TableName overrides the default table name
func (BillingAddress) TableName() string {
	return "tbl_billing_addresses"
}
