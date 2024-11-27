package entities

import "time"

type StaticProxy struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	IP string `json:"ip"`
}

// TableName overrides the default table name
func (StaticProxy) TableName() string {
	return "tbl_static_proxies"
}
