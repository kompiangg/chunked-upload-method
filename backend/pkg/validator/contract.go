package validator

type ValidatorItf interface {
	Validate(s interface{}) error
}
