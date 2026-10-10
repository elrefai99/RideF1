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
	g.JSON(
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

func Unauthorized(c *gin.Context) {
	c.JSON(
		http.StatusUnauthorized,
		Response{
			Code:    400,
			Status:  "unauthorized",
			Success: false,
			Message: "unauthorized",
		})
}

func NotFound(c *gin.Context) {
	c.JSON(404, Response{
		Success: false,
		Error:   "not found",
	})
}

func Internal(c *gin.Context, err error) {
	c.JSON(500, Response{
		Success: false,
		Error:   err.Error(),
	})
}
