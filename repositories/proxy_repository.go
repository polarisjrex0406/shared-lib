package repositories

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

// ProxyRepository is an interface that defines methods for performing CRUD operations on Proxy entity in the database.
type ProxyRepository interface {
	// BeginTx starts a new database transaction.
	BeginTx() *gorm.DB

	// Create inserts a new proxy record into the database.
	Create(tx *gorm.DB, proxy *entities.Proxy) error

	// FindOneByPurchaseID retrieves a proxy by its purchase ID.
	FindOneByPurchaseID(purchaseId uint) (*entities.Proxy, error)
}

type proxyRepository struct {
	DB *gorm.DB
}

func NewProxyRepository(db *gorm.DB) ProxyRepository {
	return &proxyRepository{DB: db}
}

func (r *proxyRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *proxyRepository) Create(tx *gorm.DB, proxy *entities.Proxy) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(proxy)
	return result.Error
}

func (r *proxyRepository) FindOneByPurchaseID(purchaseId uint) (*entities.Proxy, error) {
	proxy := entities.Proxy{}
	result := r.DB.Where("purchase_id = ?", purchaseId).First(&proxy)
	if result.Error != nil {
		return nil, result.Error
	}
	return &proxy, nil
}
