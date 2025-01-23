package entities

import "time"

// Newsletter is a struct that represents a newsletter.
type Newsletter struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	TargetEmails []string `json:"target_emails,omitempty" gorm:"serializer:json"`

	Title       string `json:"title"`
	Description string `json:"description"`

	EmailTemplateID uint `json:"email_template_id"`
}

// TableName overrides the default table name
func (Newsletter) TableName() string {
	return "tbl_newsletters"
}
