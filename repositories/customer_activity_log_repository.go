package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// CustomerActivityLogRepository is an interface that defines methods for performing CRUD operations on Category entity in the database.
type CustomerActivityLogRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new category record into the database.
	Create(tx *gorm.DB, category *entities.CustomerActivityLog) error

	CheckByCustomerIDAndEventType(customerId uint, eventType string) (*bool, error)
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

func (r *customerActivityLogRepository) CheckByCustomerIDAndEventType(customerId uint, eventType string) (*bool, error) {
	customerActivityLog := entities.CustomerActivityLog{}
	result := r.DB.
		Model(&entities.CustomerActivityLog{}).
		Where("customer_id = ? AND event_type = ?", customerId, eventType).
		Find(&customerActivityLog)

	isExist := result.Error == nil
	if result.Error == nil || result.Error == gorm.ErrRecordNotFound {
		return &isExist, nil
	}
	return nil, result.Error
}
