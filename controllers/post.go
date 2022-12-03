package controllers

import (
	"context"
	"dot-crud-redis-go-api/configs"
	"dot-crud-redis-go-api/models"
	"dot-crud-redis-go-api/responses"
	"dot-crud-redis-go-api/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetPosts() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		posts, err := models.GetPosts()

		if err != nil {
			return err
		}

		response := responses.FineResponse{Status: http.StatusOK, Message: "success", Data: posts}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func GetPost() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		redisClient := configs.GetRedis()

		defer cancel()

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			return err
		}

		rw.Header().Add("Content-Type", "application/json")

		post, err := models.GetPost(uint(id))

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.NewHTTPError(err, 404, "record not found")
			} else {
				return err
			}
		}

		// Get JSON blob from Redis
		redisResult, err := redisClient.Do("GET", "post:"+strconv.Itoa(post.ID))

		if err != nil {
			return utils.NewHTTPError(err, 500, "Failed getting data from redis")
		}

		if redisResult == nil {
			postJSON, err := json.Marshal(post)
			if err != nil {
				return err
			}

			// Save JSON blob to Redis
			_, saveRedisError := redisClient.Do("SET", "post:"+strconv.Itoa(post.ID), postJSON)

			if saveRedisError != nil {
				return utils.NewHTTPError(saveRedisError, 500, "Failed saving data to redis")
			}
		} else {
			json.Unmarshal(redisResult.([]byte), &post)
		}

		response := responses.FineResponse{Status: http.StatusOK, Message: "success", Data: post}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func CreatePost() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		redisClient := configs.GetRedis()

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

		result, err := post.Create()

		if err != nil {
			return err
		}

		postJSON, err := json.Marshal(result)
		if err != nil {
			return err
		}

		// Save JSON blob to Redis
		reply, err := redisClient.Do("SET", "post:"+strconv.Itoa(result.ID), postJSON)

		if reply != "OK" {
			return utils.NewHTTPError(nil, 500, "Failed saving data to redis")
		}

		response := responses.FineResponse{Status: http.StatusCreated, Message: "success", Data: result}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func UpdatePost() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		redisClient := configs.GetRedis()

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

		_, oldPostErr := models.GetPost(uint(postId))

		if oldPostErr != nil {
			if errors.Is(oldPostErr, gorm.ErrRecordNotFound) {
				return utils.NewHTTPError(oldPostErr, 404, "record not found")
			} else {
				return oldPostErr
			}
		}

		_, updatePostErr := post.Update()

		if updatePostErr != nil {
			return updatePostErr
		}

		updatedPost, err := models.GetPost(uint(postId))

		if err != nil {
			return err
		}

		// Delete JSON blob from Redis
		_, redisDeleteErr := redisClient.Do("DEL", "post:"+strconv.Itoa(updatedPost.ID))

		if redisDeleteErr != nil {
			return utils.NewHTTPError(err, 500, "Failed deleting data from redis")
		}

		response := responses.FineResponse{Status: http.StatusOK, Message: "success", Data: updatedPost}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func DeletePost() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		redisClient := configs.GetRedis()
		defer cancel()

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["postId"])

		if err != nil {
			return err
		}

		rw.Header().Add("Content-Type", "application/json")

		// Check for existing record
		_, existingPostErr := models.GetPost(uint(id))
		if existingPostErr != nil {
			if errors.Is(existingPostErr, gorm.ErrRecordNotFound) {
				return utils.NewHTTPError(existingPostErr, 404, "record not found")
			} else {
				return existingPostErr
			}
		}

		post, err := models.Delete(id)

		if err != nil {
			return err
		}

		// Delete JSON blob from Redis
		_, redisDeleteErr := redisClient.Do("DEL", "post:"+strconv.Itoa(id))

		if redisDeleteErr != nil {
			return utils.NewHTTPError(err, 500, "Failed deleting data from redis")
		}

		response := responses.FineResponse{Status: http.StatusOK, Message: "success", Data: post}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}
