package repository

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	GetByStoreAndSKU(storeID, skuID uuid.UUID) (*domain.Inventory, error)
	Create(inventory *domain.Inventory) error
	Update(inventory *domain.Inventory) error
	List(storeID uuid.UUID, page, pageSize int) ([]*domain.Inventory, int64, error)
	GetLowStock(storeID uuid.UUID) ([]*domain.Inventory, error)
	CreateLog(log *domain.InventoryLog) error
	GetLogs(storeID, skuID uuid.UUID, page, pageSize int) ([]*domain.InventoryLog, int64, error)
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) GetByStoreAndSKU(storeID, skuID uuid.UUID) (*domain.Inventory, error) {
	var inventory domain.Inventory
	err := r.db.Preload("Store").Preload("SKU").Where("store_id = ? AND sku_id = ? AND deleted_at IS NULL", storeID, skuID).First(&inventory).Error
	return &inventory, err
}

func (r *inventoryRepository) Create(inventory *domain.Inventory) error {
	return r.db.Create(inventory).Error
}

func (r *inventoryRepository) Update(inventory *domain.Inventory) error {
	return r.db.Save(inventory).Error
}

func (r *inventoryRepository) List(storeID uuid.UUID, page, pageSize int) ([]*domain.Inventory, int64, error) {
	var inventories []*domain.Inventory
	var total int64

	query := r.db.Model(&domain.Inventory{}).Where("store_id = ? AND deleted_at IS NULL", storeID)

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("SKU").Preload("SKU.Product").Offset(offset).Limit(pageSize).Find(&inventories).Error

	return inventories, total, err
}

func (r *inventoryRepository) GetLowStock(storeID uuid.UUID) ([]*domain.Inventory, error) {
	var inventories []*domain.Inventory
	err := r.db.Preload("SKU").Preload("SKU.Product").Where("store_id = ? AND quantity <= alert_qty AND deleted_at IS NULL", storeID).Find(&inventories).Error
	return inventories, err
}

func (r *inventoryRepository) CreateLog(log *domain.InventoryLog) error {
	return r.db.Create(log).Error
}

func (r *inventoryRepository) GetLogs(storeID, skuID uuid.UUID, page, pageSize int) ([]*domain.InventoryLog, int64, error) {
	var logs []*domain.InventoryLog
	var total int64

	query := r.db.Model(&domain.InventoryLog{}).Where("store_id = ? AND sku_id = ?", storeID, skuID)

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error

	return logs, total, err
}
