package upload

import (
	"github.com/kompiangg/shipper-fp/business/service/file"
	"github.com/kompiangg/shipper-fp/cmd/webservice/router/upload/handler"
	"github.com/kompiangg/shipper-fp/config"
	"github.com/labstack/echo/v4"
)

func InitHandler(
	e *echo.Echo,
	fileService file.ServiceItf,
	config config.Config,
) {
	uploadHandler := handler.InitUploadHandler(e, fileService, config)

	uploadRouterV1 := e.Group("/v1")
	uploadRouterV2 := e.Group("/v2")

	uploadRouterV1.POST(UploadPath, uploadHandler.OldMethodUpload())
	uploadRouterV1.GET(GetUploadedFilePath, uploadHandler.GetUploadedFile())

	uploadRouterV2.POST(RequestUploadPath, uploadHandler.CreateUploadMetadata())
	uploadRouterV2.POST(UploadPath, uploadHandler.InsertChunkByteData())
	uploadRouterV2.POST(FinishUploadPath, uploadHandler.AssembleByteData())
}
