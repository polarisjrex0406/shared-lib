package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// LoyaltyTierRepository is an interface that defines methods for performing CRUD operations on LoyaltyTier entity in the database.
type LoyaltyTierRepository interface {
	Create(loyaltyTier *entities.LoyaltyTier) error

	FindAll() ([]entities.LoyaltyTier, error)

	FindOneByCustomerID(customerId uint) (*entities.LoyaltyTier, error)

	FindOneByPoints(points int) (*entities.LoyaltyTier, error)

	Update(id uint, loyaltyTier *entities.LoyaltyTier) error

	Delete(id uint) error
}

type loyaltyTierRepository struct {
	DB *gorm.DB
}

func NewLoyaltyTierRepository(db *gorm.DB) LoyaltyTierRepository {
	return &loyaltyTierRepository{DB: db}
}

func (r *loyaltyTierRepository) Create(loyaltyTier *entities.LoyaltyTier) error {
	result := r.DB.Create(loyaltyTier)
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

func (r *loyaltyTierRepository) FindOneByCustomerID(customerId uint) (*entities.LoyaltyTier, error) {
	loyaltyTier := entities.LoyaltyTier{}

	if err := r.DB.Raw(
		"SELECT * FROM tbl_loyalty_tiers lt LEFT JOIN tbl_customers c ON lt.points <= c.points WHERE c.id = ? AND lt.deleted_at IS NULL ORDER BY lt.points DESC, lt.id LIMIT 1",
		customerId,
	).Scan(&loyaltyTier).Error; err != nil {
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

func (r *loyaltyTierRepository) Update(id uint, loyaltyTier *entities.LoyaltyTier) error {
	result := r.DB.Model(&entities.LoyaltyTier{}).
		Clauses(clause.Returning{}).
		Select("commission_rate", "daily_spins", "discount_rate", "discount_cap", "points", "rank").
		Where("id = ?", id).
		Updates(*loyaltyTier)

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
