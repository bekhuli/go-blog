package user

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

func NewValidator() *Validator {
	v := validator.New()
	v.RegisterValidation("username", validateUsername)
	return &Validator{validate: v}
}

func validateUsername(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
}

func (v *Validator) Validate(s interface{}) error {
	return v.validate.Struct(s)
}
