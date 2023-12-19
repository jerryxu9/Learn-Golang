package router

import (
	"fmt"
	"mongoapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	fmt.Println("Starting API")
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
