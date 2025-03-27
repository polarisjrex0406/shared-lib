package repositories

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// PurchaseRepository is an interface that defines methods for performing CRUD operations on Purchase entity in the database.
type (
	PurchaseRepository interface {
		// BeginTx starts a new database transaction.
		BeginTx() *gorm.DB

		// Create inserts a new purchase record into the database.
		Create(tx *gorm.DB, purchase *entities.Purchase) error

		// CountActiveByCustomerID counts number of active purchases by their customer ID.
		CountActiveByCustomerID(customerId uint) (*int64, error)

		// FindByCustomerID retrieves purchases by their customer ID.
		FindByCustomerID(customerId uint) ([]entities.Purchase, error)

		// FindByCustomerIDAndExpireAtWithPagination retrieves purchases by their customer ID and expiring time with pagination.
		FindByCustomerIDAndExpireAtWithPagination(customerId uint, expired bool, expireAt time.Time, pageNum int, pageSize int) ([]entities.Purchase, error)

		// FindByCustomerIDAndStartAt retrieves purchases by their customer ID and starting time.
		FindByCustomerIDAndStartAt(customerId uint, startAt time.Time) ([]entities.Purchase, error)

		// FindByProductIDAndCustomerID retrieves purchases by their product ID and customer ID.
		FindByProductIDAndCustomerID(productId uint, customerId uint) ([]entities.Purchase, error)

		// FindByBandwidthAndStartAtAndExpireAt retrieves purchases by bandwidth, starting time and expiring time.
		FindByBandwidthAndStartAtAndExpireAt(bandwidth int, startAt time.Time, expireAt time.Time) ([]entities.Purchase, error)

		// FindByDurationAndStartAtAndExpireAt retrieves purchases by duration, starting time and expiring time.
		FindByDurationAndStartAtAndExpireAt(duration *int, startAt time.Time, expireAt time.Time) ([]entities.Purchase, error)

		FindByExpireAtWithRange(beginAt, endAt time.Time) ([]entities.Purchase, error)

		FindByCustomerIDAndStartAtWithRange(customerIds []uint, beginAt, endAt time.Time) ([]entities.Purchase, error)

		// FindOneByID retrieves a purchase identified by its ID.
		FindOneByID(id uint) (*entities.Purchase, error)

		// Update modifies an existing purchase record in the database.
		Update(tx *gorm.DB, purchase *entities.Purchase) error

		// Delete removes a purchase record from the database using its ID.
		Delete(tx *gorm.DB, id uint) error
	}

	purchaseRepository struct {
		DB *gorm.DB
	}
)

func NewPurchaseRepository(db *gorm.DB) PurchaseRepository {
	return &purchaseRepository{DB: db}
}

// BeginTx starts a new transaction
func (r *purchaseRepository) BeginTx() *gorm.DB {
	return r.DB.Begin()
}

func (r *purchaseRepository) Create(tx *gorm.DB, purchase *entities.Purchase) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Create(purchase)
	return result.Error
}

func (r *purchaseRepository) CountActiveByCustomerID(customerId uint) (*int64, error) {
	var activePurchaseCnt int64
	result := r.DB.Model(&entities.Purchase{}).
		Where("customer_id = ? AND expire_at <= ?", customerId, time.UTC.String()).
		Count(&activePurchaseCnt)
	if result.Error != nil {
		return nil, result.Error
	}
	return &activePurchaseCnt, nil
}

func (r *purchaseRepository) FindByCustomerID(customerId uint) ([]entities.Purchase, error) {
	purchases := []entities.Purchase{}
	result := r.DB.Where("customer_id = ?", customerId).Find(&purchases)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchases, nil
}

func (r *purchaseRepository) FindByCustomerIDAndExpireAtWithPagination(
	customerId uint,
	expired bool,
	expireAt time.Time,
	pageNum int,
	pageSize int,
) ([]entities.Purchase, error) {
	purchases := []entities.Purchase{}
	// Calculate offset
	offset := (pageNum - 1) * pageSize
	// Conditional query based on expired
	condition := "customer_id = ?"
	if expired {
		condition += " AND expire_at < ?"
	} else {
		condition += " AND expire_at >= ?"
	}
	result := r.DB.Where(condition, customerId, expireAt).
		Order("id ASC").
		Limit(pageSize).
		Offset(offset).
		Find(&purchases)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchases, nil
}

