package main

import (
	"log"
	"net/http"

	"github.com/aayanmtn/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) // Use RegisterBookStoreRoutes instead of RegisterBookStoreRouter
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
