package handler

import (
	"github.com/unmanned-store/backend/internal/service"
)

type Handlers struct {
	Auth       *AuthHandler
	Store      *StoreHandler
	Product    *ProductHandler
	Category   *CategoryHandler
	Inventory  *InventoryHandler
	Order      *OrderHandler
	Payment    *PaymentHandler
	Member     *MemberHandler
	Access     *AccessHandler
	Monitoring *MonitoringHandler
	Finance    *FinanceHandler
	Analytics  *AnalyticsHandler
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		Auth:       NewAuthHandler(services.Auth),
		Store:      NewStoreHandler(services.Store),
		Product:    NewProductHandler(services.Product),
		Category:   NewCategoryHandler(services.Category),
		Inventory:  NewInventoryHandler(services.Inventory),
		Order:      NewOrderHandler(services.Order),
		Payment:    NewPaymentHandler(services.Payment),
		Member:     NewMemberHandler(services.Member),
		Access:     NewAccessHandler(services.Access),
		Monitoring: NewMonitoringHandler(services.Monitoring),
		Finance:    NewFinanceHandler(services.Finance),
		Analytics:  NewAnalyticsHandler(services.Analytics),
	}
}
