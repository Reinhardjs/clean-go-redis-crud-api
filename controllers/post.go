package controllers

import (
	"context"
	"dot-crud-redis-go-api/models"
	"dot-crud-redis-go-api/responses"
	"dot-crud-redis-go-api/usecases"
	"dot-crud-redis-go-api/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type PostController struct {
	postUsecase usecases.PostUsecase
}

func CreatePostController(postUsecase usecases.PostUsecase) PostController {
	return PostController{postUsecase}
}

func (e *PostController) GetPosts() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		posts, err := e.postUsecase.ReadAll()

		if err != nil {
			return err
		}

		response := responses.FineResponse{Status: http.StatusOK, Message: "success", Data: posts}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func (e *PostController) GetPost() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			return utils.NewHTTPError(err, 400, "Invalid post id")
		}

		rw.Header().Add("Content-Type", "application/json")

		post, err := e.postUsecase.ReadById(int(id))

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.NewHTTPError(err, 404, "record not found")
			} else {
				return err
			}
		}

		response := responses.FineResponse{Status: http.StatusOK, Message: "success", Data: post}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func (e *PostController) CreatePost() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		post := &models.Post{}
		decodeError := json.NewDecoder(r.Body).Decode(post)

		if decodeError != nil {
			return utils.NewHTTPError(nil, 400, "Invalid request body format")
		}

		if message, ok := post.Validate(); !ok {
			return utils.NewHTTPError(nil, 400, message)
		}

		result, err := e.postUsecase.Create(post)

		if err != nil {
			return err
		}

		response := responses.FineResponse{Status: http.StatusCreated, Message: "success", Data: result}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func (e *PostController) UpdatePost() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		params := mux.Vars(r)
		postId, err := strconv.Atoi(params["postId"])

		if err != nil {
			return utils.NewHTTPError(nil, 400, "Invalid post id")
		}

		post := &models.Post{}
		post.ID = postId
		decodeError := json.NewDecoder(r.Body).Decode(post)
		if decodeError != nil {
			return utils.NewHTTPError(nil, 400, "Invalid request body format")
		}

		if r.Method == "PUT" {
			if message, ok := post.Validate(); !ok {
				return utils.NewHTTPError(nil, 400, message)
			}
		}

		oldPost, oldPostErr := e.postUsecase.ReadById(int(postId))

		if oldPostErr != nil {
			if errors.Is(oldPostErr, gorm.ErrRecordNotFound) {
				return utils.NewHTTPError(oldPostErr, 404, "record not found")
			} else {
				return oldPostErr
			}
		}

		updatedPost, updatePostErr := e.postUsecase.Update(post.ID, post)
		updatedPost.CreatedAt = oldPost.CreatedAt

		if updatePostErr != nil {
			return updatePostErr
		}

		if err != nil {
			return err
		}

		response := responses.FineResponse{Status: http.StatusOK, Message: "success", Data: updatedPost}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func (e *PostController) DeletePost() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["postId"])

		if err != nil {
			return utils.NewHTTPError(err, 400, "Invalid post id")
		}

		rw.Header().Add("Content-Type", "application/json")

		// Check for existing record
		_, existingPostErr := e.postUsecase.ReadById(int(id))
		if existingPostErr != nil {
			if errors.Is(existingPostErr, gorm.ErrRecordNotFound) {
				return utils.NewHTTPError(existingPostErr, 404, "record not found")
			} else {
				return existingPostErr
			}
		}

		post, err := e.postUsecase.Delete(id)

		if err != nil {
			return err
		}

		response := responses.FineResponse{Status: http.StatusOK, Message: "success", Data: post}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}
