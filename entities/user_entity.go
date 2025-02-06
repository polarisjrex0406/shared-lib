package entities

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Email            string `json:"email" gorm:"unique"`
	Firstname        string `json:"firstname" gorm:"default:''"`
	Lastname         string `json:"lastname" gorm:"default:''"`
	Password         string `json:"pswd" gorm:"column:pswd"`
	Role             Role   `json:"role" gorm:"default:'user'"`
	InChargeFieldIDs []uint `json:"incharge_field_ids,omitempty" gorm:"serializer:json"`
}

// TableName overrides the default table name
func (User) TableName() string {
	return "tbl_users"
}
