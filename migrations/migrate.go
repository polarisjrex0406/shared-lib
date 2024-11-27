package migrations

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	tables := []interface{}{
		&entities.AuthInfo{},
		&entities.Balance{},
		&entities.BasePrice{},
		&entities.BillingAddress{},
		&entities.Category{},
		&entities.ClaimedPrize{},
		&entities.Coupon{},
		&entities.CryptomusTransaction{},
		&entities.Customer{},
		&entities.DataImpulseSubuser{},
		&entities.Invoice{},
		&entities.LoyaltyPointsHistory{},
		&entities.LoyaltyTier{},
		&entities.Prize{},
		&entities.PrizeGroup{},
		&entities.Product{},
		&entities.Provider{},
		&entities.Purchase{},
		&entities.ReferralEarning{},
		&entities.Transaction{},
		&entities.TTProxySubuser{},
		&entities.User{},
	}
	tableList, err := db.Migrator().GetTables()
	if err != nil {
		return err
	}
	for i := range tableList {
		db.Migrator().DropTable(tableList[i])
	}
	for _, table := range tables {
		if err := db.Migrator().CreateTable(table); err != nil {
			return err
		}
	}
	return nil
}
