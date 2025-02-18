package entities

import (
	"time"
)

type SupportTicket struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CustomerID uint                `json:"customer_id"`
	ReplyEmail string              `json:"reply_email"`
	Topic      string              `json:"topic"`
	Subject    string              `json:"subject"`
	Status     SupportTicketStatus `json:"status"`
	OpenedAt   time.Time           `json:"opened_at"`
	ClosedAt   time.Time           `json:"closed_at"`
}

// TableName overrides the default table name
func (SupportTicket) TableName() string {
	return "tbl_support_tickets"
}
