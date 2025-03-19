package recovery

import (
	"net/http"
	"runtime/debug"

	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logging.Error("panic 恢复:", err)
				logging.Error("堆栈信息:", string(debug.Stack()))

				c.JSON(http.StatusInternalServerError, gin.H{
					"code": e.ERROR,
					"msg":  e.GetMsg(e.ERROR),
					"data": nil,
				})
				c.Abort()
			}
		}()

		c.Next()
	}
}
