package file

import (
	"context"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/kompiangg/shipper-fp/business/entity"
	x "github.com/kompiangg/shipper-fp/pkg/errors"
	"github.com/kompiangg/shipper-fp/utils/compress"
)

func (r *repository) insertCacheMetadata(ctx context.Context, v entity.UploadFileMetadata) error {
	marshalledMetadata, err := json.Marshal(v)
	if err != nil {
		return x.Wrap(err, "insertCacheMetadata")
	}

	encodedJSON := compress.Encode(marshalledMetadata)
	key := fmt.Sprintf(FileMetadata, v.UniqueName)

	err = r.redis.Set(ctx, key, encodedJSON, DurationFileMetadataExpiration).Err()
	if err != nil {
		return x.Wrap(err, "insertCacheMetadata")
	}

	return nil
}

func (r *repository) getCacheMetadata(ctx context.Context, param entity.GetFileMetadata, remove bool) (res entity.GetFileMetadataResponse, err error) {
	key := fmt.Sprintf(FileMetadata, param.IdentityName)

	encodedValue, err := r.redis.Get(ctx, key).Bytes()
	if err != nil {
		return res, x.Wrap(err, "getCacheMetadata")
	}

	decodedJSON, err := compress.Decode(encodedValue)
	if err != nil {
		return res, x.Wrap(err, "getCacheMetadata")
	}

	err = json.Unmarshal(decodedJSON, &res)
	if err != nil {
		return res, x.Wrap(err, "getCacheMetadata")
	}

	// if remove {
	// 	if err := r.redis.Del(ctx, key).Err(); err != nil {
	// 		errors.ErrorStack(err)
	// 	}
	// }

	return res, nil
}

func (r *repository) insertByteCode(ctx context.Context, v entity.UploadBinaryFile) error {
	encodedBytes := compress.Encode(v.ByteCodeData)
	key := fmt.Sprintf(FileBinary, v.IdentityName, v.Order)

	err := r.redis.Set(ctx, key, encodedBytes, DurationFileMetadataExpiration).Err()
	if err != nil {
		return x.Wrap(err, "insertByteCode")
	}

	return nil
}

func (r *repository) getByteCodeByUniqueNameKey(ctx context.Context, key string, remove bool) (res []byte, err error) {
	bytesData, err := r.redis.Get(ctx, key).Bytes()
	if err != nil {
		return res, x.Wrap(err, "getByteCode")
	}

	bytes, err := compress.Decode(bytesData)
	if err != nil {
		return res, x.Wrap(err, "getByteCode")
	}

	// if remove {
	// 	if err := r.redis.Del(ctx, key).Err(); err != nil {
	// 		errors.ErrorStack(err)
	// 	}
	// }

	return bytes, nil
}
