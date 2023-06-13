package services

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"resqiar.com-server/entities"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var authService = AuthServiceImpl{}

func TestConvertToken(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockWrongToken := "INVALID_TOKEN"
	mockWrongURL := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", mockWrongToken)
	mockWrongResponse := `{"error": "invalid token"}`

	httpmock.RegisterResponder(http.MethodGet, mockWrongURL, httpmock.NewStringResponder(http.StatusBadRequest, mockWrongResponse))

	t.Run("should return error when the token is not valid", func(t *testing.T) {
		_, err := authService.ConvertToken(mockWrongToken)

		expectedErr := errors.New("Invalid token")

		assert.Error(t, err)
		assert.EqualError(t, err, expectedErr.Error())
	})

	mockCorrectToken := "VALID_TOKEN"
	mockCorrectURL := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", mockCorrectToken)
	mockCorrectResponse := `{
        "sub": "00120219999",
        "name": "example name",
        "given_name": "example",
        "family_name": "example name test",
        "picture": "https://example.com",
        "email": "example@gmail.com",
        "email_verified": true,
        "locale": "en-US"
      }`

	t.Run("should return error when HTTP failed", func(t *testing.T) {
		// mock failed HTTP call
		httpmock.RegisterResponder(http.MethodGet, mockCorrectURL, httpmock.NewErrorResponder(errors.New("something went wrong")))
		result, err := authService.ConvertToken(mockCorrectToken)

		assert.Nil(t, result)
		assert.Error(t, err)
	})

	t.Run("should return error when the response is not parseable", func(t *testing.T) {
		mockNonJSON := "this is indeed not a JSON"

		// mock success HTTP call but with non-parseable response
		httpmock.RegisterResponder(http.MethodGet, mockCorrectURL, httpmock.NewStringResponder(http.StatusOK, mockNonJSON))

		result, err := authService.ConvertToken(mockCorrectToken)

		assert.Nil(t, result)
		assert.Error(t, err)
	})

	// mock success HTTP call
	httpmock.RegisterResponder(http.MethodGet, mockCorrectURL, httpmock.NewStringResponder(http.StatusOK, mockCorrectResponse))

	t.Run("should not return any error when token is valid", func(t *testing.T) {
		_, err := authService.ConvertToken(mockCorrectToken)
		assert.Nil(t, err)
	})

	t.Run("should return a pointer of GooglePayload when the token is valid", func(t *testing.T) {
		result, _ := authService.ConvertToken(mockCorrectToken)
		assert.IsType(t, &entities.GooglePayload{}, result)
	})

	t.Run("should return non-nil value when token is valid", func(t *testing.T) {
		result, _ := authService.ConvertToken(mockCorrectToken)

		assert.NotNil(t, result.SUB)
		assert.NotNil(t, result.Email)
		assert.NotNil(t, result.EmailVerified)
		assert.NotNil(t, result.Name)
		assert.NotNil(t, result.GivenName)
		assert.NotNil(t, result.FamilyName)
		assert.NotNil(t, result.Picture)
		assert.NotNil(t, result.Locale)
	})
}

func TestSignIK(t *testing.T) {
	// Set up the environment variables required for the function
	os.Setenv("IMAGE_KIT_KEY", "dummy_key")
	os.Setenv("IMAGE_KIT_KEY_PUBLIC", "dummy_key")
	os.Setenv("IMAGE_KIT_URL", "dummy_url")

	// Create a mock fiber.Ctx object
	ctx := new(fiber.Ctx)

	t.Run("should return a valid signed token", func(t *testing.T) {
		signedToken := authService.SignIK(ctx)

		// Assert that the signedToken is not nil
		assert.NotNil(t, signedToken)

		// Assert that the signedToken contains the expected fields (Token, Signature, Expire)
		assert.NotEmpty(t, signedToken.Token)
		assert.NotEmpty(t, signedToken.Signature)
		assert.NotEmpty(t, signedToken.Expires)
	})

	// Clean up the environment variables
	os.Unsetenv("IMAGE_KIT_KEY")
	os.Unsetenv("IMAGE_KIT_KEY_PUBLIC")
	os.Unsetenv("IMAGE_KIT_URL")
}
