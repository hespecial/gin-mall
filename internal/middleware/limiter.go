package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/internal/common/constant"
	"github.com/hespecial/gin-mall/pkg/limiter"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func Limiter(r time.Duration, b int) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, exists := c.Get(constant.Username)
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "登录信息有误"})
			c.Abort()
			return
		}

		l := limiter.GetLimiter(rate.Every(r), b, username.(string))
		if !l.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "请勿频繁操作"})
			c.Abort()
			return
		}

		c.Next()
	}
}
