package controllers

import (
	"context"
	"dot-crud-redis-go-api/models"
	"dot-crud-redis-go-api/responses"
	"encoding/json"
	"net/http"
	"time"
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
