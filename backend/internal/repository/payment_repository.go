package repository

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *domain.Payment) error
	GetByID(id uuid.UUID) (*domain.Payment, error)
	GetByOrderID(orderID uuid.UUID) (*domain.Payment, error)
	GetByPaymentNo(paymentNo string) (*domain.Payment, error)
	Update(payment *domain.Payment) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(payment *domain.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) GetByID(id uuid.UUID) (*domain.Payment, error) {
	var payment domain.Payment
	err := r.db.Preload("Order").First(&payment, "id = ? AND deleted_at IS NULL", id).Error
	return &payment, err
}

func (r *paymentRepository) GetByOrderID(orderID uuid.UUID) (*domain.Payment, error) {
	var payment domain.Payment
	err := r.db.Preload("Order").Where("order_id = ? AND deleted_at IS NULL", orderID).First(&payment).Error
	return &payment, err
}

func (r *paymentRepository) GetByPaymentNo(paymentNo string) (*domain.Payment, error) {
	var payment domain.Payment
	err := r.db.Preload("Order").Where("payment_no = ? AND deleted_at IS NULL", paymentNo).First(&payment).Error
	return &payment, err
}

func (r *paymentRepository) Update(payment *domain.Payment) error {
	return r.db.Save(payment).Error
}
