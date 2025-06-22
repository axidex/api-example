package controller

import (
	"context"
	"github.com/axidex/api-example/server/internal/dto"
	"github.com/axidex/api-example/server/internal/storage"
	"github.com/axidex/api-example/server/pkg/utils"
)

type Controller interface {
	CreateUser(ctx context.Context, user *dto.User) error
	DeleteUser(ctx context.Context, user *dto.User) error
	GetUser(ctx context.Context, name string) (*dto.User, error)
	ListUsers(ctx context.Context, limit, offset int) ([]*dto.User, error)
}

type ApiController struct {
	storage storage.ApiStorage
}

func NewApiController(storage storage.ApiStorage) Controller {
	return &ApiController{
		storage: storage,
	}
}

func (c *ApiController) CreateUser(ctx context.Context, user *dto.User) error {
	if err := c.storage.UserRepository.Create(ctx, user.Storage()); err != nil {
		return err
	}

	return nil
}

func (c *ApiController) DeleteUser(ctx context.Context, user *dto.User) error {
	if err := c.storage.UserRepository.Delete(ctx, user.Storage()); err != nil {
		return err
	}

	return nil
}

func (c *ApiController) GetUser(ctx context.Context, name string) (*dto.User, error) {
	user, err := c.storage.UserRepository.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	return dto.UserFromStorage(user), nil
}

func (c *ApiController) ListUsers(ctx context.Context, limit, offset int) ([]*dto.User, error) {
	users, err := c.storage.UserRepository.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return utils.Map(users, dto.UserFromStorage), nil
}
