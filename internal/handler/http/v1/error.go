package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidRequestBody = errors.New("invalid request body")
)

type errorResponse[T string | map[string]string] struct {
	Error  string `json:"error"`
	Detail T      `json:"detail"`
}

func abortWithError[T string | map[string]string](c *gin.Context, code int, err error, detail T) {
	c.AbortWithStatusJSON(code, errorResponse[T]{
		Error:  err.Error(),
		Detail: detail,
	})
}
