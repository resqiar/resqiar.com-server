package repositories

import (
	"resdev-server/entities"

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
	CreateUser(*entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindByID(ID string) (*entities.SafeUser, error)
}

func (repo *UserRepoImpl) CreateUser(user *entities.User) error {
	result := repo.db.Clauses(clause.Returning{}).Create(user)
	return result.Error
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

	result := repo.db.First(&user, "id = ?", ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
