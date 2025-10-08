package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
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

type StockOutRequest struct {
	SKUID    string `json:"sku_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Reason   string `json:"reason" binding:"required"`
}

type AdjustInventoryRequest struct {
	SKUID    string `json:"sku_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Reason   string `json:"reason" binding:"required"`
}

func (h *InventoryHandler) GetByStore(c *gin.Context) {
	storeIDStr := c.Query("store_id")
	storeID, err := uuid.Parse(storeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的门店ID"})
		return
	}

	page := 1
	pageSize := 20
	
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}
	
	if ps := c.Query("page_size"); ps != "" {
		if parsed, err := strconv.Atoi(ps); err == nil && parsed > 0 && parsed <= 100 {
			pageSize = parsed
		}
	}

	inventories, total, err := h.inventoryService.GetByStore(storeID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      inventories,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

func (h *InventoryHandler) GetBySKU(c *gin.Context) {
	skuIDStr := c.Param("skuId")
	skuID, err := uuid.Parse(skuIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的SKU ID"})
		return
	}

	var storeID *uuid.UUID
	if storeIDStr := c.Query("store_id"); storeIDStr != "" {
		if parsed, err := uuid.Parse(storeIDStr); err == nil {
			storeID = &parsed
		}
	}

	inventories, err := h.inventoryService.GetBySKU(skuID, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": inventories})
}

func (h *InventoryHandler) GetByProduct(c *gin.Context) {
	productIDStr := c.Param("productId")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的商品ID"})
		return
	}

	var storeID *uuid.UUID
	if storeIDStr := c.Query("store_id"); storeIDStr != "" {
		if parsed, err := uuid.Parse(storeIDStr); err == nil {
			storeID = &parsed
		}
	}

	inventories, err := h.inventoryService.GetByProduct(productID, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": inventories})
}

func (h *InventoryHandler) AdjustInventory(c *gin.Context) {
	var req AdjustInventoryRequest
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

	err = h.inventoryService.AdjustInventory(storeID.(uuid.UUID), skuID, req.Quantity, req.Reason, userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "库存调整成功"})
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

	err = h.inventoryService.StockIn(storeID.(uuid.UUID), skuID, req.Quantity, userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "入库成功"})
}

func (h *InventoryHandler) StockOut(c *gin.Context) {
	var req StockOutRequest
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

	err = h.inventoryService.StockOut(storeID.(uuid.UUID), skuID, req.Quantity, req.Reason, userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "出库成功"})
}

func (h *InventoryHandler) GetLowStockItems(c *gin.Context) {
	threshold := 10
	if t := c.Query("threshold"); t != "" {
		if parsed, err := strconv.Atoi(t); err == nil && parsed > 0 {
			threshold = parsed
		}
	}

	var storeID *uuid.UUID
	if storeIDStr := c.Query("store_id"); storeIDStr != "" {
		if parsed, err := uuid.Parse(storeIDStr); err == nil {
			storeID = &parsed
		}
	}

	items, err := h.inventoryService.GetLowStockItems(storeID, threshold)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h *InventoryHandler) GetInventoryLogs(c *gin.Context) {
	page := 1
	pageSize := 20
	
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}
	
	if ps := c.Query("page_size"); ps != "" {
		if parsed, err := strconv.Atoi(ps); err == nil && parsed > 0 && parsed <= 100 {
			pageSize = parsed
		}
	}

	var storeID, skuID *uuid.UUID
	
	if storeIDStr := c.Query("store_id"); storeIDStr != "" {
		if parsed, err := uuid.Parse(storeIDStr); err == nil {
			storeID = &parsed
		}
	}
	
	if skuIDStr := c.Query("sku_id"); skuIDStr != "" {
		if parsed, err := uuid.Parse(skuIDStr); err == nil {
			skuID = &parsed
		}
	}

	logs, total, err := h.inventoryService.GetInventoryLogs(storeID, skuID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      logs,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// 库存盘点相关方法
func (h *InventoryHandler) CreateInventoryCount(c *gin.Context) {
	var count domain.InventoryCount
	if err := c.ShouldBindJSON(&count); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.inventoryService.CreateInventoryCount(&count); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, count)
}

func (h *InventoryHandler) GetInventoryCounts(c *gin.Context) {
	storeIDStr := c.Query("store_id")
	storeID, err := uuid.Parse(storeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的门店ID"})
		return
	}

	status := c.Query("status")
	counts, err := h.inventoryService.GetInventoryCounts(storeID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": counts})
}

func (h *InventoryHandler) SubmitInventoryCount(c *gin.Context) {
	countIDStr := c.Param("id")
	countID, err := uuid.Parse(countIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的盘点ID"})
		return
	}

	var req struct {
		Items []domain.InventoryCountItem `json:"items" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	if err := h.inventoryService.SubmitInventoryCount(countID, req.Items, userID.(uuid.UUID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "盘点提交成功"})
}
