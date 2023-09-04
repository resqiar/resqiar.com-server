package services

import (
	"bytes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"regexp"
	"resqiar.com-server/entities"
	"resqiar.com-server/protobuf"
	"strings"

	"github.com/go-playground/validator/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"resqiar.com-server/config"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	html "github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/anchor"
)

var (
	validate                  = validator.New()
	removeNonAlphaNumRegex    = regexp.MustCompile("[^ a-zA-Z0-9]")
	removeMultipleSpacesRegex = regexp.MustCompile(`\s+`)
	engine                    = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("paraiso-dark"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
			&anchor.Extender{
				Texter:   anchor.Text("#"),
				Position: anchor.After,
			},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		),
	)
	sanitizePolicy = bluemonday.UGCPolicy().AllowAttrs("style").OnElements("p", "span", "pre")
)

type UtilService interface {
	FormatUsername(name string) string
	FormatToURL(value string) string
	GenerateRandomID(length int) string
	ValidateInput(payload any) string

	// ParseMD converts Markdown content into safe & sanitized HTML.
	// If error happens, it will merely returns empty string.
	ParseMD(s string) string

	ConvertToPROTO(in *entities.SafeBlogAuthor) *protobuf.SafeBlogAuthor
}

type UtilServiceImpl struct{}

func InitUtilService() UtilService {
	config.InitCustomValidation(validate)

	return &UtilServiceImpl{}
}

func (service *UtilServiceImpl) FormatUsername(name string) string {
	// remove any non-alphanumeric characters from the string
	// example "?-_!" should be ""
	// example "a?!;';';'b" should be "ab"
	validChars := removeNonAlphaNumRegex.ReplaceAllString(name, "")
	formatted := validChars

	// trim spaces
	formatted = strings.TrimSpace(formatted)

	// trim spaces between chars to maxed only one space
	// example "a       b" should be "a b"
	singleSpace := removeMultipleSpacesRegex.ReplaceAllString(formatted, " ")
	formatted = singleSpace

	// format name to lowercase
	formatted = strings.ToLower(formatted)

	// format name to replace all spaces into _ (underscore)
	formatted = strings.ReplaceAll(formatted, " ", "_")

	return formatted
}

func (service *UtilServiceImpl) FormatToURL(value string) string {
	formatted := removeNonAlphaNumRegex.ReplaceAllString(value, "")
	formatted = strings.ToLower(formatted)
	formatted = strings.TrimSpace(formatted)
	formatted = removeMultipleSpacesRegex.ReplaceAllString(formatted, " ")
	formatted = strings.ReplaceAll(formatted, " ", "-")
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

			if err.Tag() == "username" {
				errMessage = err.StructField() + " contains illegal characters"
				break
			}

			if err.Tag() == "min" {
				errMessage = err.StructField() + " field does not meet minimum characters"
				break
			}

			if err.Tag() == "max" {
				errMessage = err.StructField() + " field exceed max characters"
				break
			}

			if err.Tag() == "url" || err.Tag() == "media_url" {
				errMessage = err.StructField() + " field is not a valid URL"
				break
			}

			// raw error which is not covered above
			errMessage = "Error on field " + err.StructField()
		}
	}

	return errMessage
}

func (service *UtilServiceImpl) ParseMD(s string) string {
	var buf bytes.Buffer

	if err := engine.Convert([]byte(s), &buf); err != nil {
		log.Println("Error parsing MD:", err)
		return ""
	}

	sanitized := string(sanitizePolicy.SanitizeBytes(buf.Bytes()))
	return sanitized
}

func (service *UtilServiceImpl) ConvertToPROTO(in *entities.SafeBlogAuthor) *protobuf.SafeBlogAuthor {
	return &protobuf.SafeBlogAuthor{
		ID:          in.ID,
		Slug:        in.Slug,
		CreatedAt:   timestamppb.New(in.CreatedAt),
		UpdatedAt:   timestamppb.New(in.UpdatedAt),
		PublishedAt: timestamppb.New(in.PublishedAt),

		Title:    in.Title,
		Summary:  in.Summary,
		Content:  in.Content,
		CoverURL: in.CoverURL,

		Prev: in.Prev,
		Next: in.Next,

		AuthorID: in.AuthorID,
		Author: &protobuf.SafeUser{
			ID:        in.Author.ID,
			CreatedAt: timestamppb.New(in.Author.CreatedAt),

			Fullname:   in.Author.Fullname,
			Username:   in.Author.Username,
			Bio:        in.Author.Bio,
			PictureURL: in.Author.PictureURL,

			WebsiteURL:   in.Author.WebsiteURL,
			GithubURL:    in.Author.GithubURL,
			LinkedinURL:  in.Author.LinkedinURL,
			InstagramURL: in.Author.InstagramURL,
			TwitterURL:   in.Author.TwitterURL,
			YoutubeURL:   in.Author.YoutubeURL,

			IsTester: in.Author.IsTester,
		},
	}
}
