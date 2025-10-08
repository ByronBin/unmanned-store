package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/service"
)

type MonitoringHandler struct {
	monitoringService service.MonitoringService
}

func NewMonitoringHandler(monitoringService service.MonitoringService) *MonitoringHandler {
	return &MonitoringHandler{monitoringService: monitoringService}
}

func (h *MonitoringHandler) GetStreams(c *gin.Context) {
	storeIDStr := c.Query("store_id")
	storeID, err := uuid.Parse(storeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的门店ID"})
		return
	}

	streams, err := h.monitoringService.GetStreams(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, streams)
}

func (h *MonitoringHandler) GetAlerts(c *gin.Context) {
	alerts, total, err := h.monitoringService.GetAlerts(nil, 1, 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  alerts,
		"total": total,
	})
}

func (h *MonitoringHandler) GetDevices(c *gin.Context) {
	storeIDStr := c.Query("store_id")
	storeID, err := uuid.Parse(storeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的门店ID"})
		return
	}

	devices, err := h.monitoringService.GetDevices(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, devices)
}
