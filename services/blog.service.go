package services

import (
	"resdev-server/db"
	"resdev-server/entities"
	"resdev-server/inputs"

	"gorm.io/gorm/clause"
)

func GetAllBlogs(onlyPublished bool) ([]entities.Blog, error) {
	var blogs []entities.Blog
	result := db.DB

	if onlyPublished {
		result = result.Omit("content").Find(&blogs, "published = ?", true) // send only published blogs
	} else {
		result = result.Omit("content").Find(&blogs)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return blogs, nil
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

func GetCurrentUserBlogs(userID string) (*[]entities.Blog, error) {
	var blogs []entities.Blog
	result := db.DB.Find(&blogs, "author_id = ?", userID)
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

	// save back to the database
	db.DB.Save(&blog)

	return &blog, nil
}
