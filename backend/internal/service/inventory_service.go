package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"github.com/unmanned-store/backend/internal/repository"
)

type InventoryService interface {
	// 库存查询
	GetByStore(storeID uuid.UUID, page, pageSize int) ([]*domain.Inventory, int64, error)
	GetBySKU(skuID uuid.UUID, storeID *uuid.UUID) ([]*domain.Inventory, error)
	GetByProduct(productID uuid.UUID, storeID *uuid.UUID) ([]*domain.Inventory, error)
	
	// 库存调整
	AdjustInventory(storeID, skuID uuid.UUID, quantity int, reason string, operatorID uuid.UUID) error
	StockIn(storeID, skuID uuid.UUID, quantity int, operatorID uuid.UUID) error
	StockOut(storeID, skuID uuid.UUID, quantity int, reason string, operatorID uuid.UUID) error
	
	// 库存预警
	GetLowStockItems(storeID *uuid.UUID, threshold int) ([]*domain.Inventory, error)
	
	// 库存盘点
	CreateInventoryCount(count *domain.InventoryCount) error
	GetInventoryCounts(storeID uuid.UUID, status string) ([]*domain.InventoryCount, error)
	SubmitInventoryCount(countID uuid.UUID, items []domain.InventoryCountItem, operatorID uuid.UUID) error
	
	// 库存日志
	GetInventoryLogs(storeID, skuID *uuid.UUID, page, pageSize int) ([]*domain.InventoryLog, int64, error)
}

type inventoryService struct {
	inventoryRepo repository.InventoryRepository
	inventoryLogRepo repository.InventoryLogRepository
}

func NewInventoryService(inventoryRepo repository.InventoryRepository, inventoryLogRepo repository.InventoryLogRepository) InventoryService {
	return &inventoryService{
		inventoryRepo: inventoryRepo,
		inventoryLogRepo: inventoryLogRepo,
	}
}

func (s *inventoryService) GetByStore(storeID uuid.UUID, page, pageSize int) ([]*domain.Inventory, int64, error) {
	return s.inventoryRepo.GetByStore(storeID, page, pageSize)
}

func (s *inventoryService) GetBySKU(skuID uuid.UUID, storeID *uuid.UUID) ([]*domain.Inventory, error) {
	return s.inventoryRepo.GetBySKU(skuID, storeID)
}

func (s *inventoryService) GetByProduct(productID uuid.UUID, storeID *uuid.UUID) ([]*domain.Inventory, error) {
	return s.inventoryRepo.GetByProduct(productID, storeID)
}

func (s *inventoryService) AdjustInventory(storeID, skuID uuid.UUID, quantity int, reason string, operatorID uuid.UUID) error {
	// 获取当前库存
	inventory, err := s.inventoryRepo.GetByStoreAndSKU(storeID, skuID)
	if err != nil {
		return err
	}
	
	oldQuantity := inventory.Quantity
	newQuantity := oldQuantity + quantity
	
	// 更新库存
	inventory.Quantity = newQuantity
	if err := s.inventoryRepo.Update(inventory); err != nil {
		return err
	}
	
	// 记录库存日志
	log := &domain.InventoryLog{
		StoreID: storeID,
		SKUID: skuID,
		Type: "adjust",
		Quantity: quantity,
		OldQuantity: oldQuantity,
		NewQuantity: newQuantity,
		Reason: reason,
		OperatorID: operatorID,
	}
	
	return s.inventoryLogRepo.Create(log)
}

func (s *inventoryService) StockIn(storeID, skuID uuid.UUID, quantity int, operatorID uuid.UUID) error {
	return s.AdjustInventory(storeID, skuID, quantity, "入库", operatorID)
}

func (s *inventoryService) StockOut(storeID, skuID uuid.UUID, quantity int, reason string, operatorID uuid.UUID) error {
	return s.AdjustInventory(storeID, skuID, -quantity, reason, operatorID)
}

func (s *inventoryService) GetLowStockItems(storeID *uuid.UUID, threshold int) ([]*domain.Inventory, error) {
	return s.inventoryRepo.GetLowStockItems(storeID, threshold)
}

func (s *inventoryService) CreateInventoryCount(count *domain.InventoryCount) error {
	return s.inventoryRepo.CreateInventoryCount(count)
}

func (s *inventoryService) GetInventoryCounts(storeID uuid.UUID, status string) ([]*domain.InventoryCount, error) {
	return s.inventoryRepo.GetInventoryCounts(storeID, status)
}

func (s *inventoryService) SubmitInventoryCount(countID uuid.UUID, items []domain.InventoryCountItem, operatorID uuid.UUID) error {
	// 更新盘点状态
	count, err := s.inventoryRepo.GetInventoryCountByID(countID)
	if err != nil {
		return err
	}
	
	count.Status = "completed"
	count.CompletedAt = &time.Time{}
	*count.CompletedAt = time.Now()
	count.OperatorID = operatorID
	
	if err := s.inventoryRepo.UpdateInventoryCount(count); err != nil {
		return err
	}
	
	// 创建盘点项目
	for _, item := range items {
		item.CountID = countID
		if err := s.inventoryRepo.CreateInventoryCountItem(&item); err != nil {
			return err
		}
	}
	
	return nil
}

func (s *inventoryService) GetInventoryLogs(storeID, skuID *uuid.UUID, page, pageSize int) ([]*domain.InventoryLog, int64, error) {
	return s.inventoryLogRepo.GetLogs(storeID, skuID, page, pageSize)
}
