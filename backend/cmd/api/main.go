package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unmanned-store/backend/internal/handler"
	"github.com/unmanned-store/backend/internal/middleware"
	"github.com/unmanned-store/backend/internal/repository"
	"github.com/unmanned-store/backend/internal/service"
	"github.com/unmanned-store/backend/pkg/config"
	"github.com/unmanned-store/backend/pkg/database"
	"github.com/unmanned-store/backend/pkg/logger"
	"github.com/unmanned-store/backend/pkg/redis"
)

// @title 无人超市系统API
// @version 1.0
// @description 24小时无人值守超市管理系统
// @host localhost:8080
// @BasePath /api/v1
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化日志
	logger.Init(cfg.Log.Level, cfg.Log.FilePath)
	defer logger.Sync()

	// 初始化数据库
	db, err := database.InitDB(cfg.Database)
	if err != nil {
		logger.Fatal("Failed to initialize database", "error", err)
	}

	// 初始化Redis
	rdb, err := redis.InitRedis(cfg.Redis)
	if err != nil {
		logger.Fatal("Failed to initialize redis", "error", err)
	}

	// 初始化仓储层
	repos := repository.NewRepositories(db)

	// 初始化服务层
	services := service.NewServices(repos, rdb, cfg)

	// 初始化处理器层
	handlers := handler.NewHandlers(services)

	// 设置Gin模式
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建路由
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API路由组
	v1 := router.Group("/api/v1")
	{
		// 公开接口
		auth := v1.Group("/auth")
		{
			auth.POST("/login", handlers.Auth.Login)
			auth.POST("/register", handlers.Auth.Register)
			auth.POST("/refresh", handlers.Auth.RefreshToken)
		}

		// 需要认证的接口
		authorized := v1.Group("")
		authorized.Use(middleware.JWTAuth(cfg.JWT.Secret))
		{
			// 门店管理
			stores := authorized.Group("/stores")
			stores.Use(middleware.RequireRole("admin", "store_manager"))
			{
				stores.GET("", handlers.Store.List)
				stores.POST("", handlers.Store.Create)
				stores.GET("/:id", handlers.Store.Get)
				stores.PUT("/:id", handlers.Store.Update)
				stores.DELETE("/:id", handlers.Store.Delete)
			}

			// 商品管理
			products := authorized.Group("/products")
			{
				products.GET("", handlers.Product.List)
				products.GET("/search", handlers.Product.Search)
				products.GET("/hot", handlers.Product.GetHotProducts)
				products.GET("/:id", handlers.Product.Get)
				products.POST("", middleware.RequireRole("admin", "store_manager"), handlers.Product.Create)
				products.PUT("/:id", middleware.RequireRole("admin", "store_manager"), handlers.Product.Update)
				products.PUT("/:id/status", middleware.RequireRole("admin", "store_manager"), handlers.Product.UpdateStatus)
				products.DELETE("/:id", middleware.RequireRole("admin", "store_manager"), handlers.Product.Delete)

				// SKU管理
				products.POST("/:id/skus", middleware.RequireRole("admin", "store_manager"), handlers.Product.CreateSKU)
				products.PUT("/skus/:id", middleware.RequireRole("admin", "store_manager"), handlers.Product.UpdateSKU)
				products.DELETE("/skus/:id", middleware.RequireRole("admin", "store_manager"), handlers.Product.DeleteSKU)
				products.GET("/skus/:id", handlers.Product.GetSKU)
			}

		// 分类管理
		categories := authorized.Group("/categories")
		{
			categories.GET("", handlers.Category.List)
			categories.GET("/tree", handlers.Category.GetTree)
			categories.GET("/:id", handlers.Category.Get)
			categories.POST("", middleware.RequireRole("admin"), handlers.Category.Create)
			categories.PUT("/:id", middleware.RequireRole("admin"), handlers.Category.Update)
			categories.DELETE("/:id", middleware.RequireRole("admin"), handlers.Category.Delete)
		}

			// 库存管理
			inventory := authorized.Group("/inventory")
			inventory.Use(middleware.RequireRole("admin", "store_manager"))
			{
				inventory.GET("", handlers.Inventory.GetByStore)
				inventory.GET("/sku/:skuId", handlers.Inventory.GetBySKU)
				inventory.GET("/product/:productId", handlers.Inventory.GetByProduct)
				inventory.POST("/adjust", handlers.Inventory.AdjustInventory)
				inventory.POST("/stock-in", handlers.Inventory.StockIn)
				inventory.POST("/stock-out", handlers.Inventory.StockOut)
				inventory.GET("/low-stock", handlers.Inventory.GetLowStockItems)
				inventory.GET("/logs", handlers.Inventory.GetInventoryLogs)

				// 库存盘点
				inventory.POST("/counts", handlers.Inventory.CreateInventoryCount)
				inventory.GET("/counts", handlers.Inventory.GetInventoryCounts)
				inventory.POST("/counts/:id/submit", handlers.Inventory.SubmitInventoryCount)
			}

			// 订单管理
			orders := authorized.Group("/orders")
			{
				orders.GET("", handlers.Order.List)
				orders.POST("", handlers.Order.Create)
				orders.GET("/:id", handlers.Order.Get)
				orders.POST("/:id/cancel", handlers.Order.Cancel)
				orders.POST("/:id/refund", middleware.RequireRole("admin", "store_manager"), handlers.Order.Refund)
			}

			// 支付管理
			payments := authorized.Group("/payments")
			{
				payments.POST("/wechat", handlers.Payment.WechatPay)
				payments.POST("/alipay", handlers.Payment.AlipayPay)
			}

			// 会员管理
			members := authorized.Group("/members")
			{
				members.GET("/profile", handlers.Member.GetProfile)
				members.PUT("/profile", handlers.Member.UpdateProfile)
				members.GET("/points", handlers.Member.GetPoints)
				members.GET("/coupons", handlers.Member.GetCoupons)
			}

			// 门禁管理
			access := authorized.Group("/access")
			{
				access.POST("/open", handlers.Access.OpenDoor)
				access.GET("/logs", middleware.RequireRole("admin", "store_manager"), handlers.Access.GetLogs)
				access.POST("/blacklist", middleware.RequireRole("admin"), handlers.Access.AddToBlacklist)
			}

			// 监控管理
			monitoring := authorized.Group("/monitoring")
			monitoring.Use(middleware.RequireRole("admin", "store_manager"))
			{
				monitoring.GET("/streams", handlers.Monitoring.GetStreams)
				monitoring.GET("/alerts", handlers.Monitoring.GetAlerts)
				monitoring.GET("/devices", handlers.Monitoring.GetDevices)
			}

			// 财务报表
			finance := authorized.Group("/finance")
			finance.Use(middleware.RequireRole("admin", "store_manager"))
			{
				finance.GET("/daily", handlers.Finance.DailyReport)
				finance.GET("/monthly", handlers.Finance.MonthlyReport)
				finance.GET("/summary", handlers.Finance.Summary)
			}

			// 数据分析
			analytics := authorized.Group("/analytics")
			analytics.Use(middleware.RequireRole("admin"))
			{
				analytics.GET("/sales", handlers.Analytics.SalesStats)
				analytics.GET("/products/hot", handlers.Analytics.HotProducts)
				analytics.GET("/customers", handlers.Analytics.CustomerStats)
			}
		}
	}

	// 支付回调（无需认证）
	callback := v1.Group("/callback")
	{
		callback.POST("/wechat", handlers.Payment.WechatCallback)
		callback.POST("/alipay", handlers.Payment.AlipayCallback)
	}

	// 创建HTTP服务器
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 启动服务器
	go func() {
		logger.Info("Starting server", "port", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", "error", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", "error", err)
	}

	logger.Info("Server exited")
}
