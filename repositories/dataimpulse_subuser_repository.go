package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// DataImpulseSubuserRepository is an interface that defines methods for performing CRUD operations on DataImpulseSubuser entity in the database.
type DataImpulseSubuserRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new DataImpulse sub-user record into the database.
	Create(tx *gorm.DB, dataimpulseSubuser *entities.DataImpulseSubuser) error

	// FindOneBySubuserID retrieves a DataImpulse sub-user by its sub-user ID.
	FindOneBySubuserID(subuserId int) (*entities.DataImpulseSubuser, error)

	// FindOneByProxyID retrieves a DataImpulse sub-user by proxy ID.
	FindOneByProxyID(proxyId uint) (*entities.DataImpulseSubuser, error)

	// UpdateTraffic changes the traffic of this sub-user by its sub-user ID.
	UpdateTraffic(tx *gorm.DB, proxyId uint, subuserId int, traffic int) (*entities.DataImpulseSubuser, error)
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

func (r *dataimpulseSubuserRepository) FindOneBySubuserID(subuserId int) (*entities.DataImpulseSubuser, error) {
	dataimpulseSubuser := entities.DataImpulseSubuser{}
	result := r.DB.Where("subuser_id = ?", subuserId).First(&dataimpulseSubuser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dataimpulseSubuser, nil
}

func (r *dataimpulseSubuserRepository) FindOneByProxyID(proxyId uint) (*entities.DataImpulseSubuser, error) {
	dataimpulseSubuser := entities.DataImpulseSubuser{}
	result := r.DB.Where("proxy_id = ?", proxyId).First(&dataimpulseSubuser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dataimpulseSubuser, nil
}

func (r *dataimpulseSubuserRepository) UpdateTraffic(tx *gorm.DB, proxyId uint, subuserId int, traffic int) (*entities.DataImpulseSubuser, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	dataimpulseSubuser := entities.DataImpulseSubuser{}
	result := dbInst.Model(&dataimpulseSubuser).
		Where("proxy_id = ? AND subuser_id = ?", proxyId, subuserId).
		Update("total_balance", traffic)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &dataimpulseSubuser, nil
}
