package repositories

import (
	"fmt"

	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// CustomerActivityLogRepository is an interface that defines methods for performing CRUD operations on Category entity in the database.
type CustomerActivityLogRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new category record into the database.
	Create(tx *gorm.DB, category *entities.CustomerActivityLog) error

	CountByCustomerIDsAndEventTypeAndMetaData(customerIds []uint, eventType string, metaData string) (*int64, error)
}

type customerActivityLogRepository struct {
	DB *gorm.DB
}

func NewCustomerActivityLogRepository(db *gorm.DB) CustomerActivityLogRepository {
	return &customerActivityLogRepository{DB: db}
}

func (r *customerActivityLogRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *customerActivityLogRepository) Create(tx *gorm.DB, category *entities.CustomerActivityLog) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(category)
	return result.Error
}

func (r *customerActivityLogRepository) CountByCustomerIDsAndEventTypeAndMetaData(customerIds []uint, eventType string, metaData string) (*int64, error) {
	var customerCount int64
	result := r.DB.
		Model(&entities.CustomerActivityLog{}).
		Select("COUNT(DISTINCT customer_id)").
		Where(
			"customer_id IN ? AND event_type = ? AND meta_data LIKE ?",
			customerIds,
			eventType,
			fmt.Sprintf("%%%s%%", metaData),
		).
		Group("customer_id").
		Scan(&customerCount)

	if result.Error != nil {
		return nil, result.Error
	}
	return &customerCount, nil
}
