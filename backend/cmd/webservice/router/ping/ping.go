package ping

import (
	"github.com/kompiangg/shipper-fp/business/service"
	"github.com/kompiangg/shipper-fp/cmd/webservice/router/ping/handler"
	"github.com/kompiangg/shipper-fp/config"
	"github.com/labstack/echo/v4"
)

func InitHandler(
	e *echo.Echo,
	service service.Service,
	config config.Config,
) {
	pingHandler := handler.InitPingHandler(e, service.Ping, config)

	e.GET(PingPathV1, pingHandler.Ping())
}
