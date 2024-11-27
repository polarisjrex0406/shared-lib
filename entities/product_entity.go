package entities

import (
	"time"
)

type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Abbr            string                  `json:"abbr"`
	CategoryID      uint                    `json:"category_id"`
	Name            string                  `json:"name" gorm:"unique"`
	DisplayFeatures []ProductDisplayFeature `json:"display_features,omitempty" gorm:"serializer:json"`
	// Services
	ProxyServiceType ProxyServiceType `json:"proxy_service_type"`
	ProviderIDs      []uint           `json:"provider_ids,omitempty" gorm:"serializer:json"`
	// Attributes
	CountryTargeting bool       `json:"country_targeting"`
	IPVersion        IPVersion  `json:"ip_version"`
	Protocols        []Protocol `json:"protocols" gorm:"serializer:json"`
	StickySession    bool       `json:"sticky_session"`
	// Range of numeric settings. nil means not mentioned
	BandwidthUnlimitedMax *bool                `json:"bandwidth_unlimited_max,omitempty"`
	BandwidthRange        *NumericSettingRange `json:"bandwidth_range,omitempty" gorm:"serializer:json"`
	DurationRange         *NumericSettingRange `json:"duration_range,omitempty" gorm:"serializer:json"`
	IPCountRange          *NumericSettingRange `json:"ip_count_range,omitempty" gorm:"serializer:json"`
	ThreadsRange          *NumericSettingRange `json:"threads_range,omitempty" gorm:"serializer:json"`
	// Range of non-numeric settings. nil means not mentioned
	RegionRange []Region `json:"region_range,omitempty" gorm:"serializer:json"`
	// Indexes of base price
	BasePriceRow   BasePriceIndex `json:"base_price_row,omitempty"`
	BasePriceCol   BasePriceIndex `json:"base_price_col,omitempty"`
	BasePriceDepth BasePriceIndex `json:"base_price_depth,omitempty"`
	// Pricing formula
	PriceFormula string `json:"price_formula"`
}

// TableName overrides the default table name
func (Product) TableName() string {
	return "tbl_products"
}