func (r *purchaseRepository) FindByCustomerIDAndStartAt(customerId uint, startAt time.Time) ([]entities.Purchase, error) {
	purchases := []entities.Purchase{}
	// Conditional query based on expired
	condition := "customer_id = ? AND start_at = ?"
	result := r.DB.Where(condition, customerId, startAt).Find(&purchases)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchases, nil
}

func (r *purchaseRepository) FindByProductIDAndCustomerID(productId uint, customerId uint) ([]entities.Purchase, error) {
	purchases := []entities.Purchase{}
	// Start with the base query
	query := r.DB.Model(&entities.Purchase{})
	// Add condition only if productId is not 0
	if productId != 0 {
		query = query.Where("product_id = ?", productId)
	}
	// Add condition only if customerId is not 0
	if customerId != 0 {
		query = query.Where("customer_id = ?", customerId)
	}
	// Execute the query
	result := query.Find(&purchases)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchases, nil
}

func (r *purchaseRepository) FindByBandwidthAndStartAtAndExpireAt(bandwidth int, startAt time.Time, expireAt time.Time) ([]entities.Purchase, error) {
	purchases := []entities.Purchase{}
	condition := "bandwidth > ? AND start_at > ? AND expire_at > ?"
	result := r.DB.Where(condition, bandwidth, startAt, expireAt).Find(&purchases)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchases, nil
}

func (r *purchaseRepository) FindByDurationAndStartAtAndExpireAt(duration *int, startAt time.Time, expireAt time.Time) ([]entities.Purchase, error) {
	purchases := []entities.Purchase{}
	var condition string
	if duration == nil {
		condition = "duration IS NULL"
	} else {
		condition = fmt.Sprintf("duration = %d", *duration)
	}
	condition += " AND start_at > ? AND expire_at < ?"
	result := r.DB.Where(condition, startAt, expireAt).Find(&purchases)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchases, nil
}

func (r *purchaseRepository) FindByExpireAtWithRange(beginAt, endAt time.Time) ([]entities.Purchase, error) {
	purchases := []entities.Purchase{}
	result := r.DB.Model(&entities.Purchase{}).
		Where("expire_at >= ? AND expire_at <= ?", beginAt, endAt).
		Order("expire_at ASC").
		Find(&purchases)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchases, nil
}

func (r *purchaseRepository) FindByCustomerIDAndStartAtWithRange(customerIds []uint, beginAt, endAt time.Time) ([]entities.Purchase, error) {
	purchases := []entities.Purchase{}
	result := r.DB.Model(&entities.Purchase{}).
		Where("customer_id IN ? AND start_at >= ? AND start_at <= ?", customerIds, beginAt, endAt).
		Order("start_at ASC").
		Find(&purchases)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchases, nil
}

func (r *purchaseRepository) FindOneByID(id uint) (*entities.Purchase, error) {
	purchase := entities.Purchase{}
	result := r.DB.Where("id = ?", id).First(&purchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return &purchase, nil
}

func (r *purchaseRepository) Update(tx *gorm.DB, purchase *entities.Purchase) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}

	jsonData, err := json.Marshal(*purchase)
	if err != nil {
		return err
	}
	var purchaseMap map[string]interface{}
	err = json.Unmarshal(jsonData, &purchaseMap)
	if err != nil {
		return err
	}
	delete(purchaseMap, "id")
	delete(purchaseMap, "_enabled")
	delete(purchaseMap, "_removed")
	delete(purchaseMap, "created_at")
	delete(purchaseMap, "updated_at")

	result := dbInst.Clauses(clause.Returning{}).Model(purchase).Updates(purchaseMap)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *purchaseRepository) Delete(tx *gorm.DB, id uint) error {
	dbInst := r.DB
	if tx != nil {
		dbInst = tx
	}
	result := dbInst.Delete(&entities.Purchase{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
