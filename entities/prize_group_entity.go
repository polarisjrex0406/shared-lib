package entities

import "time"

type PrizeGroup struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Enabled   bool      `json:"_enabled" gorm:"default:true"`
	Removed   bool      `json:"_removed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Rarity     PrizeRarity `json:"rarity" gorm:"unique"`
	ChanceRate float64     `json:"chance_rate" gorm:"unique"`
}

// TableName overrides the default table name
func (PrizeGroup) TableName() string {
	return "tbl_prize_groups"
}
