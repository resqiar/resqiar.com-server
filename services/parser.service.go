package services

import (
	"bytes"
	"log"
	"sync"

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
	parserEngine = goldmark.New(
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
	parserSanitization = bluemonday.UGCPolicy().AllowAttrs("style").OnElements("p", "span", "pre")
	bufferPool         = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

type ParserService interface {
	ParseMDByte(s []byte) []byte
}

type ParserServiceImpl struct{}

func InitParserService() ParserService {
	return &ParserServiceImpl{}
}

func (service *ParserServiceImpl) ParseMDByte(s []byte) []byte {
	// use buffer from the Pool
	buf := bufferPool.Get().(*bytes.Buffer)

	// restore resources back to the pool when done
	defer bufferPool.Put(buf)

	// since we are reusing resources, better reset it first
	buf.Reset()

	if err := parserEngine.Convert(s, buf); err != nil {
		log.Println("Error parsing MD:", err)
		return nil
	}

	return parserSanitization.SanitizeBytes(buf.Bytes())
}
