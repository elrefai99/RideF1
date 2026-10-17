package user

import (
	"github.com/cloudinary/cloudinary-go/v2"
	worker "github.com/elrefai99/RideF1/app/api/cmd/Worker"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(
	api *gin.RouterGroup,
	db *gorm.DB,
	worker *worker.Queue,
	cld *cloudinary.Cloudinary,
) {
	userRouter := api.Group("/user")
	userRouter.Use(ProfileMiddleware(db))
	{
		userRouter.GET("/me", UserProfileHandler(db, worker))
		userRouter.PUT("/edit/:id", UpdateAccountHandler(db, cld))
	}
}
