package domain

import (
	"time"

	"github.com/google/uuid"
)

// BaseModel 基础模型
type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

// Store 门店模型
type Store struct {
	BaseModel
	Name      string     `gorm:"size:100;not null" json:"name"`
	Code      string     `gorm:"size:50;uniqueIndex;not null" json:"code"`
	Address   string     `gorm:"size:255" json:"address"`
	Phone     string     `gorm:"size:20" json:"phone"`
	Longitude float64    `json:"longitude"`
	Latitude  float64    `json:"latitude"`
	OpenTime  string     `gorm:"size:10" json:"open_time"`
	CloseTime string     `gorm:"size:10" json:"close_time"`
	Status    string     `gorm:"size:20;default:'active'" json:"status"` // active, inactive, maintenance
	ManagerID *uuid.UUID `gorm:"type:uuid" json:"manager_id"`
	Manager   *User      `gorm:"foreignKey:ManagerID" json:"manager,omitempty"`
}

// User 用户模型
type User struct {
	BaseModel
	Username     string     `gorm:"size:50;uniqueIndex;not null" json:"username"`
	Password     string     `gorm:"size:255;not null" json:"-"`
	Nickname     string     `gorm:"size:50" json:"nickname"`
	Phone        string     `gorm:"size:20;uniqueIndex" json:"phone"`
	Email        string     `gorm:"size:100" json:"email"`
	Avatar       string     `gorm:"size:255" json:"avatar"`
	Role         string     `gorm:"size:20;not null;default:'customer'" json:"role"` // admin, store_manager, staff, customer
	StoreID      *uuid.UUID `gorm:"type:uuid" json:"store_id"`
	Store        *Store     `gorm:"foreignKey:StoreID" json:"store,omitempty"`
	Status       string     `gorm:"size:20;default:'active'" json:"status"` // active, inactive, banned
	IsVerified   bool       `gorm:"default:false" json:"is_verified"`
	WechatOpenID string     `gorm:"size:100;uniqueIndex" json:"wechat_openid,omitempty"`
	MemberLevel  int        `gorm:"default:0" json:"member_level"`
	Points       int        `gorm:"default:0" json:"points"`
}

// Category 分类模型
type Category struct {
	BaseModel
	Name     string      `gorm:"size:50;not null" json:"name"`
	ParentID *uuid.UUID  `gorm:"type:uuid" json:"parent_id"`
	Parent   *Category   `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children []*Category `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Sort     int         `gorm:"default:0" json:"sort"`
	Icon     string      `gorm:"size:255" json:"icon"`
	Status   string      `gorm:"size:20;default:'active'" json:"status"`
}

// Product 商品模型
type Product struct {
	BaseModel
	Name        string       `gorm:"size:100;not null" json:"name"`
	Code        string       `gorm:"size:50;uniqueIndex;not null" json:"code"`
	CategoryID  uuid.UUID    `gorm:"type:uuid;not null" json:"category_id"`
	Category    *Category    `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Description string       `gorm:"type:text" json:"description"`
	Images      []string     `gorm:"type:jsonb" json:"images"`
	Price       float64      `gorm:"type:decimal(10,2);not null" json:"price"`
	MemberPrice float64      `gorm:"type:decimal(10,2)" json:"member_price"`
	CostPrice   float64      `gorm:"type:decimal(10,2)" json:"cost_price"`
	Unit        string       `gorm:"size:20" json:"unit"` // 件、kg、瓶等
	Status      string       `gorm:"size:20;default:'active'" json:"status"`
	SKUs        []ProductSKU `gorm:"foreignKey:ProductID" json:"skus,omitempty"`
}

// ProductSKU 商品SKU模型
type ProductSKU struct {
	BaseModel
	ProductID  uuid.UUID         `gorm:"type:uuid;not null" json:"product_id"`
	Product    *Product          `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Name       string            `gorm:"size:100;not null" json:"name"`
	Code       string            `gorm:"size:50;uniqueIndex;not null" json:"code"`
	Barcode    string            `gorm:"size:50" json:"barcode"`
	Price      float64           `gorm:"type:decimal(10,2);not null" json:"price"`
	Attributes map[string]string `gorm:"type:jsonb" json:"attributes"` // 规格属性，如：{"颜色":"红色","尺寸":"L"}
	Status     string            `gorm:"size:20;default:'active'" json:"status"`
}

// Inventory 库存模型
type Inventory struct {
	BaseModel
	StoreID  uuid.UUID   `gorm:"type:uuid;not null;index:idx_store_sku" json:"store_id"`
	Store    *Store      `gorm:"foreignKey:StoreID" json:"store,omitempty"`
	SKUID    uuid.UUID   `gorm:"type:uuid;not null;index:idx_store_sku" json:"sku_id"`
	SKU      *ProductSKU `gorm:"foreignKey:SKUID" json:"sku,omitempty"`
	Quantity int         `gorm:"not null;default:0" json:"quantity"`
	AlertQty int         `gorm:"default:10" json:"alert_qty"`            // 预警数量
	Status   string      `gorm:"size:20;default:'normal'" json:"status"` // normal, low, out_of_stock
}

