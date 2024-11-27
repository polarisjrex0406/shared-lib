package seeds

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

func LoyaltyTier(db *gorm.DB) error {
	commissionRates := [...]float64{2.0, 4.0, 6.0, 8.0, 10.0}
	dailySpins := [...]int{1, 2, 2, 3, 3}
	discountRates := [...]float64{0, 2.5, 5, 7.5, 10}
	discountCaps := [...]float64{100, 125, 150, 200}
	points := [...]int{0, 500, 1500, 3000, 6000}
	ranks := [...]string{"Bronze", "Silver", "Gold", "Platinum", "Diamond"}

	for i := 0; i < len(points); i++ {
		var discountCap *float64
		if i == 0 {
			discountCap = nil
		} else {
			discountCap = &discountCaps[i-1]
		}
		loyaltyTier := &entities.LoyaltyTier{
			CommissionRate: commissionRates[i],
			DailySpins:     dailySpins[i],
			DiscountRate:   discountRates[i],
			DiscountCap:    discountCap,
			Points:         points[i],
			Rank:           ranks[i],
		}
		if err := db.FirstOrCreate(loyaltyTier, *loyaltyTier).Error; err != nil {
			return err
		}
	}

	return nil
}
