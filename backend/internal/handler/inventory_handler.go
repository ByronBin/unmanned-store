package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/service"
)

type InventoryHandler struct {
	inventoryService service.InventoryService
}

func NewInventoryHandler(inventoryService service.InventoryService) *InventoryHandler {
	return &InventoryHandler{inventoryService: inventoryService}
}

type StockInRequest struct {
	SKUID    string `json:"sku_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Remark   string `json:"remark"`
}

func (h *InventoryHandler) List(c *gin.Context) {
	storeIDStr := c.Query("store_id")
	storeID, err := uuid.Parse(storeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的门店ID"})
		return
	}

	inventories, total, err := h.inventoryService.List(storeID, 1, 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  inventories,
		"total": total,
	})
}

func (h *InventoryHandler) StockIn(c *gin.Context) {
	var req StockInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storeID, _ := c.Get("store_id")
	userID, _ := c.Get("user_id")
	skuID, err := uuid.Parse(req.SKUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的SKU ID"})
		return
	}

	err = h.inventoryService.StockIn(storeID.(uuid.UUID), skuID, req.Quantity, userID.(uuid.UUID), req.Remark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "入库成功"})
}

func (h *InventoryHandler) StockOut(c *gin.Context) {
	var req StockInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storeID, _ := c.Get("store_id")
	userID, _ := c.Get("user_id")
	skuID, err := uuid.Parse(req.SKUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的SKU ID"})
		return
	}

	err = h.inventoryService.StockOut(storeID.(uuid.UUID), skuID, req.Quantity, userID.(uuid.UUID), req.Remark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "出库成功"})
}

func (h *InventoryHandler) Transfer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "调拨功能"})
}

func (h *InventoryHandler) Alerts(c *gin.Context) {
	storeID, _ := c.Get("store_id")
	alerts, err := h.inventoryService.GetAlerts(storeID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alerts)
}
