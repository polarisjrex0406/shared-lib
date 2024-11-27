package entities

import "time"

type TTProxySubuser struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	PurchaseID uint   `json:"purchase_id"`
	Key        string `json:"key"`
	Secret     string `json:"secret"`
	Traffic    int    `json:"traffic"`
}

// TableName overrides the default table name
func (TTProxySubuser) TableName() string {
	return "tbl_ttproxy_subusers"
}
