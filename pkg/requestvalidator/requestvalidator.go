package requestvalidator

import "github.com/go-playground/validator/v10"

type RequestValidator struct {
	Validator *validator.Validate
}

func (rv *RequestValidator) Validate(i interface{}) error {
	return rv.Validator.Struct(i)
}
