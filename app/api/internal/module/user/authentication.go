package user

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/elrefai99/RideF1/app/api/internal/Common/network"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func ProfileMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ""
		if authorization := ctx.GetHeader("Authorization"); strings.HasPrefix(authorization, "Bearer ") {
			token = strings.TrimSpace(strings.TrimPrefix(authorization, "Bearer "))
		}

		if token == "" {
			token, _ = ctx.Cookie("access_token")
		}

		if token == "" {
			network.Unauthorized(ctx, "Authorization header is required")
			return
		}

		secret := os.Getenv("ACCESS_TOKEN")
		if secret == "" {
			network.Forbidden(ctx, "This token is invalid or has expired. Please request a new one")
			return
		}
		claims := jwt.MapClaims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !parsedToken.Valid || claims["type"] != "access_token" {
			network.Forbidden(ctx, "This token is invalid or has expired. Please request a new one")
			return
		}

		data, ok := claims["data"].(map[string]any)
		if !ok {
			network.Forbidden(ctx, "The server is refusing to give the requested resource")
			return
		}
		idValue, ok := data["User_id"]
		if !ok {
			network.Forbidden(ctx, "The server is refusing to give the requested resource")
			return
		}
		id, err := strconv.ParseUint(fmt.Sprint(idValue), 10, 64)
		if err != nil || id == 0 {
			network.Forbidden(ctx, "The server is refusing to give the requested resource")
			return
		}

		var userData User
		if err := db.WithContext(ctx.Request.Context()).Where("id = ? AND status = ?", id, UserActive).First(&userData).Error; err != nil {
			network.Forbidden(ctx, "The server is refusing to give the requested resource")
			return
		}

		ctx.Set("user", userData)
		ctx.Next()
	}
}
