package storage

import (
	"context"
	"errors"
	"github.com/axidex/api-example/pkg/db"
	"github.com/axidex/api-example/pkg/tables"
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *tables.User) error
	Update(ctx context.Context, user *tables.User) error
	Get(ctx context.Context, name string) (*tables.User, error)
	Delete(ctx context.Context, user *tables.User) error
	List(ctx context.Context, limit, offset int) ([]*tables.User, error)
	Exists(ctx context.Context, user string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *tables.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	if err := r.db.WithContext(ctx).Where(tables.User{Name: user.Name}).FirstOrCreate(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Update(ctx context.Context, user *tables.User) error {
	user.UpdatedAt = time.Now()

	result := r.db.WithContext(ctx).Model(user).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return db.ErrRecordNotFound
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, user *tables.User) error {
	result := r.db.WithContext(ctx).Model(&tables.User{}).Where("name = ?", user.Name).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*tables.User, error) {
	var users []*tables.User
	if err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Limit(limit).
		Offset(offset).
		Find(&users).
		Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Get(ctx context.Context, name string) (*tables.User, error) {
	var user tables.User
	if err := r.db.WithContext(ctx).Where("name = ? AND deleted_at IS NULL", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, db.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Exists(ctx context.Context, name string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&tables.User{}).Where("name = ? AND deleted_at IS NULL", name).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
