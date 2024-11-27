package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// LoyaltyTierRepository is an interface that defines methods for performing CRUD operations on LoyaltyTier entity in the database.
type LoyaltyTierRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new loyalty tier into the database.
	Create(tx *gorm.DB, loyaltyTier *entities.LoyaltyTier) error

	// FindAll retrieves all loyalty tiers.
	FindAll() ([]entities.LoyaltyTier, error)

	// FindOneByID retrieves a loyalty tier identified by its ID.
	FindOneByID(id uint) (*entities.LoyaltyTier, error)

	// FindOneByPoints retrieves a loyalty tier by its points.
	FindOneByPoints(points int) (*entities.LoyaltyTier, error)

	// Update modifies an existing loyalty tier record in the database.
	Update(tx *gorm.DB, loyaltyTier *entities.LoyaltyTier) error

	// Delete removes a loyalty tier record from the database using its ID.
	Delete(id uint) error
}

type loyaltyTierRepository struct {
	DB *gorm.DB
}

func NewLoyaltyTierRepository(db *gorm.DB) LoyaltyTierRepository {
	return &loyaltyTierRepository{DB: db}
}

func (r *loyaltyTierRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *loyaltyTierRepository) Create(tx *gorm.DB, loyaltyTier *entities.LoyaltyTier) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(loyaltyTier)
	return result.Error
}

func (r *loyaltyTierRepository) FindAll() ([]entities.LoyaltyTier, error) {
	loyaltyTiers := []entities.LoyaltyTier{}
	if err := r.DB.Order("points ASC").
		Find(&loyaltyTiers).Error; err != nil {
		return nil, err
	}
	return loyaltyTiers, nil
}

func (r *loyaltyTierRepository) FindOneByID(id uint) (*entities.LoyaltyTier, error) {
	loyaltyTier := entities.LoyaltyTier{}
	if err := r.DB.Where("id = ?", id).
		First(&loyaltyTier).Error; err != nil {
		return nil, err
	}
	return &loyaltyTier, nil
}

func (r *loyaltyTierRepository) FindOneByPoints(points int) (*entities.LoyaltyTier, error) {
	loyaltyTier := entities.LoyaltyTier{}
	if err := r.DB.Where("points <= ?", points).
		Order("points DESC").
		First(&loyaltyTier).Error; err != nil {
		return nil, err
	}
	return &loyaltyTier, nil
}

func (r *loyaltyTierRepository) Update(tx *gorm.DB, loyaltyTier *entities.LoyaltyTier) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(loyaltyTier)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *loyaltyTierRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.LoyaltyTier{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
