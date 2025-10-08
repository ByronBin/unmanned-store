package repository

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"gorm.io/gorm"
)

type MonitoringRepository interface {
	CreateDevice(device *domain.MonitoringDevice) error
	GetDevice(id uuid.UUID) (*domain.MonitoringDevice, error)
	GetDevicesByStore(storeID uuid.UUID) ([]*domain.MonitoringDevice, error)
	UpdateDevice(device *domain.MonitoringDevice) error
	CreateAlert(alert *domain.MonitoringAlert) error
	GetAlerts(storeID *uuid.UUID, status string, page, pageSize int) ([]*domain.MonitoringAlert, int64, error)
	UpdateAlert(alert *domain.MonitoringAlert) error
}

type monitoringRepository struct {
	db *gorm.DB
}

func NewMonitoringRepository(db *gorm.DB) MonitoringRepository {
	return &monitoringRepository{db: db}
}

func (r *monitoringRepository) CreateDevice(device *domain.MonitoringDevice) error {
	return r.db.Create(device).Error
}

func (r *monitoringRepository) GetDevice(id uuid.UUID) (*domain.MonitoringDevice, error) {
	var device domain.MonitoringDevice
	err := r.db.Preload("Store").First(&device, "id = ? AND deleted_at IS NULL", id).Error
	return &device, err
}

func (r *monitoringRepository) GetDevicesByStore(storeID uuid.UUID) ([]*domain.MonitoringDevice, error) {
	var devices []*domain.MonitoringDevice
	err := r.db.Where("store_id = ? AND deleted_at IS NULL", storeID).Find(&devices).Error
	return devices, err
}

func (r *monitoringRepository) UpdateDevice(device *domain.MonitoringDevice) error {
	return r.db.Save(device).Error
}

func (r *monitoringRepository) CreateAlert(alert *domain.MonitoringAlert) error {
	return r.db.Create(alert).Error
}

func (r *monitoringRepository) GetAlerts(storeID *uuid.UUID, status string, page, pageSize int) ([]*domain.MonitoringAlert, int64, error) {
	var alerts []*domain.MonitoringAlert
	var total int64

	query := r.db.Model(&domain.MonitoringAlert{})

	if storeID != nil {
		query = query.Where("store_id = ?", *storeID)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("Store").Preload("Device").Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&alerts).Error

	return alerts, total, err
}

func (r *monitoringRepository) UpdateAlert(alert *domain.MonitoringAlert) error {
	return r.db.Save(alert).Error
}
