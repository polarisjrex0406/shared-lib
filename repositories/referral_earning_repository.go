package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// ReferralEarningRepository is an interface that defines methods for performing CRUD operations on ReferralEarning entity in the database.
type ReferralEarningRepository interface {
	// Create inserts a new referral earning record into the database.
	Create(referralEarning *entities.ReferralEarning) error

	// FindByCustomerID retrieves referral earnings by their customer ID.
	FindByCustomerID(customerId uint) ([]entities.ReferralEarning, error)
}

type referralEarningRepository struct {
	DB *gorm.DB
}

func NewReferralEarningRepository(db *gorm.DB) ReferralEarningRepository {
	return &referralEarningRepository{DB: db}
}

func (r *referralEarningRepository) Create(referralEarning *entities.ReferralEarning) error {
	return r.DB.Create(referralEarning).Error
}

func (r *referralEarningRepository) FindByCustomerID(customerId uint) ([]entities.ReferralEarning, error) {
	var referralEarnings []entities.ReferralEarning

	result := r.DB.Where("customer_id = ?", customerId).
		Order("order_date ASC").
		Find(&referralEarnings)
	if result.Error != nil {
		return nil, result.Error
	}

	return referralEarnings, nil
}
