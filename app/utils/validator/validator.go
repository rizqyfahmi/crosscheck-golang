package validatorutil

import "github.com/go-playground/validator/v10"

type ValidatorUtil struct {
	validator *validator.Validate
}

func New(validator *validator.Validate) *ValidatorUtil {
	return &ValidatorUtil{
		validator: validator,
	}
}

func (s *ValidatorUtil) Validate(i interface{}) error {
	return s.validator.Struct(i)
}
