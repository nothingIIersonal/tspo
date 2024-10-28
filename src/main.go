package main

import (
	"log"
	"net/http"
	_ "proj/docs"
	"proj/services/books"
	"proj/services/cities"
	"proj/services/genres"
	"proj/services/publishingHouses"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

var r *mux.Router

func init() {
	r = mux.NewRouter()

	r.HandleFunc("/cities", cities.GetCities).Methods(http.MethodGet)
	r.HandleFunc("/publishingHouses", publishingHouses.GetPublishingHouses).Methods(http.MethodGet)
	r.HandleFunc("/genres", genres.GetGenres).Methods(http.MethodGet)
	r.HandleFunc("/books", books.GetBooks).Methods(http.MethodGet)
	r.HandleFunc("/cities/", cities.GetCities).Methods(http.MethodGet)
	r.HandleFunc("/publishingHouses/", publishingHouses.GetPublishingHouses).Methods(http.MethodGet)
	r.HandleFunc("/genres/", genres.GetGenres).Methods(http.MethodGet)
	r.HandleFunc("/books/", books.GetBooks).Methods(http.MethodGet)

	r.HandleFunc("/cities/{id}", cities.GetCity).Methods(http.MethodGet)
	r.HandleFunc("/publishingHouses/{id}", publishingHouses.GetPublishingHouse).Methods(http.MethodGet)
	r.HandleFunc("/genres/{id}", genres.GetGenre).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", books.GetBook).Methods(http.MethodGet)

	r.HandleFunc("/cities", cities.CreateCity).Methods(http.MethodPost)
	r.HandleFunc("/publishingHouses", publishingHouses.CreatePublishingHouse).Methods(http.MethodPost)
	r.HandleFunc("/genres", genres.CreateGenre).Methods(http.MethodPost)
	r.HandleFunc("/books", books.CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/cities/", cities.CreateCity).Methods(http.MethodPost)
	r.HandleFunc("/publishingHouses/", publishingHouses.CreatePublishingHouse).Methods(http.MethodPost)
	r.HandleFunc("/genres/", genres.CreateGenre).Methods(http.MethodPost)
	r.HandleFunc("/books/", books.CreateBook).Methods(http.MethodPost)

	r.HandleFunc("/cities/{id}", cities.UpdateCity).Methods(http.MethodPut)
	r.HandleFunc("/publishingHouses/{id}", publishingHouses.UpdatePublishingHouse).Methods(http.MethodPut)
	r.HandleFunc("/genres/{id}", genres.UpdateGenre).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", books.UpdateBook).Methods(http.MethodPut)

	r.HandleFunc("/cities/{id}", cities.DeleteCity).Methods(http.MethodDelete)
	r.HandleFunc("/publishingHouses/{id}", publishingHouses.DeletePublishingHouse).Methods(http.MethodDelete)
	r.HandleFunc("/genres/{id}", genres.DeleteGenre).Methods(http.MethodDelete)
	r.HandleFunc("/books/{id}", books.DeleteBook).Methods(http.MethodDelete)

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler()).Methods(http.MethodGet)
}

func main() {
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
