package repositories

import (
	"time"

	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ClaimedPrizeRepository is an interface that defines methods for performing CRUD operations on ClaimedPrize entity in the database.
type ClaimedPrizeRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new claimedPrize record into the database.
	Create(tx *gorm.DB, claimedPrize *entities.ClaimedPrize) error

	// FindAll retrieves all claimedPrizes.
	FindAll() ([]entities.ClaimedPrize, error)

	// FindByCustomerIDAndExpireAtAndRedeemAt retrieves claimedPrizes by its customer ID, expiring and redeemed date.
	FindByCustomerIDAndExpireAtAndRedeemAt(
		customerId uint, expireAt time.Time, redeemAt time.Time,
	) ([]entities.ClaimedPrize, error)

	// FindOneByID retrieves one claimedPrize identified by its ID.
	FindOneByID(id uint) (*entities.ClaimedPrize, error)

	// Update modifies an existing claimedPrize record in the database.
	Update(tx *gorm.DB, claimedPrize *entities.ClaimedPrize) error

	// Delete removes a claimedPrize record from the database using its ID.
	Delete(id uint) error
}

type claimedPrizeRepository struct {
	DB *gorm.DB
}

func NewClaimedPrizeRepository(db *gorm.DB) ClaimedPrizeRepository {
	return &claimedPrizeRepository{DB: db}
}

func (r *claimedPrizeRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *claimedPrizeRepository) Create(tx *gorm.DB, claimedPrize *entities.ClaimedPrize) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(claimedPrize)
	return result.Error
}

func (r *claimedPrizeRepository) FindAll() ([]entities.ClaimedPrize, error) {
	claimedPrizes := []entities.ClaimedPrize{}
	result := r.DB.Order("id ASC").Find(&claimedPrizes)
	if result.Error != nil {
		return nil, result.Error
	}
	return claimedPrizes, nil
}

func (r *claimedPrizeRepository) FindByCustomerIDAndExpireAtAndRedeemAt(
	customerId uint, expireAt time.Time, redeemAt time.Time,
) ([]entities.ClaimedPrize, error) {
	claimedPrizes := []entities.ClaimedPrize{}
	result := r.DB.Where(
		"customer_id = ? AND expire_at > ? AND redeem_at = ?",
		customerId, expireAt, redeemAt,
	).Find(&claimedPrizes)
	if result.Error != nil {
		return nil, result.Error
	}
	return claimedPrizes, nil
}

func (r *claimedPrizeRepository) FindOneByID(id uint) (*entities.ClaimedPrize, error) {
	claimedPrize := entities.ClaimedPrize{}
	result := r.DB.Where("id = ?", id).First(&claimedPrize)
	if result.Error != nil {
		return nil, result.Error
	}
	return &claimedPrize, nil
}

func (r *claimedPrizeRepository) Update(tx *gorm.DB, claimedPrize *entities.ClaimedPrize) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(claimedPrize)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *claimedPrizeRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.ClaimedPrize{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
