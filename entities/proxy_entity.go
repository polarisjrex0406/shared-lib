package entities

import (
	"time"
)

type Proxy struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Type         ProxyServiceType `json:"type"`
	Protocol     Protocol         `json:"protocol"`
	Username     string           `json:"username"`
	Password     string           `json:"pswd" gorm:"column:pswd"`
	Host         string           `json:"host"`
	Port         uint             `json:"port"`
	Region       Region           `json:"region"`
	PurchaseID   *uint            `json:"purchase_id,omitempty"`
	ProviderName string           `json:"provider_name"`
}

// TableName overrides the default table name
func (Proxy) TableName() string {
	return "tbl_proxies"
}
