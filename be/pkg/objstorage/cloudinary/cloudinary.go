package cloudinary

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryConfig struct {
	APIKey    string
	APISecret string
	CloudName string
}

type cloudinaryObjStorage struct {
	cld *cloudinary.Cloudinary
}

func InitCloudinary(config CloudinaryConfig) (cloudinaryObjStorage, error) {
	var err error
	var cld cloudinaryObjStorage

	url := fmt.Sprintf("cloudinary://%s:%s@%s", config.APIKey, config.APISecret, config.CloudName)

	cld.cld, err = cloudinary.NewFromURL(url)
	if err != nil {
		return cld, err
	}

	cld.cld.Config.URL.Secure = true

	return cld, nil
}

func (c cloudinaryObjStorage) Upload(ctx context.Context, filePath string) (string, error) {
	resp, err := c.cld.Upload.Upload(ctx, filePath, uploader.UploadParams{
		UniqueFilename: api.Bool(true),
		ResourceType:   "auto",
	})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
