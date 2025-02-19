package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// AttachmentRepository is an interface that defines methods for performing CRUD operations on Attachment entity in the database.
type AttachmentRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new attachment record into the database.
	Create(tx *gorm.DB, attachment *entities.Attachment) error

	// FindOneByID retrieves an attachment identified by its ID.
	FindOneByID(id uint) ([]entities.Attachment, error)
}

type attachmentRepository struct {
	DB *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) AttachmentRepository {
	return &attachmentRepository{DB: db}
}

func (r *attachmentRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *attachmentRepository) Create(tx *gorm.DB, attachment *entities.Attachment) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(attachment)
	return result.Error
}

func (r *attachmentRepository) FindOneByID(id uint) ([]entities.Attachment, error) {
	attachments := []entities.Attachment{}
	result := r.DB.Where("id = ?", id).First(&attachments)
	if result.Error != nil {
		return nil, result.Error
	}
	return attachments, nil
}
