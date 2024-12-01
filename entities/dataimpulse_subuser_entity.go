package entities

import "time"

// DataImpulseSubuser is a struct that represents a sub-user for
// of DataImpulse provider.
type DataImpulseSubuser struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// PurchaseID is a unique identifier of a purchase which uses
	// this sub-user.
	PurchaseID uint `json:"purchase_id"`
	// SubuserID is a unique identifier of this sub-user.
	SubuserID int `json:"subuser_id"`
	// Login is the "login" of sub-user to access reseller API.
	Login string `json:"login"`
	// Password is the password of sub-user to access reseller API.
	Password string `json:"pswd" gorm:"column:pswd"`
	// TotalBalance indicates amount of balance for this sub-user.
	TotalBalance int `json:"total_balance"`
}

// TableName overrides the default table name
func (DataImpulseSubuser) TableName() string {
	return "tbl_dataimpulse_subusers"
}
