package file

import (
	"context"
	"io"

	"github.com/kompiangg/shipper-fp/business/entity"
	"github.com/kompiangg/shipper-fp/business/repository/file"
	"github.com/kompiangg/shipper-fp/business/repository/upload"
	"github.com/kompiangg/shipper-fp/config"
	"github.com/kompiangg/shipper-fp/pkg/validator"
)

type ServiceItf interface {
	PreProcessUploadFile(ctx context.Context, v entity.UploadFileMetadata) (res entity.UploadFileMetaDataResponse, err error)
	InsertingChunkByteCode(ctx context.Context, v entity.UploadBinaryFile) error
	AssembleByteCodeToFile(ctx context.Context, v entity.AssembleByteCode) error
	GetUploadedFile(ctx context.Context) ([]entity.File, error)
	OldMethodUploadFile(ctx context.Context, fileName string, src io.Reader) error
}

type service struct {
	config     config.Config
	validator  validator.ValidatorItf
	fileRepo   file.RepositoryItf
	uploadRepo upload.RepositoryItf
}

func InitService(
	config config.Config,
	validator validator.ValidatorItf,
	fileRepo file.RepositoryItf,
	uploadRepo upload.RepositoryItf,
) (service, error) {
	return service{
		config:     config,
		validator:  validator,
		fileRepo:   fileRepo,
		uploadRepo: uploadRepo,
	}, nil
}
