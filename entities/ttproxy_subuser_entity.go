package entities

import "time"

type TTProxySubuser struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	ProxyID uint `json:"proxy_id"`

	SublicenseID uint   `json:"sublicense_id"`
	Key          string `json:"key"`
	Secret       string `json:"secret"`
	ObtainLimit  int    `json:"obtain_limit"`
	TrafficLeft  int64  `json:"traffic_left"`
	IPDuration   int    `json:"ip_duration"`
	Remark       string `json:"remark"`
	TotalTraffic int64  `json:"total_traffic"`
	IPUsed       int    `json:"ip_used"`
}

// TableName overrides the default table name
func (TTProxySubuser) TableName() string {
	return "tbl_ttproxy_subusers"
}
