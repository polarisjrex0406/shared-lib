package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BillingAddressRepository is an interface that defines methods for performing CRUD operations on BillingAddress entity in the database.
type BillingAddressRepository interface {
	Create(billingAddr *entities.BillingAddress) error

	FindOneByCustomerID(customerId uint) (*entities.BillingAddress, error)

	Update(customerId uint, billingAddr *entities.BillingAddress) error
}

type billingAddressRepository struct {
	DB *gorm.DB
}

func NewBillingAddressRepository(db *gorm.DB) BillingAddressRepository {
	return &billingAddressRepository{DB: db}
}

func (r *billingAddressRepository) Create(billingAddr *entities.BillingAddress) error {
	result := r.DB.Create(billingAddr)
	return result.Error
}

func (r *billingAddressRepository) FindOneByCustomerID(customerId uint) (*entities.BillingAddress, error) {
	billingAddr := entities.BillingAddress{}

	result := r.DB.Where("customer_id = ?", customerId).First(&billingAddr)
	if result.Error != nil {
		return nil, result.Error
	}

	return &billingAddr, nil
}

func (r *billingAddressRepository) Update(customerId uint, billingAddr *entities.BillingAddress) error {
	result := r.DB.Model(&entities.BillingAddress{}).
		Clauses(clause.Returning{}).
		Where("customer_id = ?", customerId).
		Select(
			"firstname", "lastname", "country", "street_address",
			"state_abbr", "city", "zipcode",
		).
		Updates(billingAddr)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
