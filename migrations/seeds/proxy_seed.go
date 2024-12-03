package seeds

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

func Proxy(db *gorm.DB) error {
	proxies := make([]entities.Proxy, 0)
	// 1. Backconnect gateways
	//// Mixed region
	host := "169.197.83.75"
	ports := []uint{
		6230, 6351, 6769, 8073, 15978, 15979, 15980, 15981, 15982, 16006, 16007, 16008, 16009, 16010, 16011,
	}
	username := "6la1gj"
	password := "pguaj172a"
	proxy := entities.Proxy{Type: entities.ProxyBackconnect, Host: host, Username: username, Password: password, Region: entities.RegionMixed, PurchaseID: nil}
	for _, port := range ports {
		proxy.Port = port
		proxies = append(proxies, proxy)
	}
	//// USA region
	host = "169.197.83.74"
	ports = []uint{
		6004, 18813, 18814, 18815, 18816, 18817, 18818, 18819, 18820, 18821, 18822, 18823, 18824, 18825, 18826,
	}
	username = "ga12a"
	password = "haug82hf"
	proxy = entities.Proxy{Type: entities.ProxyBackconnect, Host: host, Username: username, Password: password, Region: entities.RegionUSA, PurchaseID: nil}
	for _, port := range ports {
		proxy.Port = port
		proxies = append(proxies, proxy)
	}

	// 2. Insert to DB
	for _, proxy := range proxies {
		if err := db.FirstOrCreate(&proxy, proxy).Error; err != nil {
			return err
		}
	}
	return nil
}
