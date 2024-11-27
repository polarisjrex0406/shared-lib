package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BillingAddressRepository is an interface that defines methods for performing CRUD operations on BillingAddress entity in the database.
type BillingAddressRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new billing address record into the database.
	Create(tx *gorm.DB, billingAddr *entities.BillingAddress) error

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

func (r *billingAddressRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *billingAddressRepository) Create(tx *gorm.DB, billingAddr *entities.BillingAddress) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(billingAddr)
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
