package services

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"resqiar.com-server/constants"
	"resqiar.com-server/dto"
	"resqiar.com-server/entities"
	"resqiar.com-server/inputs"
	"resqiar.com-server/repositories"
)

var blogRepoTest = repositories.BlogRepoMock{}
var blogServiceTest = BlogServiceImpl{
	UtilService: &utilService,
	Repository:  &blogRepoTest,
}

func TestGetBlogs(t *testing.T) {
	t.Run("Should return an array of published blogs", func(t *testing.T) {
		published := true

		expected := []entities.SafeBlogAuthor{
			{
				SafeBlog: entities.SafeBlog{
					PublishedAt: time.Now(),
				},
			},
			{
				SafeBlog: entities.SafeBlog{
					PublishedAt: time.Now(),
				},
			},
		}

		mock := blogRepoTest.Mock.On("GetBlogs", published, true).Return(expected, nil)

		results, err := blogServiceTest.GetAllBlogs(published, constants.DESC)

		assert.Nil(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, results, expected)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})

	t.Run("Should return an array of unpublished blogs", func(t *testing.T) {
		published := false

		expected := []entities.SafeBlogAuthor{
			{
				SafeBlog: entities.SafeBlog{
					PublishedAt: time.Time{},
				},
			},
			{
				SafeBlog: entities.SafeBlog{
					PublishedAt: time.Time{},
				},
			},
		}

		mock := blogRepoTest.Mock.On("GetBlogs", published, true).Return(expected, nil)

		results, err := blogServiceTest.GetAllBlogs(published, constants.DESC)

		assert.Nil(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, results, expected)

		for _, result := range results {
			// the published at date should be invalid
			assert.Zero(t, result.PublishedAt)
		}

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})

	t.Run("Should return an error query fails", func(t *testing.T) {
		published := false

		blogRepoTest.Mock.On("GetBlogs", published, true).Return(nil, errors.New("Something went wrong"))

		results, err := blogServiceTest.GetAllBlogs(published, constants.DESC)

		assert.Nil(t, results)
		assert.NotNil(t, err)
		assert.Error(t, err)

		blogRepoTest.Mock.AssertCalled(t, "GetBlogs", published, true)
	})

	t.Run("Should return result in DESC order", func(t *testing.T) {
		published := true

		expected := []entities.SafeBlogAuthor{
			{
				SafeBlog: entities.SafeBlog{
					UpdatedAt:   time.Now().AddDate(0, 0, -1), // yesterday
					PublishedAt: time.Now(),
				},
			},
			{
				SafeBlog: entities.SafeBlog{
					UpdatedAt:   time.Now().AddDate(0, 0, -7), // last week
					PublishedAt: time.Now(),
				},
			},
		}

		mock := blogRepoTest.Mock.On("GetBlogs", published, true).Return(expected, nil)

		results, err := blogServiceTest.GetAllBlogs(published, constants.DESC)

		assert.Nil(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, results, expected)

		// assert if the sortedResults are indeed DESC order
		assert.True(t, results[0].UpdatedAt.After(results[1].UpdatedAt))

		// assert if the mock function is called with DESC == true
		blogRepoTest.Mock.AssertCalled(t, "GetBlogs", published, true)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})

	t.Run("Should return result in ASC order", func(t *testing.T) {
		published := true

		expected := []entities.SafeBlogAuthor{
			{
				SafeBlog: entities.SafeBlog{
					UpdatedAt:   time.Now().AddDate(0, 0, -7), // last week
					PublishedAt: time.Now(),
				},
			},
			{
				SafeBlog: entities.SafeBlog{
					UpdatedAt:   time.Now().AddDate(0, 0, -1), // yesterday
					PublishedAt: time.Now(),
				},
			},
		}

		mock := blogRepoTest.Mock.On("GetBlogs", published, false).Return(expected, nil)

		results, err := blogServiceTest.GetAllBlogs(published, constants.ASC)

		assert.Nil(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, results, expected)

		// assert if the sortedResults are indeed ASC order
		assert.True(t, results[0].UpdatedAt.Before(results[1].UpdatedAt))

		// assert if the mock function is called with DESC == false
		blogRepoTest.Mock.AssertCalled(t, "GetBlogs", published, false)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})
}

