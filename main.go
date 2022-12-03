package main

import (
	"dot-crud-redis-go-api/controllers"
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa kabar?")
}

func main() {

	http.HandleFunc("/", home)
	http.Handle("/posts", controllers.GetPosts())

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
