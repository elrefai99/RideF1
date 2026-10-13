package user

import (
	models "github.com/elrefai99/RideF1/Database/model"
	"github.com/elrefai99/RideF1/app/api/internal/Common/network"
	"github.com/elrefai99/RideF1/app/api/internal/Queue"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// User Profile
func UserProfileHandler(db *gorm.DB, worker *Queue.Queue) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, ok := ctx.Get("user")
		if !ok {
			network.Unauthorized(ctx, "Authentication failed")
			return
		}

		currentUser, ok := value.(models.User)
		if !ok {
			network.Internal(ctx, gorm.ErrInvalidData)
			return
		}

		network.OK(ctx, currentUser)
	}
}
