package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// SupportTicketRepository is an interface that defines methods for performing CRUD operations on SupportTicket entity in the database.
type SupportTicketRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new support ticket record into the database.
	Create(tx *gorm.DB, supportTicket *entities.SupportTicket) error

	// FindByCustomerID retrieves all support tickets by their customer ID.
	FindByCustomerID(customerId uint) ([]entities.SupportTicket, error)
}

type supportTicketRepository struct {
	DB *gorm.DB
}

func NewSupportTicketRepository(db *gorm.DB) SupportTicketRepository {
	return &supportTicketRepository{DB: db}
}

func (r *supportTicketRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *supportTicketRepository) Create(tx *gorm.DB, supportTicket *entities.SupportTicket) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(supportTicket)
	return result.Error
}

func (r *supportTicketRepository) FindByCustomerID(customerId uint) ([]entities.SupportTicket, error) {
	supportTickets := []entities.SupportTicket{}
	result := r.DB.Where("customer_id = ?", customerId).Find(&supportTickets)
	if result.Error != nil {
		return nil, result.Error
	}
	return supportTickets, nil
}