func TestGetBlogsID(t *testing.T) {
	t.Run("Should return an array of published blog IDs", func(t *testing.T) {
		published := true

		expected := []entities.SafeBlogAuthor{
			{
				SafeBlog: entities.SafeBlog{
					ID:          "example-of-id",
					PublishedAt: time.Now(),
				},
			},
			{
				SafeBlog: entities.SafeBlog{
					ID:          "example-of-id",
					PublishedAt: time.Time{},
				},
			},
		}

		mock := blogRepoTest.Mock.On("GetBlogs", published, true).Return(expected, nil)

		results, err := blogServiceTest.GetAllBlogsID()

		assert.Nil(t, err)
		assert.NotEmpty(t, results)
		assert.IsType(t, []dto.SitemapOutput{}, results)

		for _, result := range results {
			assert.Equal(t, "example-of-id", result.ID)
			assert.NotNil(t, result.UpdatedAt)
		}

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})

	t.Run("Should return an error if query fails", func(t *testing.T) {
		published := true
		blogRepoTest.Mock.On("GetBlogs", published, true).Return(nil, errors.New("Something went wrong"))

		results, err := blogServiceTest.GetAllBlogsID()

		assert.Nil(t, results)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})
}

func TestGetBlogDetail(t *testing.T) {
	t.Run("Should return published blog detail using slug", func(t *testing.T) {
		authorUsername := "user123"
		slug := "example-of-slug"
		published := true

		expectedBlog := &entities.SafeBlogAuthor{
			SafeBlog: entities.SafeBlog{
				Slug:        slug,
				PublishedAt: time.Now(),
			},
		}

		mock := blogRepoTest.Mock.On("GetBlog", "", authorUsername, slug, published).Return(expectedBlog, nil)

		result, err := blogServiceTest.GetBlogDetail("", authorUsername, slug, published)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.NotZero(t, result.PublishedAt)
		assert.Equal(t, expectedBlog, result)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})

	t.Run("Should return unpublished blog detail using ID", func(t *testing.T) {
		blogID := "example-of-id"
		published := false

		expectedBlog := &entities.SafeBlogAuthor{
			SafeBlog: entities.SafeBlog{
				ID:          blogID,
				PublishedAt: time.Time{},
			},
		}

		mock := blogRepoTest.Mock.On("GetBlog", blogID, "", "", published).Return(expectedBlog, nil)

		result, err := blogServiceTest.GetBlogDetail(blogID, "", "", published)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedBlog, result)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})

	t.Run("Should error if author username wrong", func(t *testing.T) {
		authorUsername := "Wrong_Author"
		slug := "example-of-slug"
		published := true

		mock := blogRepoTest.Mock.On("GetBlog", "", authorUsername, slug, published).Return(nil, errors.New("Record not found"))

		result, err := blogServiceTest.GetBlogDetail("", authorUsername, slug, published)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "Record not found")

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})

	t.Run("Should error if slug wrong", func(t *testing.T) {
		authorUsername := "User123"
		slug := "example-of-wrong-slug"
		published := true

		mock := blogRepoTest.Mock.On("GetBlog", "", authorUsername, slug, published).Return(nil, errors.New("Record not found"))

		result, err := blogServiceTest.GetBlogDetail("", authorUsername, slug, published)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "Record not found")

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})
}

