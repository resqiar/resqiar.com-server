package repositories

import (
	"errors"

	"github.com/stretchr/testify/mock"
	"resqiar.com-server/entities"
)

type UserRepoMock struct {
	Mock mock.Mock
}

func (repo *UserRepoMock) CreateUser(user *entities.User) error {
	args := repo.Mock.Called(user)

	if args.Get(0) == "" {
		return nil
	}

	return errors.New(args.Get(0).(string))
}

func (repo *UserRepoMock) FindByEmail(email string) (*entities.User, error) {
	args := repo.Mock.Called(email)
	return nil, errors.New(args.Get(0).(string))
}

func (repo *UserRepoMock) FindByID(ID string) (*entities.SafeUser, error) {
	args := repo.Mock.Called(ID)
	return nil, errors.New(args.Get(0).(string))
}
