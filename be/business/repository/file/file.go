package file

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kompiangg/shipper-fp/business/entity"
	"github.com/redis/go-redis/v9"
)

type RepositoryItf interface {
	GetMetadata(ctx context.Context, param entity.GetFileMetadata, remove bool) (res entity.GetFileMetadataResponse, err error)
	InsertMetadata(ctx context.Context, v entity.UploadFileMetadata) error
	GetAllByteCode(ctx context.Context, v entity.AssembleByteCode) (res entity.AssembleByteCode, err error)
	InsertByteCode(ctx context.Context, v entity.UploadBinaryFile) error
	InsertUploadedFile(ctx context.Context, fileName string, url string) error
	GetUploadedFile(ctx context.Context) ([]entity.File, error)
}

type repository struct {
	db    *sqlx.DB
	redis *redis.Client
}

func InitRepository(
	db *sqlx.DB,
	redis *redis.Client,
) (repository, error) {
	return repository{
		db:    db,
		redis: redis,
	}, nil
}
