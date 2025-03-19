package logger

import (
	"bytes"
	"io"
	"time"

	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 记录请求信息
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		method := c.Request.Method

		// 读取请求体
		var body []byte
		if c.Request.Body != nil {
			body, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		c.Next()

		// 结束时间
		end := time.Now()
		latency := end.Sub(start)

		// 记录响应信息
		status := c.Writer.Status()

		logging.Info("请求日志",
			"path", path,
			"method", method,
			"query", raw,
			"body", string(body),
			"status", status,
			"latency", latency,
		)
	}
}
