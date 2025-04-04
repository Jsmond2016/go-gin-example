package jwt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		tokenStr := c.GetHeader("Authorization")
		token := strings.TrimPrefix(tokenStr, "Bearer ")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				if err.Error() == "Token is expired" {
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				} else {
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			} else {
				c.Set("claims", claims)
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
