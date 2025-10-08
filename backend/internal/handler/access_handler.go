package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/service"
)

type AccessHandler struct {
	accessService service.AccessService
}

func NewAccessHandler(accessService service.AccessService) *AccessHandler {
	return &AccessHandler{accessService: accessService}
}

type OpenDoorRequest struct {
	StoreID string `json:"store_id" binding:"required"`
}

func (h *AccessHandler) OpenDoor(c *gin.Context) {
	var req OpenDoorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storeID, err := uuid.Parse(req.StoreID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的门店ID"})
		return
	}

	userID, _ := c.Get("user_id")
	if err := h.accessService.OpenDoor(storeID, userID.(uuid.UUID)); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "开门成功"})
}

func (h *AccessHandler) GetLogs(c *gin.Context) {
	logs, total, err := h.accessService.GetLogs(nil, 1, 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  logs,
		"total": total,
	})
}

func (h *AccessHandler) AddToBlacklist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "添加黑名单"})
}
