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
	err := e.DB.Save(&post).Error
	if err != nil {
		fmt.Printf("[PostRepoImpl.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}
	return post, nil
}

func (e *PostRepoImpl) ReadAll() (*[]models.Post, error) {
	var posts []models.Post
	err := e.DB.Find(&posts).Error
	if err != nil {
		fmt.Printf("[PostRepoImpl.ReadAll] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data")
	}
	return &posts, nil
}

func (e *PostRepoImpl) ReadById(id int) (*models.Post, error) {
	var post = models.Post{}
	err := e.DB.Table("post").Where("id = ?", id).First(&post).Error
	if err != nil {
		fmt.Printf("[PostRepoImpl.ReadById] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}
	return &post, nil
}

func (e *PostRepoImpl) Update(id int, post *models.Post) (*models.Post, error) {
	var upPost = models.Post{}
	err := e.DB.Table("post").Where("id = ?", id).First(&upPost).Update(&post).Error
	if err != nil {
		fmt.Printf("[PostRepoImpl.Update] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data")
	}
	return &upPost, nil
}

func (e *PostRepoImpl) Delete(id int) error {
	var post = models.Post{}
	err := e.DB.Table("post").Where("id = ?", id).First(&post).Delete(&post).Error
	if err != nil {
		fmt.Printf("[PostRepoImpl.Delete] error execute query %v \n", err)
		return fmt.Errorf("id is not exsis")
	}
	return nil
}
