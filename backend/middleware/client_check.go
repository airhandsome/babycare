package middleware

import "github.com/gin-gonic/gin"

func ClientCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			clientType := c.GetHeader("X-Client-Type")
			if clientType != "web-frontend" {
				c.AbortWithStatus(403)
				return
			}
		}
		c.Next()
	}
} 