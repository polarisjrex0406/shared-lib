package seeds

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

func Product(db *gorm.DB) error {
	// 1. Display features
	// IP
	upTo5000IPs := entities.ProductDisplayFeature{Supported: true, Description: "Up to 5,000 IPs"}
	ips15000 := entities.ProductDisplayFeature{Supported: true, Description: "15,000 IPs"}
	ips12000000Plus := entities.ProductDisplayFeature{Supported: true, Description: "12,000,000+ IPs"}
	ips65000000Plus := entities.ProductDisplayFeature{Supported: true, Description: "65,000,000+ IPs"}
	// Bandwidth
	unlimitedBandwidth := entities.ProductDisplayFeature{Supported: true, Description: "Unlimited Bandwidth"}
	// Threads
	unlimitedThreads := entities.ProductDisplayFeature{Supported: true, Description: "Unlimited Threads"}
	// Protocol
	httpsSocks5 := entities.ProductDisplayFeature{Supported: true, Description: "HTTP(S)/SOCKS5"}
	httpsSocks5h := entities.ProductDisplayFeature{Supported: true, Description: "HTTP(S)/SOCKS5h"}
	// Authentication
	userPassAuthentication := entities.ProductDisplayFeature{Supported: true, Description: "User:Pass Authentication"}
	// Rotating & Sticky
	rotatingIPOnEachRequest := entities.ProductDisplayFeature{Supported: true, Description: "Rotating IP on each request"}
	stickySessions := entities.ProductDisplayFeature{Supported: false, Description: "Sticky Sessions"}
	rotatingAndStickyProxies := entities.ProductDisplayFeature{Supported: true, Description: "Rotating & Sticky Proxies"}
	// Region & Location
	mixedRegions := entities.ProductDisplayFeature{Supported: true, Description: "Mixed Regions"}
	usaRegion := entities.ProductDisplayFeature{Supported: true, Description: "USA Region"}
	locations60Plus := entities.ProductDisplayFeature{Supported: true, Description: "60+ Locations"}
	locations195Plus := entities.ProductDisplayFeature{Supported: true, Description: "195+ Locations"}
	countryTargeting := entities.ProductDisplayFeature{Supported: false, Description: "Country Targeting"}
	// IPv6
	datacenterIPv6 := entities.ProductDisplayFeature{Supported: true, Description: "Datacenter IPv6"}
	residentialIPv6 := entities.ProductDisplayFeature{Supported: true, Description: "Residential IPv6"}
	// ISP
	tier1ISP := entities.ProductDisplayFeature{Supported: true, Description: "Tier 1 ISP"}
	tier1SharedISP := entities.ProductDisplayFeature{Supported: true, Description: "Tier 1 SharedISP"}
	datacenterISP := entities.ProductDisplayFeature{Supported: true, Description: "Datacenter ISP"}
	residentialSharedISP := entities.ProductDisplayFeature{Supported: true, Description: "Residential SharedISP"}

	// 2. Product list
	products := []entities.Product{}
	category := entities.Category{}
	var flagPtr bool
	intPtrThreadsRange := 50
	// Find categories
	// Datacenter
	if err := db.Where("abbr = ?", "dc").First(&category).Error; err != nil {
		return err
	}
	// Standard Datacenter
	productStdDc := entities.Product{
		Abbr:                  "std",
		CategoryID:            category.ID,
		Name:                  "Standard Datacenter",
		ProxyServiceType:      entities.ProxyStatic,
		DisplayFeatures:       []entities.ProductDisplayFeature{upTo5000IPs, unlimitedBandwidth, httpsSocks5, userPassAuthentication, rotatingIPOnEachRequest, mixedRegions},
		CountryTargeting:      false,
		IPVersion:             entities.IPVersion4,
		Protocols:             []entities.Protocol{entities.ProtocolHTTP, entities.ProtocolHTTPS, entities.ProtocolSOCKS5},
		StickySession:         false,
		BandwidthUnlimitedMax: nil,
		BandwidthRange:        nil,
		DurationRange:         &entities.NumericSettingRange{Min: 1, Max: 30, InBetween: []int{7}, Interval: nil},
		IPCountRange:          &entities.NumericSettingRange{Min: 1000, Max: 5000, InBetween: []int{2500}, Interval: nil},
		ThreadsRange:          &entities.NumericSettingRange{Min: 200, Max: 500, InBetween: nil, Interval: &intPtrThreadsRange},
		RegionRange:           nil,
		BasePriceRow:          entities.IPCountIndex,
		BasePriceCol:          entities.DurationIndex,
		BasePriceDepth:        entities.ThreadsIndex,
		PriceFormula:          "CEILMULTI((1+0.1*(DEPTH-DEPTHMIN)/DEPTHINTERVAL)*BASEPRICE, 1.0)",
	}
	products = append(products, productStdDc)
	// Premium Datacenter
	productPremDc := entities.Product{
		Abbr:                  "prem",
		CategoryID:            category.ID,
		Name:                  "Premium Datacenter",
		ProxyServiceType:      entities.ProxyBackconnect,
		DisplayFeatures:       []entities.ProductDisplayFeature{ips15000, unlimitedBandwidth, httpsSocks5, userPassAuthentication, rotatingIPOnEachRequest, mixedRegions},
		CountryTargeting:      false,
		IPVersion:             entities.IPVersion4,
		Protocols:             []entities.Protocol{entities.ProtocolHTTP, entities.ProtocolHTTPS, entities.ProtocolSOCKS5},
		StickySession:         false,
		BandwidthUnlimitedMax: nil,
		BandwidthRange:        nil,
		DurationRange:         &entities.NumericSettingRange{Min: 1, Max: 30, InBetween: []int{7}, Interval: nil},
		IPCountRange:          nil,
		ThreadsRange:          &entities.NumericSettingRange{Min: 200, Max: 500, InBetween: nil, Interval: &intPtrThreadsRange},
		RegionRange:           []entities.Region{entities.RegionMixed, entities.RegionUSA},
		BasePriceRow:          entities.RegionIndex,
		BasePriceCol:          entities.DurationIndex,
		BasePriceDepth:        entities.ThreadsIndex,
		PriceFormula:          "CEILMULTI((1+0.1*(DEPTH-DEPTHMIN)/DEPTHINTERVAL)*BASEPRICE, 10.0)",
	}
	products = append(products, productPremDc)

	// Residential
	category = entities.Category{}
	if err := db.Where("abbr = ?", "resi").First(&category).Error; err != nil {
		return err
	}
	// Standard Residential
	ttProxy := entities.Provider{}
	if err := db.Where("name = ?", "ttproxy").First(&ttProxy).Error; err != nil {
		return err
	}
	dataimpulse := entities.Provider{}
	if err := db.Where("name = ?", "dataimpulse").First(&dataimpulse).Error; err != nil {
		return err
	}
	flagPtr = false
	intPtrBandwidthRange := 1
	productStdResi := entities.Product{
		Abbr:                  "std",
		CategoryID:            category.ID,
		Name:                  "Standard Residential",
		ProxyServiceType:      entities.ProxyProvider,
		DisplayFeatures:       []entities.ProductDisplayFeature{ips12000000Plus, unlimitedThreads, httpsSocks5h, locations60Plus, rotatingIPOnEachRequest, userPassAuthentication, countryTargeting, stickySessions},
		CountryTargeting:      false,
		IPVersion:             entities.IPVersion4,
		Protocols:             []entities.Protocol{entities.ProtocolHTTP, entities.ProtocolHTTPS, entities.ProtocolSOCKS5h},
		StickySession:         false,
		BandwidthUnlimitedMax: &flagPtr,
		BandwidthRange:        &entities.NumericSettingRange{Min: 1, Max: 1000, InBetween: nil, Interval: &intPtrBandwidthRange},
		DurationRange:         nil,
		IPCountRange:          nil,
		ThreadsRange:          nil,
		RegionRange:           nil,
		BasePriceRow:          entities.BandwidthIndex,
		BasePriceCol:          "",
		BasePriceDepth:        "",
		PriceFormula:          "CEILMULTI(BASEPRICE*ROW, 1.0)",
	}
	productStdResi.ProviderIDs = append(productStdResi.ProviderIDs, ttProxy.ID, dataimpulse.ID)
	products = append(products, productStdResi)
	// Premium Residential
	proxyverse := entities.Provider{}
	if err := db.Where("name = ?", "proxyverse").First(&proxyverse).Error; err != nil {
		return err
	}
	databay := entities.Provider{}
	if err := db.Where("name = ?", "databay").First(&databay).Error; err != nil {
		return err
	}
	productPremResi := entities.Product{
		Abbr:                  "prem",
		CategoryID:            category.ID,
		Name:                  "Premium Residential",
		ProxyServiceType:      entities.ProxyProvider,
		DisplayFeatures:       []entities.ProductDisplayFeature{ips65000000Plus, unlimitedThreads, httpsSocks5h, locations195Plus, rotatingAndStickyProxies, userPassAuthentication},
		CountryTargeting:      true,
		IPVersion:             entities.IPVersion4,
		Protocols:             []entities.Protocol{entities.ProtocolHTTP, entities.ProtocolHTTPS, entities.ProtocolSOCKS5h},
		StickySession:         true,
		BandwidthUnlimitedMax: &flagPtr,
		BandwidthRange:        &entities.NumericSettingRange{Min: 1, Max: 1000, InBetween: nil, Interval: &intPtrBandwidthRange},
		DurationRange:         nil,
		IPCountRange:          nil,
		ThreadsRange:          nil,
		RegionRange:           nil,
		BasePriceRow:          entities.BandwidthIndex,
		BasePriceCol:          "",
		BasePriceDepth:        "",
		PriceFormula:          "CEILMULTI(BASEPRICE*ROW, 0.5)",
	}
	productPremResi.ProviderIDs = []uint{}
	productPremResi.ProviderIDs = append(productPremResi.ProviderIDs, proxyverse.ID, databay.ID)
	products = append(products, productPremResi)

	// IPv6
	category = entities.Category{}
	if err := db.Where("abbr = ?", "ipv6").First(&category).Error; err != nil {
		return err
	}
	// Standard IPv6
	productStdIPv6 := entities.Product{
		Abbr:                  "std",
		CategoryID:            category.ID,
		Name:                  "Standard IPv6",
		ProxyServiceType:      entities.ProxySubnet,
		DisplayFeatures:       []entities.ProductDisplayFeature{datacenterIPv6, unlimitedBandwidth, httpsSocks5, usaRegion, rotatingIPOnEachRequest, userPassAuthentication},
		CountryTargeting:      false,
		IPVersion:             entities.IPVersion6,
		Protocols:             []entities.Protocol{entities.ProtocolHTTP, entities.ProtocolHTTPS, entities.ProtocolSOCKS5},
		StickySession:         false,
		BandwidthUnlimitedMax: nil,
		BandwidthRange:        nil,
		DurationRange:         &entities.NumericSettingRange{Min: 1, Max: 30, InBetween: []int{7}, Interval: nil},
		IPCountRange:          nil,
		ThreadsRange:          &entities.NumericSettingRange{Min: 100, Max: 500, InBetween: []int{250}, Interval: nil},
		RegionRange:           nil,
		BasePriceRow:          entities.ThreadsIndex,
		BasePriceCol:          entities.DurationIndex,
		BasePriceDepth:        "",
		PriceFormula:          "BASEPRICE",
	}
	products = append(products, productStdIPv6)
	// Premium IPv6
	productPremIPv6 := entities.Product{
		Abbr:                  "prem",
		CategoryID:            category.ID,
		Name:                  "Premium IPv6",
		ProxyServiceType:      entities.ProxySubnet,
		DisplayFeatures:       []entities.ProductDisplayFeature{residentialIPv6, unlimitedBandwidth, tier1ISP, usaRegion, rotatingAndStickyProxies, userPassAuthentication},
		CountryTargeting:      false,
		IPVersion:             entities.IPVersion6,
		Protocols:             []entities.Protocol{entities.ProtocolHTTP, entities.ProtocolHTTPS, entities.ProtocolSOCKS5},
		StickySession:         true,
		BandwidthUnlimitedMax: nil,
		BandwidthRange:        nil,
		DurationRange:         &entities.NumericSettingRange{Min: 1, Max: 30, InBetween: []int{7}, Interval: nil},
		IPCountRange:          nil,
		ThreadsRange:          &entities.NumericSettingRange{Min: 100, Max: 500, InBetween: []int{250}, Interval: nil},
		RegionRange:           nil,
		BasePriceRow:          entities.ThreadsIndex,
		BasePriceCol:          entities.DurationIndex,
		BasePriceDepth:        "",
		PriceFormula:          "BASEPRICE",
	}
	products = append(products, productPremIPv6)

	// Shared ISP
	category = entities.Category{}
	if err := db.Where("abbr = ?", "sisp").First(&category).Error; err != nil {
		return err
	}
	infiniteGB := true
	// Standard Shared ISP
	productStdSharedIsp := entities.Product{
		Abbr:                  "std",
		CategoryID:            category.ID,
		Name:                  "Standard Shared ISP",
		ProxyServiceType:      entities.ProxyISPPool,
		DisplayFeatures:       []entities.ProductDisplayFeature{datacenterISP, unlimitedBandwidth, httpsSocks5, userPassAuthentication, rotatingIPOnEachRequest, usaRegion},
		CountryTargeting:      false,
		IPVersion:             entities.IPVersion4,
		Protocols:             []entities.Protocol{entities.ProtocolHTTP, entities.ProtocolHTTPS, entities.ProtocolSOCKS5},
		StickySession:         false,
		BandwidthUnlimitedMax: &infiniteGB,
		BandwidthRange:        &entities.NumericSettingRange{Min: 250, Max: 5000, InBetween: []int{1000}, Interval: nil},
		DurationRange:         nil,
		IPCountRange:          &entities.NumericSettingRange{Min: 10, Max: 1000, InBetween: []int{25, 50, 75, 100, 250, 500}, Interval: nil},
		ThreadsRange:          nil,
		RegionRange:           nil,
		BasePriceRow:          entities.IPCountIndex,
		BasePriceCol:          entities.BandwidthIndex,
		BasePriceDepth:        "",
		PriceFormula:          "BASEPRICE",
	}
	products = append(products, productStdSharedIsp)
	// Premium Shared ISP
	productPremSharedIsp := entities.Product{
		Abbr:                  "prem",
		CategoryID:            category.ID,
		Name:                  "Premium Shared ISP",
		ProxyServiceType:      entities.ProxyISPPool,
		DisplayFeatures:       []entities.ProductDisplayFeature{residentialSharedISP, unlimitedBandwidth, tier1SharedISP, userPassAuthentication, rotatingAndStickyProxies, usaRegion},
		CountryTargeting:      false,
		IPVersion:             entities.IPVersion4,
		Protocols:             []entities.Protocol{entities.ProtocolHTTP, entities.ProtocolHTTPS, entities.ProtocolSOCKS5},
		StickySession:         false,
		BandwidthUnlimitedMax: &infiniteGB,
		BandwidthRange:        &entities.NumericSettingRange{Min: 250, Max: 5000, InBetween: []int{1000}, Interval: nil},
		DurationRange:         nil,
		IPCountRange:          &entities.NumericSettingRange{Min: 10, Max: 1000, InBetween: []int{25, 50, 75, 100, 250, 500}, Interval: nil},
		ThreadsRange:          nil,
		RegionRange:           nil,
		BasePriceRow:          entities.IPCountIndex,
		BasePriceCol:          entities.BandwidthIndex,
		BasePriceDepth:        "",
		PriceFormula:          "BASEPRICE",
	}
	products = append(products, productPremSharedIsp)

	// Seed
	for _, product := range products {
		if err := db.FirstOrCreate(&product, product).Error; err != nil {
			return err
		}
	}
	return nil
}
