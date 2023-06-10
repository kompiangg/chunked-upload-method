package file

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kompiangg/shipper-fp/business/entity"
	"github.com/kompiangg/shipper-fp/pkg/errors"
	dbutils "github.com/kompiangg/shipper-fp/utils/db"
)

func (r *repository) InsertMetadata(ctx context.Context, v entity.UploadFileMetadata) error {
	err := r.insertCacheMetadata(ctx, v)
	if err != nil {
		return errors.Wrap(err, "InsertMetadata")
	}

	return nil
}

func (r *repository) GetMetadata(ctx context.Context, param entity.GetFileMetadata, remove bool) (res entity.GetFileMetadataResponse, err error) {
	res, err = r.getCacheMetadata(ctx, param, remove)
	if err != nil {
		return res, errors.Wrap(err, "GetMetadata")
	}

	return res, nil
}

func (r *repository) InsertByteCode(ctx context.Context, v entity.UploadBinaryFile) error {
	err := r.insertByteCode(ctx, v)
	if err != nil {
		return errors.Wrap(err, "InsertBinaryByOrder")
	}

	return nil
}

func (r *repository) GetAllByteCode(ctx context.Context, v entity.AssembleByteCode) (res entity.AssembleByteCode, err error) {
	keyPrefix := fmt.Sprintf(PrefixFileBinary, v.UniqueName)

	iter := r.redis.Scan(ctx, 0, keyPrefix, 0).Iterator()

	for iter.Next(ctx) {
		key := iter.Val()
		splittedKey := strings.Split(key, ":")

		order, err := strconv.Atoi(splittedKey[len(splittedKey)-1])
		if err != nil {
			return res, errors.Wrap(err, "getAllValueByUniqueName")
		}

		bytes, err := r.getByteCodeByUniqueNameKey(ctx, key, true)
		if err != nil {
			return res, errors.Wrap(err, "getAllValueByUniqueName")
		}

		v.FileByteCode = append(v.FileByteCode, entity.FileByteCode{
			ByteCode: bytes,
			Order:    order,
		})
	}
	if err := iter.Err(); err != nil {
		return res, errors.Wrap(err, "getAllValueByUniqueName")
	}

	return v, nil
}

func (r *repository) InsertUploadedFile(ctx context.Context, fileName string, url string) error {
	if err := dbutils.TrxWrapper(ctx, r.db, "InsertUploadedFile", &sql.TxOptions{Isolation: sql.LevelDefault}, func(tx *sqlx.Tx) error {
		err := r.insertUploadedFile(tx, fileName, url)
		if err != nil {
			return errors.Wrap(err, "insertUploadedFile")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "InsertUploadedFile")
	}

	return nil
}

func (r *repository) GetUploadedFile(ctx context.Context) ([]entity.File, error) {
	res, err := r.getUploadedFile(ctx)
	if err != nil {
		return res, errors.Wrap(err, "GetUploadedFile")
	}

	return res, nil
}
