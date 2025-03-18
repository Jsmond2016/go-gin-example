package middleware

import (
	"time"

	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/gin-gonic/gin"
)

// 添加请求耗时中间件
func RequestTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		logging.Info("请求耗时:", c.Request.URL.Path, latency)
	}
}
