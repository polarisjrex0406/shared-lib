package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// LoyaltyPointsHistoryRepository is an interface that defines methods for performing CRUD operations on LoyaltyPointsHistory entity in the database.
type LoyaltyPointsHistoryRepository interface {
	Create(loyaltyPointsHistory *entities.LoyaltyPointsHistory) error

	FindByCustomerID(customerId uint) ([]entities.LoyaltyPointsHistory, error)
}

type loyaltyPointsHistoryRepository struct {
	DB *gorm.DB
}

func NewLoyaltyPointsHistoryRepository(db *gorm.DB) LoyaltyPointsHistoryRepository {
	return &loyaltyPointsHistoryRepository{DB: db}
}

func (r *loyaltyPointsHistoryRepository) Create(loyaltyPointsHistory *entities.LoyaltyPointsHistory) error {
	result := r.DB.Create(loyaltyPointsHistory)
	return result.Error
}

func (r *loyaltyPointsHistoryRepository) FindByCustomerID(customerId uint) ([]entities.LoyaltyPointsHistory, error) {
	loyaltyPointsHistories := []entities.LoyaltyPointsHistory{}
	result := r.DB.Where("customer_id = ?", customerId).
		Order("spent_date ASC").
		Find(&loyaltyPointsHistories)
	if result.Error != nil {
		return nil, result.Error
	}
	return loyaltyPointsHistories, nil
}
