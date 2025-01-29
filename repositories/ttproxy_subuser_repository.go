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

	FindAll() ([]entities.TTProxySubuser, error)

	FindWithPagination(pageNum int, pageSize int) ([]entities.TTProxySubuser, error)

	FindOneByID(id uint) (*entities.TTProxySubuser, error)

	// FindOneByProxyID retrieves a ttproxy sub-user by its proxy ID.
	FindOneByProxyID(proxyId uint) (*entities.TTProxySubuser, error)

	Update(tx *gorm.DB, id uint, proxyId uint, obtainLimit int, trafficLeft int64, ipDuration int, remark string, totalTraffic int64, ipUsed int) (*entities.TTProxySubuser, error)

	// UpdateTraffic changes the left traffic of this sub-user by key.
	UpdateTrafficLeftByKey(tx *gorm.DB, key string, trafficLeft int64) (*entities.TTProxySubuser, error)

	Delete(tx *gorm.DB, id uint) error
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

func (r *ttproxySubuserRepository) FindAll() ([]entities.TTProxySubuser, error) {
	ttproxySubusers := []entities.TTProxySubuser{}
	result := r.DB.Order("id ASC").Find(&ttproxySubusers)
	if result.Error != nil {
		return nil, result.Error
	}
	return ttproxySubusers, nil
}

func (r *ttproxySubuserRepository) FindWithPagination(pageNum int, pageSize int) ([]entities.TTProxySubuser, error) {
	ttproxySubusers := []entities.TTProxySubuser{}
	offset := (pageNum - 1) * pageSize
	result := r.DB.Where("1 = 1").
		Limit(pageSize).
		Offset(offset).
		Find(&ttproxySubusers)

	if result.Error != nil {
		return nil, result.Error
	}

	return ttproxySubusers, nil
}

func (r *ttproxySubuserRepository) FindOneByID(id uint) (*entities.TTProxySubuser, error) {
	ttproxySubuser := entities.TTProxySubuser{}
	result := r.DB.Where("id = ?", id).First(&ttproxySubuser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ttproxySubuser, nil
}

func (r *ttproxySubuserRepository) FindOneByProxyID(proxyId uint) (*entities.TTProxySubuser, error) {
	ttproxySubuser := entities.TTProxySubuser{}
	result := r.DB.Where("proxy_id = ?", proxyId).First(&ttproxySubuser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ttproxySubuser, nil
}

func (r *ttproxySubuserRepository) Update(tx *gorm.DB, id uint, proxyId uint,
	obtainLimit int, trafficLeft int64, ipDuration int, remark string, totalTraffic int64, ipUsed int,
) (*entities.TTProxySubuser, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	ttproxySubuser := entities.TTProxySubuser{}
	result := dbInst.Model(&ttproxySubuser).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"proxy_id":      proxyId,
			"obtain_limit":  obtainLimit,
			"traffic_left":  trafficLeft,
			"ip_duration":   ipDuration,
			"remark":        remark,
			"total_traffic": totalTraffic,
			"ip_used":       ipUsed,
		})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &ttproxySubuser, nil
}

func (r *ttproxySubuserRepository) UpdateTrafficLeftByKey(tx *gorm.DB, key string, trafficLeft int64) (*entities.TTProxySubuser, error) {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	ttproxySubuser := entities.TTProxySubuser{}
	result := dbInst.Model(&ttproxySubuser).
		Clauses(clause.Returning{}).
		Where("key = ?", key).
		Update("traffic_left", trafficLeft)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &ttproxySubuser, nil
}

func (r *ttproxySubuserRepository) Delete(tx *gorm.DB, id uint) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Delete(&entities.TTProxySubuser{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
