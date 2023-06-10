package handler

import (
	"net/http"

	"github.com/kompiangg/shipper-fp/business/service/ping"
	"github.com/kompiangg/shipper-fp/config"
	httputils "github.com/kompiangg/shipper-fp/utils/http"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type pingHandler struct {
	service ping.PingItf
	config  config.Config
	e       *echo.Echo
	log     *logrus.Entry
}

func InitPingHandler(
	e *echo.Echo,
	service ping.PingItf,
	config config.Config,
) pingHandler {
	return pingHandler{
		e:       e,
		service: service,
		config:  config,
	}
}

// Ping
//
//	@Tags			Ping
//	@Description	Ping
//	@Summary		Ping
//
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	handler.HTTPPingResp
//	@Router			/v1/ping [get]
func (p *pingHandler) Ping() echo.HandlerFunc {
	return func(c echo.Context) error {
		return httputils.WriteResponse(c, http.StatusOK, p.service.Ping())
	}
}
