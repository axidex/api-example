package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *GinHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "I'm alive"})
}
