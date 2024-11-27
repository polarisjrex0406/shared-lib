package seeds

import (
	"fmt"

	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

func BasePrice(db *gorm.DB) error {
	// Get Product ID
	productStdDC := entities.Product{}
	if err := db.Where("name = ?", "Standard Datacenter").First(&productStdDC).Error; err != nil {
		return err
	}
	// Seed Base Price for Standard Datacenter Product
	ipCounts := [...]int{1000, 2500, 5000}
	durations := [...]int{1, 7, 30}
	var stdDCPrice [3][3]float64 = [3][3]float64{
		{3.00, 9.00, 25.00},
		{4.00, 12.00, 35.00},
		{5.00, 15.00, 50.00},
	}
	for i := 0; i < len(ipCounts); i++ {
		rowIndex := fmt.Sprintf("%d", ipCounts[i])
		for j := 0; j < len(durations); j++ {
			colIndex := fmt.Sprintf("%d", durations[j])
			basePriceStdDC := &entities.BasePrice{
				ProductID:  productStdDC.ID,
				RowIndex:   rowIndex,
				ColIndex:   colIndex,
				PriceValue: stdDCPrice[i][j],
			}
			if err := db.FirstOrCreate(basePriceStdDC, *basePriceStdDC).Error; err != nil {
				return err
			}
		}
	}
	// Seed Base Price for Premium Datacenter Product
	productPremDC := entities.Product{}
	if err := db.Where("name = ?", "Premium Datacenter").First(&productPremDC).Error; err != nil {
		return err
	}
	regions := [...]entities.Region{entities.RegionMixed, entities.RegionUSA}
	var premDCPrice [2][3]float64 = [2][3]float64{
		{6.00, 25.00, 80.00},
		{7.00, 30.00, 90.00},
	}
	for i := 0; i < len(regions); i++ {
		rowIndex := regions[i]
		for j := 0; j < len(durations); j++ {
			colIndex := fmt.Sprintf("%d", durations[j])
			basePricePremDC := &entities.BasePrice{
				ProductID:  productPremDC.ID,
				RowIndex:   string(rowIndex),
				ColIndex:   colIndex,
				PriceValue: premDCPrice[i][j],
			}
			if err := db.FirstOrCreate(basePricePremDC, *basePricePremDC).Error; err != nil {
				return err
			}
		}
	}
	// Seed Base Price for Standard/Premium Residential Products
	productStdResi := entities.Product{}
	if err := db.Where("name = ?", "Standard Residential").First(&productStdResi).Error; err != nil {
		return err
	}
	productPremResi := entities.Product{}
	if err := db.Where("name = ?", "Premium Residential").First(&productPremResi).Error; err != nil {
		return err
	}
	gigaBytes := [...]int{1, 10, 25, 50, 100, 250, 500, 1000}
	stdResiPrice := [...]float64{3.00, 2.90, 2.80, 2.70, 2.60, 2.40, 2.20, 2.00}
	premResiPrice := [...]float64{6.00, 5.50, 5.00, 4.50, 4.25, 4.00, 3.75, 3.50}
	for i := 0; i < len(gigaBytes); i++ {
		rowIndex := fmt.Sprintf("%d", gigaBytes[i])
		basePriceStdResi := &entities.BasePrice{
			ProductID:  productStdResi.ID,
			RowIndex:   rowIndex,
			ColIndex:   "",
			PriceValue: stdResiPrice[i],
		}
		if err := db.FirstOrCreate(basePriceStdResi, *basePriceStdResi).Error; err != nil {
			return err
		}

		basePricePremResi := &entities.BasePrice{
			ProductID:  productPremResi.ID,
			RowIndex:   rowIndex,
			ColIndex:   "",
			PriceValue: premResiPrice[i],
		}
		if err := db.FirstOrCreate(basePricePremResi, *basePricePremResi).Error; err != nil {
			return err
		}
	}
	// Seed Base Price for Standard/Premium IPv6 Products
	productStdIPv6 := entities.Product{}
	if err := db.Where("name = ?", "Standard IPv6").First(&productStdIPv6).Error; err != nil {
		return err
	}
	productPremIPv6 := entities.Product{}
	if err := db.Where("name = ?", "Premium IPv6").First(&productPremIPv6).Error; err != nil {
		return err
	}
	threads := [...]int{100, 250, 500}
	durations = [...]int{1, 7, 30}
	var stdIPv6Price [3][3]float64 = [3][3]float64{
		{5.00, 15.00, 50.00},
		{10.00, 30.00, 100.00},
		{15.00, 50.00, 150.00},
	}
	var premIPv6Price [3][3]float64 = [3][3]float64{
		{30.00, 90.00, 300.00},
		{60.00, 180.00, 600.00},
		{100.00, 300.00, 1000.00},
	}
	for i := 0; i < len(threads); i++ {
		rowIndex := fmt.Sprintf("%d", threads[i])
		for j := 0; j < len(durations); j++ {
			colIndex := fmt.Sprintf("%d", durations[j])
			basePriceIPv6Std := &entities.BasePrice{
				ProductID:  productStdIPv6.ID,
				RowIndex:   rowIndex,
				ColIndex:   colIndex,
				PriceValue: stdIPv6Price[i][j],
			}
			if err := db.FirstOrCreate(basePriceIPv6Std, *basePriceIPv6Std).Error; err != nil {
				return err
			}

			basePriceIPv6Prem := &entities.BasePrice{
				ProductID:  productPremIPv6.ID,
				RowIndex:   rowIndex,
				ColIndex:   colIndex,
				PriceValue: premIPv6Price[i][j],
			}
			if err := db.FirstOrCreate(basePriceIPv6Prem, *basePriceIPv6Prem).Error; err != nil {
				return err
			}
		}
	}

	// Seed Base Price for Standard/Premium Shared ISP Products
	productStdSharedIsp := entities.Product{}
	if err := db.Where("name = ?", "Standard Shared ISP").First(&productStdSharedIsp).Error; err != nil {
		return err
	}
	productPremSharedIsp := entities.Product{}
	if err := db.Where("name = ?", "Premium Shared ISP").First(&productPremSharedIsp).Error; err != nil {
		return err
	}
	ispIpCounts := [...]int{10, 25, 50, 75, 100, 250, 500, 1000}
	bandwidth := [...]string{"250", "1000", "5000", "infinite"}
	var stdSharedIspPrice [8][4]float64 = [8][4]float64{
		{3.500, 5.000, 6.500, 9.000},
		{8.000, 11.25, 15.00, 21.25},
		{15.50, 21.00, 28.50, 41.00},
		{22.50, 30.75, 42.00, 60.75},
		{29.00, 40.00, 55.00, 80.00},
		{70.00, 97.50, 135.0, 197.5},
		{135.0, 190.0, 265.0, 390.0},
		{250.0, 360.0, 510.0, 760.0},
	}
	var premSharedIspPrice [8][4]float64 = [8][4]float64{
		{5.2500, 7.2500, 9.1000, 12.150},
		{12.000, 16.310, 21.000, 28.690},
		{23.250, 30.450, 39.900, 55.350},
		{33.750, 44.590, 58.800, 82.010},
		{43.500, 58.000, 77.000, 108.00},
		{105.00, 141.38, 189.00, 266.63},
		{202.50, 275.50, 371.00, 526.50},
		{375.00, 522.00, 714.00, 1026.0},
	}
	for i := 0; i < len(ispIpCounts); i++ {
		rowIndex := fmt.Sprintf("%d", ispIpCounts[i])
		for j := 0; j < len(bandwidth); j++ {
			colIndex := bandwidth[j]
			basePriceSharedIspStd := &entities.BasePrice{
				ProductID:  productStdSharedIsp.ID,
				RowIndex:   rowIndex,
				ColIndex:   colIndex,
				PriceValue: stdSharedIspPrice[i][j],
			}
			if err := db.FirstOrCreate(basePriceSharedIspStd, *basePriceSharedIspStd).Error; err != nil {
				return err
			}

			basePriceSharedIspPrem := &entities.BasePrice{
				ProductID:  productPremSharedIsp.ID,
				RowIndex:   rowIndex,
				ColIndex:   colIndex,
				PriceValue: premSharedIspPrice[i][j],
			}
			if err := db.FirstOrCreate(basePriceSharedIspPrem, *basePriceSharedIspPrem).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
