package service

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/repository"
)

type FinanceService interface {
	GetDailyReport(storeID uuid.UUID, date string) (map[string]interface{}, error)
	GetMonthlyReport(storeID uuid.UUID, month string) (map[string]interface{}, error)
	GetSummary(storeID *uuid.UUID) (map[string]interface{}, error)
}

type financeService struct {
	orderRepo repository.OrderRepository
	storeRepo repository.StoreRepository
}

func NewFinanceService(orderRepo repository.OrderRepository, storeRepo repository.StoreRepository) FinanceService {
	return &financeService{
		orderRepo: orderRepo,
		storeRepo: storeRepo,
	}
}

func (s *financeService) GetDailyReport(storeID uuid.UUID, date string) (map[string]interface{}, error) {
	// TODO: 实现日报表
	return make(map[string]interface{}), nil
}

func (s *financeService) GetMonthlyReport(storeID uuid.UUID, month string) (map[string]interface{}, error) {
	// TODO: 实现月报表
	return make(map[string]interface{}), nil
}

func (s *financeService) GetSummary(storeID *uuid.UUID) (map[string]interface{}, error) {
	// TODO: 实现汇总报表
	return make(map[string]interface{}), nil
}
