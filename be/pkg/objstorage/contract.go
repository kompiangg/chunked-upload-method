package objstorage

import "context"

type ObjectStorageItf interface {
	Upload(ctx context.Context, FileName string) (string, error)
}
