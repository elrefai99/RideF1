package upload

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
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

type SignedImageUpload struct {
	CloudName    string `json:"cloud_name"`
	APIKey       string `json:"api_key"`
	Timestamp    int64  `json:"timestamp"`
	Folder       string `json:"folder"`
	PublicID     string `json:"public_id"`
	ImageURL     string `json:"image_url"`
	UploadPreset string `json:"upload_preset"`
	Overwrite    bool   `json:"overwrite"`
	Signature    string `json:"signature"`
	UploadURL    string `json:"upload_url"`
}

func CreateSignedImageUpload(
	cld *cloudinary.Cloudinary,
	folder string,
	uploadPreset string,
	publicID string,
) (SignedImageUpload, error) {
	timestamp := time.Now().Unix()
	params := url.Values{
		"folder":    {folder},
		"overwrite": {"true"},
		"public_id": {publicID},
		"timestamp": {strconv.FormatInt(timestamp, 10)},
	}
	if uploadPreset != "" {
		params.Set("upload_preset", uploadPreset)
	}
	signature, err := api.SignParameters(params, cld.Config.Cloud.APISecret)

	if err != nil {
		return SignedImageUpload{}, err
	}

	cloudName := cld.Config.Cloud.CloudName
	return SignedImageUpload{
		CloudName:    cloudName,
		APIKey:       cld.Config.Cloud.APIKey,
		Timestamp:    timestamp,
		Folder:       folder,
		PublicID:     publicID,
		ImageURL:     fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/%s", cloudName, publicID),
		UploadPreset: uploadPreset,
		Overwrite:    true,
		Signature:    signature,
		UploadURL:    fmt.Sprintf("https://api.cloudinary.com/v1_1/%s/image/upload", cloudName),
	}, nil
}
