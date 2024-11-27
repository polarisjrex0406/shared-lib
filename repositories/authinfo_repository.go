package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// AuthInfoRepository is an interface that defines methods for performing CRUD operations on AuthInfo entity in the database.
type AuthInfoRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new auth info record into the database.
	Create(tx *gorm.DB, authInfo *entities.AuthInfo) error

	// FindOneByCustomerID retrieves an auth info by its customer ID.
	FindOneByCustomerID(customerId uint) (*entities.AuthInfo, error)

	// UpdateEmailVerified changes the email verification status of this auth info identified by customer ID.
	UpdateEmailVerified(tx *gorm.DB, customerId uint, isVerified bool) (*entities.AuthInfo, error)

	// UpdateOneByCustomerID modifies an existing auth info by customer ID.
	UpdateOneByCustomerID(tx *gorm.DB, customerId uint, authInfo *entities.AuthInfo) error

	// UpdateTFAPassed changes the 2FA passing status of this auth info identified by customer ID.
	UpdateTFAPassed(tx *gorm.DB, customerId uint, isPassed bool) (*entities.AuthInfo, error)
}

type authInfoRepository struct {
	DB *gorm.DB
}

func NewAuthInfoRepository(db *gorm.DB) AuthInfoRepository {
	return &authInfoRepository{DB: db}
}

func (r *authInfoRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *authInfoRepository) Create(tx *gorm.DB, authInfo *entities.AuthInfo) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(authInfo)
	return result.Error
}

func (r *authInfoRepository) FindOneByCustomerID(customerId uint) (*entities.AuthInfo, error) {
	authInfo := entities.AuthInfo{}
	result := r.DB.Where("customer_id = ?", customerId).First(&authInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &authInfo, nil
}

func (r *authInfoRepository) UpdateEmailVerified(tx *gorm.DB, customerId uint, isVerified bool) (*entities.AuthInfo, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	authInfo := entities.AuthInfo{}
	result := dbInst.Model(&authInfo).
		Clauses(clause.Returning{}).
		Where("customer_id = ?", customerId).
		Update("email_verified", isVerified)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &authInfo, nil
}

func (r *authInfoRepository) UpdateOneByCustomerID(tx *gorm.DB, customerId uint, authInfo *entities.AuthInfo) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).
		Where("customer_id = ?", customerId).
		Updates(authInfo)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *authInfoRepository) UpdateTFAPassed(tx *gorm.DB, customerId uint, isPassed bool) (*entities.AuthInfo, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	authInfo := entities.AuthInfo{}
	result := dbInst.Model(&authInfo).
		Clauses(clause.Returning{}).
		Where("customer_id = ?", customerId).
		Update("tfa_passed", isPassed)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &authInfo, nil
}
