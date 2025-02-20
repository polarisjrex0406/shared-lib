package entities

import (
	"time"
)

type Attachment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	SupportMessageID uint   `json:"support_message_id"`
	FileName         string `json:"file_name"`
	FileURL          string `json:"file_url"`
	FileSize         int64  `json:"file_size"`
	FileType         string `json:"file_type"`
}

// TableName overrides the default table name
func (Attachment) TableName() string {
	return "tbl_attachments"
}
