package user

import (
	worker "github.com/elrefai99/RideF1/app/api/cmd/Worker"
	"github.com/elrefai99/RideF1/app/api/internal/Common/network"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// User Profile
func UserProfileHandler(db *gorm.DB, worker *worker.Queue) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, ok := ctx.Get("user")
		if !ok {
			network.Unauthorized(ctx, "Authentication failed")
			return
		}

		currentUser, ok := value.(User)
		if !ok {
			network.Internal(ctx, gorm.ErrInvalidData)
			return
		}

		network.OK(ctx, currentUser)
	}
}

func UpdateAccountHandler(db *gorm.DB, worker *worker.Queue) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		image, err := ctx.FormFile("img")
		if err != nil {
			ctx.JSON(400, gin.H{"error": "image field is required"})
			return
		}

		ctx.JSON(200, gin.H{
			"filename": image.Filename,
			"size":     image.Size,
		})
	}
}
