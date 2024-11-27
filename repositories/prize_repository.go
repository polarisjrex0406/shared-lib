package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// PrizeRepository is an interface that defines methods for performing CRUD operations on Prize entity in the database.
type PrizeRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new prize record into the database.
	Create(tx *gorm.DB, prize *entities.Prize) error

	// FindAll retrieves all prizes.
	FindAll() ([]entities.Prize, error)

	// FindByGroupID retrieves prizes by its group ID.
	FindByGroupID(groupId uint) ([]entities.Prize, error)

	// FindOneByID retrieves one prize identified by its ID.
	FindOneByID(id uint) (*entities.Prize, error)

	// Update modifies an existing prize record in the database.
	Update(tx *gorm.DB, prize *entities.Prize) error

	// Delete removes a prize record from the database using its ID.
	Delete(id uint) error
}

type prizeRepository struct {
	DB *gorm.DB
}

func NewPrizeRepository(db *gorm.DB) PrizeRepository {
	return &prizeRepository{DB: db}
}

func (r *prizeRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *prizeRepository) Create(tx *gorm.DB, prize *entities.Prize) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(prize)
	return result.Error
}

func (r *prizeRepository) FindAll() ([]entities.Prize, error) {
	prizes := []entities.Prize{}
	result := r.DB.Order("id ASC").Find(&prizes)
	if result.Error != nil {
		return nil, result.Error
	}
	return prizes, nil
}

func (r *prizeRepository) FindByGroupID(groupId uint) ([]entities.Prize, error) {
	prizes := []entities.Prize{}
	result := r.DB.Where("group_id = ?", groupId).Find(&prizes)
	if result.Error != nil {
		return nil, result.Error
	}
	return prizes, nil
}

func (r *prizeRepository) FindOneByID(id uint) (*entities.Prize, error) {
	prize := entities.Prize{}
	result := r.DB.Where("id = ?", id).First(&prize)
	if result.Error != nil {
		return nil, result.Error
	}
	return &prize, nil
}

func (r *prizeRepository) Update(tx *gorm.DB, prize *entities.Prize) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(prize)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *prizeRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.Prize{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
