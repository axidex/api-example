package api

import (
	"github.com/axidex/api-example/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CustomRecoveryFunc is a custom recovery function
func CustomRecoveryFunc(logger logger.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		logger.Error(
			c.Request.Context(),
			"Panic recovered | %v | path: %s | method: %s | ip: %s",
			recovered,
			c.Request.URL.Path,
			c.Request.Method,
			c.ClientIP(),
		)

		// Abort the request and return a 500 Internal Server Error
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
