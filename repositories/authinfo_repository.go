package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// AuthInfoRepository is an interface that defines methods for performing CRUD operations on AuthInfo entity in the database.
type AuthInfoRepository interface {
	Create(authInfo *entities.AuthInfo) error

	FindByCustomerID(customerId uint) (*entities.AuthInfo, error)

	UpdateAPIKey(customerId uint, apiKey string) error

	UpdateEmailVerified(customerId uint, emailVerified bool) error

	UpdateMFAPassed(customerId uint, mfaPassed bool) error

	UpdatePassword(customerId uint, password string) error
}

type authInfoRepository struct {
	DB *gorm.DB
}

func NewAuthInfoRepository(db *gorm.DB) AuthInfoRepository {
	return &authInfoRepository{DB: db}
}

func (r *authInfoRepository) Create(authInfo *entities.AuthInfo) error {
	result := r.DB.Create(authInfo)
	return result.Error
}

func (r *authInfoRepository) FindByCustomerID(customerId uint) (*entities.AuthInfo, error) {
	authInfo := entities.AuthInfo{}

	result := r.DB.Where("customer_id = ?", customerId).First(&authInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &authInfo, nil
}

func (r *authInfoRepository) UpdateAPIKey(customerId uint, apiKey string) error {
	authInfo := entities.AuthInfo{
		APIKey: apiKey,
	}

	if err := r.update(customerId, []string{"api_key"}, &authInfo); err != nil {
		return err
	}

	return nil
}

func (r *authInfoRepository) UpdateEmailVerified(customerId uint, emailVerified bool) error {
	authInfo := entities.AuthInfo{
		EmailVerified: emailVerified,
	}

	if err := r.update(customerId, []string{"email_verified"}, &authInfo); err != nil {
		return err
	}

	return nil
}

func (r *authInfoRepository) UpdateMFAPassed(customerId uint, mfaPassed bool) error {
	authInfo := entities.AuthInfo{
		MFAPassed: mfaPassed,
	}

	if err := r.update(customerId, []string{"mfa_passed"}, &authInfo); err != nil {
		return err
	}

	return nil
}

func (r *authInfoRepository) UpdatePassword(customerId uint, password string) error {
	authInfo := entities.AuthInfo{
		Password: password,
	}

	if err := r.update(customerId, []string{"password"}, &authInfo); err != nil {
		return err
	}

	return nil
}

func (r *authInfoRepository) update(customerId uint, fields []string, authInfo *entities.AuthInfo) error {
	result := r.DB.Model(authInfo).
		Clauses(clause.Returning{}).
		Where("customer_id = ?", customerId).
		Select(fields).
		Updates(*authInfo)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