func TestCreateBlog(t *testing.T) {
	t.Run("Should successfully called and return the created blog", func(t *testing.T) {
		userID := "example-of-id"

		payload := inputs.CreateBlogInput{
			Title: "Example Title",
		}
		input := entities.Blog{
			Title:    payload.Title,
			AuthorID: userID,
		}

		mock := blogRepoTest.Mock.On("CreateBlog", &input).Return(&input, nil)

		result, err := blogServiceTest.CreateBlog(&payload, userID)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &input, result)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})

	t.Run("Should return error when the query fails", func(t *testing.T) {
		userID := "example-of-id"

		payload := inputs.CreateBlogInput{
			Title: "Example Title",
		}
		input := entities.Blog{
			Title:    payload.Title,
			AuthorID: userID,
		}

		mock := blogRepoTest.Mock.On("CreateBlog", &input).Return(nil, errors.New("Something went wrong"))

		result, err := blogServiceTest.CreateBlog(&payload, userID)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})
}

func TestEditBlog(t *testing.T) {
	t.Run("Should successfully called and return the created blog", func(t *testing.T) {
		blogID := "example-of-id"
		userID := "example-of-id"

		payload := &inputs.UpdateBlogInput{
			ID: blogID,
		}

		expectedBlog := &entities.Blog{
			ID: payload.ID,
		}

		expected := inputs.SafeUpdateBlogInput{}

		firstMock := blogRepoTest.Mock.On("GetByIDAndAuthor", payload.ID, userID).Return(expectedBlog, nil)
		secondMock := blogRepoTest.Mock.On("UpdateBlog", payload.ID, &expected).Return(nil)

		err := blogServiceTest.EditBlog(payload, userID)

		assert.Nil(t, err)

		t.Cleanup(func() {
			// Cleanup mocking
			firstMock.Unset()
			secondMock.Unset()
		})
	})

	t.Run("Should return error when the blog is not found", func(t *testing.T) {
		blogID := "example-of-wrong-id"
		userID := "example-of-id"

		payload := &inputs.UpdateBlogInput{
			ID: blogID,
		}

		expected := inputs.SafeUpdateBlogInput{}

		firstMock := blogRepoTest.Mock.On("GetByIDAndAuthor", payload.ID, userID).Return(nil, errors.New("Record not found"))
		secondMock := blogRepoTest.Mock.On("UpdateBlog", payload.ID, &expected).Return(nil)

		err := blogServiceTest.EditBlog(payload, userID)

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.EqualError(t, err, "Record not found")

		t.Cleanup(func() {
			// Cleanup mocking
			firstMock.Unset()
			secondMock.Unset()
		})
	})

	t.Run("Should return error when the blog failed to be edited", func(t *testing.T) {
		blogID := "example-of-id"
		userID := "example-of-id"

		payload := &inputs.UpdateBlogInput{
			ID: blogID,
		}

		expectedBlog := &entities.Blog{
			ID: payload.ID,
		}

		expected := inputs.SafeUpdateBlogInput{}

		firstMock := blogRepoTest.Mock.On("GetByIDAndAuthor", payload.ID, userID).Return(expectedBlog, nil)
		secondMock := blogRepoTest.Mock.On("UpdateBlog", payload.ID, &expected).Return(errors.New("Error updating blog"))

		err := blogServiceTest.EditBlog(payload, userID)

		assert.NotNil(t, err)
		assert.Error(t, err)

		t.Cleanup(func() {
			// Cleanup mocking
			firstMock.Unset()
			secondMock.Unset()
		})
	})
}

