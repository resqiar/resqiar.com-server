package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"resdev-server/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/imagekit-developer/imagekit-go"
)

func ConvertToken(accessToken string) (*entities.GooglePayload, error) {
	resp, httpErr := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken))
	if httpErr != nil {
		return nil, httpErr
	}

	// clean up when this function returns (destroyed)
	defer resp.Body.Close()

	respBody, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	// Bind JSON into struct
	var data entities.GooglePayload
	err := json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func SignIK(c *fiber.Ctx) imagekit.SignedToken {
	IMAGE_KIT_KEY := os.Getenv("IMAGE_KIT_KEY")
	IMAGE_KIT_KEY_PUBLIC := os.Getenv("IMAGE_KIT_KEY_PUBLIC")
	IMAGE_KIT_URL := os.Getenv("IMAGE_KIT_URL")

	// Initialize image kit with provided params
	ik := imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey:  IMAGE_KIT_KEY,
		PublicKey:   IMAGE_KIT_KEY_PUBLIC,
		UrlEndpoint: IMAGE_KIT_URL,
	})

	// return an Object containing Token, Signature and Expire
	signed := ik.SignToken(imagekit.SignTokenParam{})
	return signed
}
