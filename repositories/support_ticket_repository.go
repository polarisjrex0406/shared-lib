package repositories

import (
	"time"

	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SupportTicketRepository is an interface that defines methods for performing CRUD operations on SupportTicket entity in the database.
type SupportTicketRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new support ticket record into the database.
	Create(tx *gorm.DB, supportTicket *entities.SupportTicket) error

	// FindByCustomerID retrieves all support tickets opened by a customer.
	FindByCustomerID(customerId uint) ([]entities.SupportTicket, error)

	// FindByCustomerIDAndStatus retrieves all support tickets of a status opened by a customer.
	FindByCustomerIDAndStatus(customerId uint, status entities.SupportTicketStatus) ([]entities.SupportTicket, error)
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
	result := r.DB.Where("customer_id = ?", customerId).
		Order("opened_at ASC").
		Find(&supportTickets)
	if result.Error != nil {
		return nil, result.Error
	}
	return supportTickets, nil
}

func (r *supportTicketRepository) FindByCustomerIDAndStatus(customerId uint, status entities.SupportTicketStatus) ([]entities.SupportTicket, error) {
	supportTickets := []entities.SupportTicket{}
	result := r.DB.Where("customer_id = ? AND status = ?", customerId, status).
		Order("opened_at ASC").
		Find(&supportTickets)
	if result.Error != nil {
		return nil, result.Error
	}
	return supportTickets, nil
}

func (r *supportTicketRepository) UpdateStatusByID(
	tx *gorm.DB,
	id uint,
	status entities.SupportTicketStatus,
) (*entities.SupportTicket, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	supportTicket := entities.SupportTicket{}
	result := dbInst.Model(&supportTicket).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Update("status", status)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &supportTicket, nil
}

func (r *supportTicketRepository) UpdateStatusAndClosedAtByID(
	tx *gorm.DB,
	id uint,
	status entities.SupportTicketStatus,
	closedAt time.Time,
) (*entities.SupportTicket, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	supportTicket := entities.SupportTicket{}
	result := dbInst.Model(&supportTicket).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":    status,
			"closed_at": closedAt,
		})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &supportTicket, nil
}
