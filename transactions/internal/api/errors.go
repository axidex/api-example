package api

import (
	"github.com/axidex/api-example/server/pkg/db"
	"github.com/axidex/api-example/transactions/internal/api/errors"
	"net/http"
)

var ErrorsStatusCodes = map[error]int{
	db.ErrRecordNotFound:           http.StatusNotFound,
	errors.ErrInvalidRequestParams: http.StatusUnprocessableEntity,
}

// GetErrorResponse return status code and error info
func GetErrorResponse(err error) (int, string) {
	code, ok := ErrorsStatusCodes[err]
	if !ok {
		return http.StatusInternalServerError, "Internal Server Error"
	}

	return code, err.Error()
}
