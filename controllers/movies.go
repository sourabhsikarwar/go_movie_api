package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sourabhsikarwar/go_movie_api/helpers"
	"github.com/sourabhsikarwar/go_movie_api/models"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := helpers.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	res, err := helpers.InsertMovie(movie)

	jsonEncoder := json.NewEncoder(w)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonEncoder.Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)

	// using interface for different type of datatypes
	jsonEncoder.Encode(map[string]interface{}{
		"message": res,
		"movie":   movie,
	})
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	movieId := params["id"]

	helpers.UpdateMovie(movieId)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Movie updated Successfully",
		"id":      movieId,
	})
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	movieId := params["id"]

	helpers.DeleteMovie(movieId)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Movie deleted Successfully",
		"id":      movieId,
	})
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	helpers.DeleteAllMovies()

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "All Movies deleted Successfully",
	})
}
