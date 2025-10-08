package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"github.com/unmanned-store/backend/internal/repository"
	"github.com/unmanned-store/backend/pkg/config"
)

type AccessService interface {
	OpenDoor(storeID, userID uuid.UUID) error
	GetLogs(storeID *uuid.UUID, page, pageSize int) ([]*domain.AccessLog, int64, error)
	AddToBlacklist(userID, operatorID uuid.UUID, reason string) error
}

type accessService struct {
	accessRepo repository.AccessRepository
	userRepo   repository.UserRepository
	cfg        *config.Config
}

func NewAccessService(accessRepo repository.AccessRepository, userRepo repository.UserRepository, cfg *config.Config) AccessService {
	return &accessService{
		accessRepo: accessRepo,
		userRepo:   userRepo,
		cfg:        cfg,
	}
}

func (s *accessService) OpenDoor(storeID, userID uuid.UUID) error {
	// 检查黑名单
	isBlacklisted, err := s.accessRepo.IsInBlacklist(userID)
	if err != nil {
		return err
	}
	if isBlacklisted {
		return errors.New("用户已被加入黑名单")
	}

	// TODO: 调用硬件接口开门

	// 记录日志
	log := &domain.AccessLog{
		StoreID: storeID,
		UserID:  userID,
		Action:  "enter",
		Method:  "qrcode",
		Status:  "success",
	}
	return s.accessRepo.CreateLog(log)
}

func (s *accessService) GetLogs(storeID *uuid.UUID, page, pageSize int) ([]*domain.AccessLog, int64, error) {
	return s.accessRepo.GetLogs(storeID, page, pageSize)
}

func (s *accessService) AddToBlacklist(userID, operatorID uuid.UUID, reason string) error {
	blacklist := &domain.Blacklist{
		UserID:   userID,
		Reason:   reason,
		Operator: operatorID,
		Status:   "active",
	}
	return s.accessRepo.CreateBlacklist(blacklist)
}
