package main

import (
	"dot-crud-redis-go-api/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.Handle("/posts", controllers.GetPosts()).Methods("GET")
	router.Handle("/posts/{id}", controllers.GetPost()).Methods("GET")
	router.Handle("/posts", controllers.CreatePost()).Methods("POST")
	router.Handle("/posts/{postId}", controllers.UpdatePost()).Methods("PUT")
	router.Handle("/posts/{postId}", controllers.UpdatePost()).Methods("PATCH")
	router.Handle("/posts/{postId}", controllers.DeletePost()).Methods("DELETE")

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
