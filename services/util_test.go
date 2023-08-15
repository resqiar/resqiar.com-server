package services

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var utilService = UtilServiceImpl{}

func TestFormatUsername(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{" ", ""},
		{"   _!?     ", ""},
		{"    A    ", "a"},
		{"TesT UsErNAmE", "test_username"},
		{"Another User", "another_user"},
		{"UPPERCASE USER", "uppercase_user"},
		{"a                   b     user", "a_b_user"},
		{"L 0 ? ] / Y ;  <script>", "l_0_y_script"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Should format: %s INTO %s", tc.input, tc.expected), func(t *testing.T) {
			generated := utilService.FormatUsername(tc.input)

			// if the value is Nil, then stop the test
			require.NotNil(t, generated)

			// assert if the value is as expected
			assert.Equal(t, tc.expected, generated)
		})
	}
}

func TestFormatToURL(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Exploring OAuth2 Protocol: What, Why, and How?", "exploring-oauth2-protocol-what-why-and-how"},
		{"					Exploring OAuth2 Protocol: What, Why, and How?		", "exploring-oauth2-protocol-what-why-and-how"},
		{"					Exploring OAuth2				 Protocol: What, Why, and How?		", "exploring-oauth2-protocol-what-why-and-how"},
		{"					Exploring OAuth2				 Protocol: What, Why, and How?		", "exploring-oauth2-protocol-what-why-and-how"},
		{"Asynchronous Programming in JavaScript Ecosystem", "asynchronous-programming-in-javascript-ecosystem"},
		{"<><><?><><)_(*&&^%^&%^$%#$%%^&(()##$%^&*(&", ""},
		{"TROUBLE MAKER IS ON THE 666 W4Y", "trouble-maker-is-on-the-666-w4y"},
		{"Wh  y  ? ev 3       erything    mu    s-t be not th e    best			         ", "wh-y-ev-3-erything-mu-st-be-not-th-e-best"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Should format: %s INTO %s", tc.input, tc.expected), func(t *testing.T) {
			generated := utilService.FormatToURL(tc.input)

			// if the value is Nil, then stop the test
			require.NotNil(t, generated)

			// assert if the value is as expected
			assert.Equal(t, tc.expected, generated)
		})
	}
}

func TestGenerateRandomID(t *testing.T) {
	t.Run("Should not return Nil or Empty", func(t *testing.T) {
		generated := utilService.GenerateRandomID(8)

		require.NotNil(t, generated)
		require.NotEmpty(t, generated)
	})

	t.Run("Should generate with expected length", func(t *testing.T) {
		expectedLen := 12
		generated := utilService.GenerateRandomID(expectedLen)
		assert.Len(t, generated, expectedLen)
	})

	t.Run("Should generate unique IDs for > 1000x", func(t *testing.T) {
		// Generate multiple random IDs
		count := 1000

		// we use hash map to help keep track of
		// unique keys, if the generated ids have duplicate,
		// then the len will not match the count
		ids := make(map[string]bool)

		for i := 0; i < count; i++ {
			generated := utilService.GenerateRandomID(8)
			ids[generated] = true
		}

		// Check if all generated IDs are unique
		assert.Len(t, ids, count)
	})

	t.Run("Should generate valid characters", func(t *testing.T) {
		invalidCharacter := regexp.MustCompile("[^a-zA-Z0-9_-]")

		generated := utilService.GenerateRandomID(8)

		assert.False(t, invalidCharacter.MatchString(generated), "Generated ID contains invalid characters")
	})
}

func TestValidateInput(t *testing.T) {
	t.Run("Should return error if the payload is Nil", func(t *testing.T) {
		result := utilService.ValidateInput(nil)
		assert.Equal(t, "Invalid Payload", result)
	})

	type Payload struct {
		Title    string `validate:"required,max=10"`
		Summary  string `validate:"max=10"`
		Content  string `validate:"max=10"`
		CoverURL string `validate:"omitempty,url"`
		AuthorID string `validate:"omitempty,uuid"`
	}

	testCases := []struct {
		payload  Payload
		expected string
	}{
		{
			Payload{},
			"Title field is required",
		},
		{
			Payload{
				Summary: "Summary Example",
			},
			"Title field is required",
		},
		{
			Payload{
				Content: "Content Example",
			},
			"Title field is required",
		},
		{
			Payload{
				CoverURL: "https://google.com",
			},
			"Title field is required",
		},
		{
			Payload{
				Title: "123456789ABCDEF",
			},
			"Title field exceed max characters",
		},
		{
			Payload{
				Title:   "123456789",
				Summary: "123456789ABCDEF",
				Content: "123456789ABCDEF",
			},
			"Summary field exceed max characters",
		},
		{
			Payload{
				Title:   "123456789",
				Summary: "123456789",
				Content: "123456789ABCDEF",
			},
			"Content field exceed max characters",
		},
		{
			Payload{
				Title:    "123456789",
				CoverURL: "123456789ABCDEF",
			},
			"CoverURL field is not a valid URL",
		},
		{
			Payload{
				Title:    "123456789",
				CoverURL: "google.com",
			},
			"CoverURL field is not a valid URL",
		},
		{
			Payload{
				Title:    "123456789",
				CoverURL: "www.google.com",
			},
			"CoverURL field is not a valid URL",
		},
		{
			Payload{
				Title:    "123456789",
				CoverURL: "https://google.com",
			},
			"",
		},
		{
			Payload{
				Title:    "123456789",
				CoverURL: "http://github.com/stuff/RaNDoMStTuuFFF.spec.go.png",
			},
			"",
		},
		{
			Payload{
				Title:    "123456789",
				AuthorID: "not-a-uuid",
			},
			"Error on field AuthorID",
		},
		{
			Payload{
				Title:    "123456789",
				AuthorID: "2312313-xnciauwgd-3123if-c",
			},
			"Error on field AuthorID",
		},
		{
			Payload{
				Title:    "123456789",
				AuthorID: "2e2f02c5-d432-449e-b00b-557cc580ef3e",
			},
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Should return message of '%s'", tc.expected), func(t *testing.T) {
			result := utilService.ValidateInput(&tc.payload)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestParseMD(t *testing.T) {
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
			generated := utilService.ParseMD(tc.input)

			if tc.expected == "" {
				require.Empty(t, generated)
			} else {
				// if the value is Empty (Nil / empty string), then stop the test
				require.NotEmpty(t, generated)

				// assert if the value is as expected
				assert.Equal(t, tc.expected, generated)
			}
		})
	}
}
