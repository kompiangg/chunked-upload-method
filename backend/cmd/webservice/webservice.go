package webservice

import (
	"fmt"

	"github.com/kompiangg/shipper-fp/business/service"
	inmiddleware "github.com/kompiangg/shipper-fp/cmd/webservice/middleware"
	"github.com/kompiangg/shipper-fp/cmd/webservice/router/ping"
	"github.com/kompiangg/shipper-fp/cmd/webservice/router/upload"
	"github.com/kompiangg/shipper-fp/config"
	"github.com/kompiangg/shipper-fp/docs"
	"github.com/kompiangg/shipper-fp/pkg/http"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitWebService(
	service service.Service,
	config config.Config,
) error {
	echoServer, err := http.InitEchoServer(&http.ServerConfig{
		Port:                 config.ServerConfig.Port,
		Environment:          config.ServerConfig.Environment,
		WhiteListAllowOrigin: config.ServerConfig.WhiteListAllowOrigin,
	})
	if err != nil {
		return err
	}

	echoServer.Echo.Use(middleware.CORS())

	middleware := inmiddleware.InitMiddleware(config)

	// Swagger
	initSwagger(config)
	echoServer.Echo.GET("/swagger/*", echoSwagger.WrapHandler, middleware.NotRunInProd())

	// Init All Router
	ping.InitHandler(echoServer.Echo, service, config)
	upload.InitHandler(echoServer.Echo, service.File, config)

	// Start HTTP Server
	err = echoServer.ServeHTTP()
	if err != nil {
		return err
	}
	return nil
}

func initSwagger(config config.Config) {
	docs.SwaggerInfo.Title = fmt.Sprintf("%s API", config.SwaggerConfig.Title)
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", config.SwaggerConfig.Host, config.SwaggerConfig.Port)
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Version = config.SwaggerConfig.Version
	docs.SwaggerInfo.Description = config.SwaggerConfig.Description
	docs.SwaggerInfo.Schemes = config.SwaggerConfig.Schemes
}
