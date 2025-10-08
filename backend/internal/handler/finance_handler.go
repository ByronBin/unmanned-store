package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unmanned-store/backend/internal/service"
)

type FinanceHandler struct {
	financeService service.FinanceService
}

func NewFinanceHandler(financeService service.FinanceService) *FinanceHandler {
	return &FinanceHandler{financeService: financeService}
}

func (h *FinanceHandler) DailyReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "日报"})
}

func (h *FinanceHandler) MonthlyReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "月报"})
}

func (h *FinanceHandler) Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "汇总"})
}
