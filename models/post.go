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

func GetPosts() ([]*Post, error) {

	DB := configs.GetDB()

	posts := make([]*Post, 0)
	err := DB.Table("posts").Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("DB error : %v", err)
	}

	return posts, nil
}
