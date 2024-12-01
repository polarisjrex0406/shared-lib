package entities

import "time"

type Backconnect struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Scheme   string `json:"scheme"`
	Username string `json:"username"`
	Password string `json:"pswd" gorm:"column:pswd"`
	Host     string `json:"host"`
	Port     uint   `json:"port"`

	Region Region `json:"region"`
}

// TableName overrides the default table name
func (Backconnect) TableName() string {
	return "tbl_backconnect"
}
