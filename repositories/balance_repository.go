package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BalanceRepository is an interface that defines methods for performing CRUD operations on Balance entity in the database.
type BalanceRepository interface {
	// FindOneByCustomerID retrieves a balance by customer ID.
	FindOneByCustomerID(customerId uint) (*entities.Balance, error)

	// AddCurrentAndTotal adds amount (always positive) to current and total by customer ID.
	AddCurrentAndTotal(tx *gorm.DB, customerId uint, amount float64) (*entities.Balance, error)

	// DeductCurrentAndTotal subtracts amount (always positive) to current and total by customer ID.
	DeductCurrentAndTotal(tx *gorm.DB, customerId uint, amount float64) (*entities.Balance, error)

	// MovePendingToCurrent moves amount (always positive) from pending to current by customer ID.
	MovePendingToCurrent(tx *gorm.DB, customerId uint, movingAmount float64, pendingAmount float64) (*entities.Balance, error)

	// UpdatePendingAndTotal adds amount (can be either positive or negative) to pending and total by customer ID.
	UpdatePendingAndTotal(tx *gorm.DB, customerId uint, amount float64) (*entities.Balance, error)
}

type balanceRepository struct {
	DB *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) BalanceRepository {
	return &balanceRepository{DB: db}
}

func (r *balanceRepository) FindOneByCustomerID(customerId uint) (*entities.Balance, error) {
	balance := entities.Balance{}
	result := r.DB.Where("customer_id = ?", customerId).First(&balance)
	if result.Error != nil {
		return nil, result.Error
	}
	return &balance, nil
}

func (r *balanceRepository) AddCurrentAndTotal(tx *gorm.DB, customerId uint, amount float64) (*entities.Balance, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	balance := entities.Balance{}
	result := dbInst.Model(&balance).
		Clauses(clause.Returning{}).
		Where("customer_id = ?", customerId).
		Updates(map[string]interface{}{
			"current": gorm.Expr("current + ?", amount),
			"total":   gorm.Expr("total + ?", amount),
		})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &balance, nil
}

func (r *balanceRepository) DeductCurrentAndTotal(tx *gorm.DB, customerId uint, amount float64) (*entities.Balance, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	balance := entities.Balance{}
	result := dbInst.Model(&balance).
		Clauses(clause.Returning{}).
		Where("customer_id = ? AND current >= ? AND total >= ?",
			customerId,
			amount,
			amount,
		).
		Updates(map[string]interface{}{
			"current": gorm.Expr("current - ?", amount),
			"total":   gorm.Expr("total - ?", amount),
		})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &balance, nil
}

func (r *balanceRepository) MovePendingToCurrent(tx *gorm.DB, customerId uint, movingAmount float64, pendingAmount float64) (*entities.Balance, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	balance := entities.Balance{}
	result := dbInst.Model(&balance).
		Clauses(clause.Returning{}).
		Where("customer_id = ? AND pending >= ?", customerId, pendingAmount).
		Updates(map[string]interface{}{
			"pending": gorm.Expr("pending - ?", pendingAmount),
			"current": gorm.Expr("current + ?", movingAmount),
			"total":   gorm.Expr("total - ?", pendingAmount-movingAmount),
		})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &balance, nil
}

func (r *balanceRepository) UpdatePendingAndTotal(tx *gorm.DB, customerId uint, amount float64) (*entities.Balance, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	balance := entities.Balance{}
	result := dbInst.Model(&balance).
		Clauses(clause.Returning{}).
		Where("customer_id = ?", customerId).
		Updates(map[string]interface{}{
			"pending": gorm.Expr("pending + ?", amount),
			"total":   gorm.Expr("total + ?", amount),
		})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &balance, nil
}
