package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"github.com/omimic12/shared-lib/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CustomerRepository is an interface that defines methods for performing CRUD operations on Customer entity in the database.
type CustomerRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new customer record into the database.
	Create(tx *gorm.DB, customer *entities.Customer) error

	FindAll() ([]entities.Customer, error)

	// FindByReferrerID retrieves all customers by their referrer ID.
	FindByReferrerID(referrerID uint) ([]entities.Customer, error)

	// FindAllIDs retrieves ID of all customers.
	FindAllIDs() ([]uint, error)

	// FindOneByEmail retrieves a customer by its email.
	FindOneByEmail(email string) (*entities.Customer, error)

	// FindOneByID retrieves a customer identified by its ID.
	FindOneByID(id uint) (*entities.Customer, error)

	// FindOneByProfileName retrieves a customer by its profile name.
	FindOneByProfileName(profileName string) (*entities.Customer, error)

	// Update modifies an existing customer record in the database.
	Update(tx *gorm.DB, customer *entities.Customer) error

	// UpdatePoints adds points of this customer identified by its ID.
	UpdatePoints(tx *gorm.DB, id uint, points int) (*entities.Customer, error)

	// UpdateUsedSpins sets used spin count of this customer identified by its ID.
	UpdateUsedSpins(tx *gorm.DB, id *uint, usedSpins int) (*entities.Customer, error)
}

type customerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{DB: db}
}

func (r *customerRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *customerRepository) Create(tx *gorm.DB, customer *entities.Customer) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(customer)
	return result.Error
}

func (r *customerRepository) FindAll() ([]entities.Customer, error) {
	customers := []entities.Customer{}
	result := r.DB.Order("id ASC").Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (r *customerRepository) FindByReferrerID(referrerID uint) ([]entities.Customer, error) {
	var customers []entities.Customer
	result := r.DB.Where("referrer_id = ?", referrerID).
		Order("created_at ASC").
		Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (r *customerRepository) FindAllIDs() ([]uint, error) {
	var ids []uint
	result := r.DB.Model(&entities.Customer{}).
		Select("id").
		Find(&ids)
	if result.Error != nil {
		return nil, result.Error
	}
	return ids, nil
}

func (r *customerRepository) FindOneByEmail(email string) (*entities.Customer, error) {
	customer := entities.Customer{}
	result := r.DB.Where("email = ?", email).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (r *customerRepository) FindOneByID(id uint) (*entities.Customer, error) {
	customer := entities.Customer{}
	result := r.DB.Where("id = ?", id).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (r *customerRepository) FindOneByProfileName(profileName string) (*entities.Customer, error) {
	customer := entities.Customer{}
	result := r.DB.Where("profile_name = ?", profileName).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (r *customerRepository) Update(tx *gorm.DB, customer *entities.Customer) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}

	customerMap, err := utils.RemUnwanted(*customer)
	if err != nil {
		return err
	}

	result := dbInst.Clauses(clause.Returning{}).
		Model(customer).
		Updates(customerMap)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *customerRepository) UpdatePoints(tx *gorm.DB, id uint, points int) (*entities.Customer, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	customer := entities.Customer{}
	result := dbInst.Model(&customer).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Update("points", gorm.Expr("points + ?", points))
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &customer, nil
}

func (r *customerRepository) UpdateUsedSpins(tx *gorm.DB, id *uint, usedSpins int) (*entities.Customer, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	customer := entities.Customer{}
	result := dbInst.Model(&customer).
		Clauses(clause.Returning{})
	if id != nil {
		result = result.Where("id = ?", *id)
	} else {
		result = result.Where("1 = 1")
	}
	result = result.Update("used_spins", usedSpins)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &customer, nil
}
