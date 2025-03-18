package jwt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/jwt"
	"github.com/EDDYCJY/go-gin-example/pkg/response"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		token := strings.TrimPrefix(tokenStr, "Bearer ")
		if token == "" {
			response.Error(c, http.StatusUnauthorized, response.CodeInvalidParams, nil)
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			var code int
			if err.Error() == "Token is expired" {
				code = response.CodeTokenExpired
			} else {
				code = response.CodeTokenInvalid
			}
			response.Error(c, http.StatusUnauthorized, code, nil)
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("claims", claims)
		c.Next()
	}
}