package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/babycare/utils/logger"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.WithField("stack", string(debug.Stack())).Error(err)

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()
		c.Next()
	}
} 