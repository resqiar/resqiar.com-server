package repositories

import (
	"resqiar.com-server/entities"
	"resqiar.com-server/inputs"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func InitUserRepo(db *gorm.DB) UserRepository {
	return &UserRepoImpl{
		db: db,
	}
}

type UserRepository interface {
	CreateUser(*entities.User) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindByID(ID string) (*entities.SafeUser, error)
	FindByUsername(username string) (*entities.SafeUser, error)
	UpdateUser(ID string, payload *inputs.UpdateUserInput) error
}

func (repo *UserRepoImpl) CreateUser(user *entities.User) (*entities.User, error) {
	input := user
	err := repo.db.Clauses(clause.Returning{}).Create(input).Error
	if err != nil {
		return nil, err
	}
	return input, err
}

func (repo *UserRepoImpl) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	result := repo.db.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepoImpl) FindByID(ID string) (*entities.SafeUser, error) {
	var user entities.SafeUser

	result := repo.db.Model(&entities.User{}).First(&user, "id = ?", ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepoImpl) FindByUsername(username string) (*entities.SafeUser, error) {
	var user entities.SafeUser

	result := repo.db.Model(&entities.User{}).First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepoImpl) UpdateUser(ID string, payload *inputs.UpdateUserInput) error {
	if err := repo.db.Model(&entities.User{}).Where("id = ?", ID).Updates(&payload).Error; err != nil {
		return err
	}

	return nil
}
