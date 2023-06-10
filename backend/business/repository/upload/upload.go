package upload

import (
	"context"

	"github.com/kompiangg/shipper-fp/pkg/objstorage"
)

type RepositoryItf interface {
	UploadFile(ctx context.Context, filePath string) (url string, err error)
}

type repository struct {
	cld objstorage.ObjectStorageItf
}

func InitRepository(
	cld objstorage.ObjectStorageItf,
) (repository, error) {
	return repository{
		cld: cld,
	}, nil
}
