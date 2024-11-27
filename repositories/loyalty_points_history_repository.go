package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// LoyaltyPointsHistoryRepository is an interface that defines methods for performing CRUD operations on LoyaltyPointsHistory entity in the database.
type LoyaltyPointsHistoryRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new loyalty points history record into the database.
	Create(tx *gorm.DB, loyaltyPointsHistory *entities.LoyaltyPointsHistory) error

	// FindByCustomerID retrieves loyalty points histories by their customer ID.
	FindByCustomerID(customerId uint) ([]entities.LoyaltyPointsHistory, error)
}

type loyaltyPointsHistoryRepository struct {
	DB *gorm.DB
}

func NewLoyaltyPointsHistoryRepository(db *gorm.DB) LoyaltyPointsHistoryRepository {
	return &loyaltyPointsHistoryRepository{DB: db}
}

func (r *loyaltyPointsHistoryRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *loyaltyPointsHistoryRepository) Create(tx *gorm.DB, loyaltyPointsHistory *entities.LoyaltyPointsHistory) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(loyaltyPointsHistory)
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
