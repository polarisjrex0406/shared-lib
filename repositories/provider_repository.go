package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ProviderRepository is an interface that defines methods for performing CRUD operations on Provider entity in the database.
type ProviderRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new provider record into the database.
	Create(tx *gorm.DB, provider *entities.Provider) error

	// FindAll retrieves all providers.
	FindAll() ([]entities.Provider, error)

	// FindOneByID retrieves a provider identified by its ID.
	FindOneByID(id uint) (*entities.Provider, error)

	// Update modifies an existing provider record in the database.
	Update(tx *gorm.DB, provider *entities.Provider) error

	// Delete removes a provider record from the database using its ID.
	Delete(id uint) error
}

type providerRepository struct {
	DB *gorm.DB
}

func NewProviderRepository(db *gorm.DB) ProviderRepository {
	return &providerRepository{DB: db}
}

func (r *providerRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *providerRepository) Create(tx *gorm.DB, provider *entities.Provider) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(provider)
	return result.Error
}

func (r *providerRepository) FindAll() ([]entities.Provider, error) {
	providers := []entities.Provider{}
	result := r.DB.Order("id ASC").Find(&providers)
	if result.Error != nil {
		return nil, result.Error
	}
	return providers, nil
}

func (r *providerRepository) FindOneByID(id uint) (*entities.Provider, error) {
	provider := entities.Provider{}
	result := r.DB.Where("id = ?", id).First(&provider)
	if result.Error != nil {
		return nil, result.Error
	}
	return &provider, nil
}

func (r *providerRepository) Update(tx *gorm.DB, provider *entities.Provider) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(provider)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *providerRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.Provider{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
