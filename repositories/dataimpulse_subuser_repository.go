package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// DataImpulseSubuserRepository is an interface that defines methods for performing CRUD operations on DataImpulseSubuser entity in the database.
type DataImpulseSubuserRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new dataimpulse sub-user record into the database.
	Create(tx *gorm.DB, dataimpulseSubuser *entities.DataImpulseSubuser) error

	// FindOneByPurchaseID retrieves a dataimpulse sub-user by its purchase ID.
	FindOneByPurchaseID(purchaseId uint) (*entities.DataImpulseSubuser, error)

	// FindOneBySubuserID retrieves a dataimpulse sub-user by its sub-user ID.
	FindOneBySubuserID(subuserId int) (*entities.DataImpulseSubuser, error)

	// UpdatePassword changes the password of this sub-user by its sub-user ID.
	UpdatePassword(tx *gorm.DB, purchaseId uint, subuserId int, password string) (*entities.DataImpulseSubuser, error)

	// UpdateTraffic changes the traffic of this sub-user by its sub-user ID.
	UpdateTraffic(tx *gorm.DB, purchaseId uint, subuserId int, traffic int) (*entities.DataImpulseSubuser, error)
}

type dataimpulseSubuserRepository struct {
	DB *gorm.DB
}

func NewDataImpulseSubuserRepository(db *gorm.DB) DataImpulseSubuserRepository {
	return &dataimpulseSubuserRepository{DB: db}
}

func (r *dataimpulseSubuserRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *dataimpulseSubuserRepository) Create(tx *gorm.DB, dataimpulseSubuser *entities.DataImpulseSubuser) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(dataimpulseSubuser)
	return result.Error
}

func (r *dataimpulseSubuserRepository) FindOneByPurchaseID(purchaseId uint) (*entities.DataImpulseSubuser, error) {
	dataimpulseSubuser := entities.DataImpulseSubuser{}
	result := r.DB.Where("purchase_id = ?", purchaseId).First(&dataimpulseSubuser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dataimpulseSubuser, nil
}

func (r *dataimpulseSubuserRepository) FindOneBySubuserID(subuserId int) (*entities.DataImpulseSubuser, error) {
	dataimpulseSubuser := entities.DataImpulseSubuser{}
	result := r.DB.Where("subuser_id = ?", subuserId).First(&dataimpulseSubuser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dataimpulseSubuser, nil
}

func (r *dataimpulseSubuserRepository) UpdatePassword(tx *gorm.DB, purchaseId uint, subuserId int, password string) (*entities.DataImpulseSubuser, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	dataimpulseSubuser := entities.DataImpulseSubuser{}
	result := dbInst.Model(&dataimpulseSubuser).
		Where("purchase_id = ? AND subuser_id = ?", purchaseId, subuserId).
		Update("password", password)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &dataimpulseSubuser, nil
}

func (r *dataimpulseSubuserRepository) UpdateTraffic(tx *gorm.DB, purchaseId uint, subuserId int, traffic int) (*entities.DataImpulseSubuser, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	dataimpulseSubuser := entities.DataImpulseSubuser{}
	result := dbInst.Model(&dataimpulseSubuser).
		Where("purchase_id = ? AND subuser_id = ?", purchaseId, subuserId).
		Update("total_balance", traffic)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &dataimpulseSubuser, nil
}
