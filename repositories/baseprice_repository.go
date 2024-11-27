package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// BasePriceRepository is an interface that defines methods for performing CRUD operations on BasePrice entity in the database.
type BasePriceRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new base price record into the database.
	Create(tx *gorm.DB, basePrice *entities.BasePrice) error

	// FindIndexesByProductID retrieves values of indexes by product ID and index name.
	FindIndexesByProductID(productID uint, indexName string) ([]string, error)

	// FindOneByProductIDAndRowIndexAndColIndex retrieves a base price by product ID, row and column indexes.
	FindOneByProductIDAndRowIndexAndColIndex(productID uint, rowIndex string, colIndex string) (*entities.BasePrice, error)
}

type basePriceRepository struct {
	DB *gorm.DB
}

func NewBasePriceRepository(db *gorm.DB) BasePriceRepository {
	return &basePriceRepository{DB: db}
}

func (r *basePriceRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *basePriceRepository) Create(tx *gorm.DB, basePrice *entities.BasePrice) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(basePrice)
	return result.Error
}

func (r *basePriceRepository) FindIndexesByProductID(productID uint, indexName string) ([]string, error) {
	var indexes []string
	result := r.DB.Model(&entities.BasePrice{}).
		Where("product_id = ?", productID).
		Select(indexName).
		Find(&indexes)
	if result.Error != nil {
		return nil, result.Error
	}
	return indexes, nil
}

func (r *basePriceRepository) FindOneByProductIDAndRowIndexAndColIndex(productID uint, rowIndex string, colIndex string) (*entities.BasePrice, error) {
	basePrice := entities.BasePrice{}
	result := r.DB.Where("product_id = ? AND row_index = ? AND col_index = ?", productID, rowIndex, colIndex).First(&basePrice)
	if result.Error != nil {
		return nil, result.Error
	}
	return &basePrice, nil
}
