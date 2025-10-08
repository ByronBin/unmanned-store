package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	User       UserRepository
	Store      StoreRepository
	Product    ProductRepository
	Category   CategoryRepository
	Inventory  InventoryRepository
	Order      OrderRepository
	Payment    PaymentRepository
	Coupon     CouponRepository
	Access     AccessRepository
	Monitoring MonitoringRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:       NewUserRepository(db),
		Store:      NewStoreRepository(db),
		Product:    NewProductRepository(db),
		Category:   NewCategoryRepository(db),
		Inventory:  NewInventoryRepository(db),
		Order:      NewOrderRepository(db),
		Payment:    NewPaymentRepository(db),
		Coupon:     NewCouponRepository(db),
		Access:     NewAccessRepository(db),
		Monitoring: NewMonitoringRepository(db),
	}
}
