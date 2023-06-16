package handlers

import (
	"log"
	"os"

	"resqiar.com-server/config"
	"resqiar.com-server/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SendAuthGoogle(c *fiber.Ctx) error
	SendGoogleCallback(c *fiber.Ctx) error
	SendLogout(c *fiber.Ctx) error
	SendAuthIK(c *fiber.Ctx) error
}

type AuthHandlerImpl struct {
	UserService services.UserService
	AuthService services.AuthService
	UtilService services.UtilService
}

func (handler *AuthHandlerImpl) SendAuthGoogle(c *fiber.Ctx) error {
	// create a config for google config
	conf := config.GoogleConfig()

	// generate random 32-long for state identification
	generated := handler.UtilService.GenerateRandomID(32)

	sess, _ := config.StateStore.Get(c)
	sess.Set("session_state", generated)
	sess.Save()

	// create url for auth process.
	// we can pass state as someway to identify
	// and validate the login process.
	URL := conf.AuthCodeURL(generated)

	// redirect to the google authentication URL
	return c.Redirect(URL)
}

func (handler *AuthHandlerImpl) SendGoogleCallback(c *fiber.Ctx) error {
	// get session store for current context
	sess, sessErr := config.SessionStore.Get(c)
	stateSess, stateErr := config.StateStore.Get(c)
	if sessErr != nil || stateErr != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// state from the session storage
	savedState := stateSess.Get("session_state")

	conf := config.GoogleConfig()

	// get state and  from the google callback
	state := c.Query("state")
	code := c.Query("code")

	// compare the state that is coming from the callback
	// with the one that is stored inside the session storage.
	if state != savedState {
		// Handle the invalid state error
		return c.Status(fiber.StatusBadRequest).SendString("Invalid state")
	}

	// exchange code that retrieved from google via
	// URL query parameter into token, this token
	// can be used later to query information of current user
	// from respective provider.
	token, err := conf.Exchange(c.Context(), code)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	profile, err := handler.AuthService.ConvertToken(token.AccessToken)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// find current user by provided email,
	// if the user found in the database, then we can just logged in,
	// if not, then register that user.
	isExist, err := handler.UserService.FindUserByEmail(profile.Email)
	// this error indicates user not found
	if err != nil {
		// register user and save their data into database
		result, err := handler.UserService.RegisterUser(profile)
		if err != nil {
			log.Printf("Failed to register user: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to register user")
		}

		// Store the user's id in the session
		sess.Set("ID", result.ID)

		// Save into memory session and.
		// saving also set a session cookie containing session_id
		if err := sess.Save(); err != nil {
			log.Printf("Failed to save user session: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to save user session")
		}

		// return immediately
		return c.Status(fiber.StatusOK).Redirect(os.Getenv("CLIENT_URL"))
	}

	// Store the existed user's id in the session
	sess.Set("ID", isExist.ID)

	if err := sess.Save(); err != nil {
		log.Printf("Failed to save user session: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save user session")
	}

	return c.Status(fiber.StatusOK).Redirect(os.Getenv("CLIENT_URL"))
}

func (handler *AuthHandlerImpl) SendLogout(c *fiber.Ctx) error {
	sess, err := config.SessionStore.Get(c)
	if err != nil {
		log.Println(err.Error())
	}

	// destroy current user session
	sess.Destroy()

	return c.SendStatus(fiber.StatusOK)
}

func (handler *AuthHandlerImpl) SendAuthIK(c *fiber.Ctx) error {
	signed := handler.AuthService.SignIK(c)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"token":     signed.Token,
		"signature": signed.Signature,
		"expire":    signed.Expires,
	})
}
