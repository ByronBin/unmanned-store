package service

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"github.com/unmanned-store/backend/internal/repository"
)

type StoreService interface {
	Create(store *domain.Store) error
	GetByID(id uuid.UUID) (*domain.Store, error)
	Update(store *domain.Store) error
	Delete(id uuid.UUID) error
	List(page, pageSize int) ([]*domain.Store, int64, error)
}

type storeService struct {
	storeRepo repository.StoreRepository
}

func NewStoreService(storeRepo repository.StoreRepository) StoreService {
	return &storeService{storeRepo: storeRepo}
}

func (s *storeService) Create(store *domain.Store) error {
	return s.storeRepo.Create(store)
}

func (s *storeService) GetByID(id uuid.UUID) (*domain.Store, error) {
	return s.storeRepo.GetByID(id)
}

func (s *storeService) Update(store *domain.Store) error {
	return s.storeRepo.Update(store)
}

func (s *storeService) Delete(id uuid.UUID) error {
	return s.storeRepo.Delete(id)
}

func (s *storeService) List(page, pageSize int) ([]*domain.Store, int64, error) {
	return s.storeRepo.List(page, pageSize)
}
