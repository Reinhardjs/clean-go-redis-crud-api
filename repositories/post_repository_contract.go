package repositories

import (
	"dot-crud-redis-go-api/models"
)

type PostRepo interface {
	Create(post *models.Post) (*models.Post, error)
	ReadAll() (*[]models.Post, error)
	ReadById(id int) (*models.Post, error)
	Update(id int, post *models.Post) (*models.Post, error)
	Delete(id int) (map[string]interface{}, error)
}
