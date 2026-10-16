package upload

import (
	"context"
	"errors"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func ExtImages(filename string) (string, error) {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".webp":
	default:
		return "", errors.New("Image must be JPG, PNG, or WEBP")
	}

	return ext, nil
}

func UploadImage(
	ctx context.Context,
	cld *cloudinary.Cloudinary,
	folder string,
	publicID string,
	file any,
) (string, error) {
	result, err := cld.Upload.Upload(
		ctx,
		file,
		uploader.UploadParams{
			Folder:       folder,
			PublicID:     publicID,
			ResourceType: "image",
		},
	)
	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}
