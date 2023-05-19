package services

import (
	"strings"

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
