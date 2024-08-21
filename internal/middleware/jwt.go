package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/internal/common/enum"
	"github.com/hespecial/gin-mall/pkg/jwt"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求未携带token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 尝试解析访问令牌
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			// 如果访问令牌无效或过期，检查是否存在刷新令牌
			refreshToken := c.GetHeader("X-Refresh-Token")
			if refreshToken == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "访问令牌无效或过期，且无刷新令牌"})
				c.Abort()
				return
			}

			// 解析刷新令牌并生成新令牌
			newAccessToken, newRefreshToken, err := jwt.ParseRefreshToken(refreshToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "刷新令牌无效或过期"})
				c.Abort()
				return
			}

			// 将新令牌添加到响应头
			c.Header("New-Access-Token", newAccessToken)
			c.Header("New-Refresh-Token", newRefreshToken)

			// 尝试解析新的访问令牌
			claims, err = jwt.ParseToken(newAccessToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "新访问令牌无效"})
				c.Abort()
				return
			}
		}

		// 将用户信息存储在上下文中
		c.Set(enum.UserID, claims.UserID)
		c.Set(enum.Username, claims.Username)

		// 继续处理请求
		c.Next()
	}
}
