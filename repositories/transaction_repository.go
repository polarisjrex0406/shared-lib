package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// TransactionRepository is an interface that defines methods for performing CRUD operations on Transaction entity in the database.
type TransactionRepository interface {
	// Create inserts a new transaction record into the database.
	Create(transaction *entities.Transaction) error

	// FindByCustomerIDWithPagination retrieves transactions identified by customer ID and pagination.
	FindByCustomerIDWithPagination(customerId uint, pageNum, pageSize int) ([]entities.Transaction, int, error)

	// FindOneByID retrieves a transaction identified by its ID.
	FindOneByID(id uint) (*entities.Transaction, error)

	// UpdateStatus changes the payment status of this transaction identified by its ID.
	UpdateStatus(id uint, status entities.TransactionStatus) (*entities.Transaction, error)
}

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{DB: db}
}

func (r *transactionRepository) Create(transaction *entities.Transaction) error {
	return r.DB.Create(transaction).Error
}

func (r *transactionRepository) FindByCustomerIDWithPagination(customerId uint, pageNum, pageSize int) ([]entities.Transaction, int, error) {
	paginatedResults := []struct {
		Transaction entities.Transaction
		TotalCount  int `gorm:"column:total_count"`
	}{}
	// Calculate offset
	offset := (pageNum - 1) * pageSize
	// Conditional query based on expired
	condition := "customer_id = ?"

	result := r.DB.
		Model(&entities.Transaction{}).
		Select("*, COUNT(*) OVER() AS total_count").
		Where(condition, customerId).
		Order("id ASC").
		Limit(pageSize).
		Offset(offset).
		Find(&paginatedResults)
	if result.Error != nil {
		return nil, -1, result.Error
	}

	var total int
	if len(paginatedResults) > 0 {
		total = paginatedResults[0].TotalCount
	}

	// Convert to Transaction slice
	transactions := make([]entities.Transaction, len(paginatedResults))
	for i, r := range paginatedResults {
		transactions[i] = r.Transaction
	}

	return transactions, total, nil
}

func (r *transactionRepository) FindOneByID(id uint) (*entities.Transaction, error) {
	transaction := &entities.Transaction{}

	result := r.DB.Where("id = ?", id).First(transaction)
	if result.Error != nil {
		return nil, result.Error
	}

	return transaction, nil
}

func (r *transactionRepository) UpdateStatus(id uint, status entities.TransactionStatus) (*entities.Transaction, error) {
	transaction := entities.Transaction{}

	result := r.DB.Model(&transaction).
		Where("id = ?", id).
		Update("status", status)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &transaction, nil
}
