package service

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/repository"
)

type AnalyticsService interface {
	GetSalesStats(storeID *uuid.UUID, period string) (map[string]interface{}, error)
	GetHotProducts(storeID *uuid.UUID, limit int) ([]map[string]interface{}, error)
	GetCustomerStats(storeID *uuid.UUID) (map[string]interface{}, error)
}

type analyticsService struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
}

func NewAnalyticsService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository) AnalyticsService {
	return &analyticsService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *analyticsService) GetSalesStats(storeID *uuid.UUID, period string) (map[string]interface{}, error) {
	// TODO: 实现销售统计
	return make(map[string]interface{}), nil
}

func (s *analyticsService) GetHotProducts(storeID *uuid.UUID, limit int) ([]map[string]interface{}, error) {
	// TODO: 实现热销商品
	return []map[string]interface{}{}, nil
}

func (s *analyticsService) GetCustomerStats(storeID *uuid.UUID) (map[string]interface{}, error) {
	// TODO: 实现顾客统计
	return make(map[string]interface{}), nil
}
