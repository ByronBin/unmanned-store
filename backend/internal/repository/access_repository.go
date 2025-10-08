package repository

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"gorm.io/gorm"
)

type AccessRepository interface {
	CreateLog(log *domain.AccessLog) error
	GetLogs(storeID *uuid.UUID, page, pageSize int) ([]*domain.AccessLog, int64, error)
	CreateBlacklist(blacklist *domain.Blacklist) error
	IsInBlacklist(userID uuid.UUID) (bool, error)
	GetBlacklist(page, pageSize int) ([]*domain.Blacklist, int64, error)
}

type accessRepository struct {
	db *gorm.DB
}

func NewAccessRepository(db *gorm.DB) AccessRepository {
	return &accessRepository{db: db}
}

func (r *accessRepository) CreateLog(log *domain.AccessLog) error {
	return r.db.Create(log).Error
}

func (r *accessRepository) GetLogs(storeID *uuid.UUID, page, pageSize int) ([]*domain.AccessLog, int64, error) {
	var logs []*domain.AccessLog
	var total int64

	query := r.db.Model(&domain.AccessLog{})

	if storeID != nil {
		query = query.Where("store_id = ?", *storeID)
	}

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("Store").Preload("User").Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error

	return logs, total, err
}

func (r *accessRepository) CreateBlacklist(blacklist *domain.Blacklist) error {
	return r.db.Create(blacklist).Error
}

func (r *accessRepository) IsInBlacklist(userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&domain.Blacklist{}).Where("user_id = ? AND status = 'active' AND deleted_at IS NULL", userID).Count(&count).Error
	return count > 0, err
}

func (r *accessRepository) GetBlacklist(page, pageSize int) ([]*domain.Blacklist, int64, error) {
	var blacklist []*domain.Blacklist
	var total int64

	query := r.db.Model(&domain.Blacklist{}).Where("deleted_at IS NULL")

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("User").Offset(offset).Limit(pageSize).Find(&blacklist).Error

	return blacklist, total, err
}
