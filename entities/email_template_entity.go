package entities

import "time"

type EmailTemplate struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Name     string            `json:"name"`
	MetaData map[string]string `json:"meta_data,omitempty" gorm:"serializer:json"`
}

// TableName overrides the default table name
func (EmailTemplate) TableName() string {
	return "tbl_email_templates"
}
