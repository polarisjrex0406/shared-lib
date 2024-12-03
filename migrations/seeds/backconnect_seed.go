package seeds

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

func Backconnect(db *gorm.DB) error {
	proxies := []entities.Proxy{}
	if err := db.Where("type = ?", string(entities.ProxyBackconnect)).Find(&proxies).Error; err != nil {
		return err
	}

	// Insert to DB
	for _, proxy := range proxies {
		gateway := entities.Backconnect{
			ProxyID: proxy.ID,
		}
		if err := db.FirstOrCreate(&gateway, gateway).Error; err != nil {
			return err
		}
	}
	return nil
}
