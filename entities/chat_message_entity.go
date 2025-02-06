package entities

import (
	"time"
)

type ChatMessage struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	ChatSessionID uint      `json:"chat_session_id"`
	SenderType    string    `json:"sender_type"`
	SenderID      uint      `json:"sender_id"`
	Content       string    `json:"content"`
	SentAt        time.Time `json:"sent_at"`
	IsRead        bool      `json:"is_read"`
}

// TableName overrides the default table name
func (ChatMessage) TableName() string {
	return "tbl_chat_messages"
}
