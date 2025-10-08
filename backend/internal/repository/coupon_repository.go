package repository

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"gorm.io/gorm"
)

type CouponRepository interface {
	Create(coupon *domain.Coupon) error
	GetByID(id uuid.UUID) (*domain.Coupon, error)
	List(page, pageSize int) ([]*domain.Coupon, int64, error)
	CreateUserCoupon(userCoupon *domain.UserCoupon) error
	GetUserCoupons(userID uuid.UUID, status string) ([]*domain.UserCoupon, error)
}

type couponRepository struct {
	db *gorm.DB
}

func NewCouponRepository(db *gorm.DB) CouponRepository {
	return &couponRepository{db: db}
}

func (r *couponRepository) Create(coupon *domain.Coupon) error {
	return r.db.Create(coupon).Error
}

func (r *couponRepository) GetByID(id uuid.UUID) (*domain.Coupon, error) {
	var coupon domain.Coupon
	err := r.db.First(&coupon, "id = ? AND deleted_at IS NULL", id).Error
	return &coupon, err
}

func (r *couponRepository) List(page, pageSize int) ([]*domain.Coupon, int64, error) {
	var coupons []*domain.Coupon
	var total int64

	query := r.db.Model(&domain.Coupon{}).Where("deleted_at IS NULL")

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&coupons).Error

	return coupons, total, err
}

func (r *couponRepository) CreateUserCoupon(userCoupon *domain.UserCoupon) error {
	return r.db.Create(userCoupon).Error
}

func (r *couponRepository) GetUserCoupons(userID uuid.UUID, status string) ([]*domain.UserCoupon, error) {
	var userCoupons []*domain.UserCoupon
	query := r.db.Preload("Coupon").Where("user_id = ? AND deleted_at IS NULL", userID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Find(&userCoupons).Error
	return userCoupons, err
}
