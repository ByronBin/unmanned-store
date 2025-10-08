package service

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"github.com/unmanned-store/backend/internal/repository"
	"github.com/unmanned-store/backend/pkg/config"
)

type MonitoringService interface {
	GetDevices(storeID uuid.UUID) ([]*domain.MonitoringDevice, error)
	GetStreams(storeID uuid.UUID) ([]map[string]interface{}, error)
	GetAlerts(storeID *uuid.UUID, page, pageSize int) ([]*domain.MonitoringAlert, int64, error)
}

type monitoringService struct {
	monitoringRepo repository.MonitoringRepository
	cfg            *config.Config
}

func NewMonitoringService(monitoringRepo repository.MonitoringRepository, cfg *config.Config) MonitoringService {
	return &monitoringService{
		monitoringRepo: monitoringRepo,
		cfg:            cfg,
	}
}

func (s *monitoringService) GetDevices(storeID uuid.UUID) ([]*domain.MonitoringDevice, error) {
	return s.monitoringRepo.GetDevicesByStore(storeID)
}

func (s *monitoringService) GetStreams(storeID uuid.UUID) ([]map[string]interface{}, error) {
	// TODO: 实现视频流获取
	return []map[string]interface{}{}, nil
}

func (s *monitoringService) GetAlerts(storeID *uuid.UUID, page, pageSize int) ([]*domain.MonitoringAlert, int64, error) {
	return s.monitoringRepo.GetAlerts(storeID, "", page, pageSize)
}
