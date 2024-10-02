package router

import (
	"github.com/gorilla/mux"
	"github.com/sourabhsikarwar/go_movie_api/controllers"
)

// return pointer so that the router reference can be accessed
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/delete-all-movies", controllers.DeleteAllMovies).Methods("DELETE")

	return router
}
