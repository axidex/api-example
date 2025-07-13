package api

import "github.com/gin-gonic/gin"

// swagger:enum ResponseType
type ResponseType string

const (
	ResponseSuccess ResponseType = "SUCCESS"
	ResponseError   ResponseType = "ERROR"
)

type Response struct {
	Status ResponseType `json:"status"`
	Info   string       `json:"info,omitempty"`
	Data   interface{}  `json:"data,omitempty"`
}

func ResponseUserSuccess(c *gin.Context, obj interface{}, statusCode int) {
	response := Response{Status: ResponseSuccess, Data: obj}

	c.JSON(statusCode, response)
}

func ResponseUserError(c *gin.Context, err error) {
	statusCode, errorInfo := GetErrorResponse(err)
	response := Response{Status: ResponseError, Data: errorInfo}

	c.JSON(statusCode, response)
}
