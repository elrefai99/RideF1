package auth

import (
	worker "github.com/elrefai99/RideF1/app/api/cmd/Worker"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(
	api *gin.RouterGroup,
	db *gorm.DB,
	worker *worker.Queue,
) {
	authRouter := api.Group("/auth")
	{
		authRouter.POST("/register", RegisterHandler(db, worker))
		authRouter.POST("/login", LoginHandler(db, worker))
		authRouter.POST("/refresh", RefreshHandler(db, worker))
		authRouter.GET("/logout", LogoutHandler())
	}
}
