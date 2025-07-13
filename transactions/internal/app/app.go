package app

import "context"

type initFunc func(context.Context) error

type IApp interface {
	Run(ctx context.Context) error
}
