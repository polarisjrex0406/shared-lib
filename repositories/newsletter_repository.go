package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// NewsletterRepository is an interface that defines methods for performing CRUD operations on Newsletter entity in the database.
type NewsletterRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new newsletter record into the database.
	Create(tx *gorm.DB, newsletter *entities.Newsletter) error

	// FindAll retrieves all newsletters.
	FindAll() ([]entities.Newsletter, error)

	// FindOneByID retrieves one newsletter identified by its ID.
	FindOneByID(id uint) (*entities.Newsletter, error)

	// Update modifies an existing newsletter record in the database.
	Update(tx *gorm.DB, newsletter *entities.Newsletter) error

	// Delete removes a newsletter record from the database using its ID.
	Delete(id uint) error
}

type newsletterRepository struct {
	DB *gorm.DB
}

func NewNewsletterRepository(db *gorm.DB) NewsletterRepository {
	return &newsletterRepository{DB: db}
}

func (r *newsletterRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *newsletterRepository) Create(tx *gorm.DB, newsletter *entities.Newsletter) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(newsletter)
	return result.Error
}

func (r *newsletterRepository) FindAll() ([]entities.Newsletter, error) {
	newsletters := []entities.Newsletter{}
	result := r.DB.Order("id ASC").Find(&newsletters)
	if result.Error != nil {
		return nil, result.Error
	}
	return newsletters, nil
}

func (r *newsletterRepository) FindOneByID(id uint) (*entities.Newsletter, error) {
	newsletter := entities.Newsletter{}
	result := r.DB.Where("id = ?", id).First(&newsletter)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newsletter, nil
}

func (r *newsletterRepository) Update(tx *gorm.DB, newsletter *entities.Newsletter) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(newsletter)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *newsletterRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.Newsletter{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
