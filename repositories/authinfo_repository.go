package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// AuthInfoRepository is an interface that defines methods for performing CRUD operations on AuthInfo entity in the database.
type AuthInfoRepository interface {
	FindOneByCustomerID(customerId uint) (*entities.AuthInfo, error)

	UpdateAPIKey(customerId uint, apiKey string) (*entities.AuthInfo, error)

	UpdateEmailVerified(customerId uint, emailVerified bool) (*entities.AuthInfo, error)

	UpdateMFAPassed(customerId uint, mfaPassed bool) (*entities.AuthInfo, error)

	UpdatePassword(customerId uint, password string) (*entities.AuthInfo, error)
}

type authInfoRepository struct {
	DB *gorm.DB
}

func NewAuthInfoRepository(db *gorm.DB) AuthInfoRepository {
	return &authInfoRepository{DB: db}
}

func (r *authInfoRepository) FindOneByCustomerID(customerId uint) (*entities.AuthInfo, error) {
	authInfo := entities.AuthInfo{}

	result := r.DB.Where("customer_id = ?", customerId).First(&authInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &authInfo, nil
}

func (r *authInfoRepository) UpdateAPIKey(customerId uint, apiKey string) (*entities.AuthInfo, error) {
	authInfo := entities.AuthInfo{
		APIKey: apiKey,
	}

	if err := r.update(customerId, []string{"api_key"}, &authInfo); err != nil {
		return nil, err
	}

	return &authInfo, nil
}

func (r *authInfoRepository) UpdateEmailVerified(customerId uint, emailVerified bool) (*entities.AuthInfo, error) {
	authInfo := entities.AuthInfo{
		EmailVerified: emailVerified,
	}

	if err := r.update(customerId, []string{"email_verified"}, &authInfo); err != nil {
		return nil, err
	}

	return &authInfo, nil
}

func (r *authInfoRepository) UpdateMFAPassed(customerId uint, mfaPassed bool) (*entities.AuthInfo, error) {
	authInfo := entities.AuthInfo{
		MFAPassed: mfaPassed,
	}

	if err := r.update(customerId, []string{"mfa_passed"}, &authInfo); err != nil {
		return nil, err
	}

	return &authInfo, nil
}

func (r *authInfoRepository) UpdatePassword(customerId uint, password string) (*entities.AuthInfo, error) {
	authInfo := entities.AuthInfo{
		Password: password,
	}

	if err := r.update(customerId, []string{"password"}, &authInfo); err != nil {
		return nil, err
	}

	return &authInfo, nil
}

func (r *authInfoRepository) update(customerId uint, fields []string, authInfo *entities.AuthInfo) error {
	result := r.DB.Model(&entities.AuthInfo{}).
		Clauses(clause.Returning{}).
		Where("customer_id = ?", customerId).
		Select(fields).
		Updates(authInfo)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
