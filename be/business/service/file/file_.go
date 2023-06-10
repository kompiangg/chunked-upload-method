package file

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/kompiangg/shipper-fp/business/entity"
	"github.com/kompiangg/shipper-fp/pkg/errors"
)

func (f *service) PreProcessUploadFile(ctx context.Context, v entity.UploadFileMetadata) (res entity.UploadFileMetaDataResponse, err error) {
	err = f.validator.Validate(&v)
	if err != nil {
		return res, errors.Wrap(err, "PreProcessUploadFile")
	}

	v.ChunkSize = f.config.UploadFileConfig.ChunkSize
	v.ChunkCount = (v.Size / f.config.UploadFileConfig.ChunkSize) + 1

	uniqueName, err := uuid.NewUUID()
	if err != nil {
		return res, errors.Wrap(err, "PreProcessUploadFile")
	}

	extensions := filepath.Ext(v.OriginName)
	v.UniqueName = uniqueName.String() + extensions

	err = f.fileRepo.InsertMetadata(ctx, v)
	if err != nil {
		return res, errors.Wrap(err, "PreProcessUploadFile")
	}

	res.IdentityName = uniqueName.String()

	res = entity.UploadFileMetaDataResponse{
		IdentityName:  v.UniqueName,
		ChunkByteSize: v.ChunkSize,
		ChunkCount:    v.ChunkCount,
	}

	return res, nil
}

func (f *service) InsertingChunkByteCode(ctx context.Context, v entity.UploadBinaryFile) error {
	err := f.fileRepo.InsertByteCode(ctx, v)
	if err != nil {
		return errors.Wrap(err, "InsertingChunkBinaryCode")
	}

	return nil
}

func (f *service) AssembleByteCodeToFile(ctx context.Context, v entity.AssembleByteCode) error {
	metadata, err := f.fileRepo.GetMetadata(ctx, entity.GetFileMetadata{
		IdentityName: v.UniqueName,
	}, true)
	if err != nil {
		return errors.Wrap(err, "AssembleByteCodeToFile")
	}

	byteCode, err := f.fileRepo.GetAllByteCode(ctx, v)
	if err != nil {
		return errors.Wrap(err, "AssembleByteCodeToFile")
	}

	if len(byteCode.FileByteCode) != metadata.ChunkCount {
		return errors.New(errors.ErrValidation)
	}

	orderedByteCode := make([]entity.FileByteCode, len(byteCode.FileByteCode))

	for idx := range byteCode.FileByteCode {
		orderedByteCode[byteCode.FileByteCode[idx].Order] = byteCode.FileByteCode[idx]
	}

	filePath := fmt.Sprintf("%s/%s", f.config.UploadFolderPath, metadata.UniqueName)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Wrap(err, "AssembleByteCodeToFile")
	}

	for idx := range orderedByteCode {
		file.Write(orderedByteCode[idx].ByteCode)
	}

	defer file.Close()

	url, err := f.uploadRepo.UploadFile(ctx, filePath)
	if err != nil {
		return errors.Wrap(err, "AssembleByteCodeToFile")
	}

	err = os.Remove(filePath)
	if err != nil {
		errors.Wrap(err, "AssembleByteCodeToFile")
		errors.ErrorStack(err)
		err = nil
	}

	err = f.fileRepo.InsertUploadedFile(ctx, metadata.OriginName, url)
	if err != nil {
		return errors.Wrap(err, "AssembleByteCodeToFile")
	}

	return nil
}

func (f *service) GetUploadedFile(ctx context.Context) ([]entity.File, error) {
	res, err := f.fileRepo.GetUploadedFile(ctx)
	if err != nil {
		return res, errors.Wrap(err, "GetUploadedFile")
	}

	return res, nil
}

func (f *service) OldMethodUploadFile(ctx context.Context, fileName string, src io.Reader) error {
	filePath := fmt.Sprintf("%s/%s", f.config.UploadFolderPath, fileName)

	dstFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Wrap(err, "OldMethodUploadFile")
	}

	_, err = io.Copy(dstFile, src)
	if err != nil {
		return errors.Wrap(err, "OldMethodUploadFile")
	}

	dstFile.Close()

	url, err := f.uploadRepo.UploadFile(ctx, filePath)
	if err != nil {
		return errors.Wrap(err, "AssembleByteCodeToFile")
	}

	err = os.Remove(filePath)
	if err != nil {
		errors.Wrap(err, "AssembleByteCodeToFile")
		errors.ErrorStack(err)
		err = nil
	}

	err = f.fileRepo.InsertUploadedFile(ctx, fileName, url)
	if err != nil {
		return errors.Wrap(err, "AssembleByteCodeToFile")
	}

	return nil
}
