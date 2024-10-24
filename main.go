// main.go
package main

import (
	"DevRank/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/developer", handlers.GetDeveloper).Methods("GET")

	http.ListenAndServe(":8080", r)
}
