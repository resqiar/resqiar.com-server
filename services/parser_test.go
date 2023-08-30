package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var parserService = ParserServiceImpl{}

func TestParseMDByte(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello paragraph", "<p>hello paragraph</p>\n"},
		{"# Hello World", "<h1 id=\"hello-world\">Hello World <a href=\"#hello-world\" rel=\"nofollow\">#</a></h1>\n"},
		{"## Hello World", "<h2 id=\"hello-world\">Hello World <a href=\"#hello-world\" rel=\"nofollow\">#</a></h2>\n"},
		{"### Hello World", "<h3 id=\"hello-world\">Hello World <a href=\"#hello-world\" rel=\"nofollow\">#</a></h3>\n"},
		{"#### Hello World", "<h4 id=\"hello-world\">Hello World <a href=\"#hello-world\" rel=\"nofollow\">#</a></h4>\n"},
		{"##### Hello World", "<h5 id=\"hello-world\">Hello World <a href=\"#hello-world\" rel=\"nofollow\">#</a></h5>\n"},
		{"###### Hello World", "<h6 id=\"hello-world\">Hello World <a href=\"#hello-world\" rel=\"nofollow\">#</a></h6>\n"},
		{"* Item 1\n* Item 2", "<ul>\n<li>Item 1</li>\n<li>Item 2</li>\n</ul>\n"},
		{"> Blockquote", "<blockquote>\n<p>Blockquote</p>\n</blockquote>\n"},
		{"[Link](https://www.example.com)", "<p><a href=\"https://www.example.com\" rel=\"nofollow\">Link</a></p>\n"},
		{"**Bold Text**", "<p><strong>Bold Text</strong></p>\n"},
		{"*Italic Text*", "<p><em>Italic Text</em></p>\n"},
		{"`Code Snippet`", "<p><code>Code Snippet</code></p>\n"},
		{"1. Item 1\n2. Item 2", "<ol>\n<li>Item 1</li>\n<li>Item 2</li>\n</ol>\n"},
		{"![Image](https://www.example.com/image.jpg)", "<p><img src=\"https://www.example.com/image.jpg\" alt=\"Image\"></p>\n"},

		{"<script>alert('XSS Attack');</script>", ""},
		{"<img src=\"x\" onerror=\"alert('XSS Attack')\">", "<img src=\"x\">"},
		{"<a href=\"javascript:alert('XSS Attack')\">Click Me</a>", "<p>Click Me</p>\n"},
		{"<p onmouseover=alert(\"XSS Attack!\")>click me!</p>", "<p>click me!</p>"},
		{"<img src=\"https://url.to.file.which/not.exist\" onerror=alert(document.cookie);/>", "<img src=\"https://url.to.file.which/not.exist\"/>"},
		{"<img src=j&#X41vascript:alert(\"XSS Attack\")/>", "<p>&lt;img src=j&amp;#X41vascript:alert(&#34;XSS Attack&#34;)/&gt;</p>\n"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Should parse MD from: %s TO %s", tc.input, tc.expected), func(t *testing.T) {
			generated := parserService.ParseMDByte([]byte(tc.input))

			if tc.expected == "" {
				require.Empty(t, generated)
			} else {
				// if the value is Empty (Nil / empty string), then stop the test
				require.NotEmpty(t, generated)

				// assert if the value is as expected
				assert.Equal(t, []byte(tc.expected), generated)
			}
		})
	}
}
