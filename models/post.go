package models

import (
	"dot-crud-redis-go-api/configs"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetPosts() (error, []*Post) {

	DB := configs.GetDB()

	posts := make([]*Post, 0)
	err := DB.Table("posts").Find(&posts).Error
	if err != nil {
		return err, nil
	}

	return nil, posts
}
