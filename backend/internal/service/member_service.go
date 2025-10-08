package service

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"github.com/unmanned-store/backend/internal/repository"
)

type MemberService interface {
	GetProfile(userID uuid.UUID) (*domain.User, error)
	UpdateProfile(user *domain.User) error
	GetCoupons(userID uuid.UUID) ([]*domain.UserCoupon, error)
}

type memberService struct {
	userRepo   repository.UserRepository
	couponRepo repository.CouponRepository
}

func NewMemberService(userRepo repository.UserRepository, couponRepo repository.CouponRepository) MemberService {
	return &memberService{
		userRepo:   userRepo,
		couponRepo: couponRepo,
	}
}

func (s *memberService) GetProfile(userID uuid.UUID) (*domain.User, error) {
	return s.userRepo.GetByID(userID)
}

func (s *memberService) UpdateProfile(user *domain.User) error {
	return s.userRepo.Update(user)
}

func (s *memberService) GetCoupons(userID uuid.UUID) ([]*domain.UserCoupon, error) {
	return s.couponRepo.GetUserCoupons(userID, "")
}
