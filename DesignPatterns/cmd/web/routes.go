package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)                 // recover after bad request
	mux.Use(middleware.Timeout(60 * time.Second)) // request will max open for 60 sec

	mux.Get("/", app.ShowHome)
	return mux
}
