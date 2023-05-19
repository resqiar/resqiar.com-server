package services

import (
	"fmt"
	"resdev-server/constants"
	"resdev-server/db"
	"resdev-server/entities"

	"gorm.io/gorm/clause"
)

func RegisterUser(profile *entities.GooglePayload) (*entities.User, error) {
	// format the given name from the provider
	formattedName := FormatUsername(profile.GivenName)

	// concatenate formatted name with the nano id
	formattedName = fmt.Sprintf("%s_%s", formattedName, GenerateRandomID(7))

	newUser := entities.User{
		Username:   formattedName,
		Email:      profile.Email,
		Provider:   constants.Google,
		ProviderID: profile.SUB,
	}

	result := db.DB.Clauses(clause.Returning{}).Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}

func FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User

	result := db.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
