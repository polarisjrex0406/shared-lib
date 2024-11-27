package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// TransactionRepository is an interface that defines methods for performing CRUD operations on Transaction entity in the database.
type TransactionRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new transaction record into the database.
	Create(tx *gorm.DB, transaction *entities.Transaction) error

	// FindByCustomerIDWithPagination retrieves transactions identified by customer ID and pagination.
	FindByCustomerIDWithPagination(
		customerId uint,
		pageNum int,
		pageSize int,
	) ([]entities.Transaction, error)

	// FindOneByID retrieves a transaction identified by its ID.
	FindOneByID(id uint) (*entities.Transaction, error)

	// UpdateStatus changes the payment status of this transaction identified by its ID.
	UpdateStatus(tx *gorm.DB, id uint, status entities.TransactionStatus) (*entities.Transaction, error)
}

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{DB: db}
}

func (r *transactionRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *transactionRepository) Create(tx *gorm.DB, transaction *entities.Transaction) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(transaction)
	return result.Error
}

func (r *transactionRepository) FindByCustomerIDWithPagination(
	customerId uint,
	pageNum int,
	pageSize int,
) ([]entities.Transaction, error) {
	transactions := []entities.Transaction{}
	// Calculate offset
	offset := (pageNum - 1) * pageSize
	// Conditional query based on expired
	condition := "customer_id = ?"
	result := r.DB.Where(condition, customerId).
		Limit(pageSize).
		Offset(offset).
		Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (r *transactionRepository) FindOneByID(id uint) (*entities.Transaction, error) {
	transaction := &entities.Transaction{}
	result := r.DB.Where("id = ?", id).First(transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	return transaction, nil
}

func (r *transactionRepository) UpdateStatus(tx *gorm.DB, id uint, status entities.TransactionStatus) (*entities.Transaction, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	transaction := entities.Transaction{}
	result := dbInst.Model(&transaction).
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
