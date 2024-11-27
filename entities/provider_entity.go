package entities

import (
	"time"
)

type Provider struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Name string `json:"name" gorm:"unique"`
	// Attributes
	CountryTargeting bool       `json:"country_targeting"`
	IPVersion        IPVersion  `json:"ip_version"`
	Protocols        []Protocol `json:"protocols" gorm:"serializer:json"`
	StickySession    bool       `json:"sticky_session"`
	// Max values of numeric settings
	BandwidthMax int `json:"bandwidth_max"`
	IPCountMax   int `json:"ip_count_max"`
	ThreadsMax   int `json:"threads_max"`
	// Range of non-numeric settings. nil means not mentioned
	RegionRange []Region `json:"region_range,omitempty" gorm:"serializer:json"`
}

// TableName overrides the default table name
func (Provider) TableName() string {
	return "tbl_providers"
}
