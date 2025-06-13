package provider

import (
	"context"
	"gorm.io/gorm"
)

type initFunc = func(context.Context) error

type Dependencies struct {
	db *gorm.DB
}

func (d *Dependencies) Stop() {}
