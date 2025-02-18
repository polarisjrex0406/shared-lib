package entities

import (
	"time"
)

// IssueTopic is a struct that represents a topic of a customer support issue.
type IssueTopic struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Name     string   `json:"name" gorm:"unique"`
	Keywords []string `json:"keywords,omitempty" gorm:"serializer:json"`
}

// TableName overrides the default table name
func (IssueTopic) TableName() string {
	return "tbl_issue_topics"
}
