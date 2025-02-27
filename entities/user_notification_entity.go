package entities

import "time"

// UserNotification is a struct that represents a notification.
type UserNotification struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	TargetUserIDs []uint `json:"target_user_ids,omitempty" gorm:"serializer:json"`
	ReadUserIDs   []uint `json:"read_user_ids,omitempty" gorm:"serializer:json"`

	Title   string `json:"title"`
	Content string `json:"content"`
}

// TableName overrides the default table name
func (UserNotification) TableName() string {
	return "tbl_user_notifications"
}
