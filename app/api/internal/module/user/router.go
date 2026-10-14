package user

import (
	"github.com/elrefai99/RideF1/app/api/internal/Queue"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(
	api *gin.RouterGroup,
	db *gorm.DB,
	worker *Queue.Queue,
) {
	userRouter := api.Group("/user")
	userRouter.Use(ProfileMiddleware(db))
	{
		userRouter.GET("/me", UserProfileHandler(db, worker))
	}
}
