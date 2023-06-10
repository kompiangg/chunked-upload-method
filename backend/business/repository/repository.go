package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kompiangg/shipper-fp/business/repository/file"
	"github.com/kompiangg/shipper-fp/business/repository/upload"
	"github.com/kompiangg/shipper-fp/config"
	"github.com/kompiangg/shipper-fp/pkg/objstorage"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	Upload upload.RepositoryItf
	File   file.RepositoryItf
}

func InitRepository(
	config config.Config,
	db *sqlx.DB,
	redis *redis.Client,
	cld objstorage.ObjectStorageItf,
) (Repository, error) {
	uploadDom, err := upload.InitRepository(cld)
	if err != nil {
		return Repository{}, err
	}

	fileDom, err := file.InitRepository(db, redis)
	if err != nil {
		return Repository{}, err
	}

	return Repository{
		Upload: &uploadDom,
		File:   &fileDom,
	}, nil
}
