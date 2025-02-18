package entities

import (
	"time"
)

type SupportMessage struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	SupportTicketID uint      `json:"support_ticket_id"`
	SenderType      string    `json:"sender_type"`
	SenderID        uint      `json:"sender_id"`
	HtmlContent     string    `json:"html_content"`
	AttachmentID    *uint     `json:"attachment_id"`
	SentAt          time.Time `json:"sent_at"`
}

// TableName overrides the default table name
func (SupportMessage) TableName() string {
	return "tbl_support_messages"
}
