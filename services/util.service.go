package services

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type UtilService interface {
	FormatUsername(name string) string
	GenerateRandomID(length int) string
	ValidateInput(payload any) string
}

type UtilServiceImpl struct{}

func (service *UtilServiceImpl) FormatUsername(name string) string {
	// remove any non-alphanumeric characters from the string
	// example "?-_!" should be ""
	// example "a?!;';';'b" should be "ab"
	validChars := regexp.MustCompile("[^ a-zA-Z0-9]").ReplaceAllString(name, "")
	formatted := validChars

	// trim spaces
	formatted = strings.TrimSpace(formatted)

	// trim spaces between chars to maxed only one space
	// example "a       b" should be "a b"
	singleSpace := regexp.MustCompile(`\s+`).ReplaceAllString(formatted, " ")
	formatted = singleSpace

	// format name to lowercase
	formatted = strings.ToLower(formatted)

	// format name to replace all spaces into _ (underscore)
	formatted = strings.ReplaceAll(formatted, " ", "_")

	return formatted
}

func (service *UtilServiceImpl) GenerateRandomID(length int) string {
	// generate random string id using nanoid package
	id, _ := gonanoid.New(length)
	return id
}

func (service *UtilServiceImpl) ValidateInput(payload any) string {
	if payload == nil {
		return "Invalid Payload"
	}

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
			errMessage = "Error on field " + err.StructField()
		}
	}

	return errMessage
}
