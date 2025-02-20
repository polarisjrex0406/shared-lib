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

	// FindByCustomerIDAndClosedAt retrieves all support tickets by closed time and a customer.
	FindByCustomerIDAndClosedAt(customerId uint, closedAt time.Time) ([]entities.SupportTicket, error)

	// FindByCustomerIDAndClosedAtAndIssueTopicID retrieves all support tickets with closed time and topic opened by a customer.
	FindByCustomerIDAndClosedAtAndIssueTopicID(customerId uint, closedAt time.Time, issueTopicId uint) ([]entities.SupportTicket, error)

	// FindOneByIDAndCustomerID retrieves a support ticket identified by its ID for a customer
	FindOneByIDAndCustomerID(id uint, customerId uint) (*entities.SupportTicket, error)

	// UpdateStatusByID modifies the status of a ticket identified by its ID.
	UpdateStatusByID(tx *gorm.DB, id uint, status entities.SupportTicketStatus) (*entities.SupportTicket, error)

	// UpdateStatusAndClosedAtByID modifies the status and closed time of a ticket identified by its ID.
	UpdateStatusAndClosedAtByID(tx *gorm.DB, id uint, status entities.SupportTicketStatus, closedAt time.Time) (*entities.SupportTicket, error)
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

func (r *supportTicketRepository) FindByCustomerIDAndClosedAt(customerId uint, closedAt time.Time) ([]entities.SupportTicket, error) {
	supportTickets := []entities.SupportTicket{}
	result := r.DB.Where("customer_id = ? AND closed_at = ?", customerId, closedAt).
		Order("opened_at ASC").
		Find(&supportTickets)
	if result.Error != nil {
		return nil, result.Error
	}
	return supportTickets, nil
}

func (r *supportTicketRepository) FindByCustomerIDAndClosedAtAndIssueTopicID(
	customerId uint,
	closedAt time.Time,
	issueTopicId uint,
) ([]entities.SupportTicket, error) {
	supportTickets := []entities.SupportTicket{}
	result := r.DB.Where("customer_id = ? AND closed_at = ? AND issue_topic_id = ?", customerId, closedAt, issueTopicId).
		Order("opened_at ASC").
		Find(&supportTickets)
	if result.Error != nil {
		return nil, result.Error
	}
	return supportTickets, nil
}

func (r *supportTicketRepository) FindOneByIDAndCustomerID(id uint, customerId uint) (*entities.SupportTicket, error) {
	supportTicket := entities.SupportTicket{}
	result := r.DB.Where("id = ? AND customer_id = ?", id, customerId).
		First(&supportTicket)
	if result.Error != nil {
		return nil, result.Error
	}
	return &supportTicket, nil
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
