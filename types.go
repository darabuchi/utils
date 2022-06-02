package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(m interface{}) error {
	return validate.Struct(m)
}
