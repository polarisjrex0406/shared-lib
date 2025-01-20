package entities

import "time"

type TTProxySubuser struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	ProxyID uint `json:"proxy_id"`
	Traffic int64  `json:"traffic"`
}

// TableName overrides the default table name
func (TTProxySubuser) TableName() string {
	return "tbl_ttproxy_subusers"
}
