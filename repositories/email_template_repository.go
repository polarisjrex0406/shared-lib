package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// EmailTemplateRepository is an interface that defines methods for performing CRUD operations on EmailTemplate entity in the database.
type EmailTemplateRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new email template record into the database.
	Create(tx *gorm.DB, emailTemplate *entities.EmailTemplate) error

	// FindAll retrieves all email templates.
	FindAll() ([]entities.EmailTemplate, error)

	// FindOneByID retrieves one email template identified by its ID.
	FindOneByID(id uint) (*entities.EmailTemplate, error)

	// Update modifies an existing email template record in the database.
	Update(tx *gorm.DB, emailTemplate *entities.EmailTemplate) error

	// Delete removes an email template record from the database using its ID.
	Delete(id uint) error
}

type emailTemplateRepository struct {
	DB *gorm.DB
}

func NewEmailTemplateRepository(db *gorm.DB) EmailTemplateRepository {
	return &emailTemplateRepository{DB: db}
}

func (r *emailTemplateRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *emailTemplateRepository) Create(tx *gorm.DB, emailTemplate *entities.EmailTemplate) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(emailTemplate)
	return result.Error
}

func (r *emailTemplateRepository) FindAll() ([]entities.EmailTemplate, error) {
	emailTemplates := []entities.EmailTemplate{}
	result := r.DB.Order("id ASC").Find(&emailTemplates)
	if result.Error != nil {
		return nil, result.Error
	}
	return emailTemplates, nil
}

func (r *emailTemplateRepository) FindOneByID(id uint) (*entities.EmailTemplate, error) {
	emailTemplate := entities.EmailTemplate{}
	result := r.DB.Where("id = ?", id).First(&emailTemplate)
	if result.Error != nil {
		return nil, result.Error
	}
	return &emailTemplate, nil
}

func (r *emailTemplateRepository) Update(tx *gorm.DB, emailTemplate *entities.EmailTemplate) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Clauses(clause.Returning{}).Updates(emailTemplate)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *emailTemplateRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.EmailTemplate{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
