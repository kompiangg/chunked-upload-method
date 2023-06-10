package http

import (
	"log"
	"net/http"
	"strings"

	x "github.com/kompiangg/shipper-fp/pkg/errors"
	httppkg "github.com/kompiangg/shipper-fp/pkg/http"
	"github.com/kompiangg/shipper-fp/pkg/validator"
	"github.com/labstack/echo/v4"
)

type LoggerItf interface {
	Println(v ...any)
}

func WriteResponse(c echo.Context, code int, data interface{}) error {
	if data == nil {
		data = http.StatusText(code)
	}

	err := c.JSON(code, httppkg.HTTPBaseResponse{
		Error: nil,
		Data:  data,
	})
	if err != nil {
		log.Println("[WriteResponse] FATAL ERROR on send response to client:", err)
		return err
	}

	return nil
}

func WriteErrorResponse(c echo.Context, errParam error, detail interface{}) error {
	e := httppkg.GetResponseErr(errParam)

	if x.Is(errParam, x.ErrValidation) {
		unwrapErr := x.Unwrap(errParam)
		detail = strings.Split(unwrapErr.Error(), validator.MessageSeparator)
	} else {
		x.ErrorStack(errParam)
	}

	err := c.JSON(e.HTTPErrorCode, httppkg.HTTPBaseResponse{
		Error: &httppkg.HTTPErrorBaseResponse{
			Message: e.Message,
			Detail:  detail,
		},
		Data: nil,
	})

	if err != nil {
		log.Println("[WriteErrorResponse] FATAL ERROR on send response to client:", err)
		return err
	}

	return nil
}
