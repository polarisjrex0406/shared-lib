package entities

import (
	"time"
)

type ChatSession struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CustomerSupportTicketID uint              `json:"customer_support_ticket_id"`
	StartedAt               time.Time         `json:"started_at"`
	ClosedAt                time.Time         `json:"closed_at"`
	Status                  ChatSessionStatus `json:"status"`
}

// TableName overrides the default table name
func (ChatSession) TableName() string {
	return "tbl_chat_sessions"
}
