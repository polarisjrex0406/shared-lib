package entities

import "time"

// Notification is a struct that represents a notification.
type Notification struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	TargetCustomerIDs []uint `json:"target_customer_ids,omitempty" gorm:"serializer:json"`
	ReadCustomerIDs   []uint `json:"read_customer_ids,omitempty" gorm:"serializer:json"`

	Title   string `json:"title"`
	Content string `json:"content"`
}

// TableName overrides the default table name
func (Notification) TableName() string {
	return "tbl_notifications"
}
