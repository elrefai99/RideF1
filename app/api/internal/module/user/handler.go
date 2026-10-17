package user

import (
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	worker "github.com/elrefai99/RideF1/app/api/cmd/Worker"
	"github.com/elrefai99/RideF1/app/api/internal/Common/function/upload"
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

func UpdateAccountHandler(db *gorm.DB, cld *cloudinary.Cloudinary) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, exists := ctx.Get("user")
		currentUser, ok := value.(User)
		if !exists || !ok {
			network.Unauthorized(ctx, "Authentication failed")
			return
		}

		file, err := ctx.FormFile("img")
		if err != nil {
			network.BadRequest(ctx, "Image file is required")
			return
		}
		if file.Size > 5<<20 {
			network.BadRequest(ctx, "Image must be 5 MB or smaller")
			return
		}

		if _, err := upload.ExtImages(file.Filename); err != nil {
			network.BadRequest(ctx, err.Error())
			return
		}

		src, err := file.Open()
		if err != nil {
			network.BadRequest(ctx, "Could not read image")
			return
		}
		defer src.Close()

		result, err := cld.Upload.Upload(
			ctx.Request.Context(),
			src,
			uploader.UploadParams{
				Folder:   "users",
				PublicID: fmt.Sprintf("user-%s", currentUser.Uuid),
			},
		)
		if err != nil {
			network.Internal(ctx, err)
			return
		}

		network.OK(ctx, gin.H{"image_url": result.SecureURL})
	}
}
