package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unmanned-store/backend/internal/service"
)

type AnalyticsHandler struct {
	analyticsService service.AnalyticsService
}

func NewAnalyticsHandler(analyticsService service.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{analyticsService: analyticsService}
}

func (h *AnalyticsHandler) SalesStats(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "销售统计"})
}

func (h *AnalyticsHandler) HotProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "热销商品"})
}

func (h *AnalyticsHandler) CustomerStats(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "顾客统计"})
}
