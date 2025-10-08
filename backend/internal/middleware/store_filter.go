package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// StoreFilter 多租户数据过滤中间件
// 自动在查询条件中添加 store_id 过滤
func StoreFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")

		// admin 角色可以访问所有门店数据
		if role == "admin" {
			c.Next()
			return
		}

		// 其他角色只能访问所属门店的数据
		storeID, exists := c.Get("store_id")
		if exists {
			c.Set("filter_store_id", storeID.(uuid.UUID))
		}

		c.Next()
	}
}
