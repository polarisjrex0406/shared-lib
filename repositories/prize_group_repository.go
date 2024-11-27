package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// PrizeGroupRepository is an interface that defines methods for performing CRUD operations on PrizeGroup entity in the database.
type PrizeGroupRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new prize group record into the database.
	Create(tx *gorm.DB, prizeGroup *entities.PrizeGroup) error

	// FindAll retrieves all prize groups.
	FindAll() ([]entities.PrizeGroup, error)

	// FindOneByID retrieves one prize group identified by its ID.
	FindOneByID(id uint) (*entities.PrizeGroup, error)

	// Update modifies an existing prize group record in the database.
	Update(tx *gorm.DB, prizeGroup *entities.PrizeGroup) error

	// Delete removes a prize group record from the database using its ID.
	Delete(id uint) error
}

type prizeGroupRepository struct {
	DB *gorm.DB
}

func NewPrizeGroupRepository(db *gorm.DB) PrizeGroupRepository {
	return &prizeGroupRepository{DB: db}
}

func (r *prizeGroupRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *prizeGroupRepository) Create(tx *gorm.DB, prizeGroup *entities.PrizeGroup) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(prizeGroup)
	return result.Error
}

func (r *prizeGroupRepository) FindAll() ([]entities.PrizeGroup, error) {
	prizeGroups := []entities.PrizeGroup{}
	result := r.DB.Order("chance_rate ASC").Find(&prizeGroups)
	if result.Error != nil {
		return nil, result.Error
	}
	return prizeGroups, nil
}

func (r *prizeGroupRepository) FindOneByID(id uint) (*entities.PrizeGroup, error) {
	prizeGroup := entities.PrizeGroup{}
	result := r.DB.Where("id = ?", id).First(&prizeGroup)
	if result.Error != nil {
		return nil, result.Error
	}
	return &prizeGroup, nil
}

func (r *prizeGroupRepository) Update(tx *gorm.DB, prizeGroup *entities.PrizeGroup) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(prizeGroup)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *prizeGroupRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.PrizeGroup{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
