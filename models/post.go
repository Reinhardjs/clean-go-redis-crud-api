package models

import (
	"dot-crud-redis-go-api/configs"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (post *Post) Validate() (string, bool) {

	if post.Title == "" {
		return "Title should be on the payload", false
	}

	if post.Description == "" {
		return "Description should be on the payload", false
	}

	return "Payload is valid", true
}

func GetPosts() ([]*Post, error) {
	DB := configs.GetDB()

	posts := make([]*Post, 0)
	err := DB.Table("posts").Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("DB error : %v", err)
	}

	return posts, nil
}

func GetPost(id uint) (*Post, error) {
	DB := configs.GetDB()

	post := &Post{}
	err := DB.Table("posts").Where("id = ?", id).First(post).Error

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (post *Post) Create() (interface{}, error) {
	DB := configs.GetDB()

	result := DB.Create(post)

	if result.Error != nil {
		return nil, fmt.Errorf("DB error : %v", result.Error)
	}

	return result.Value, nil
}
