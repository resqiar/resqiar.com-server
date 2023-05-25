package services

import (
	"strings"

	"github.com/go-playground/validator/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func FormatUsername(name string) string {
	// format name to lowercase
	formatted := strings.ToLower(name)

	// format name to replace all spaces into _ (underscore)
	formatted = strings.ReplaceAll(formatted, " ", "_")

	return formatted
}

func GenerateRandomID(length int) string {
	// generate random string id using nanoid package
	id, _ := gonanoid.New(length)
	return id
}

func ValidateInput(payload any) string {
	// instantiate new instance
	validate := validator.New()

	// save error messages here
	var errMessage string

	errors := validate.Struct(payload)
	if errors != nil {
		// loop through all possible errors,
		// then give appropriate message based on
		// defined error tag, StructField, etc
		for _, err := range errors.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				errMessage = err.StructField() + " field is required"
				break
			}

			if err.Tag() == "max" {
				errMessage = err.StructField() + " field exceed max characters"
				break
			}

			if err.Tag() == "url" {
				errMessage = err.StructField() + " field is not a valid URL"
				break
			}

			// raw error which is not covered above
			errMessage = err.Error()
		}
	}

	return errMessage
}
