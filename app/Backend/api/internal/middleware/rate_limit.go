package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimit(r rate.Limit, burst int) gin.HandlerFunc {
	limiter := rate.NewLimiter(
		r,
		burst,
	)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				gin.H{
					"error": "rate limit exceeded",
				},
			)
			return
		}

		c.Next()
	}
}
