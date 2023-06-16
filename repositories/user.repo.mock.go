package repositories

import (
	"errors"

	"github.com/stretchr/testify/mock"
	"resqiar.com-server/entities"
)

type UserRepoMock struct {
	Mock mock.Mock
}

func (repo *UserRepoMock) CreateUser(user *entities.User) (*entities.User, error) {
	args := repo.Mock.Called(user)

	if args.Get(0) != nil {
		return args.Get(0).(*entities.User), nil
	}

	return nil, errors.New(args.Get(1).(string))
}

func (repo *UserRepoMock) FindByEmail(email string) (*entities.User, error) {
	args := repo.Mock.Called(email)

	validEmail := "valid@example.com"

	if args.Get(0) == validEmail {
		user := entities.User{
			Email: email,
		}

		return &user, nil
	}

	return nil, errors.New("Record not found")
}

func (repo *UserRepoMock) FindByID(ID string) (*entities.SafeUser, error) {
	args := repo.Mock.Called(ID)

	validID := "example-of-valid-id"

	if args.Get(0) == validID {
		user := entities.SafeUser{
			ID: ID,
		}

		return &user, nil
	}

	return nil, errors.New("Record not found")
}
