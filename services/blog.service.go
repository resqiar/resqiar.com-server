package services

import (
	"errors"
	"resdev-server/db"
	"resdev-server/entities"
	"resdev-server/inputs"
	"time"

	"gorm.io/gorm/clause"
)

// GetAllBlogs retrieves a list of SafeBlogAuthor entities from the database.
// If onlyPublished is true, it retrieves only the published blogs, otherwise everything.
// It returns the list of blogs and any error encountered during the process.
func GetAllBlogs(onlyPublished bool) ([]entities.SafeBlogAuthor, error) {
	var blogs []entities.SafeBlogAuthor

	query := db.DB.Model(&entities.Blog{})

	// Define SELECT and JOIN for database query operations
	BLOG_SELECT_SQL := "blogs.id, blogs.created_at, blogs.updated_at, blogs.published_at, blogs.title, blogs.summary, blogs.cover_url, blogs.author_id, "
	AUTHOR_SELECT_SQL := "users.id AS author_id, users.username AS author_username, users.created_at AS author_created_at, users.bio AS author_bio, users.picture_url AS author_picture_url"
	JOIN_SQL := "JOIN users ON blogs.author_id = users.id"

	// Add the SELECT and JOIN statements to the query
	query = query.Select(BLOG_SELECT_SQL + AUTHOR_SELECT_SQL).Joins(JOIN_SQL)

	// If onlyPublished is true, add a condition to retrieve only published blogs
	if onlyPublished {
		query.Where("blogs.published = ?", true)
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
		}

		// Scan the rows and bind them into the temp struct
		err := db.DB.ScanRows(rows, &temp)
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
			},
		}

		// append back to blogs array
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

// GetPublishedBlogDetail retrieves a single SafeBlogAuthor entity from the database
// based on the provided blogID.
// It returns the retrieved blog and any error encountered during the process.
// If no blog is found or an error occurs, it returns an appropriate error.
func GetBlogDetail(blogID string, published bool) (*entities.SafeBlogAuthor, error) {
	var blog entities.SafeBlogAuthor

	// Define SELECT and JOIN for database query operations
	BLOG_SELECT_SQL := "blogs.id, blogs.created_at, blogs.updated_at, blogs.published_at, blogs.title, blogs.summary, blogs.content, blogs.cover_url, blogs.author_id, "
	AUTHOR_SELECT_SQL := "users.id AS author_id, users.username AS author_username, users.created_at AS author_created_at, users.bio AS author_bio, users.picture_url AS author_picture_url"
	JOIN_SQL := "JOIN users ON blogs.author_id = users.id"

	// Execute the query and retrieve the rows
	result := db.DB.Model(&entities.Blog{}).
		Select(BLOG_SELECT_SQL + AUTHOR_SELECT_SQL).
		Joins(JOIN_SQL)

	if published {
		result.Where("blogs.id = ? AND published = ?", blogID, true)
	} else {
		result.Where("blogs.id = ?", blogID)
	}

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
	}

	// Scan the rows and bind them into the temp struct
	err = db.DB.ScanRows(rows, &temp)
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
		},
	}

	return &blog, nil
}

func CreateBlog(payload *inputs.CreateBlogInput, userID string) (*entities.Blog, error) {
	newBlog := entities.Blog{
		Title:   payload.Title,
		Summary: payload.Summary,
		Content: payload.Content,

		// when creating blog, always set published to false.
		// although the default value in database is false,
		// we still want to ensure the published value here-
		// is NOT coming from the payload, but rather hardcoded.
		Published: false,

		CoverURL: payload.CoverURL,
		AuthorID: userID,
	}

	result := db.DB.Clauses(clause.Returning{}).Create(&newBlog)
	if result.Error != nil {
		return nil, result.Error
	}

	return &newBlog, nil
}

func EditBlog(payload *inputs.UpdateBlogInput, userID string) (*inputs.SafeUpdateBlogInput, error) {
	var blog entities.Blog
	error := db.DB.First(&blog, "id = ? AND author_id = ?", payload.ID, userID).Error
	if error != nil {
		return nil, error
	}

	safe := &inputs.SafeUpdateBlogInput{
		Title:    payload.Title,
		Summary:  payload.Summary,
		Content:  payload.Content,
		CoverURL: payload.CoverURL,
	}

	if err := db.DB.Model(&entities.Blog{}).Where("id = ?", blog.ID).Updates(&safe).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func GetCurrentUserBlogs(userID string) (*[]entities.Blog, error) {
	var blogs []entities.Blog
	result := db.DB.Omit("content").Find(&blogs, "author_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &blogs, nil
}

func ChangeBlogPublish(payload *inputs.BlogIDInput, userID string, publishState bool) (*entities.Blog, error) {
	var blog entities.Blog
	result := db.DB.First(&blog, "ID = ? AND author_id = ?", payload.ID, userID)
	if result.Error != nil {
		return nil, result.Error
	}

	// update published state based on given param
	blog.Published = publishState

	// if publish state is true
	// then we need to update the PublishedAt field
	if publishState {
		blog.PublishedAt = time.Now()
	} else {
		// otherwise, reset the PublishedAt field to "January 1, year 1, 00:00:00 UTC" (invalid date)
		blog.PublishedAt = time.Time{}
	}

	// save back to the database
	db.DB.Save(&blog)

	return &blog, nil
}
