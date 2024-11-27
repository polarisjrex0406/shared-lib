package entities

import (
	"time"
)

// Category is a struct that represents a category of a proxy
// product.
type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Name stores the name of this category to be displayed,
	// e.g. "Datacenter" or "Shared ISP".
	Name string `json:"name" gorm:"unique"`
	// Abbr stores abbreviation of this category to be used in
	// proxy credentials, e.g. "dc" in "user-prem-dc-1111".
	Abbr string `json:"abbr" gorm:"unique"`
}

// TableName overrides the default table name
func (Category) TableName() string {
	return "tbl_categories"
}
