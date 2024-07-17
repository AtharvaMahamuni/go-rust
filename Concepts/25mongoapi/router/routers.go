package router

import (
	"github.com/atharvamahamuni/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAll).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteAll).Methods("DELETE")

	return router
}
