package entities

import (
	"time"
)

type CustomerSupportTicket struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CustomerID      uint                               `json:"customer_id"`
	ReplyEmail      string                             `json:"reply_email"`
	Topic           string                             `json:"topic"`
	Subject         string                             `json:"subject"`
	Description     string                             `json:"description"`
	AttachmentIDs   []uint                             `json:"attachment_ids,omitempty" gorm:"serializer:json"`
	RequestMethod   CustomerSupportTicketRequestMethod `json:"request_method"`    // "form", "email"
	UserID          *uint                              `json:"user_id,omitempty"` // null until assigned
	Status          CustomerSupportTicketStatus        `json:"status"`            // "open", "pending", "in_progress" "closed"
	IsAdvanced      bool                               `json:"is_advanced"`
	RequestedAt     time.Time                          `json:"requested_at"`
	RepliedAt       time.Time                          `json:"replied_at"`
	EmailTemplateID *uint                              `json:"email_template_id,omitempty"`
	ChatSessionID   *uint                              `json:"chat_session_id,omitempty"` // null until chat starts
}

// TableName overrides the default table name
func (CustomerSupportTicket) TableName() string {
	return "tbl_customer_support_tickets"
}
