package repositories

import (
	"encoding/json"

	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CouponRepository is an interface that defines methods for performing CRUD operations on Coupon entity in the database.
type CouponRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new coupon record into the database.
	Create(tx *gorm.DB, coupon *entities.Coupon) error

	// FindAll retrieves all coupons.
	FindAll() ([]entities.Coupon, error)

	// FindOneByCode retrieves one coupon by its code.
	FindOneByCode(code string) (*entities.Coupon, error)

	// FindOneByID retrieves one coupon identified by its ID.
	FindOneByID(id uint) (*entities.Coupon, error)

	// UpdateRedeemingCustomerIDs modifies the list of redeeming customers' ID for this coupon.
	UpdateRedeemingCustomerIDs(tx *gorm.DB, id uint, redeemingCustomerIds []uint) (*entities.Coupon, error)

	// Delete removes a coupon record from the database using its ID.
	Delete(id uint) error
}

type couponRepository struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) CouponRepository {
	return &couponRepository{DB: db}
}

func (r *couponRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *couponRepository) Create(tx *gorm.DB, coupon *entities.Coupon) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(coupon)
	return result.Error
}

func (r *couponRepository) FindAll() ([]entities.Coupon, error) {
	coupons := []entities.Coupon{}
	result := r.DB.Order("id ASC").Find(&coupons)
	if result.Error != nil {
		return nil, result.Error
	}
	return coupons, nil
}

func (r *couponRepository) FindOneByCode(code string) (*entities.Coupon, error) {
	coupon := entities.Coupon{}
	result := r.DB.Where("code = ?", code).First(&coupon)
	if result.Error != nil {
		return nil, result.Error
	}
	return &coupon, nil
}

func (r *couponRepository) FindOneByID(id uint) (*entities.Coupon, error) {
	coupon := entities.Coupon{}
	result := r.DB.Where("id = ?", id).First(&coupon)
	if result.Error != nil {
		return nil, result.Error
	}
	return &coupon, nil
}

func (r *couponRepository) UpdateRedeemingCustomerIDs(tx *gorm.DB, id uint, redeemingCustomerIds []uint) (*entities.Coupon, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	coupon := entities.Coupon{}
	jsonData, err := json.Marshal(redeemingCustomerIds) // Convert to JSON
	if err != nil {
		return nil, err
	}
	result := dbInst.Model(&coupon).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Update("redeeming_customer_ids", jsonData)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &coupon, nil
}

func (r *couponRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.Coupon{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
