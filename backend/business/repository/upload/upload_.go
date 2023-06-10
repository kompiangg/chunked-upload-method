package upload

import (
	"context"

	"github.com/kompiangg/shipper-fp/pkg/errors"
)

func (u *repository) UploadFile(ctx context.Context, filePath string) (url string, err error) {
	url, err = u.cld.Upload(ctx, filePath)
	if err != nil {
		return "", errors.Wrap(err, "UploadFile")
	}

	return url, nil
}
