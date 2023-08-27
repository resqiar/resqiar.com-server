package repositories

import (
	"errors"
	"fmt"
	"time"

	"resqiar.com-server/entities"
	"resqiar.com-server/inputs"
	"resqiar.com-server/types"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BlogRepository interface {
	GetBlogs(onlyPublished bool, desc bool, username string) ([]entities.SafeBlogAuthor, error)
	GetBlog(opts *types.GetBlogOpts) (*entities.SafeBlogAuthor, error)
	CreateBlog(input *entities.Blog) (*entities.Blog, error)
	UpdateBlog(blogID string, safe *inputs.SafeUpdateBlogInput) error
	GetByIDAndAuthor(blogID string, userID string) (*entities.Blog, error)
	GetCurrentUserBlogs(userID string, desc bool) ([]entities.Blog, error)
	GetCurrentUserSlugs(slug string, userID string) ([]entities.Blog, error)
	GetCurrentUserBlog(blogID string, userID string) (*entities.Blog, error)
	SaveBlog(blog *entities.Blog) error
}

type BlogRepoImpl struct {
	db *gorm.DB
}

func InitBlogRepo(db *gorm.DB) BlogRepository {
	return &BlogRepoImpl{
		db: db,
	}
}

func (repo *BlogRepoImpl) GetBlogs(onlyPublished bool, orderDesc bool, username string) ([]entities.SafeBlogAuthor, error) {
	var blogs []entities.SafeBlogAuthor

	query := repo.db.Model(&entities.Blog{})

	// Define SELECT and JOIN for database query operations
	BLOG_SELECT_SQL := "blogs.id, blogs.slug, blogs.created_at, blogs.updated_at, blogs.published_at, blogs.title, blogs.summary, blogs.cover_url, blogs.author_id, blogs.prev, blogs.next, "
	AUTHOR_SELECT_SQL := "users.id AS author_id, users.username AS author_username, users.created_at AS author_created_at, users.bio AS author_bio, users.picture_url AS author_picture_url, users.is_tester AS author_is_tester"
	JOIN_SQL := "JOIN users ON blogs.author_id = users.id"

	// Add the SELECT and JOIN statements to the query
	query = query.Select(BLOG_SELECT_SQL + AUTHOR_SELECT_SQL).Joins(JOIN_SQL)

	// If onlyPublished is true, add a condition to retrieve only published blogs
	if onlyPublished {
		query.Where("blogs.published = ?", true)
	}

	if username != "" {
		query.Where("users.username", username)
	}

	if orderDesc {
		query.Order("updated_at DESC")
	} else {
		query.Order("updated_at ASC")
	}

	// Execute the query and retrieve the rows
	rows, err := query.Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close() // close the row when the function out

	// Loop through the result rows and populate the blogs slice
	for rows.Next() {
		var temp struct {
			entities.SafeBlog
			AuthorID         string    `gorm:"column:author_id"`
			AuthorUsername   string    `gorm:"column:author_username"`
			AuthorCreatedAt  time.Time `gorm:"column:author_created_at"`
			AuthorBio        string    `gorm:"column:author_bio"`
			AuthorPictureURL string    `gorm:"column:author_picture_url"`
			AuthorIsTester   bool      `gorm:"column:author_is_tester"`
		}

		// Scan the rows and bind them into the temp struct
		err := repo.db.ScanRows(rows, &temp)
		if err != nil {
			return nil, err
		}

		// Create a SafeBlogAuthor entity and append it to the blogs slice
		blog := entities.SafeBlogAuthor{
			SafeBlog: temp.SafeBlog,
			Author: entities.SafeUser{
				ID:         temp.AuthorID,
				Username:   temp.AuthorUsername,
				CreatedAt:  temp.AuthorCreatedAt,
				Bio:        temp.AuthorBio,
				PictureURL: temp.AuthorPictureURL,
				IsTester:   temp.AuthorIsTester,
			},
		}

		// append back to blogs array
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (repo *BlogRepoImpl) GetBlog(opts *types.GetBlogOpts) (*entities.SafeBlogAuthor, error) {
	var blog entities.SafeBlogAuthor
	var condition string
	var args []interface{}

	var CONTENT_SELECT_SQL string

	if opts.IncludeContent {
		CONTENT_SELECT_SQL = "blogs.content, "
	}

	// Define SELECT and JOIN for database query operations
	BLOG_SELECT_SQL := "blogs.id, blogs.slug, blogs.created_at, blogs.updated_at, blogs.published_at, blogs.title, blogs.summary, blogs.cover_url, blogs.author_id, blogs.prev, blogs.next, "
	AUTHOR_SELECT_SQL := "users.id AS author_id, users.username AS author_username, users.created_at AS author_created_at, users.bio AS author_bio, users.picture_url AS author_picture_url, users.is_tester AS author_is_tester"
	JOIN_SQL := "JOIN users ON blogs.author_id = users.id"

	result := repo.db.Model(&entities.Blog{}).
		Select(BLOG_SELECT_SQL + CONTENT_SELECT_SQL + AUTHOR_SELECT_SQL).
		Joins(JOIN_SQL)

	if opts.UseID != "" {
		condition = "blogs.ID = ?" // use ID instead of slug
		args = []interface{}{opts.UseID}
	} else {
		condition = "blogs.slug = ? AND users.username = ?" // use slug and username
		args = []interface{}{opts.BlogSlug, opts.BlogAuthor}
	}

	if opts.Published {
		condition += " AND published = ?"
		args = append(args, true)
	}

	result.Where(condition, args...)

	// Check for any errors during query execution
	if result.Error != nil {
		return nil, result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// check if no rows are returned
	if !rows.Next() {
		// return nil when no rows are found
		return nil, errors.New("404")
	}

	var temp struct {
		entities.SafeBlog
		AuthorID         string    `gorm:"column:author_id"`
		AuthorUsername   string    `gorm:"column:author_username"`
		AuthorCreatedAt  time.Time `gorm:"column:author_created_at"`
		AuthorBio        string    `gorm:"column:author_bio"`
		AuthorPictureURL string    `gorm:"column:author_picture_url"`
		AuthorIsTester   bool      `gorm:"column:author_is_tester"`
	}

	// Scan the rows and bind them into the temp struct
	err = repo.db.ScanRows(rows, &temp)
	if err != nil {
		return nil, err
	}

	blog = entities.SafeBlogAuthor{
		SafeBlog: temp.SafeBlog,
		Author: entities.SafeUser{
			ID:         temp.AuthorID,
			Username:   temp.AuthorUsername,
			CreatedAt:  temp.AuthorCreatedAt,
			Bio:        temp.AuthorBio,
			PictureURL: temp.AuthorPictureURL,
			IsTester:   temp.AuthorIsTester,
		},
	}

	return &blog, nil
}

func (repo *BlogRepoImpl) CreateBlog(input *entities.Blog) (*entities.Blog, error) {
	newBlog := input

	result := repo.db.Clauses(clause.Returning{}).Create(newBlog)
	if result.Error != nil {
		return nil, result.Error
	}

	return newBlog, nil
}

func (repo *BlogRepoImpl) GetByIDAndAuthor(blogID string, userID string) (*entities.Blog, error) {
	var blog entities.Blog

	err := repo.db.First(&blog, "id = ? AND author_id = ?", blogID, userID).Error
	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (repo *BlogRepoImpl) UpdateBlog(blogID string, safe *inputs.SafeUpdateBlogInput) error {
	if err := repo.db.Model(&entities.Blog{}).Where("id = ?", blogID).Updates(&safe).Error; err != nil {
		return err
	}

	return nil
}

func (repo *BlogRepoImpl) GetCurrentUserBlogs(userID string, desc bool) ([]entities.Blog, error) {
	var blogs []entities.Blog

	// set default query order as desc (newest first)
	queryOrder := "DESC"

	if !desc {
		queryOrder = "ASC"
	}

	if err := repo.db.
		Omit("content").
		Order(fmt.Sprintf("updated_at %s", queryOrder)).
		Find(&blogs, "author_id = ?", userID).
		Error; err != nil {
		return nil, err
	}

	return blogs, nil
}

func (repo *BlogRepoImpl) GetCurrentUserSlugs(slug string, userID string) ([]entities.Blog, error) {
	var blogs []entities.Blog

	if err := repo.db.Find(&blogs, "slug = ? AND author_id = ?", slug, userID).Error; err != nil {
		return nil, err
	}

	return blogs, nil
}

func (repo *BlogRepoImpl) GetCurrentUserBlog(blogID string, userID string) (*entities.Blog, error) {
	var blog entities.Blog

	if err := repo.db.First(&blog, "id = ? AND author_id = ?", blogID, userID).Error; err != nil {
		return nil, err
	}

	return &blog, nil
}

func (repo *BlogRepoImpl) SaveBlog(blog *entities.Blog) error {
	if err := repo.db.Save(&blog).Error; err != nil {
		return err
	}

	return nil
}
