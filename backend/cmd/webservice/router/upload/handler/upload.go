package handler

import (
	"github.com/kompiangg/shipper-fp/business/service/file"
	"github.com/kompiangg/shipper-fp/config"
	"github.com/labstack/echo/v4"
)

type uploadHandler struct {
	fileService file.ServiceItf
	config      config.Config
	e           *echo.Echo
}

func InitUploadHandler(
	e *echo.Echo,
	fileService file.ServiceItf,
	config config.Config,
) uploadHandler {
	return uploadHandler{
		e:           e,
		fileService: fileService,
		config:      config,
	}
}
