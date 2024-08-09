package main

import (
	"github.com/go-chi/chi/v5"
	"golang-resume-web/internal/handlers"
	"net/http"
)

// routes handles application routes
func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/contact", handlers.Contact)
	mux.Get("/projects", handlers.Projects)
	mux.Get("/resume", handlers.Resume)
	fileServer := http.FileServer(http.Dir("./statics"))
	mux.Handle("/statics/*", http.StripPrefix("/statics/", fileServer))

	return mux
}
