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
	Data      any       `json:"data,omitempty"`
	Error     string    `json:"error,omitempty"`
}

func Created(g *gin.Context, data any) {
	g.JSON(
		http.StatusCreated,
		Response{
			Code:      201,
			Status:    "Created",
			Success:   true,
			TimeStamp: time.Now(),
			Data:      data,
		},
	)
}

func BadRequest(g *gin.Context, message string) {
	g.JSON(
		http.StatusBadRequest,
		Response{
			Code:      400,
			Status:    "Bad Request",
			Success:   false,
			TimeStamp: time.Now(),
			Error:     message,
		},
	)
}

func Unauthorized(c *gin.Context) {
	c.JSON(
		http.StatusUnauthorized,
		Response{
			Code:    400,
			Status:  "unauthorized",
			Success: false,
			Error:   "unauthorized",
		})
}

func NotFound(c *gin.Context) {
	c.JSON(404, Response{
		Success: false,
		Error:   "not found",
	})
}

func Internal(c *gin.Context) {
	c.JSON(500, Response{
		Success: false,
		Error:   "internal server error",
	})
}
