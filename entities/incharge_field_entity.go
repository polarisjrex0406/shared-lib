package entities

import (
	"time"
)

type InChargeField struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Topic       string `json:"topic"`
	Description string `json:"description"`
}

// TableName overrides the default table name
func (InChargeField) TableName() string {
	return "tbl_incharge_fields"
}
