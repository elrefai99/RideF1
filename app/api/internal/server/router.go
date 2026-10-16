package server

import (
	worker "github.com/elrefai99/RideF1/app/api/cmd/Worker"
	"github.com/elrefai99/RideF1/app/api/internal/module/auth"
	"github.com/elrefai99/RideF1/app/api/internal/module/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouters(
	r *gin.Engine,
	db *gorm.DB,
	worker *worker.Queue,
) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"status":  "ok",
		})
	})
	api := r.Group("/api")

	auth.AuthRouter(api, db, worker)
	user.UserRouter(api, db, worker)

}
