package implementations

import (
	"dot-crud-redis-go-api/models"
	"dot-crud-redis-go-api/repositories"
	"dot-crud-redis-go-api/usecases"
)

type PostUsecaseImpl struct {
	postRepo repositories.PostRepo
}

func CreatePostUsecase(postRepo repositories.PostRepo) usecases.PostUsecase {
	return &PostUsecaseImpl{postRepo}
}

func (e *PostUsecaseImpl) Create(post *models.Post) (*models.Post, error) {
	return e.postRepo.Create(post)
}

func (e *PostUsecaseImpl) ReadAll() (*[]models.Post, error) {
	return e.postRepo.ReadAll()
}

func (e *PostUsecaseImpl) ReadById(id int) (*models.Post, error) {
	return e.postRepo.ReadById(id)
}

func (e *PostUsecaseImpl) Update(id int, post *models.Post) (*models.Post, error) {
	return e.postRepo.Update(id, post)
}

func (e *PostUsecaseImpl) Delete(id int) (map[string]interface{}, error) {
	return e.postRepo.Delete(id)
}
