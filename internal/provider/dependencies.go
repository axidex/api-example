package provider

import "context"

type initFunc = func(context.Context) error

type Dependencies struct {
}

func (d *Dependencies) Stop() {}