func TestGetCurrentUserBlogs(t *testing.T) {
	t.Run("Should return an array of published blogs", func(t *testing.T) {
		userID := "example-of-id"

		expected := []entities.Blog{
			{
				AuthorID: userID,
			},
			{
				AuthorID: userID,
			},
		}

		mock := blogRepoTest.Mock.On("GetCurrentUserBlogs", userID, true).Return(expected, nil)

		results, err := blogServiceTest.GetCurrentUserBlogs(userID, constants.DESC)

		assert.Nil(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, results, expected)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})

	t.Run("Should return error of if query fails", func(t *testing.T) {
		userID := "example-of-wrong-id"

		mock := blogRepoTest.Mock.On("GetCurrentUserBlogs", userID, true).Return(nil, errors.New("Record not found"))

		results, err := blogServiceTest.GetCurrentUserBlogs(userID, constants.DESC)

		assert.Nil(t, results)
		assert.NotNil(t, err)
		assert.Error(t, err)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})

	t.Run("Should return result in DESC order", func(t *testing.T) {
		userID := "example-of-id"

		expected := []entities.Blog{
			{
				UpdatedAt: time.Now().AddDate(0, 0, -1), // yesterday
			},
			{
				UpdatedAt: time.Now().AddDate(0, 0, -7), // last week
			},
		}

		mock := blogRepoTest.Mock.On("GetCurrentUserBlogs", userID, true).Return(expected, nil)

		results, err := blogServiceTest.GetCurrentUserBlogs(userID, constants.DESC)

		assert.Nil(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, results, expected)

		// assert if the sortedResults are indeed DESC order
		assert.True(t, results[0].UpdatedAt.After(results[1].UpdatedAt))

		// assert if the mock function is called with DESC == true
		blogRepoTest.Mock.AssertCalled(t, "GetCurrentUserBlogs", userID, true)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})

	t.Run("Should return result in ASC order", func(t *testing.T) {
		userID := "example-of-id"

		expected := []entities.Blog{
			{
				UpdatedAt: time.Now().AddDate(0, 0, -7), // last week
			},
			{
				UpdatedAt: time.Now().AddDate(0, 0, -1), // yesterday
			},
		}

		mock := blogRepoTest.Mock.On("GetCurrentUserBlogs", userID, false).Return(expected, nil)

		results, err := blogServiceTest.GetCurrentUserBlogs(userID, constants.ASC)

		assert.Nil(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, results, expected)

		// assert if the sortedResults are indeed ASC order
		assert.True(t, results[0].UpdatedAt.Before(results[1].UpdatedAt))

		// assert if the mock function is called with DESC == false
		blogRepoTest.Mock.AssertCalled(t, "GetCurrentUserBlogs", userID, false)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})
}

func TestGetCurrentUserSlugs(t *testing.T) {
	t.Run("Should return array of current user slugs", func(t *testing.T) {
	})
}

func TestGetCurrentUserBlog(t *testing.T) {
	t.Run("Should return a blog own by current user", func(t *testing.T) {
		userID := "example-of-id"
		blogID := "example-of-blog-id"

		expected := &entities.Blog{
			AuthorID: userID,
		}

		mock := blogRepoTest.Mock.On("GetCurrentUserBlog", blogID, userID).Return(expected, nil)

		results, err := blogServiceTest.GetCurrentUserBlog(blogID, userID)

		assert.Nil(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, results, expected)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})

	t.Run("Should return error of if query fails", func(t *testing.T) {
		userID := "example-of-wrong-id"
		blogID := "example-of-blog-id"

		mock := blogRepoTest.Mock.On("GetCurrentUserBlog", blogID, userID).Return(nil, errors.New("Record not found"))

		results, err := blogServiceTest.GetCurrentUserBlog(blogID, userID)

		assert.Nil(t, results)
		assert.NotNil(t, err)
		assert.Error(t, err)

		t.Cleanup(func() {
			// Cleanup mocking
			mock.Unset()
		})
	})
}

