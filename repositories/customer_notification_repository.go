package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CustomerNotificationRepository is an interface that defines methods for performing CRUD operations on CustomerNotification entity in the database.
type CustomerNotificationRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new customer notification record into the database.
	Create(tx *gorm.DB, customerNotification *entities.CustomerNotification) error

	// FindAll retrieves all customer notifications.
	FindAll() ([]entities.CustomerNotification, error)

	// FindOneByID retrieves one customer notification identified by its ID.
	FindOneByID(id uint) (*entities.CustomerNotification, error)

	// FindByTargetCustomerID retrieves customer notifications identified by target customer ID.
	FindByTargetCustomerID(targetCustomerId uint) ([]entities.CustomerNotification, error)

	// FindByTargetCustomerIDAndReadCustomerID retrieves customer notifications identified by target customer ID and read customer ID.
	FindByTargetCustomerIDAndReadCustomerID(targetCustomerId, readCustomerId uint) ([]entities.CustomerNotification, error)

	// Update modifies an existing customer notification record in the database.
	Update(tx *gorm.DB, customerNotification *entities.CustomerNotification) error

	// UpdateReadCustomerIDs modifies read customer ID of an existing customer notification record in the database.
	UpdateReadCustomerIDs(tx *gorm.DB, id, readCustomerId uint) error

	// Delete removes a customer notification record from the database using its ID.
	Delete(id uint) error
}

type customerNotificationRepository struct {
	DB *gorm.DB
}

func NewCustomerNotificationRepository(db *gorm.DB) CustomerNotificationRepository {
	return &customerNotificationRepository{DB: db}
}

func (r *customerNotificationRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *customerNotificationRepository) Create(tx *gorm.DB, customerNotification *entities.CustomerNotification) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(customerNotification)
	return result.Error
}

func (r *customerNotificationRepository) FindAll() ([]entities.CustomerNotification, error) {
	customerNotifications := []entities.CustomerNotification{}
	result := r.DB.Order("id ASC").Find(&customerNotifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return customerNotifications, nil
}

func (r *customerNotificationRepository) FindOneByID(id uint) (*entities.CustomerNotification, error) {
	customerNotification := entities.CustomerNotification{}
	result := r.DB.Where("id = ?", id).First(&customerNotification)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customerNotification, nil
}

func (r *customerNotificationRepository) FindByTargetCustomerID(targetCustomerId uint) ([]entities.CustomerNotification, error) {
	customerNotifications := []entities.CustomerNotification{}
	result := r.DB.Where("target_customer_ids::jsonb @> ?", fmt.Sprintf("[%d]", targetCustomerId)).
		Order("id DESC").
		Find(&customerNotifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return customerNotifications, nil
}

func (r *customerNotificationRepository) FindByTargetCustomerIDAndReadCustomerID(targetCustomerId, readCustomerId uint) ([]entities.CustomerNotification, error) {
	customerNotifications := []entities.CustomerNotification{}
	result := r.DB.Where(
		"target_customer_ids::jsonb @> ? AND NOT read_customer_ids::jsonb @> ?",
		fmt.Sprintf("[%d]", targetCustomerId),
		fmt.Sprintf("[%d]", readCustomerId),
	).
		Order("id DESC").
		Find(&customerNotifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return customerNotifications, nil
}

func (r *customerNotificationRepository) Update(tx *gorm.DB, customerNotification *entities.CustomerNotification) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(customerNotification)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *customerNotificationRepository) UpdateReadCustomerIDs(tx *gorm.DB, id, readCustomerId uint) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}

	readCustomerIdJson, _ := json.Marshal([]uint{readCustomerId})

	result := dbInst.Model(&entities.UserNotification{}).
		Where("id = ?", id).
		Update("read_customer_ids",
			gorm.Expr(`
			CASE
				WHEN NOT read_customer_ids::jsonb @> ?::jsonb 
				THEN read_customer_ids::jsonb || ?::jsonb 
				ELSE read_customer_ids::jsonb
			END`, readCustomerIdJson, readCustomerIdJson),
		)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *customerNotificationRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.CustomerNotification{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
