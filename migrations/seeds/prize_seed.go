package seeds

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

func Prize(db *gorm.DB) error {
	// All
	prizes := []entities.Prize{}
	// Common
	commonPrizes := []entities.Prize{
		productDcPrize(db, "Standard", 1, 1000),
		loyaltyPointsPrize(5),
		loyaltyPointsPrize(10),
	}
	if err := setGroup(db, commonPrizes, entities.PrizeRarityCommon); err != nil {
		return err
	}
	prizes = append(prizes, commonPrizes...)
	// Uncommon
	uncommonPrizes := []entities.Prize{
		productDcPrize(db, "Standard", 1, 2500),
		productResiPrize(db, "Standard", 1),
		creditPrize(5),
		discountPrize(5),
	}
	if err := setGroup(db, uncommonPrizes, entities.PrizeRarityUncommon); err != nil {
		return err
	}
	prizes = append(prizes, uncommonPrizes...)
	// Rare
	rarePrizes := []entities.Prize{
		productDcPrize(db, "Standard", 1, 5000),
		productResiPrize(db, "Premium", 1),
		creditPrize(10),
		discountPrize(10),
	}
	if err := setGroup(db, rarePrizes, entities.PrizeRarityRare); err != nil {
		return err
	}
	prizes = append(prizes, rarePrizes...)
	// Epic
	epicPrizes := []entities.Prize{
		productDcPrize(db, "Standard", 30, 1000),
		productResiPrize(db, "Standard", 10),
		productResiPrize(db, "Premium", 5),
		discountPrize(20),
	}
	if err := setGroup(db, epicPrizes, entities.PrizeRarityEpic); err != nil {
		return err
	}
	prizes = append(prizes, epicPrizes...)
	// Create
	for _, prize := range prizes {
		if err := db.FirstOrCreate(&prize, prize).Error; err != nil {
			return err
		}
	}
	return nil
}

func setGroup(db *gorm.DB, prizes []entities.Prize, rarity entities.PrizeRarity) error {
	prizeGroup := entities.PrizeGroup{}
	if err := db.Where("rarity = ?", rarity).First(&prizeGroup).Error; err != nil {
		return err
	}
	for i := range prizes {
		prizes[i].GroupID = prizeGroup.ID
	}
	return nil
}

func productDcPrize(db *gorm.DB, abbrLong string, duration int, ipCount int) entities.Prize {
	product := entities.Product{}
	if err := db.Where("name = ?", abbrLong+" Datacenter").First(&product).Error; err != nil {
		return entities.Prize{}
	}
	return entities.Prize{
		Kind:      entities.PrizeKindProduct,
		ProductID: &product.ID,
		Duration:  &duration,
		IPCount:   &ipCount,
		Threads:   &product.ThreadsRange.Min,
	}
}

func productResiPrize(db *gorm.DB, abbrLong string, bandwidth int) entities.Prize {
	product := entities.Product{}
	if err := db.Where("name = ?", abbrLong+" Residential").First(&product).Error; err != nil {
		return entities.Prize{}
	}
	return entities.Prize{
		Kind:      entities.PrizeKindProduct,
		ProductID: &product.ID,
		Bandwidth: &bandwidth,
	}
}

func loyaltyPointsPrize(loyaltyPoints int) entities.Prize {
	return entities.Prize{
		Kind:          entities.PrizeKindLoyaltyPoints,
		LoyaltyPoints: &loyaltyPoints,
	}
}

func creditPrize(credit float64) entities.Prize {
	return entities.Prize{
		Kind:   entities.PrizeKindCredit,
		Credit: &credit,
	}
}

func discountPrize(discountRate float64) entities.Prize {
	return entities.Prize{
		Kind:         entities.PrizeKindDiscount,
		DiscountRate: &discountRate,
	}
}
