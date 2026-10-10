package auth

import (
	"github.com/elrefai99/RideF1/app/api/internal/Common/network"
	"github.com/elrefai99/RideF1/app/api/internal/Queue"
	"github.com/elrefai99/RideF1/app/api/internal/module/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandler(db *gorm.DB, worker *Queue.Queue) gin.HandlerFunc {
	service := AuthService(db)
	return func(ctx *gin.Context) {

		var body user.CreateUserRequest

		if err := ctx.ShouldBindJSON(&body); err != nil {
			network.BadRequest(ctx, "Pls enter currect data")
			return
		}

		if _, err := service.RegisterService(ctx, body); err != nil {
			network.BadRequest(ctx, err.Error())
			return
		}

		network.Created(ctx, "Success create new account")
	}
}
