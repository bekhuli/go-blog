package post

import "github.com/go-playground/validator/v10"

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()

	return &Validator{validate: v}
}

func (v *Validator) Validate(s interface{}) error {
	return v.validate.Struct(s)
}
