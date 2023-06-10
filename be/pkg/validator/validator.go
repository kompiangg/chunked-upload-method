package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	validator_en "github.com/go-playground/validator/v10/translations/en"
	pkgerrors "github.com/kompiangg/shipper-fp/pkg/errors"
)

const MessageSeparator = "\n --- "

type Validator struct {
	Validator           *validator.Validate
	EnglishTranslator   ut.Translator
	IndonesiaTranslator ut.Translator
}

func New() (Validator, error) {
	validator := validator.New()

	validator.RegisterValidation("notblank", validators.NotBlank)

	english := en.New()
	uni := ut.New(english, english)
	enTrans, _ := uni.GetTranslator("en")
	idTrans, _ := uni.GetTranslator("id")
	err := validator_en.RegisterDefaultTranslations(validator, enTrans)
	if err != nil {
		return Validator{}, err
	}

	return Validator{
		Validator:           validator,
		EnglishTranslator:   enTrans,
		IndonesiaTranslator: idTrans,
	}, nil
}

func (v *Validator) Validate(s interface{}) error {
	err := v.Validator.Struct(s)
	if err != nil {
		return errors.Join(pkgerrors.ErrValidation, v.translateValidationError(err, "en"))
	}

	return nil
}

func (v *Validator) translateValidationError(err error, lang string) error {
	var translator ut.Translator

	if lang == "id" {
		translator = v.IndonesiaTranslator
	} else {
		translator = v.EnglishTranslator
	}

	errsMessage := []string{}

	validatorErrs := err.(validator.ValidationErrors)
	errsMessage = append(errsMessage, "")
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(translator))
		errsMessage = append(errsMessage, translatedErr.Error())
	}

	errMessage := strings.Join(errsMessage, MessageSeparator)

	return errors.New(errMessage)
}
