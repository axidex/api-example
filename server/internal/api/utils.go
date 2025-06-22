package api

import (
	"github.com/axidex/api-example/server/internal/errors"
)

type BindFunc func(obj any) error

func BindRequestParams[T any](bindFunc BindFunc) (T, error) {
	var params T
	if err := bindFunc(&params); err != nil {
		return params, errors.ErrInvalidRequestParams
	}

	return params, nil
}
