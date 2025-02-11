package entities

import (
	"time"
)

// IssueCategory is a struct that represents a category of a customer support issue.
type IssueCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Name       string   `json:"name" gorm:"unique"`
	IsAdvanced bool     `json:"is_advanced"`
	Keywords   []string `json:"keywords,omitempty" gorm:"serializer:json"`
}

// TableName overrides the default table name
func (IssueCategory) TableName() string {
	return "tbl_issue_categories"
}
