package repository

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"gorm.io/gorm"
)

type StoreRepository interface {
	Create(store *domain.Store) error
	GetByID(id uuid.UUID) (*domain.Store, error)
	GetByCode(code string) (*domain.Store, error)
	Update(store *domain.Store) error
	Delete(id uuid.UUID) error
	List(page, pageSize int) ([]*domain.Store, int64, error)
}

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepository{db: db}
}

func (r *storeRepository) Create(store *domain.Store) error {
	return r.db.Create(store).Error
}

func (r *storeRepository) GetByID(id uuid.UUID) (*domain.Store, error) {
	var store domain.Store
	err := r.db.Preload("Manager").First(&store, "id = ? AND deleted_at IS NULL", id).Error
	return &store, err
}

func (r *storeRepository) GetByCode(code string) (*domain.Store, error) {
	var store domain.Store
	err := r.db.Where("code = ? AND deleted_at IS NULL", code).First(&store).Error
	return &store, err
}

func (r *storeRepository) Update(store *domain.Store) error {
	return r.db.Save(store).Error
}

func (r *storeRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Store{}, "id = ?", id).Error
}

func (r *storeRepository) List(page, pageSize int) ([]*domain.Store, int64, error) {
	var stores []*domain.Store
	var total int64

	query := r.db.Model(&domain.Store{}).Where("deleted_at IS NULL")

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("Manager").Offset(offset).Limit(pageSize).Find(&stores).Error

	return stores, total, err
}
