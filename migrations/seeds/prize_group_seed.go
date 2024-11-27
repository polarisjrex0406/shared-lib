package seeds

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

func PrizeGroup(db *gorm.DB) error {
	prizeGroups := []entities.PrizeGroup{
		{
			Rarity:     entities.PrizeRarityCommon,
			ChanceRate: 16.00,
		},
		{
			Rarity:     entities.PrizeRarityUncommon,
			ChanceRate: 3.00,
		},
		{
			Rarity:     entities.PrizeRarityRare,
			ChanceRate: 0.74,
		},
		{
			Rarity:     entities.PrizeRarityEpic,
			ChanceRate: 0.26,
		},
	}
	for _, prizeGroup := range prizeGroups {
		if err := db.FirstOrCreate(&prizeGroup, prizeGroup).Error; err != nil {
			return err
		}
	}
	return nil
}
