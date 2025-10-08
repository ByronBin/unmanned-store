package repository

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *domain.Order) error
	GetByID(id uuid.UUID) (*domain.Order, error)
	GetByOrderNo(orderNo string) (*domain.Order, error)
	Update(order *domain.Order) error
	List(userID *uuid.UUID, storeID *uuid.UUID, page, pageSize int) ([]*domain.Order, int64, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetByID(id uuid.UUID) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Store").Preload("User").Preload("Items").Preload("Items.SKU").Preload("Items.SKU.Product").First(&order, "id = ? AND deleted_at IS NULL", id).Error
	return &order, err
}

func (r *orderRepository) GetByOrderNo(orderNo string) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Store").Preload("User").Preload("Items").Where("order_no = ? AND deleted_at IS NULL", orderNo).First(&order).Error
	return &order, err
}

func (r *orderRepository) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) List(userID *uuid.UUID, storeID *uuid.UUID, page, pageSize int) ([]*domain.Order, int64, error) {
	var orders []*domain.Order
	var total int64

	query := r.db.Model(&domain.Order{}).Where("deleted_at IS NULL")

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	if storeID != nil {
		query = query.Where("store_id = ?", *storeID)
	}

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("Store").Preload("User").Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&orders).Error

	return orders, total, err
}
