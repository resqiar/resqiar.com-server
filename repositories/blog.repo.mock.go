package repositories

import (
	"github.com/stretchr/testify/mock"
	"resqiar.com-server/entities"
	"resqiar.com-server/inputs"
)

type BlogRepoMock struct {
	Mock mock.Mock
}

func (repo *BlogRepoMock) GetBlogs(onlyPublished bool) ([]entities.SafeBlogAuthor, error) {
	args := repo.Mock.Called(onlyPublished)

	if args.Get(0) != nil {
		return args.Get(0).([]entities.SafeBlogAuthor), args.Error(1)
	}

	return nil, args.Error(1)
}

func (repo *BlogRepoMock) GetBlog(blogID string, published bool) (*entities.SafeBlogAuthor, error) {
	args := repo.Mock.Called(blogID, published)

	if args.Get(0) != nil {
		return args.Get(0).(*entities.SafeBlogAuthor), args.Error(1)
	}

	return nil, args.Error(1)
}

func (repo *BlogRepoMock) CreateBlog(input *entities.Blog) (*entities.Blog, error) {
	args := repo.Mock.Called(input)

	if args.Get(0) != nil {
		return args.Get(0).(*entities.Blog), args.Error(1)
	}

	return nil, args.Error(1)
}

func (repo *BlogRepoMock) UpdateBlog(blogID string, safe *inputs.SafeUpdateBlogInput) error {
	args := repo.Mock.Called(blogID, safe)

	if args.Get(0) == nil {
		return nil
	}

	return args.Error(0)
}

func (repo *BlogRepoMock) GetByIDAndAuthor(blogID string, userID string) (*entities.Blog, error) {
	args := repo.Mock.Called(blogID, userID)

	if args.Get(0) != nil {
		return args.Get(0).(*entities.Blog), args.Error(1)
	}

	return nil, args.Error(1)
}

func (repo *BlogRepoMock) GetCurrentUserBlogs(userID string) ([]entities.Blog, error) {
	args := repo.Mock.Called(userID)

	if args.Get(0) != nil {
		return args.Get(0).([]entities.Blog), args.Error(1)
	}

	return nil, args.Error(1)
}

func (repo *BlogRepoMock) GetCurrentUserBlog(blogID string, userID string) (*entities.Blog, error) {
	args := repo.Mock.Called(blogID, userID)

	if args.Get(0) != nil {
		return args.Get(0).(*entities.Blog), args.Error(1)
	}

	return nil, args.Error(1)
}

func (repo *BlogRepoMock) SaveBlog(blog *entities.Blog) error {
	args := repo.Mock.Called(blog)

	if args.Get(0) == nil {
		return nil
	}

	return args.Error(0)
}
