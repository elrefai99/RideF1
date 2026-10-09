package server

import "github.com/gin-gonic/gin"

func NewRouters(
	r *gin.Engine,
) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"status":  "ok",
		})
	})

	api := r.Group("/api")
	auth := api.Group("/auth")
	{
		auth.POST("/login")
		auth.POST("/logout")
	}
}
