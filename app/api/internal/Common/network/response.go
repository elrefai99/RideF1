package network

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Success   bool      `json:"success"`
	TimeStamp time.Time `json:"timeStamp"`
	Message   string    `json:"message"`
	Data      any       `json:"data,omitempty"`
	Error     string    `json:"error,omitempty"`
}

func OK(g *gin.Context, data any) {
	g.JSON(
		http.StatusOK,
		Response{
			Code:      200,
			Status:    "OK",
			Success:   true,
			TimeStamp: time.Now(),
			Message:   "Success Request",
			Data:      data,
		},
	)
}

func Created(g *gin.Context, data any) {
	g.JSON(
		http.StatusCreated,
		Response{
			Code:      201,
			Status:    "Created",
			Success:   true,
			TimeStamp: time.Now(),
			Message:   "Success Create",
			Data:      data,
		},
	)
}

func BadRequest(g *gin.Context, message string) {
	g.AbortWithStatusJSON(
		http.StatusBadRequest,
		Response{
			Code:      400,
			Status:    "Bad Request",
			Success:   false,
			TimeStamp: time.Now(),
			Message:   message,
		},
	)
}

func Unauthorized(c *gin.Context, message ...string) {
	responseMessage := ""
	if len(message) > 0 {
		responseMessage = message[0]
	}

	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		Response{
			Code:      http.StatusUnauthorized,
			Status:    "Unauthorized",
			Success:   false,
			TimeStamp: time.Now(),
			Message:   responseMessage,
		})
}

func Forbidden(c *gin.Context, message ...string) {
	responseMessage := ""
	if len(message) > 0 {
		responseMessage = message[0]
	}

	c.AbortWithStatusJSON(
		http.StatusForbidden,
		Response{
			Code:      http.StatusForbidden,
			Status:    "Forbidden",
			Success:   false,
			TimeStamp: time.Now(),
			Message:   responseMessage,
		})
}

func NotFound(c *gin.Context) {
	c.AbortWithStatusJSON(404, Response{
		Success: false,
		Error:   "not found",
	})
}

func Internal(c *gin.Context, err error) {
	c.AbortWithStatusJSON(500, Response{
		Success: false,
		Error:   err.Error(),
	})
}
