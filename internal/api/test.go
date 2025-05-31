package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Test
// @Summary Test endpoint
// @Description Endpoint for testing integration with telemetry
// @Tags tests
// @Accept json
// @Produce json
// @Success 200 {object} string "Success"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /v1/test [get]
func (h *GinHandler) Test(c *gin.Context) {
	ctx := c.Request.Context()

	h.logger.Info(ctx, "Test endpoint")
	time.Sleep(10 * time.Second)

	c.JSON(http.StatusOK, "success")
	return
}
