package handlers

import (
	"log"
	"os"
	"resdev-server/config"
	"resdev-server/services"

	"github.com/gofiber/fiber/v2"
)

func SendAuthGoogle(c *fiber.Ctx) error {
	// create a config for google config
	conf := config.GoogleConfig()

	// create url for auth process.
	// we can pass state as someway to identify
	// and validate the login process, for now skip it.
	URL := conf.AuthCodeURL("state")

	// redirect to the google authentication URL
	return c.Redirect(URL)
}

func SendGoogleCallback(c *fiber.Ctx) error {
	// get session store for current context
	sess, err := config.SessionStore.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	conf := config.GoogleConfig()
	code := c.Query("code")

	// exchange code that retrieved from google via
	// URL query parameter into token, this token
	// can be used later to query information of current user
	// from respective provider.
	token, err := conf.Exchange(c.Context(), code)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	profile, err := services.ConvertToken(token.AccessToken)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// find current user by provided email,
	// if the user found in the database, then we can just logged in,
	// if not, then register that user.
	isExist, err := services.FindUserByEmail(profile.Email)
	// this error indicates user not found
	if err != nil {
		// register user and save their data into database
		result, err := services.RegisterUser(profile)
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

func SendLogout(c *fiber.Ctx) error {
	sess, err := config.SessionStore.Get(c)
	if err != nil {
		log.Println(err.Error())
	}

	// destroy current user session
	sess.Destroy()

	return c.SendStatus(fiber.StatusOK)
}

func SendAuthStatus(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	if userID == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.SendStatus(fiber.StatusOK)
}
