package routes

import (
	"github.com/gorilla/mux"
	"github.com/prasaddls/go-bookmicroservice/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	// Create book
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")

	// Get all books
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")

	// Get a specific book by ID
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")

	// Update a specific book by ID
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")

	// Delete a specific book by ID
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
