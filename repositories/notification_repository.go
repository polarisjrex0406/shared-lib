package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// NotificationRepository is an interface that defines methods for performing CRUD operations on Notification entity in the database.
type NotificationRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new notification record into the database.
	Create(tx *gorm.DB, notification *entities.Notification) error

	// FindAll retrieves all notifications.
	FindAll() ([]entities.Notification, error)

	// FindOneByID retrieves one notification identified by its ID.
	FindOneByID(id uint) (*entities.Notification, error)

	// Update modifies an existing notification record in the database.
	Update(tx *gorm.DB, notification *entities.Notification) error

	// Delete removes a notification record from the database using its ID.
	Delete(id uint) error
}

type notificationRepository struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{DB: db}
}

func (r *notificationRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *notificationRepository) Create(tx *gorm.DB, notification *entities.Notification) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(notification)
	return result.Error
}

func (r *notificationRepository) FindAll() ([]entities.Notification, error) {
	notifications := []entities.Notification{}
	result := r.DB.Order("id ASC").Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return notifications, nil
}

func (r *notificationRepository) FindOneByID(id uint) (*entities.Notification, error) {
	notification := entities.Notification{}
	result := r.DB.Where("id = ?", id).First(&notification)
	if result.Error != nil {
		return nil, result.Error
	}
	return &notification, nil
}

func (r *notificationRepository) Update(tx *gorm.DB, notification *entities.Notification) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(notification)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *notificationRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.Notification{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
