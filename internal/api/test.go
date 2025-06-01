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
// @Accept application/json
// @Produce text/plain
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

// TestError
// @Summary Test endpoint with error
// @Description Endpoint for testing integration with telemetry with error
// @Tags tests
// @Accept application/json
// @Produce text/plain
// @Success 200 {object} string "Success"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /v1/test-error [get]
func (h *GinHandler) TestError(c *gin.Context) {
	ctx := c.Request.Context()

	h.logger.Info(ctx, "Test Error endpoint")
	time.Sleep(5 * time.Second)

	c.JSON(http.StatusBadRequest, "failed")
	return
}
