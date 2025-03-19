package permission

import (
	"net/http"

	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

func CheckPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 JWT 中获取用户信息
		userID := c.GetInt("user_id")
		role := c.GetString("role")

		// 检查用户权限
		if !hasPermission(userID, role, c.Request.URL.Path, c.Request.Method) {
			c.JSON(http.StatusForbidden, gin.H{
				"code": e.ERROR_NO_PERMISSION,
				"msg":  e.GetMsg(e.ERROR_NO_PERMISSION),
				"data": nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// TODO: 待实现-完善
func hasPermission(userID int, role, path, method string) bool {
	// 实现具体的权限检查逻辑
	return true
}
