package auth

import (
	"github.com/elrefai99/RideF1/app/api/internal/Queue"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(
	api *gin.RouterGroup,
	db *gorm.DB,
	worker *Queue.Queue,
) {
	authRouter := api.Group("/auth")
	{
		authRouter.POST("/register", RegisterHandler(db, worker))
		authRouter.POST("/login", LoginHandler(db, worker))
		authRouter.POST("/refresh", RefreshHandler(db, worker))
		authRouter.GET("/logout", LogoutHandler())
	}
}
