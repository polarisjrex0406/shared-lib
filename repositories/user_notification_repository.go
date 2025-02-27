package repositories

import (
	"fmt"

	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// UserNotificationRepository is an interface that defines methods for performing CRUD operations on UserNotification entity in the database.
type UserNotificationRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new user notification record into the database.
	Create(tx *gorm.DB, userNotification *entities.UserNotification) error

	// FindAll retrieves all user notifications.
	FindAll() ([]entities.UserNotification, error)

	// FindOneByID retrieves one user notification identified by its ID.
	FindOneByID(id uint) (*entities.UserNotification, error)

	// FindByTargetUserID retrieves user notifications identified by target user ID.
	FindByTargetUserID(targetUserId uint) ([]entities.UserNotification, error)

	// FindByTargetUserIDAndReadUserID retrieves user notifications identified by target user ID and read user ID.
	FindByTargetUserIDAndReadUserID(targetUserId, readUserId uint) ([]entities.UserNotification, error)

	// Delete removes a user notification record from the database using its ID.
	Delete(id uint) error
}

type userNotificationRepository struct {
	DB *gorm.DB
}

func NewUserNotificationRepository(db *gorm.DB) UserNotificationRepository {
	return &userNotificationRepository{DB: db}
}

func (r *userNotificationRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *userNotificationRepository) Create(tx *gorm.DB, userNotification *entities.UserNotification) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(userNotification)
	return result.Error
}

func (r *userNotificationRepository) FindAll() ([]entities.UserNotification, error) {
	userNotifications := []entities.UserNotification{}
	result := r.DB.Order("id ASC").Find(&userNotifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return userNotifications, nil
}

func (r *userNotificationRepository) FindOneByID(id uint) (*entities.UserNotification, error) {
	userNotification := entities.UserNotification{}
	result := r.DB.Where("id = ?", id).First(&userNotification)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userNotification, nil
}

func (r *userNotificationRepository) FindByTargetUserID(targetUserId uint) ([]entities.UserNotification, error) {
	userNotifications := []entities.UserNotification{}
	result := r.DB.Where("target_user_ids::jsonb @> ?", fmt.Sprintf("[%d]", targetUserId)).
		Order("id DESC").
		Find(&userNotifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return userNotifications, nil
}

func (r *userNotificationRepository) FindByTargetUserIDAndReadUserID(targetUserId, readUserId uint) ([]entities.UserNotification, error) {
	userNotifications := []entities.UserNotification{}
	result := r.DB.Where(
		"target_user_ids::jsonb @> ? AND NOT read_user_ids::jsonb @> ?",
		fmt.Sprintf("[%d]", targetUserId),
		fmt.Sprintf("[%d]", readUserId),
	).
		Order("id DESC").
		Find(&userNotifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return userNotifications, nil
}

func (r *userNotificationRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.UserNotification{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
