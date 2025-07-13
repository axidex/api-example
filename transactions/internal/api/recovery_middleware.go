package api

import (
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CustomRecoveryFunc is a custom recovery function
func CustomRecoveryFunc(l logger.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		l.Error(
			c.Request.Context(),
			"Panic recovered",
			logger.NewAttribute("info", recovered),
			logger.NewAttribute("path", c.Request.URL.Path),
			logger.NewAttribute("method", c.Request.Method),
			logger.NewAttribute("client_ip", c.ClientIP()),
		)

		// Abort the request and return a 500 Internal Server Error
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
