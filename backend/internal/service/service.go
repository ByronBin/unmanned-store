package service

import (
	"github.com/redis/go-redis/v9"
	"github.com/unmanned-store/backend/internal/repository"
	"github.com/unmanned-store/backend/pkg/config"
)

type Services struct {
	Auth       AuthService
	Store      StoreService
	Product    ProductService
	Category   CategoryService
	Inventory  InventoryService
	Order      OrderService
	Payment    PaymentService
	Member     MemberService
	Access     AccessService
	Monitoring MonitoringService
	Finance    FinanceService
	Analytics  AnalyticsService
}

func NewServices(repos *repository.Repositories, rdb *redis.Client, cfg *config.Config) *Services {
	return &Services{
		Auth:       NewAuthService(repos.User, rdb, cfg),
		Store:      NewStoreService(repos.Store),
		Product:    NewProductService(repos.Product, repos.Category),
		Category:   NewCategoryService(repos.Category),
		Inventory:  NewInventoryService(repos.Inventory, repos.InventoryLog),
		Order:      NewOrderService(repos.Order, repos.Product, repos.Inventory, repos.Coupon, rdb),
		Payment:    NewPaymentService(repos.Payment, repos.Order, cfg),
		Member:     NewMemberService(repos.User, repos.Coupon),
		Access:     NewAccessService(repos.Access, repos.User, cfg),
		Monitoring: NewMonitoringService(repos.Monitoring, cfg),
		Finance:    NewFinanceService(repos.Order, repos.Store),
		Analytics:  NewAnalyticsService(repos.Order, repos.Product),
	}
}
