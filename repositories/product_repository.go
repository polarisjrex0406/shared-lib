package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ProductRepository is an interface that defines methods for performing CRUD operations on Product entity in the database.
type ProductRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new product record into the database.
	Create(tx *gorm.DB, product *entities.Product) error

	// FindAllIDs retrieves ID of all products.
	FindAllIDs() ([]uint, error)

	// FindByCategoryID retrieves customers by their category ID.
	FindByCategoryID(categoryId uint) ([]entities.Product, error)

	// FindOneByID retrieves a product identified by its ID.
	FindOneByID(id uint) (*entities.Product, error)

	// Update modifies an existing product record in the database.
	Update(tx *gorm.DB, product *entities.Product) error

	// Delete removes a product record from the database using its ID.
	Delete(id uint) error
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{DB: db}
}

func (r *productRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *productRepository) Create(tx *gorm.DB, product *entities.Product) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(product)
	return result.Error
}

func (r *productRepository) FindAllIDs() ([]uint, error) {
	var ids []uint
	result := r.DB.Model(&entities.Product{}).
		Select("id").
		Find(&ids)
	if result.Error != nil {
		return nil, result.Error
	}
	return ids, nil
}

func (r *productRepository) FindByCategoryID(categoryId uint) ([]entities.Product, error) {
	products := []entities.Product{}
	result := r.DB.Where("category_id = ?", categoryId).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *productRepository) FindOneByID(id uint) (*entities.Product, error) {
	product := entities.Product{}
	result := r.DB.Where("id = ?", id).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *productRepository) Update(tx *gorm.DB, product *entities.Product) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *productRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
