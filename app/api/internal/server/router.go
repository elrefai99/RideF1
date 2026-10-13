package server

import (
	"github.com/elrefai99/RideF1/app/api/internal/Queue"
	"github.com/elrefai99/RideF1/app/api/internal/middleware"
	"github.com/elrefai99/RideF1/app/api/internal/module/auth"
	"github.com/elrefai99/RideF1/app/api/internal/module/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouters(
	r *gin.Engine,
	db *gorm.DB,
	worker *Queue.Queue,
) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"status":  "ok",
		})
	})

	api := r.Group("/api")
	authRouter := api.Group("/auth")
	{
		authRouter.POST("/register", auth.RegisterHandler(db, worker))
		authRouter.POST("/login", auth.LoginHandler(db, worker))
		authRouter.POST("/refresh", auth.RefreshHandler(db, worker))
		authRouter.GET("/logout", auth.LogoutHandler())
	}

	userRouter := api.Group("/user")
	userRouter.Use(middleware.ProfileMiddleware(db))
	{
		userRouter.GET("/me", user.UserProfileHandler(db, worker))
	}
}
