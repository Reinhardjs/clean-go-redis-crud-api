package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa kabar?")
}

func main() {
	http.HandleFunc("/", home)

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
