package entities

import "time"

// AuthInfo is a struct that represents authentication-related
// information for a customer.
type AuthInfo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// CustomerID is a unique identifier for the customer
	// associated with this authentication information.
	CustomerID uint `json:"customer_id" gorm:"unique"`
	// Password stores the hashed password for this account.
	Password string `json:"pswd" gorm:"column:pswd"`
	// APIKey is used for authenticating API requests associated
	// with this account.
	APIKey string `json:"api_key"`
	// EmailVerified indicates whether the customer's email address
	// has been verified.
	EmailVerified bool `json:"email_verified"`
	// TFAPassed indicates whether the user has successfully
	// completed 2FA.
	TFAPassed bool `json:"tfa_passed"`
}

// TableName overrides the default table name
func (AuthInfo) TableName() string {
	return "tbl_authinfo"
}
