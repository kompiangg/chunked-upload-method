package http

import (
	"net/http"

	x "github.com/kompiangg/shipper-fp/pkg/errors"
)

type errorSchema struct {
	HTTPErrorCode int
	Message       string
}

var errMap map[error]errorSchema = map[error]errorSchema{
	x.ErrInternalServer:         {HTTPErrorCode: http.StatusInternalServerError, Message: x.ErrInternalServer.Error()},
	x.ErrBadRequest:             {HTTPErrorCode: http.StatusBadRequest, Message: x.ErrBadRequest.Error()},
	x.ErrValidation:             {HTTPErrorCode: http.StatusBadRequest, Message: x.ErrBadRequest.Error()},
	x.ErrRecordNotFound:         {HTTPErrorCode: http.StatusNotFound, Message: x.ErrRecordNotFound.Error()},
	x.ErrNotFound:               {HTTPErrorCode: http.StatusNotFound, Message: x.ErrNotFound.Error()},
	x.ErrUnauthorized:           {HTTPErrorCode: http.StatusUnauthorized, Message: x.ErrUnauthorized.Error()},
	x.ErrAuthTokenExpired:       {HTTPErrorCode: http.StatusUnauthorized, Message: x.ErrAuthTokenExpired.Error()},
	x.ErrAccountDuplicated:      {HTTPErrorCode: http.StatusBadRequest, Message: x.ErrAccountDuplicated.Error()},
	x.ErrUniqueRecord:           {HTTPErrorCode: http.StatusBadRequest, Message: x.ErrUniqueRecord.Error()},
	x.ErrMerchantNameDuplicated: {HTTPErrorCode: http.StatusBadRequest, Message: x.ErrMerchantNameDuplicated.Error()},
	x.ErrUsernameDuplicated:     {HTTPErrorCode: http.StatusBadRequest, Message: x.ErrUsernameDuplicated.Error()},
	x.ErrEmailPasswordIncorrect: {HTTPErrorCode: http.StatusUnauthorized, Message: x.ErrEmailPasswordIncorrect.Error()},
}

func GetResponseErr(param error) errorSchema {
	param = x.Unwrap(param)

	res, exists := errMap[param]
	if !exists {
		return errMap[x.ErrInternalServer]
	}

	return res
}