func TestChangeBlogPublish(t *testing.T) {
	t.Run("Should change a blog publication status from FALSE to TRUE", func(t *testing.T) {
		slug := "example-of-title"
		userID := "example-of-user-id"
		payload := inputs.BlogIDInput{
			ID: "example-of-id",
		}

		unpublishedBlog := &entities.Blog{
			ID:        payload.ID,
			Title:     "Example of Title",
			Published: false,
		}

		firstMock := blogRepoTest.Mock.On("GetByIDAndAuthor", payload.ID, userID).Return(unpublishedBlog, nil)
		secondMock := blogRepoTest.Mock.On("SaveBlog", unpublishedBlog).Return(nil)
		thirdMock := blogRepoTest.Mock.On("GetCurrentUserSlugs", slug, userID).Return([]entities.Blog{}, nil)

		err := blogServiceTest.ChangeBlogPublish(&payload, userID, true)

		assert.Nil(t, err)

		t.Cleanup(func() {
			// Cleanup mocking
			firstMock.Unset()
			secondMock.Unset()
			thirdMock.Unset()
		})
	})

	t.Run("Should change a blog publication status from TRUE to FALSE", func(t *testing.T) {
		slug := "example-of-title"
		userID := "example-of-user-id"
		payload := inputs.BlogIDInput{
			ID: "example-of-id",
		}

		unpublishedBlog := &entities.Blog{
			ID:        payload.ID,
			Title:     "Example of Title",
			Published: true,
		}

		firstMock := blogRepoTest.Mock.On("GetByIDAndAuthor", payload.ID, userID).Return(unpublishedBlog, nil)
		secondMock := blogRepoTest.Mock.On("SaveBlog", unpublishedBlog).Return(nil)
		thirdMock := blogRepoTest.Mock.On("GetCurrentUserSlugs", slug, userID).Return([]entities.Blog{}, nil)

		err := blogServiceTest.ChangeBlogPublish(&payload, userID, false)

		assert.Nil(t, err)

		t.Cleanup(func() {
			// Cleanup mocking
			firstMock.Unset()
			secondMock.Unset()
			thirdMock.Unset()
		})
	})

	t.Run("Should return error if blog not found", func(t *testing.T) {
		slug := "example-of-title"
		userID := "example-of-wrong-id"
		payload := inputs.BlogIDInput{
			ID: "example-of-id",
		}

		unpublishedBlog := &entities.Blog{
			ID:        payload.ID,
			Title:     "Example of Title",
			Published: true,
		}

		firstMock := blogRepoTest.Mock.On("GetByIDAndAuthor", payload.ID, userID).Return(nil, errors.New("Record not found"))
		secondMock := blogRepoTest.Mock.On("SaveBlog", unpublishedBlog).Return(nil)
		thirdMock := blogRepoTest.Mock.On("GetCurrentUserSlugs", slug, userID).Return([]entities.Blog{}, nil)

		err := blogServiceTest.ChangeBlogPublish(&payload, userID, true)

		assert.NotNil(t, err)
		assert.Error(t, err)

		t.Cleanup(func() {
			// Cleanup mocking
			firstMock.Unset()
			secondMock.Unset()
			thirdMock.Unset()
		})
	})

	t.Run("Should return error if failed to be saved", func(t *testing.T) {
		slug := "example-of-title"
		userID := "example-of-id"
		payload := inputs.BlogIDInput{
			ID: "example-of-id",
		}

		unpublishedBlog := &entities.Blog{
			ID:        payload.ID,
			Title:     "Example of Title",
			Published: true,
		}

		firstMock := blogRepoTest.Mock.On("GetByIDAndAuthor", payload.ID, userID).Return(unpublishedBlog, nil)
		secondMock := blogRepoTest.Mock.On("SaveBlog", unpublishedBlog).Return(errors.New("Error saving blog"))
		thirdMock := blogRepoTest.Mock.On("GetCurrentUserSlugs", slug, userID).Return([]entities.Blog{}, nil)

		err := blogServiceTest.ChangeBlogPublish(&payload, userID, true)

		assert.NotNil(t, err)
		assert.Error(t, err)

		t.Cleanup(func() {
			// Cleanup mocking
			firstMock.Unset()
			secondMock.Unset()
			thirdMock.Unset()
		})
	})
}
