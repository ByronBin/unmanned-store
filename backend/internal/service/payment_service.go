package service

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/repository"
	"github.com/unmanned-store/backend/pkg/config"
)

type PaymentService interface {
	CreateWechatPayment(orderID uuid.UUID) (map[string]interface{}, error)
	CreateAlipayPayment(orderID uuid.UUID) (map[string]interface{}, error)
	HandleWechatCallback(data map[string]interface{}) error
	HandleAlipayCallback(data map[string]interface{}) error
}

type paymentService struct {
	paymentRepo repository.PaymentRepository
	orderRepo   repository.OrderRepository
	cfg         *config.Config
}

func NewPaymentService(paymentRepo repository.PaymentRepository, orderRepo repository.OrderRepository, cfg *config.Config) PaymentService {
	return &paymentService{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
		cfg:         cfg,
	}
}

func (s *paymentService) CreateWechatPayment(orderID uuid.UUID) (map[string]interface{}, error) {
	// TODO: 实现微信支付
	return make(map[string]interface{}), nil
}

func (s *paymentService) CreateAlipayPayment(orderID uuid.UUID) (map[string]interface{}, error) {
	// TODO: 实现支付宝支付
	return make(map[string]interface{}), nil
}

func (s *paymentService) HandleWechatCallback(data map[string]interface{}) error {
	// TODO: 处理微信支付回调
	return nil
}

func (s *paymentService) HandleAlipayCallback(data map[string]interface{}) error {
	// TODO: 处理支付宝回调
	return nil
}
