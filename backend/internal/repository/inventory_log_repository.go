package repository

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"gorm.io/gorm"
)

type InventoryLogRepository interface {
	Create(log *domain.InventoryLog) error
	GetLogs(storeID, skuID *uuid.UUID, page, pageSize int) ([]*domain.InventoryLog, int64, error)
}

type inventoryLogRepository struct {
	db *gorm.DB
}

func NewInventoryLogRepository(db *gorm.DB) InventoryLogRepository {
	return &inventoryLogRepository{db: db}
}

func (r *inventoryLogRepository) Create(log *domain.InventoryLog) error {
	return r.db.Create(log).Error
}

func (r *inventoryLogRepository) GetLogs(storeID, skuID *uuid.UUID, page, pageSize int) ([]*domain.InventoryLog, int64, error) {
	var logs []*domain.InventoryLog
	var total int64

	query := r.db.Model(&domain.InventoryLog{})

	if storeID != nil {
		query = query.Where("store_id = ?", *storeID)
	}
	
	if skuID != nil {
		query = query.Where("sku_id = ?", *skuID)
	}

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error

	return logs, total, err
}
