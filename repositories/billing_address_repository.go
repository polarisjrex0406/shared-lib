package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BillingAddressRepository is an interface that defines methods for performing CRUD operations on BillingAddress entity in the database.
type BillingAddressRepository interface {
	// FindOneByCustomerID retrieves a billing address by its customer.
	FindOneByCustomerID(customerId uint) (*entities.BillingAddress, error)

	// UpdateOneByCustomerID modifies an existing billing address by its customer ID.
	UpdateOneByCustomerID(tx *gorm.DB, customerId uint, billingAddr *entities.BillingAddress) error
}

type billingAddressRepository struct {
	DB *gorm.DB
}

func NewBillingAddressRepository(db *gorm.DB) BillingAddressRepository {
	return &billingAddressRepository{DB: db}
}

func (r *billingAddressRepository) FindOneByCustomerID(customerId uint) (*entities.BillingAddress, error) {
	billingAddr := entities.BillingAddress{}
	result := r.DB.Where("customer_id = ?", customerId).First(&billingAddr)
	if result.Error != nil {
		return nil, result.Error
	}
	return &billingAddr, nil
}

func (r *billingAddressRepository) UpdateOneByCustomerID(tx *gorm.DB, customerId uint, billingAddr *entities.BillingAddress) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).
		Where("customer_id = ?", customerId).
		Updates(billingAddr)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
