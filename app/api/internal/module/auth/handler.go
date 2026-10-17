package auth

import (
	"log"
	"os"
	"time"

	worker "github.com/elrefai99/RideF1/app/api/cmd/Worker"
	"github.com/elrefai99/RideF1/app/api/internal/Common/network"
	"github.com/elrefai99/RideF1/app/api/internal/Queue/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func RegisterHandler(db *gorm.DB, worker *worker.Queue) gin.HandlerFunc {
	service := AuthService(db)
	return func(ctx *gin.Context) {

		var body CreateUserRequest

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

		payload := IPayload{
			Uid:      "",
			User_id:  0,
			Email:    body.Email,
			Username: "",
		}

		token, err := GenerateTokens(ctx, db, payload, PendingToken)
		if err != nil {
			log.Println(err)
		}

		ctx.SetCookie("access_token", token, int((1 * 24 * time.Minute).Seconds()), "/", "", true, true)
		network.Created(ctx, token)
	}
}

// Login Handler
func LoginHandler(db *gorm.DB, worker *worker.Queue) gin.HandlerFunc {
	service := AuthService(db)
	return func(ctx *gin.Context) {

		var body LoginRequest

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

		payload := IPayload{
			Uid:      data.Uuid.String(),
			User_id:  data.ID,
			Email:    body.Email,
			Username: data.Username,
		}

		accessToken, err := GenerateTokens(ctx, db, payload, AccessToken)
		if err != nil {
			log.Println(err)
		}
		refreshToken, err := GenerateTokens(ctx, db, payload, RefreshToken)
		if err != nil {
			log.Println(err)
		}

		ctx.SetCookie("access_token", accessToken, int((1 * 24 * time.Minute).Seconds()), "/", "", true, true)
		ctx.SetCookie("refresh_token", refreshToken, int((1 * 24 * time.Minute).Seconds()), "/", "", true, true)
		ctx.SetCookie("_atf1", accessToken, int((1 * 24 * time.Minute).Seconds()), "/", "", true, true)
		ctx.SetCookie("_rtf1", refreshToken, int((1 * 24 * time.Minute).Seconds()), "/", "", true, true)
		network.OK(ctx, accessToken)
	}
}

// Refresh Handler
func RefreshHandler(db *gorm.DB, worker *worker.Queue) gin.HandlerFunc {
	service := AuthService(db)
	return func(ctx *gin.Context) {
		var body TokenRequest
		if err := ctx.ShouldBindJSON(&body); err != nil {
			network.BadRequest(ctx, "Please provide a token")
			return
		}
		token := body.Token
		if token == "" {
			network.Unauthorized(ctx, "Authentication failed")
			return
		}

		secret := os.Getenv("REFRESH_TOKEN")
		if secret == "" {
			network.Unauthorized(ctx, "This token is invalid or has expired. Please request a new one")
			return
		}

		claims := jwt.MapClaims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !parsedToken.Valid || claims["type"] != "refresh_token" {
			network.Unauthorized(ctx, "This token is invalid or has expired. Please request a new one")
			return
		}

		dataToken, ok := claims["data"].(map[string]any)
		if !ok {
			network.Forbidden(ctx, "The server is refusing to give the requested resource")
			return
		}

		idValue, ok := dataToken["data"]
		if !ok {
			network.Forbidden(ctx, "The server is refusing to give the requested resource")
			return
		}

		data, err := service.RefreshService(ctx, idValue)
		if err != nil {
			network.BadRequest(ctx, err.Error())
			return
		}

		payload := IPayload{
			Uid:      data.Uuid.String(),
			User_id:  data.ID,
			Email:    data.Email,
			Username: data.Username,
		}

		accessToken, err := GenerateTokens(ctx, db, payload, AccessToken)
		if err != nil {
			log.Println(err)
		}
		refreshToken, err := GenerateTokens(ctx, db, payload, RefreshToken)
		if err != nil {
			log.Println(err)
		}

		ctx.SetCookie("access_token", accessToken, int((1 * 24 * time.Minute).Seconds()), "/", "", true, true)
		ctx.SetCookie("refresh_token", refreshToken, int((1 * 24 * time.Minute).Seconds()), "/", "", true, true)
		ctx.SetCookie("_atf1", accessToken, int((1 * 24 * time.Minute).Seconds()), "/", "", true, true)
		ctx.SetCookie("_rtf1", refreshToken, int((1 * 24 * time.Minute).Seconds()), "/", "", true, true)
		network.OK(ctx, accessToken)
	}
}

// Logout Handler
func LogoutHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.SetCookie("access_token", "", -1, "/", "", false, true)
		ctx.SetCookie("refresh_token", "", -1, "/", "", false, true)
		ctx.SetCookie("_atf1", "", -1, "/", "", false, true)
		ctx.SetCookie("_rtf1", "", -1, "/", "", false, true)
		network.OK(ctx, "Success Remove Tokens")
	}
}
