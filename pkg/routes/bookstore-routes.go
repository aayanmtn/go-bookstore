package routes

import (
	"github.com/gorilla/mux"
)

var RegisterBookStoreRouter = func(router *mux.Router) {
	// Define the route for the book store API
	router.HandleFunc("/book", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.GetBookByID).Methods("GET")
	router.HandleFunc("/book", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE")
}
