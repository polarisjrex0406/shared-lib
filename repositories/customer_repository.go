package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CustomerRepository is an interface that defines methods for performing CRUD operations on Customer entity in the database.
type CustomerRepository interface {
	Create(customer *entities.Customer) error

	CheckByEmail(email string) (bool, error)

	FindByReferrerID(referrerID uint) ([]entities.Customer, error)

	FindAllIDs() ([]uint, error)

	FindEmailByID(id uint) (string, error)

	FindOneByEmail(email string) (*entities.Customer, error)

	FindOneByID(id uint) (*entities.Customer, error)

	FindOneByProfileName(profileName string) (*entities.Customer, error)

	UpdateProfile(id uint, email, profileName string) (*entities.Customer, error)

	UpdateSettings(id uint, enableMFA, subscribeNL, notifyExpire bool) (*entities.Customer, error)

	UpdatePoints(id uint, points int) (*entities.Customer, error)

	UpdateUsedSpins(id *uint, usedSpins int) (*entities.Customer, error)
}

type customerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{DB: db}
}

func (r *customerRepository) Create(customer *entities.Customer) error {
	result := r.DB.Create(customer)
	return result.Error
}

func (r *customerRepository) CheckByEmail(email string) (bool, error) {
	customer := entities.Customer{}

	result := r.DB.Where("email = ?", email).First(&customer)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (r *customerRepository) FindByReferrerID(referrerID uint) ([]entities.Customer, error) {
	var customers []entities.Customer

	result := r.DB.Where("referrer_id = ?", referrerID).
		Order("created_at ASC").
		Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}

	return customers, nil
}

func (r *customerRepository) FindAllIDs() ([]uint, error) {
	var ids []uint

	result := r.DB.Model(&entities.Customer{}).
		Select("id").
		Find(&ids)
	if result.Error != nil {
		return nil, result.Error
	}

	return ids, nil
}

func (r *customerRepository) FindEmailByID(id uint) (string, error) {
	var email string

	result := r.DB.Model(&entities.Customer{}).
		Select("email").
		Where("id = ?", id).
		First(&email)
	if result.Error != nil {
		return "", result.Error
	}

	return email, nil
}

func (r *customerRepository) FindOneByEmail(email string) (*entities.Customer, error) {
	customer := entities.Customer{}

	result := r.DB.Joins("AuthInfo").
		Where("email = ?", email).
		First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (r *customerRepository) FindOneByID(id uint) (*entities.Customer, error) {
	customer := entities.Customer{}

	result := r.DB.Joins("AuthInfo").
		Joins("BillingAddress").
		Where("id = ?", id).
		First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (r *customerRepository) FindOneByProfileName(profileName string) (*entities.Customer, error) {
	customer := entities.Customer{}

	result := r.DB.Where("profile_name = ?", profileName).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (r *customerRepository) UpdateProfile(id uint, email, profileName string) (*entities.Customer, error) {
	customer := entities.Customer{}

	result := r.DB.Model(&customer).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"email":        email,
			"profile_name": profileName,
		})

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &customer, nil
}

func (r *customerRepository) UpdateSettings(id uint, enableMFA, subscribeNL, notifyExpire bool) (*entities.Customer, error) {
	customer := entities.Customer{}

	result := r.DB.Model(&customer).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"enabled_mfa":   enableMFA,
			"subscribe_nl":  subscribeNL,
			"notify_expire": notifyExpire,
		})

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &customer, nil
}

func (r *customerRepository) UpdatePoints(id uint, points int) (*entities.Customer, error) {
	customer := entities.Customer{}

	result := r.DB.Model(&customer).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Update("points", gorm.Expr("points + ?", points))

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &customer, nil
}

func (r *customerRepository) UpdateUsedSpins(id *uint, usedSpins int) (*entities.Customer, error) {
	customer := entities.Customer{}

	result := r.DB.Model(&customer).
		Clauses(clause.Returning{})

	if id != nil {
		result = result.Where("id = ?", *id)
	} else {
		result = result.Where("1 = 1")
	}

	result = result.Update("used_spins", usedSpins)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &customer, nil
}