// InventoryLog 库存日志
type InventoryLog struct {
	BaseModel
	StoreID     uuid.UUID  `gorm:"type:uuid;not null" json:"store_id"`
	SKUID       uuid.UUID  `gorm:"type:uuid;not null" json:"sku_id"`
	Type        string     `gorm:"size:20;not null" json:"type"` // in, out, transfer, adjust
	Quantity    int        `gorm:"not null" json:"quantity"`
	OldQuantity int        `json:"old_quantity"`
	NewQuantity int        `json:"new_quantity"`
	Reason      string     `gorm:"size:255" json:"reason"`
	OperatorID  uuid.UUID  `gorm:"type:uuid" json:"operator_id"`
	RelatedID   *uuid.UUID `gorm:"type:uuid" json:"related_id"` // 关联订单或调拨单ID
}

// InventoryCount 库存盘点
type InventoryCount struct {
	BaseModel
	StoreID     uuid.UUID            `gorm:"type:uuid;not null" json:"store_id"`
	Store       *Store               `gorm:"foreignKey:StoreID" json:"store,omitempty"`
	CountNo     string               `gorm:"size:50;uniqueIndex;not null" json:"count_no"`
	Status      string               `gorm:"size:20;default:'pending'" json:"status"` // pending, in_progress, completed
	StartTime   *time.Time           `json:"start_time"`
	EndTime     *time.Time           `json:"end_time"`
	CompletedAt *time.Time           `json:"completed_at"`
	OperatorID  *uuid.UUID           `gorm:"type:uuid" json:"operator_id"`
	Operator    *User                `gorm:"foreignKey:OperatorID" json:"operator,omitempty"`
	Items       []InventoryCountItem `gorm:"foreignKey:CountID" json:"items,omitempty"`
}

// InventoryCountItem 库存盘点明细
type InventoryCountItem struct {
	BaseModel
	CountID    uuid.UUID       `gorm:"type:uuid;not null" json:"count_id"`
	Count      *InventoryCount `gorm:"foreignKey:CountID" json:"count,omitempty"`
	SKUID      uuid.UUID       `gorm:"type:uuid;not null" json:"sku_id"`
	SKU        *ProductSKU     `gorm:"foreignKey:SKUID" json:"sku,omitempty"`
	SystemQty  int             `gorm:"not null" json:"system_qty"`              // 系统库存
	CountedQty int             `gorm:"not null" json:"counted_qty"`             // 盘点数量
	Difference int             `gorm:"not null" json:"difference"`              // 差异数量
	Status     string          `gorm:"size:20;default:'pending'" json:"status"` // pending, confirmed
}

// Order 订单模型
type Order struct {
	BaseModel
	OrderNo        string      `gorm:"size:50;uniqueIndex;not null" json:"order_no"`
	StoreID        uuid.UUID   `gorm:"type:uuid;not null" json:"store_id"`
	Store          *Store      `gorm:"foreignKey:StoreID" json:"store,omitempty"`
	UserID         uuid.UUID   `gorm:"type:uuid;not null" json:"user_id"`
	User           *User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	TotalAmount    float64     `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	DiscountAmount float64     `gorm:"type:decimal(10,2);default:0" json:"discount_amount"`
	PaidAmount     float64     `gorm:"type:decimal(10,2);default:0" json:"paid_amount"`
	Status         string      `gorm:"size:20;not null;default:'pending'" json:"status"` // pending, paid, completed, cancelled, refunded
	PaymentType    string      `gorm:"size:20" json:"payment_type"`                      // wechat, alipay
	Items          []OrderItem `gorm:"foreignKey:OrderID" json:"items,omitempty"`
	CouponID       *uuid.UUID  `gorm:"type:uuid" json:"coupon_id"`
	Remark         string      `gorm:"size:255" json:"remark"`
	PaidAt         *time.Time  `json:"paid_at"`
	CompletedAt    *time.Time  `json:"completed_at"`
	CancelledAt    *time.Time  `json:"cancelled_at"`
}

// OrderItem 订单明细
type OrderItem struct {
	BaseModel
	OrderID    uuid.UUID   `gorm:"type:uuid;not null" json:"order_id"`
	Order      *Order      `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	SKUID      uuid.UUID   `gorm:"type:uuid;not null" json:"sku_id"`
	SKU        *ProductSKU `gorm:"foreignKey:SKUID" json:"sku,omitempty"`
	Quantity   int         `gorm:"not null" json:"quantity"`
	Price      float64     `gorm:"type:decimal(10,2);not null" json:"price"`
	TotalPrice float64     `gorm:"type:decimal(10,2);not null" json:"total_price"`
}

