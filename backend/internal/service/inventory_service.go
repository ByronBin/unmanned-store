package service

import (
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/unmanned-store/backend/internal/domain"
	"github.com/unmanned-store/backend/internal/repository"
)

type InventoryService interface {
	StockIn(storeID, skuID uuid.UUID, quantity int, operator uuid.UUID, remark string) error
	StockOut(storeID, skuID uuid.UUID, quantity int, operator uuid.UUID, remark string) error
	Transfer(fromStoreID, toStoreID, skuID uuid.UUID, quantity int, operator uuid.UUID) error
	GetInventory(storeID, skuID uuid.UUID) (*domain.Inventory, error)
	List(storeID uuid.UUID, page, pageSize int) ([]*domain.Inventory, int64, error)
	GetAlerts(storeID uuid.UUID) ([]*domain.Inventory, error)
}

type inventoryService struct {
	inventoryRepo repository.InventoryRepository
	productRepo   repository.ProductRepository
	rdb           *redis.Client
}

func NewInventoryService(inventoryRepo repository.InventoryRepository, productRepo repository.ProductRepository, rdb *redis.Client) InventoryService {
	return &inventoryService{
		inventoryRepo: inventoryRepo,
		productRepo:   productRepo,
		rdb:           rdb,
	}
}

func (s *inventoryService) StockIn(storeID, skuID uuid.UUID, quantity int, operator uuid.UUID, remark string) error {
	inventory, err := s.inventoryRepo.GetByStoreAndSKU(storeID, skuID)
	if err != nil {
		inventory = &domain.Inventory{
			StoreID:  storeID,
			SKUID:    skuID,
			Quantity: 0,
			AlertQty: 10,
		}
		if err := s.inventoryRepo.Create(inventory); err != nil {
			return err
		}
	}

	beforeQty := inventory.Quantity
	inventory.Quantity += quantity

	log := &domain.InventoryLog{
		StoreID:   storeID,
		SKUID:     skuID,
		Type:      "in",
		Quantity:  quantity,
		BeforeQty: beforeQty,
		AfterQty:  inventory.Quantity,
		Operator:  operator,
		Remark:    remark,
	}

	if err := s.inventoryRepo.Update(inventory); err != nil {
		return err
	}

	return s.inventoryRepo.CreateLog(log)
}

func (s *inventoryService) StockOut(storeID, skuID uuid.UUID, quantity int, operator uuid.UUID, remark string) error {
	inventory, err := s.inventoryRepo.GetByStoreAndSKU(storeID, skuID)
	if err != nil {
		return err
	}

	beforeQty := inventory.Quantity
	inventory.Quantity -= quantity

	log := &domain.InventoryLog{
		StoreID:   storeID,
		SKUID:     skuID,
		Type:      "out",
		Quantity:  quantity,
		BeforeQty: beforeQty,
		AfterQty:  inventory.Quantity,
		Operator:  operator,
		Remark:    remark,
	}

	if err := s.inventoryRepo.Update(inventory); err != nil {
		return err
	}

	return s.inventoryRepo.CreateLog(log)
}

func (s *inventoryService) Transfer(fromStoreID, toStoreID, skuID uuid.UUID, quantity int, operator uuid.UUID) error {
	// 实现库存调拨逻辑
	if err := s.StockOut(fromStoreID, skuID, quantity, operator, "调拨出库"); err != nil {
		return err
	}
	return s.StockIn(toStoreID, skuID, quantity, operator, "调拨入库")
}

func (s *inventoryService) GetInventory(storeID, skuID uuid.UUID) (*domain.Inventory, error) {
	return s.inventoryRepo.GetByStoreAndSKU(storeID, skuID)
}

func (s *inventoryService) List(storeID uuid.UUID, page, pageSize int) ([]*domain.Inventory, int64, error) {
	return s.inventoryRepo.List(storeID, page, pageSize)
}

func (s *inventoryService) GetAlerts(storeID uuid.UUID) ([]*domain.Inventory, error) {
	return s.inventoryRepo.GetLowStock(storeID)
}
