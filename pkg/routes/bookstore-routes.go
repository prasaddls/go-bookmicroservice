package routes

import (
	"github.com/gorilla/mux"
	"github.com/prasaddls/go-bookmicroservice/controllers"
)

var RegisterBookStoreRoutes = func(route *mux.Route) {
	route.HandlerFunc("/books", controllers.CreateBook).Methods("POST")
	route.HandlerFunc("/books", controllers.GetBooks).Methods("GET")
	route.HandlerFunc("/books/{id}", controllers.GetBook).Methods("GET")
	route.HandlerFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	route.HandlerFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
