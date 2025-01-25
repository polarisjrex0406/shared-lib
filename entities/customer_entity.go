package entities

import (
	"gorm.io/gorm"
)

// Customer is a struct that represents customer data like email, loyalty points, settings etc.
type Customer struct {
	gorm.Model

	// Email stores a unique email address of this customer.
	Email string `json:"email" gorm:"unique"`
	// PhoneNum stores a unique phone number of this customer or empty string.
	PhoneNum string `json:"phone_num"`
	// Points indicates loyalty points that this customer earned.
	Points int `json:"points"`
	// UsedSpins indicates spin counts that this customer used for this day.
	UsedSpins int `json:"used_spins"`
	// ReferrerID indicates the ID of a customer who has invited this customer.
	ReferrerID *uint `json:"referrer_id"`
	// EnabledTFA indicates whether enables 2FA for this customer or not.
	EnabledTFA bool `json:"enabled_tfa"`
	// SubscribeNL indicates whether this customer will receive updates, information, or content via email or not.
	SubscribeNL bool `json:"subscribe_nl"`
	// NotifyExpire indicates whether this customer will receive notifications about purchase expiring or not.
	NotifyExpire bool `json:"notify_expire"`
	// ProfileName stores a unique name of this customer, used in proxy credentials and referral.
	ProfileName string `json:"profile_name" gorm:"unique"`
}

// TableName overrides the default table name
func (Customer) TableName() string {
	return "tbl_customers"
}
