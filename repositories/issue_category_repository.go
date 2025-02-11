package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// IssueCategoryRepository is an interface that defines methods for performing CRUD operations on IssueCategory entity in the database.
type IssueCategoryRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new issue category record into the database.
	Create(tx *gorm.DB, issueCategory *entities.IssueCategory) error

	// FindAll retrieves all issue categories.
	FindAll() ([]entities.IssueCategory, error)

	// FindOneByID retrieves one issue category identified by its ID.
	FindOneByID(id uint) (*entities.IssueCategory, error)

	// FindOneByName retrieves one issue category by its name.
	FindOneByName(name string) (*entities.IssueCategory, error)

	// Delete removes a issue category record from the database using its ID.
	Delete(id uint) error
}

type issueCategoryRepository struct {
	DB *gorm.DB
}

func NewIssueCategoryRepository(db *gorm.DB) IssueCategoryRepository {
	return &issueCategoryRepository{DB: db}
}

func (r *issueCategoryRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *issueCategoryRepository) Create(tx *gorm.DB, issueCategory *entities.IssueCategory) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(issueCategory)
	return result.Error
}

func (r *issueCategoryRepository) FindAll() ([]entities.IssueCategory, error) {
	categories := []entities.IssueCategory{}
	result := r.DB.Order("id ASC").Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (r *issueCategoryRepository) FindOneByID(id uint) (*entities.IssueCategory, error) {
	issueCategory := entities.IssueCategory{}
	result := r.DB.Where("id = ?", id).First(&issueCategory)
	if result.Error != nil {
		return nil, result.Error
	}
	return &issueCategory, nil
}

func (r *issueCategoryRepository) FindOneByName(name string) (*entities.IssueCategory, error) {
	issueCategory := entities.IssueCategory{}
	result := r.DB.Where("name = ?", name).First(&issueCategory)
	if result.Error != nil {
		return nil, result.Error
	}
	return &issueCategory, nil
}

func (r *issueCategoryRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.IssueCategory{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
