package service

import (
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/unmanned-store/backend/internal/domain"
	"github.com/unmanned-store/backend/internal/repository"
)

type OrderService interface {
	Create(order *domain.Order) error
	GetByID(id uuid.UUID) (*domain.Order, error)
	List(userID *uuid.UUID, storeID *uuid.UUID, page, pageSize int) ([]*domain.Order, int64, error)
	Cancel(orderID, userID uuid.UUID) error
	Refund(orderID, operatorID uuid.UUID) error
}

type orderService struct {
	orderRepo     repository.OrderRepository
	productRepo   repository.ProductRepository
	inventoryRepo repository.InventoryRepository
	couponRepo    repository.CouponRepository
	rdb           *redis.Client
}

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository, inventoryRepo repository.InventoryRepository, couponRepo repository.CouponRepository, rdb *redis.Client) OrderService {
	return &orderService{
		orderRepo:     orderRepo,
		productRepo:   productRepo,
		inventoryRepo: inventoryRepo,
		couponRepo:    couponRepo,
		rdb:           rdb,
	}
}

func (s *orderService) Create(order *domain.Order) error {
	// TODO: 实现订单创建逻辑（包括库存检查、扣减等）
	return s.orderRepo.Create(order)
}

func (s *orderService) GetByID(id uuid.UUID) (*domain.Order, error) {
	return s.orderRepo.GetByID(id)
}

func (s *orderService) List(userID *uuid.UUID, storeID *uuid.UUID, page, pageSize int) ([]*domain.Order, int64, error) {
	return s.orderRepo.List(userID, storeID, page, pageSize)
}

func (s *orderService) Cancel(orderID, userID uuid.UUID) error {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return err
	}
	order.Status = "cancelled"
	return s.orderRepo.Update(order)
}

func (s *orderService) Refund(orderID, operatorID uuid.UUID) error {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return err
	}
	order.Status = "refunded"
	return s.orderRepo.Update(order)
}
