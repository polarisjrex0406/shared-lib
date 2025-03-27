package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CryptomusTransactionRepository is an interface that defines methods for performing CRUD operations on CryptomusTransaction entity in the database.
type CryptomusTransactionRepository interface {
	// Create inserts a new cryptomus transaction record into the database.
	Create(cryptomusTransaction *entities.CryptomusTransaction) error

	// FindAll retrieves all cryptomus transactions.
	FindAll() ([]entities.CryptomusTransaction, error)

	// FindOneByOrderID retrieves a cryptomus transaction by its order ID.
	FindOneByOrderID(orderID string) (*entities.CryptomusTransaction, error)

	// FindOneByTransactionID retrieves one cryptomus transaction identified by its transaction ID.
	FindOneByTransactionID(transactionId uint) (*entities.CryptomusTransaction, error)

	FindByTransactionIDAndPaymentStatus(transactionIds []uint, paymentStatus entities.CryptomusPaymentStatus) ([]entities.CryptomusTransaction, error)

	// Update modifies an existing cryptomus transaction by its order ID.
	Update(orderID string, cryptomusTransaction *entities.CryptomusTransaction) error
}

type cryptomusTransactionRepository struct {
	DB *gorm.DB
}

func NewCryptomusTransactionRepository(db *gorm.DB) CryptomusTransactionRepository {
	return &cryptomusTransactionRepository{DB: db}
}

func (r *cryptomusTransactionRepository) Create(cryptomusTransaction *entities.CryptomusTransaction) error {
	return r.DB.Create(cryptomusTransaction).Error
}

func (r *cryptomusTransactionRepository) FindAll() ([]entities.CryptomusTransaction, error) {
	cryptomusTransactions := []entities.CryptomusTransaction{}

	result := r.DB.Order("id ASC").Find(&cryptomusTransactions)
	if result.Error != nil {
		return nil, result.Error
	}

	return cryptomusTransactions, nil
}

func (r *cryptomusTransactionRepository) FindOneByOrderID(orderID string) (*entities.CryptomusTransaction, error) {
	cryptomusTransaction := entities.CryptomusTransaction{}

	result := r.DB.Where("order_id = ?", orderID).First(&cryptomusTransaction)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cryptomusTransaction, nil
}

func (r *cryptomusTransactionRepository) FindOneByTransactionID(transactionId uint) (*entities.CryptomusTransaction, error) {
	cryptomusTransaction := entities.CryptomusTransaction{}

	result := r.DB.Where("transaction_id = ?", transactionId).First(&cryptomusTransaction)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cryptomusTransaction, nil
}

func (r *cryptomusTransactionRepository) FindByTransactionIDAndPaymentStatus(transactionIds []uint, paymentStatus entities.CryptomusPaymentStatus) ([]entities.CryptomusTransaction, error) {
	cryptomusTransactions := []entities.CryptomusTransaction{}
	err := r.DB.Model(&entities.CryptomusTransaction{}).
		Where("transaction_id IN ? AND payment_status = ?", transactionIds, paymentStatus).
		Order("created_at ASC").
		Find(&cryptomusTransactions).Error

	return cryptomusTransactions, err
}

func (r *cryptomusTransactionRepository) Update(orderID string, cryptomusTransaction *entities.CryptomusTransaction) error {
	result := r.DB.
		Model(&entities.CryptomusTransaction{}).
		Clauses(clause.Returning{}).
		Where("order_id = ?", orderID).
		Updates(cryptomusTransaction)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
