package seeds

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

func Category(db *gorm.DB) error {
	categories := []entities.Category{
		{Name: "Datacenter", Abbr: "dc"},
		{Name: "Residential", Abbr: "resi"},
		{Name: "IPv6", Abbr: "ipv6"},
		{Name: "Shared ISP", Abbr: "sisp"},
	}
	for _, category := range categories {
		if err := db.FirstOrCreate(&category, category).Error; err != nil {
			return err
		}
	}
	return nil
}
