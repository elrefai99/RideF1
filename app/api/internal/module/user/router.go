package user

import (
	worker "github.com/elrefai99/RideF1/app/api/cmd/Worker"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(
	api *gin.RouterGroup,
	db *gorm.DB,
	worker *worker.Queue,
) {
	userRouter := api.Group("/user")
	userRouter.Use(ProfileMiddleware(db))
	{
		userRouter.GET("/me", UserProfileHandler(db, worker))
	}
}
