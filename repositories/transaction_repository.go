package repositories

import (
	"time"

	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// TransactionRepository is an interface that defines methods for performing CRUD operations on Transaction entity in the database.
type TransactionRepository interface {
	// Create inserts a new transaction record into the database.
	Create(transaction *entities.Transaction) error

	FindIDsByStatus(status entities.TransactionStatus) ([]uint, error)

	// FindByCustomerIDWithPagination retrieves transactions identified by customer ID and pagination.
	FindByCustomerIDWithPagination(customerId uint, pageNum, pageSize int) ([]entities.Transaction, int, error)

	FindAllWithPagination(pageNum, pageSize int) ([]entities.Transaction, int, error)

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

func (r *transactionRepository) FindIDsByStatus(status entities.TransactionStatus) ([]uint, error) {
	var ids []uint

	result := r.DB.Model(&entities.Transaction{}).
		Select("id").
		Where("status = ?", status).
		Find(&ids)
	if result.Error != nil {
		return nil, result.Error
	}

	return ids, nil
}

func (r *transactionRepository) FindByCustomerIDWithPagination(customerId uint, pageNum, pageSize int) ([]entities.Transaction, int, error) {
	paginatedResults := []struct {
		ID            uint                       `json:"id"`
		CreatedAt     time.Time                  `json:"created_at"`
		CustomerID    uint                       `json:"customer_id"`
		Status        entities.TransactionStatus `json:"status"`
		PaymentMethod entities.PaymentMethod     `json:"payment_method"`
		TotalCount    int                        `gorm:"column:total_count"`
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
		transactions[i] = entities.Transaction{
			CustomerID:    r.CustomerID,
			Status:        r.Status,
			PaymentMethod: r.PaymentMethod,
		}
		transactions[i].ID = r.ID
		transactions[i].CreatedAt = r.CreatedAt
	}

	return transactions, total, nil
}

func (r *transactionRepository) FindAllWithPagination(pageNum, pageSize int) ([]entities.Transaction, int, error) {
	paginatedResults := []struct {
		ID            uint                       `json:"id"`
		CreatedAt     time.Time                  `json:"created_at"`
		CustomerID    uint                       `json:"customer_id"`
		Status        entities.TransactionStatus `json:"status"`
		PaymentMethod entities.PaymentMethod     `json:"payment_method"`
		TotalCount    int                        `gorm:"column:total_count"`
	}{}
	// Calculate offset
	offset := (pageNum - 1) * pageSize

	result := r.DB.
		Model(&entities.Transaction{}).
		Select("*, COUNT(*) OVER() AS total_count").
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
		transactions[i] = entities.Transaction{
			CustomerID:    r.CustomerID,
			Status:        r.Status,
			PaymentMethod: r.PaymentMethod,
		}
		transactions[i].ID = r.ID
		transactions[i].CreatedAt = r.CreatedAt
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
