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
