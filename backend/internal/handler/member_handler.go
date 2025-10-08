package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"github.com/unmanned-store/backend/internal/service"
)

type MemberHandler struct {
	memberService service.MemberService
}

func NewMemberHandler(memberService service.MemberService) *MemberHandler {
	return &MemberHandler{memberService: memberService}
}

func (h *MemberHandler) GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	profile, err := h.memberService.GetProfile(userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *MemberHandler) UpdateProfile(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	user.ID = userID.(uuid.UUID)

	if err := h.memberService.UpdateProfile(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *MemberHandler) GetPoints(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"points": 0})
}

func (h *MemberHandler) GetCoupons(c *gin.Context) {
	userID, _ := c.Get("user_id")
	coupons, err := h.memberService.GetCoupons(userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, coupons)
}
