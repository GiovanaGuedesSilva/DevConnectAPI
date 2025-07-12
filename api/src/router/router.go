package router

import "github.com/gorilla/mux"

func Generate() *mux.Router {
	router := mux.NewRouter()

	// Define your routes here
	// Example: router.HandleFunc("/api/v1/resource", handlerFunction).Methods("GET")

	return router
}
