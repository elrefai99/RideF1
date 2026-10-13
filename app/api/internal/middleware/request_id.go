package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Request_ID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req_Id := ctx.GetHeader("X-Request-ID")

		if req_Id == "" {
			req_Id = uuid.NewString()
		}
		ctx.Set("req_Id", req_Id)
		ctx.Header("X-Request-ID", req_Id)

		ctx.Next()
	}
}
