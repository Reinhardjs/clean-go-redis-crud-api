package implementations

import (
	"dot-crud-redis-go-api/models"
	"dot-crud-redis-go-api/repositories"
	"fmt"

	"github.com/jinzhu/gorm"
)

type PostRepoImpl struct {
	DB *gorm.DB
}

func CreatePostRepo(DB *gorm.DB) repositories.PostRepo {
	return &PostRepoImpl{DB}
}

func (e *PostRepoImpl) Create(post *models.Post) (*models.Post, error) {
	result := e.DB.Model(&models.Post{}).Create(post)

	if result.Error != nil {
		return &models.Post{}, fmt.Errorf("DB error : %v", result.Error)
	}

	return post, nil
}

func (e *PostRepoImpl) ReadAll() (*[]models.Post, error) {
	posts := make([]models.Post, 0)
	err := e.DB.Table("posts").Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("DB error : %v", err)
	}

	return &posts, nil
}

func (e *PostRepoImpl) ReadById(id int) (*models.Post, error) {
	post := &models.Post{}
	err := e.DB.Table("posts").Where("id = ?", id).First(post).Error

	fmt.Println(*post)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (e *PostRepoImpl) Update(id int, post *models.Post) (*models.Post, error) {
	updatedPost := &models.Post{}
	result := e.DB.Model(updatedPost).Where("id = ?", post.ID).Updates(models.Post{Title: post.Title, Description: post.Description})

	if result.Error != nil {
		return nil, fmt.Errorf("DB error : %v", result.Error)
	}

	return updatedPost, nil
}

func (e *PostRepoImpl) Delete(id int) (map[string]interface{}, error) {
	result := e.DB.Delete(&models.Post{}, id)

	return map[string]interface{}{
		"rows_affected": result.RowsAffected,
	}, nil
}
