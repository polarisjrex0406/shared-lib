package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CategoryRepository is an interface that defines methods for performing CRUD operations on Category entity in the database.
type CategoryRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new category record into the database.
	Create(tx *gorm.DB, category *entities.Category) error

	// FindAll retrieves all categories.
	FindAll() ([]entities.Category, error)

	// FindOneByAbbr retrieves one category by its abbr.
	FindOneByAbbr(abbr string) (*entities.Category, error)

	// FindOneByID retrieves one category identified by its ID.
	FindOneByID(id uint) (*entities.Category, error)

	// FindOneByName retrieves one category by its name.
	FindOneByName(name string) (*entities.Category, error)

	// Update modifies an existing category record in the database.
	Update(tx *gorm.DB, category *entities.Category) error

	// Delete removes a category record from the database using its ID.
	Delete(id uint) error
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{DB: db}
}

func (r *categoryRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *categoryRepository) Create(tx *gorm.DB, category *entities.Category) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(category)
	return result.Error
}

func (r *categoryRepository) FindAll() ([]entities.Category, error) {
	categories := []entities.Category{}
	result := r.DB.Order("id ASC").Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (r *categoryRepository) FindOneByAbbr(abbr string) (*entities.Category, error) {
	category := entities.Category{}
	result := r.DB.Where("abbr = ?", abbr).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (r *categoryRepository) FindOneByID(id uint) (*entities.Category, error) {
	category := entities.Category{}
	result := r.DB.Where("id = ?", id).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (r *categoryRepository) FindOneByName(name string) (*entities.Category, error) {
	category := entities.Category{}
	result := r.DB.Where("name = ?", name).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (r *categoryRepository) Update(tx *gorm.DB, category *entities.Category) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(category)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *categoryRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
