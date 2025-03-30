package entities

import (
	"gorm.io/gorm"
)

type CustomerActivityLog struct {
	gorm.Model

	CustomerID uint              `json:"customer_id"`
	EventType  string            `json:"event_type"`
	MetaData   map[string]string `json:"meta_data,omitempty" gorm:"serializer:json"`
}

// TableName overrides the default table name
func (CustomerActivityLog) TableName() string {
	return "tbl_customer_activity_log"
}
