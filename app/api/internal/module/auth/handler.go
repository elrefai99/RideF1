package auth

import (
	"log"
	"os"
	"time"

	"github.com/elrefai99/RideF1/app/api/internal/Common/network"
	"github.com/elrefai99/RideF1/app/api/internal/Queue"
	"github.com/elrefai99/RideF1/app/api/internal/Queue/model"
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

		emailJob := model.IEmailJob{
			Type:    "email",
			Sender:  os.Getenv("SMTP_USERNAME"),
			To:      body.Email,
			Subject: "Hi In RideF1",
			Body:    "Wellcome",
		}

		if err := worker.Add(emailJob); err != nil {
			log.Printf("could not queue confirmation email: %v", err)
		}

		token, err := GenerateTokens(body.Email, "access_token")
		if err != nil {
			log.Println(err)
		}

		ctx.SetCookie("access_token", token, int((1 * 24 * time.Minute).Seconds()), "/", "", true, true)
		network.Created(ctx, token)
	}
}

// Login Handler
func LoginHandler(db *gorm.DB, worker *Queue.Queue) gin.HandlerFunc {
	service := AuthService(db)
	return func(ctx *gin.Context) {

		var body user.LoginRequest

		if err := ctx.ShouldBindJSON(&body); err != nil {
			network.BadRequest(ctx, "Pls enter currect data")
			return
		}
		data, err := service.LoginService(ctx, body)
		if err != nil {
			network.BadRequest(ctx, err.Error())
			return
		}

				emailJob := model.IEmailJob{
			Type:    "email",
			Sender:  os.Getenv("SMTP_USERNAME"),
			To:      body.Email,
			Subject: "Hi In RideF1",
			Body:    "Wellcome",
		}

		if err := worker.Add(emailJob); err != nil {
			log.Printf("could not queue confirmation email: %v", err)
		}

		token, err := GenerateTokens(data, "access_token")
		if err != nil {
			log.Println(err)
		}

		ctx.SetCookie("access_token", token, int((1 * 24 * time.Minute).Seconds()), "/", "", true, true)
		network.OK(ctx, token)
	}
}
