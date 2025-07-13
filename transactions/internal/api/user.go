package api

import (
	"github.com/axidex/api-example/transactions/internal/api/dto"
	"github.com/axidex/api-example/transactions/pkg/ton"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateCell
// @Summary Create cell with payload
// @Description Returns cell BOC in base64
// @Tags Users
// @Accept json
// @Produce json
// @Param payload query string true "Payload for cell"
// @Success 200 {object} Response "Successfully found user"
// @Failure 422 {object} Response "Invalid request parameters"
// @Failure 500 {object} Response "Internal server error"
// @Router /v1/cell [post]
func (h *GinHandler) CreateCell(c *gin.Context) {
	ctx := c.Request.Context()

	cellRequest, err := BindRequestParams[dto.CellRequest](c.ShouldBindQuery)
	if err != nil {
		h.logger.Error(ctx, "Can't bind request params")
		ResponseUserError(c, err)
		return
	}

	cell, err := ton.CreateCell(cellRequest.Payload)
	if err != nil {
		h.logger.Error(ctx, "Can't create cell")
		ResponseUserError(c, err)
		return
	}

	ResponseUserSuccess(c, cell, http.StatusOK)
}
