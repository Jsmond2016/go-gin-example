package ratelimit

import (
	"net/http"
	"sync"
	"time"

	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	rate   int // 限制次数
	burst  int // 突发次数
	mutex  sync.Mutex
	tokens map[string][]time.Time // 记录访问时间
}

func NewRateLimiter(rate, burst int) *RateLimiter {
	return &RateLimiter{
		rate:   rate,
		burst:  burst,
		tokens: make(map[string][]time.Time),
	}
}

func (rl *RateLimiter) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		rl.mutex.Lock()
		now := time.Now()

		// 清理过期的记录
		if times, exists := rl.tokens[ip]; exists {
			var valid []time.Time
			for _, t := range times {
				if now.Sub(t) < time.Minute {
					valid = append(valid, t)
				}
			}
			rl.tokens[ip] = valid
		}

		// 检查是否超过限制
		if len(rl.tokens[ip]) >= rl.rate {
			rl.mutex.Unlock()
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code": e.ERROR_TOO_MANY_REQUESTS,
				"msg":  e.GetMsg(e.ERROR_TOO_MANY_REQUESTS),
				"data": nil,
			})
			c.Abort()
			return
		}

		// 记录新的访问时间
		rl.tokens[ip] = append(rl.tokens[ip], now)
		rl.mutex.Unlock()

		c.Next()
	}
}
