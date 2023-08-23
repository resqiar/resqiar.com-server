package config

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	illegalUsernameRegex = regexp.MustCompile(`[^a-zA-Z0-9_]`)
)

func InitCustomValidation(validate *validator.Validate) {
	// register custom validation for validator - only do this once
	validate.RegisterValidation("media_url", func(fl validator.FieldLevel) bool {
		value := fl.Field().Interface().(string)

		// only allow "" and " " - else must be validated as a URL
		if value == "" || value == " " {
			return true
		}

		// if the value is not "" or " " - we have to validate it as a URL
		if err := validate.Var(value, "url"); err != nil {
			return false
		}

		return true
	})

	validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		value := fl.Field().Interface().(string)

		if illegalUsernameRegex.Match([]byte(value)) {
			return false
		}

		return true
	})
}
