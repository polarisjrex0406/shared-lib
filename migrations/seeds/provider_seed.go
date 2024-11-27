package seeds

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

func Provider(db *gorm.DB) error {
	name := [...]string{"ttproxy", "dataimpulse", "proxyverse", "databay"}

	countryTargeting := [...]bool{false, false, true, true}
	ipVersion := entities.IPVersion4
	protocols := [...]entities.Protocol{
		entities.ProtocolHTTP,
		entities.ProtocolHTTPS,
		entities.ProtocolSOCKS5,
		entities.ProtocolSOCKS5h}
	stickySession := [...]bool{true, true, true, true}
	bandwidth := 0
	ipCount := [...]int{10000000, 5000000, 70000000, 23700000}
	threads := 0
	regionRange := [...]entities.Region{entities.RegionMixed}

	for i := 0; i < len(name); i++ {
		provider := &entities.Provider{
			Name:             name[i],
			CountryTargeting: countryTargeting[i],
			IPVersion:        ipVersion,
			Protocols:        protocols[:],
			StickySession:    stickySession[i],
			BandwidthMax:     bandwidth,
			IPCountMax:       ipCount[i],
			ThreadsMax:       threads,
			RegionRange:      regionRange[:],
		}
		if err := db.FirstOrCreate(provider, *provider).Error; err != nil {
			return err
		}
	}

	return nil
}
