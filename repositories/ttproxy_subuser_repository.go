package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TTProxySubuserRepository is an interface that defines methods for performing CRUD operations on TTProxySubuser entity in the database.
type TTProxySubuserRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new ttproxy sub-user record into the database.
	Create(tx *gorm.DB, ttproxySubuser *entities.TTProxySubuser) error

	// FindOneByProxyID retrieves a ttproxy sub-user by its proxy ID.
	FindOneByProxyID(proxyId uint) (*entities.TTProxySubuser, error)

	// UpdateTraffic changes the traffic of this sub-user by proxy ID and key.
	UpdateTraffic(tx *gorm.DB, proxyId uint, key string, traffic int64) (*entities.TTProxySubuser, error)
}

type ttproxySubuserRepository struct {
	DB *gorm.DB
}

func NewTTProxySubuserRepository(db *gorm.DB) TTProxySubuserRepository {
	return &ttproxySubuserRepository{DB: db}
}

func (r *ttproxySubuserRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *ttproxySubuserRepository) Create(tx *gorm.DB, ttproxySubuser *entities.TTProxySubuser) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(ttproxySubuser)
	return result.Error
}

func (r *ttproxySubuserRepository) FindOneByProxyID(proxyId uint) (*entities.TTProxySubuser, error) {
	ttproxySubuser := entities.TTProxySubuser{}
	result := r.DB.Where("proxy_id = ?", proxyId).First(&ttproxySubuser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ttproxySubuser, nil
}

func (r *ttproxySubuserRepository) UpdateTraffic(tx *gorm.DB, proxyId uint, key string, traffic int64) (*entities.TTProxySubuser, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	ttproxySubuser := entities.TTProxySubuser{}
	result := dbInst.Model(&ttproxySubuser).
		Clauses(clause.Returning{}).
		Where("proxy_id = ? AND key = ?", proxyId, key).
		Update("traffic", traffic)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &ttproxySubuser, nil
}
