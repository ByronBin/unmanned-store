package repository

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	GetByStoreAndSKU(storeID, skuID uuid.UUID) (*domain.Inventory, error)
	GetByStore(storeID uuid.UUID, page, pageSize int) ([]*domain.Inventory, int64, error)
	GetBySKU(skuID uuid.UUID, storeID *uuid.UUID) ([]*domain.Inventory, error)
	GetByProduct(productID uuid.UUID, storeID *uuid.UUID) ([]*domain.Inventory, error)
	Create(inventory *domain.Inventory) error
	Update(inventory *domain.Inventory) error
	GetLowStockItems(storeID *uuid.UUID, threshold int) ([]*domain.Inventory, error)
	
	// 库存盘点
	CreateInventoryCount(count *domain.InventoryCount) error
	GetInventoryCounts(storeID uuid.UUID, status string) ([]*domain.InventoryCount, error)
	GetInventoryCountByID(id uuid.UUID) (*domain.InventoryCount, error)
	UpdateInventoryCount(count *domain.InventoryCount) error
	CreateInventoryCountItem(item *domain.InventoryCountItem) error
	
	// 库存日志
	CreateLog(log *domain.InventoryLog) error
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

func (r *inventoryRepository) GetByStore(storeID uuid.UUID, page, pageSize int) ([]*domain.Inventory, int64, error) {
	var inventories []*domain.Inventory
	var total int64

	query := r.db.Model(&domain.Inventory{}).Where("store_id = ? AND deleted_at IS NULL", storeID)

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("SKU").Preload("SKU.Product").Offset(offset).Limit(pageSize).Find(&inventories).Error

	return inventories, total, err
}

func (r *inventoryRepository) GetBySKU(skuID uuid.UUID, storeID *uuid.UUID) ([]*domain.Inventory, error) {
	var inventories []*domain.Inventory
	query := r.db.Where("sku_id = ? AND deleted_at IS NULL", skuID)
	
	if storeID != nil {
		query = query.Where("store_id = ?", *storeID)
	}
	
	err := query.Preload("Store").Preload("SKU").Find(&inventories).Error
	return inventories, err
}

func (r *inventoryRepository) GetByProduct(productID uuid.UUID, storeID *uuid.UUID) ([]*domain.Inventory, error) {
	var inventories []*domain.Inventory
	query := r.db.Joins("JOIN product_skus ON inventories.sku_id = product_skus.id").
		Where("product_skus.product_id = ? AND inventories.deleted_at IS NULL", productID)
	
	if storeID != nil {
		query = query.Where("inventories.store_id = ?", *storeID)
	}
	
	err := query.Preload("Store").Preload("SKU").Find(&inventories).Error
	return inventories, err
}

func (r *inventoryRepository) GetLowStockItems(storeID *uuid.UUID, threshold int) ([]*domain.Inventory, error) {
	var inventories []*domain.Inventory
	query := r.db.Where("quantity <= ? AND deleted_at IS NULL", threshold)
	
	if storeID != nil {
		query = query.Where("store_id = ?", *storeID)
	}
	
	err := query.Preload("Store").Preload("SKU").Preload("SKU.Product").Find(&inventories).Error
	return inventories, err
}

// 库存盘点方法
func (r *inventoryRepository) CreateInventoryCount(count *domain.InventoryCount) error {
	return r.db.Create(count).Error
}

func (r *inventoryRepository) GetInventoryCounts(storeID uuid.UUID, status string) ([]*domain.InventoryCount, error) {
	var counts []*domain.InventoryCount
	query := r.db.Where("store_id = ? AND deleted_at IS NULL", storeID)
	
	if status != "" {
		query = query.Where("status = ?", status)
	}
	
	err := query.Order("created_at DESC").Find(&counts).Error
	return counts, err
}

func (r *inventoryRepository) GetInventoryCountByID(id uuid.UUID) (*domain.InventoryCount, error) {
	var count domain.InventoryCount
	err := r.db.First(&count, "id = ? AND deleted_at IS NULL", id).Error
	return &count, err
}

func (r *inventoryRepository) UpdateInventoryCount(count *domain.InventoryCount) error {
	return r.db.Save(count).Error
}

func (r *inventoryRepository) CreateInventoryCountItem(item *domain.InventoryCountItem) error {
	return r.db.Create(item).Error
}

func (r *inventoryRepository) CreateLog(log *domain.InventoryLog) error {
	return r.db.Create(log).Error
}
