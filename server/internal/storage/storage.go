package storage

import "gorm.io/gorm"

type ApiStorage struct {
	UserRepository UserRepository
}

func NewApiStorage(db *gorm.DB) ApiStorage {
	return ApiStorage{
		UserRepository: NewUserRepository(db),
	}
}
