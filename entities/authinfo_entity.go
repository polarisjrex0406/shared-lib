package entities

import (
	"gorm.io/gorm"
)

// AuthInfo is a struct that represents authentication-related
// information for a customer.
type AuthInfo struct {
	gorm.Model

	CustomerID    uint   `gorm:"uniqueIndex:idx_customer_auth_info;not null"`
	Password      string `json:"password"`
	APIKey        string `json:"api_key"`
	EmailVerified bool   `json:"email_verified"`
	MFAPassed     bool   `json:"mfa_passed"`
}

// TableName overrides the default table name
func (AuthInfo) TableName() string {
	return "tbl_authinfo"
}
