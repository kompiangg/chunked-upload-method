package handler

import (
	"io"
	"net/http"
	"strconv"

	"github.com/kompiangg/shipper-fp/business/entity"
	"github.com/kompiangg/shipper-fp/pkg/errors"
	httputils "github.com/kompiangg/shipper-fp/utils/http"
	"github.com/labstack/echo/v4"
)

func (u *uploadHandler) CreateUploadMetadata() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.UploadFileMetadata
		if err := c.Bind(&req); err != nil {
			return httputils.WriteErrorResponse(c, errors.ErrBadRequest, nil)
		}

		res, err := u.fileService.PreProcessUploadFile(c.Request().Context(), req)
		if err != nil {
			return httputils.WriteErrorResponse(c, err, nil)
		}

		return httputils.WriteResponse(c, http.StatusCreated, res)
	}
}

func (u *uploadHandler) InsertChunkByteData() echo.HandlerFunc {
	return func(c echo.Context) error {
		body := c.Request().Body
		header := c.Request().Header

		identityName := header.Get(entity.IdentityNameHeader)
		chunkOrderStr := header.Get(entity.ChunkOrderHeader)

		chunkOrder, err := strconv.Atoi(chunkOrderStr)
		if err != nil {
			return httputils.WriteErrorResponse(c, errors.ErrBadRequest, "chunk order need to be number")
		}

		byteCodeData, err := io.ReadAll(body)
		if err != nil {
			return httputils.WriteErrorResponse(c, errors.ErrInternalServer, nil)
		}

		err = u.fileService.InsertingChunkByteCode(c.Request().Context(), entity.UploadBinaryFile{
			IdentityName: identityName,
			Order:        chunkOrder,
			ByteCodeData: byteCodeData,
		})
		if err != nil {
			return httputils.WriteErrorResponse(c, errors.ErrInternalServer, nil)
		}

		return httputils.WriteResponse(c, http.StatusCreated, nil)
	}
}

func (u *uploadHandler) AssembleByteData() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.AssembleByteCode
		if err := c.Bind(&req); err != nil {
			return httputils.WriteErrorResponse(c, errors.ErrBadRequest, nil)
		}

		err := u.fileService.AssembleByteCodeToFile(c.Request().Context(), entity.AssembleByteCode{
			UniqueName: req.UniqueName,
		})
		if err != nil {
			return httputils.WriteErrorResponse(c, err, nil)
		}

		return httputils.WriteResponse(c, http.StatusCreated, nil)
	}
}

func (u *uploadHandler) GetUploadedFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := u.fileService.GetUploadedFile(c.Request().Context())
		if err != nil {
			return httputils.WriteErrorResponse(c, err, nil)
		}

		return httputils.WriteResponse(c, http.StatusOK, res)
	}
}

func (u *uploadHandler) OldMethodUpload() echo.HandlerFunc {
	return func(c echo.Context) error {
		uploadedFile, err := c.FormFile(entity.FileFormKey)
		if err != nil || uploadedFile == nil {
			return httputils.WriteErrorResponse(c, errors.ErrBadRequest, nil)
		}

		openedUploadedFile, err := uploadedFile.Open()
		if err != nil {
			return httputils.WriteErrorResponse(c, errors.ErrInternalServer, nil)
		}

		err = u.fileService.OldMethodUploadFile(c.Request().Context(), uploadedFile.Filename, openedUploadedFile)
		if err != nil {
			return httputils.WriteErrorResponse(c, err, nil)
		}

		return httputils.WriteResponse(c, http.StatusCreated, nil)
	}
}
