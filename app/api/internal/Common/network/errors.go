package network

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("not found")
	ErrBadRequest   = errors.New("bad request")
)

func HandleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrUnauthorized):
		Unauthorized(c)

	case errors.Is(err, ErrNotFound):
		NotFound(c)

	case errors.Is(err, ErrBadRequest):
		BadRequest(c, err.Error())

	default:
		Internal(c, err)
	}
}
