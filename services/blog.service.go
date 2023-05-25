package services

import (
	"resdev-server/db"
	"resdev-server/entities"
)

func GetAllBlogs() ([]entities.Blog, error) {
	var blogs []entities.Blog
	result := db.DB.Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}

	return blogs, nil
}
