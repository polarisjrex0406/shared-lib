package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UserRepository is an interface that defines methods for performing CRUD operations on User entity in the database.
type UserRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new user record into the database.
	Create(tx *gorm.DB, user *entities.User) error

	// FindOneByEmail retrieves a user by its email.
	FindOneByEmail(email string) (*entities.User, error)

	// Update modifies an existing user record in the database.
	Update(tx *gorm.DB, user *entities.User) error

	// UpdatePassword changes the password of a user identified by its ID.
	UpdatePassword(tx *gorm.DB, id uint, password string) (*entities.User, error)

	// Delete removes a user record from the database using its ID.
	Delete(tx *gorm.DB, id uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *userRepository) Create(tx *gorm.DB, user *entities.User) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	return dbInst.Create(user).Error
}

func (r *userRepository) FindAll() ([]entities.User, error) {
	users := []entities.User{}
	result := r.DB.Order("id ASC").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *userRepository) FindOneByEmail(email string) (*entities.User, error) {
	user := entities.User{}
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) Update(tx *gorm.DB, user *entities.User) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *userRepository) UpdatePassword(tx *gorm.DB, id uint, password string) (*entities.User, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	user := entities.User{}
	result := dbInst.Model(&user).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Update("pswd", password)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (r *userRepository) Delete(tx *gorm.DB, id uint) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Delete(&entities.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
