package middleware

import (
	"github.com/kompiangg/shipper-fp/config"
	"github.com/kompiangg/shipper-fp/pkg/errors"
	httputils "github.com/kompiangg/shipper-fp/utils/http"
	"github.com/labstack/echo/v4"
)

type middleware struct {
	config config.Config
}

func InitMiddleware(
	config config.Config,
) middleware {
	return middleware{
		config: config,
	}
}

func (m *middleware) NotRunInProd() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if m.config.ServerConfig.Environment == "prod" {
				return httputils.WriteErrorResponse(c, errors.ErrUnauthorized, nil)
			}

			return next(c)
		}
	}
}
