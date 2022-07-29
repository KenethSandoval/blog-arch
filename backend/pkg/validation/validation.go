package validation

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type (
	Validation struct {
		validate *validator.Validate
	}
)

func New() *Validation {
	valid := &Validation{
		validate: validator.New(),
	}

	return valid
}

func (v *Validation) Validate(i interface{}) error {
	if i == nil {
		return &validator.InvalidValidationError{Type: reflect.TypeOf(i)}
	}

	return v.validate.Struct(i)
}
