package models

import (
	"dot-crud-redis-go-api/configs"
	"fmt"
	"time"
)

type Post struct {
	ID          int       `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
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

	fmt.Println(*post)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (post *Post) Create() (*Post, error) {
	DB := configs.GetDB()

	result := DB.Model(&Post{}).Create(post)

	if result.Error != nil {
		return &Post{}, fmt.Errorf("DB error : %v", result.Error)
	}

	return post, nil
}

func (post *Post) Update() (*Post, error) {
	DB := configs.GetDB()

	updatedPost := &Post{}
	result := DB.Model(updatedPost).Where("id = ?", post.ID).Updates(Post{Title: post.Title, Description: post.Description})

	if result.Error != nil {
		return nil, fmt.Errorf("DB error : %v", result.Error)
	}

	return updatedPost, nil
}

func Delete(postId int) (map[string]interface{}, error) {
	DB := configs.GetDB()

	result := DB.Delete(&Post{}, postId)

	return map[string]interface{}{
		"rows_affected": result.RowsAffected,
	}, nil
}
