package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unmanned-store/backend/internal/service"
)

type PaymentHandler struct {
	paymentService service.PaymentService
}

func NewPaymentHandler(paymentService service.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) WechatPay(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "微信支付"})
}

func (h *PaymentHandler) AlipayPay(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "支付宝支付"})
}

func (h *PaymentHandler) WechatCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *PaymentHandler) AlipayCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
