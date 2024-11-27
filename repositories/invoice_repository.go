package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// InvoiceRepository is an interface that defines methods for performing CRUD operations on Invoice entity in the database.
type InvoiceRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new invoice record into the database.
	Create(tx *gorm.DB, invoice *entities.Invoice) error

	// FindAll retrieves all invoices.
	FindAll() ([]entities.Invoice, error)

	// FindOneByID retrieves an invoice identified by its ID.
	FindOneByID(id uint) (*entities.Invoice, error)

	// FindByPurchaseID retrieves invoices by its purchase ID.
	FindByPurchaseID(purchaseId uint) ([]entities.Invoice, error)

	// Update modifies an existing invoice.
	Update(tx *gorm.DB, invoice *entities.Invoice) error

	// Delete removes an invoice record from the database using its ID.
	Delete(tx *gorm.DB, id uint) error
}

type invoiceRepository struct {
	DB *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) InvoiceRepository {
	return &invoiceRepository{DB: db}
}

func (r *invoiceRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *invoiceRepository) Create(tx *gorm.DB, invoice *entities.Invoice) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(invoice)
	return result.Error
}

func (r *invoiceRepository) FindAll() ([]entities.Invoice, error) {
	invoices := []entities.Invoice{}
	result := r.DB.Order("id ASC").Find(&invoices)
	if result.Error != nil {
		return nil, result.Error
	}
	return invoices, nil
}

func (r *invoiceRepository) FindOneByID(id uint) (*entities.Invoice, error) {
	invoice := entities.Invoice{}
	result := r.DB.Where("id = ?", id).First(&invoice)
	if result.Error != nil {
		return nil, result.Error
	}
	return &invoice, nil
}

func (r *invoiceRepository) FindByPurchaseID(purchaseId uint) ([]entities.Invoice, error) {
	invoices := []entities.Invoice{}
	result := r.DB.Where("purchase_id = ?", purchaseId).Find(&invoices)
	if result.Error != nil {
		return nil, result.Error
	}
	return invoices, nil
}

func (r *invoiceRepository) Update(tx *gorm.DB, invoice *entities.Invoice) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(invoice)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *invoiceRepository) Delete(tx *gorm.DB, id uint) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Delete(&entities.Invoice{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
