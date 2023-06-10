package errors

import (
	"os"

	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"
)

var (
	ErrInternalServer         = errors.New("internal server error")
	ErrBadRequest             = errors.New("bad request")
	ErrRecordNotFound         = errors.New("record not found")
	ErrNotFound               = errors.New("not found")
	ErrUnauthorized           = errors.New("unauthorized")
	ErrAuthTokenExpired       = errors.New("token expired")
	ErrAccountDuplicated      = errors.New("email or username has already exists")
	ErrUniqueRecord           = errors.New("record duplicated, must be unique")
	ErrMerchantNameDuplicated = errors.New("merchant name duplicated, must be unique")
	ErrUsernameDuplicated     = errors.New("username duplicated, must be unique")
	ErrEmailPasswordIncorrect = errors.New("email or password is incorrect")
)

var (
	ErrValidation = errors.New("validation error")
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
		DisableQuote:  true,
		PadLevelText:  false,
	})
}

func Wrap(cause error, msg string) error {
	if cause == nil {
		return nil
	}

	return errors.WrapPrefix(cause, msg, 1)
}

func Unwrap(err error) error {
	if err != nil {
		return err.(*errors.Error).Unwrap()
	}

	return nil
}

func ErrorStack(err error) {
	log.Warningln(err.(*errors.Error).ErrorStack())
}

func New(e interface{}) *errors.Error {
	return errors.New(e)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
