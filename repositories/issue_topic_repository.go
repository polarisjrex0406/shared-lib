package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// IssueTopicRepository is an interface that defines methods for performing CRUD operations on IssueTopic entity in the database.
type IssueTopicRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new issue topic record into the database.
	Create(tx *gorm.DB, issueTopic *entities.IssueTopic) error

	// FindAll retrieves all issue topics.
	FindAll() ([]entities.IssueTopic, error)

	// FindOneByID retrieves an issue topic identified by its ID.
	FindOneByID(id uint) (*entities.IssueTopic, error)

	// FindOneByName retrieves an issue topic by its name.
	FindOneByName(name string) (*entities.IssueTopic, error)

	// Delete removes an issue topic record from the database using its ID.
	Delete(id uint) error
}

type issueTopicRepository struct {
	DB *gorm.DB
}

func NewIssueTopicRepository(db *gorm.DB) IssueTopicRepository {
	return &issueTopicRepository{DB: db}
}

func (r *issueTopicRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *issueTopicRepository) Create(tx *gorm.DB, issueTopic *entities.IssueTopic) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(issueTopic)
	return result.Error
}

func (r *issueTopicRepository) FindAll() ([]entities.IssueTopic, error) {
	topics := []entities.IssueTopic{}
	result := r.DB.Order("id ASC").Find(&topics)
	if result.Error != nil {
		return nil, result.Error
	}
	return topics, nil
}

func (r *issueTopicRepository) FindOneByID(id uint) (*entities.IssueTopic, error) {
	issueTopic := entities.IssueTopic{}
	result := r.DB.Where("id = ?", id).First(&issueTopic)
	if result.Error != nil {
		return nil, result.Error
	}
	return &issueTopic, nil
}

func (r *issueTopicRepository) FindOneByName(name string) (*entities.IssueTopic, error) {
	issueTopic := entities.IssueTopic{}
	result := r.DB.Where("name = ?", name).First(&issueTopic)
	if result.Error != nil {
		return nil, result.Error
	}
	return &issueTopic, nil
}

func (r *issueTopicRepository) Delete(id uint) error {
	result := r.DB.Delete(&entities.IssueTopic{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
