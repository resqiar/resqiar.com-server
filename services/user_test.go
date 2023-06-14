package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"resqiar.com-server/constants"
	"resqiar.com-server/entities"
	"resqiar.com-server/repositories"
)

var userRepo = &repositories.UserRepoMock{}
var userService = UserServiceImpl{
	Repository:  userRepo,
	UtilService: &utilService,
}

func TestRegisterUser(t *testing.T) {
	payload := entities.GooglePayload{
		SUB:       "00231231231",
		GivenName: "user name",
		Email:     "test@example.com",
		Picture:   "image.com/example",
	}

	expectedInput := entities.User{
		Email:      payload.Email,
		Provider:   constants.Google,
		ProviderID: payload.SUB,
		PictureURL: payload.Picture,
	}

	t.Run("Should successfully register user with given input (no error)", func(t *testing.T) {
		matcher := func(user *entities.User) bool {
			return user.Email == expectedInput.Email &&
				user.Provider == expectedInput.Provider &&
				user.ProviderID == expectedInput.ProviderID &&
				user.PictureURL == expectedInput.PictureURL
		}

		userRepo.Mock.On("CreateUser",
			mock.MatchedBy(matcher)).Return("")

		result, error := userService.RegisterUser(&payload)

		assert.Nil(t, error)

		t.Run("Should return result with the type of *entities.User", func(t *testing.T) {
			assert.NotNil(t, result)
			assert.IsType(t, &entities.User{}, result)
		})
	})

	t.Run("Should return error if the required field is null", func(t *testing.T) {
		payload.Email = ""

		matcher := func(user *entities.User) bool {
			return user.Email == ""
		}

		userRepo.Mock.On("CreateUser",
			mock.MatchedBy(matcher)).Return("Email cannot be null")

		result, error := userService.RegisterUser(&payload)

		assert.Error(t, error)
		assert.Nil(t, result)
	})
}
