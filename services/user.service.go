package services

import (
	"fmt"
	"resqiar.com-server/constants"
	"resqiar.com-server/entities"
	"resqiar.com-server/repositories"
)

type UserService interface {
	RegisterUser(profile *entities.GooglePayload) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
	FindUserByID(userID string) (*entities.SafeUser, error)
}

type UserServiceImpl struct {
	UtilService UtilService
	Repository  repositories.UserRepository
}

func (service *UserServiceImpl) RegisterUser(profile *entities.GooglePayload) (*entities.User, error) {
	// format the given name from the provider
	formattedName := service.UtilService.FormatUsername(profile.GivenName)

	// concatenate formatted name with the nano id
	formattedName = fmt.Sprintf("%s_%s", formattedName, service.UtilService.GenerateRandomID(7))

	newUser := &entities.User{
		Username:   formattedName,
		Email:      profile.Email,
		Provider:   constants.Google,
		ProviderID: profile.SUB,
		PictureURL: profile.Picture,
	}

	result, err := service.Repository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *UserServiceImpl) FindUserByEmail(email string) (*entities.User, error) {
	result, err := service.Repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *UserServiceImpl) FindUserByID(userID string) (*entities.SafeUser, error) {
	safeUser, err := service.Repository.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return safeUser, nil
}
