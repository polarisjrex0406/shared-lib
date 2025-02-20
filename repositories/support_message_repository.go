package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// SupportMessageRepository is an interface that defines methods for performing CRUD operations on SupportMessage entity in the database.
type SupportMessageRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new support message record into the database.
	Create(tx *gorm.DB, supportMessage *entities.SupportMessage) error

	// FindByCustomerID retrieves all support messages belongs to a ticket.
	FindBySupportTicketID(supportTicketId uint) ([]entities.SupportMessage, error)

	// FindOneBySupportTicketIDAndSenderType retrieves a support message belongs to a ticket by sender type.
	FindOneBySupportTicketIDAndSenderType(supportTicketId uint, senderType entities.SupportMessageSenderType) (*entities.SupportMessage, error)
}

type supportMessageRepository struct {
	DB *gorm.DB
}

func NewSupportMessageRepository(db *gorm.DB) SupportMessageRepository {
	return &supportMessageRepository{DB: db}
}

func (r *supportMessageRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *supportMessageRepository) Create(tx *gorm.DB, supportMessage *entities.SupportMessage) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(supportMessage)
	return result.Error
}

func (r *supportMessageRepository) FindBySupportTicketID(supportTicketId uint) ([]entities.SupportMessage, error) {
	supportMessages := []entities.SupportMessage{}
	result := r.DB.Where("support_ticket_id = ?", supportTicketId).
		Order("sent_at DESC").
		Find(&supportMessages)
	if result.Error != nil {
		return nil, result.Error
	}
	return supportMessages, nil
}

func (r *supportMessageRepository) FindOneBySupportTicketIDAndSenderType(
	supportTicketId uint,
	senderType entities.SupportMessageSenderType,
) (*entities.SupportMessage, error) {
	supportMessage := entities.SupportMessage{}
	result := r.DB.Where("support_ticket_id = ? AND sender_type = ?", supportTicketId, senderType).
		Order("sent_at DESC").
		First(&supportMessage)
	if result.Error != nil {
		return nil, result.Error
	}
	return &supportMessage, nil
}