// Payment 支付记录
type Payment struct {
	BaseModel
	OrderID       uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex" json:"order_id"`
	Order         *Order     `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	PaymentNo     string     `gorm:"size:50;uniqueIndex;not null" json:"payment_no"`
	TransactionID string     `gorm:"size:100" json:"transaction_id"` // 第三方交易号
	Amount        float64    `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentType   string     `gorm:"size:20;not null" json:"payment_type"`
	Status        string     `gorm:"size:20;not null;default:'pending'" json:"status"` // pending, success, failed
	PaidAt        *time.Time `json:"paid_at"`
	CallbackData  string     `gorm:"type:text" json:"callback_data"`
}

// Coupon 优惠券
type Coupon struct {
	BaseModel
	Name        string    `gorm:"size:100;not null" json:"name"`
	Type        string    `gorm:"size:20;not null" json:"type"` // discount, voucher
	Value       float64   `gorm:"type:decimal(10,2);not null" json:"value"`
	MinAmount   float64   `gorm:"type:decimal(10,2);default:0" json:"min_amount"`
	MaxDiscount float64   `gorm:"type:decimal(10,2)" json:"max_discount"`
	Total       int       `gorm:"not null" json:"total"`
	Used        int       `gorm:"default:0" json:"used"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Status      string    `gorm:"size:20;default:'active'" json:"status"`
}

// UserCoupon 用户优惠券
type UserCoupon struct {
	BaseModel
	UserID   uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	User     *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CouponID uuid.UUID  `gorm:"type:uuid;not null" json:"coupon_id"`
	Coupon   *Coupon    `gorm:"foreignKey:CouponID" json:"coupon,omitempty"`
	Status   string     `gorm:"size:20;default:'unused'" json:"status"` // unused, used, expired
	UsedAt   *time.Time `json:"used_at"`
}

// MemberPointsLog 积分记录
type MemberPointsLog struct {
	BaseModel
	UserID    uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	User      *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Points    int        `gorm:"not null" json:"points"`       // 正数为增加，负数为扣减
	Type      string     `gorm:"size:20;not null" json:"type"` // order, sign_in, exchange, expire
	RelatedID *uuid.UUID `gorm:"type:uuid" json:"related_id"`
	Remark    string     `gorm:"size:255" json:"remark"`
}

// AccessLog 门禁日志
type AccessLog struct {
	BaseModel
	StoreID  uuid.UUID `gorm:"type:uuid;not null" json:"store_id"`
	Store    *Store    `gorm:"foreignKey:StoreID" json:"store,omitempty"`
	UserID   uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User     *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Action   string    `gorm:"size:20;not null" json:"action"` // enter, exit
	Method   string    `gorm:"size:20" json:"method"`          // qrcode, face
	DeviceID string    `gorm:"size:100" json:"device_id"`
	Status   string    `gorm:"size:20;not null" json:"status"` // success, denied
	Remark   string    `gorm:"size:255" json:"remark"`
}

// Blacklist 黑名单
type Blacklist struct {
	BaseModel
	UserID   uuid.UUID `gorm:"type:uuid;not null;uniqueIndex" json:"user_id"`
	User     *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Reason   string    `gorm:"size:255;not null" json:"reason"`
	Operator uuid.UUID `gorm:"type:uuid;not null" json:"operator"`
	Status   string    `gorm:"size:20;default:'active'" json:"status"`
}

// MonitoringDevice 监控设备
type MonitoringDevice struct {
	BaseModel
	StoreID   uuid.UUID `gorm:"type:uuid;not null" json:"store_id"`
	Store     *Store    `gorm:"foreignKey:StoreID" json:"store,omitempty"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Type      string    `gorm:"size:20;not null" json:"type"` // camera, sensor, refrigerator
	DeviceID  string    `gorm:"size:100;uniqueIndex;not null" json:"device_id"`
	Brand     string    `gorm:"size:50" json:"brand"`
	Model     string    `gorm:"size:50" json:"model"`
	IP        string    `gorm:"size:50" json:"ip"`
	Port      int       `json:"port"`
	StreamURL string    `gorm:"size:255" json:"stream_url"`
	Location  string    `gorm:"size:100" json:"location"`
	Status    string    `gorm:"size:20;default:'online'" json:"status"` // online, offline, fault
}

// MonitoringAlert 监控告警
type MonitoringAlert struct {
	BaseModel
	StoreID   uuid.UUID         `gorm:"type:uuid;not null" json:"store_id"`
	Store     *Store            `gorm:"foreignKey:StoreID" json:"store,omitempty"`
	DeviceID  *uuid.UUID        `gorm:"type:uuid" json:"device_id"`
	Device    *MonitoringDevice `gorm:"foreignKey:DeviceID" json:"device,omitempty"`
	Type      string            `gorm:"size:20;not null" json:"type"`  // intrusion, device_offline, temperature
	Level     string            `gorm:"size:20;not null" json:"level"` // info, warning, critical
	Message   string            `gorm:"size:255;not null" json:"message"`
	Status    string            `gorm:"size:20;default:'pending'" json:"status"` // pending, processing, resolved
	Handler   *uuid.UUID        `gorm:"type:uuid" json:"handler"`
	HandledAt *time.Time        `json:"handled_at"`
}
