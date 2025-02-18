package migrations

import (
	"fmt"

	"github.com/omimic12/shared-lib/migrations/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.LoyaltyTier(db); err != nil {
		fmt.Println("Failed to seed LoyaltyTier with error: ", err.Error())
	}
	if err := seeds.Category(db); err != nil {
		fmt.Println("Failed to seed Category with error: ", err.Error())
	}
	if err := seeds.Provider(db); err != nil {
		fmt.Println("Failed to seed Provider with error: ", err.Error())
	}
	if err := seeds.Product(db); err != nil {
		fmt.Println("Failed to seed Product with error: ", err.Error())
	}
	if err := seeds.BasePrice(db); err != nil {
		fmt.Println("Failed to seed BasePrice with error: ", err.Error())
	}
	if err := seeds.User(db); err != nil {
		fmt.Println("Failed to seed User with error: ", err.Error())
	}
	if err := seeds.PrizeGroup(db); err != nil {
		fmt.Println("Failed to seed PrizeGroup with error: ", err.Error())
	}
	if err := seeds.Prize(db); err != nil {
		fmt.Println("Failed to seed Prize with error: ", err.Error())
	}
	if err := seeds.Proxy(db); err != nil {
		fmt.Println("Failed to seed Proxy with error: ", err.Error())
	}
	if err := seeds.Backconnect(db); err != nil {
		fmt.Println("Failed to seed Backconnect with error: ", err.Error())
	}
	if err := seeds.IssueTopic(db); err != nil {
		fmt.Println("Failed to seed Backconnect with error: ", err.Error())
	}
	return nil
}
