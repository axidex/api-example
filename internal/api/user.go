package api

import (
	"github.com/axidex/api-example/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUser
// @Summary Get user by name
// @Description Returns user information by provided name
// @Tags Users
// @Accept json
// @Produce json
// @Param name query string true "Username to search"
// @Success 200 {object} Response "Successfully found user"
// @Failure 404 {object} Response "User not found"
// @Failure 422 {object} Response "Invalid request parameters"
// @Failure 500 {object} Response "Internal server error"
// @Router /v1/user [get]
func (h *GinHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	user, err := BindRequestParams[dto.User](c.ShouldBindQuery)
	if err != nil {
		h.logger.Error(ctx, "Can't bind request params")
		ResponseUserError(c, err)
		return
	}

	foundedUser, err := h.controller.GetUser(ctx, user.Name)
	if err != nil {
		h.logger.Error(ctx, "Can't get user")
		ResponseUserError(c, err)
		return
	}

	ResponseUserSuccess(c, foundedUser, http.StatusOK)
}

// CreateUser
// @Summary Create user
// @Description Creates user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.User true "User search request"
// @Success 200 {object} Response "Successfully found user"
// @Failure 404 {object} Response "User not found"
// @Failure 422 {object} Response "Invalid request parameters"
// @Failure 500 {object} Response "Internal server error"
// @Router /v1/user [Post]
func (h *GinHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	user, err := BindRequestParams[dto.User](c.ShouldBindBodyWithJSON)
	if err != nil {
		h.logger.Error(ctx, "Can't bind request params")
		return
	}

	if err := h.controller.CreateUser(ctx, &user); err != nil {
		h.logger.Error(ctx, "Can't get user")
		ResponseUserError(c, err)
		return
	}

	ResponseUserSuccess(c, nil, http.StatusOK)
}
