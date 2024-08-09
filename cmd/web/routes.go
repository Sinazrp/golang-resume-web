package main

import (
	"github.com/bmizerany/pat"
	"golang-resume-web/internal/handlers"
	"net/http"
)

// routes handles application routes
func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/contact", http.HandlerFunc(handlers.Contact))
	mux.Get("/projects", http.HandlerFunc(handlers.Projects))
	mux.Get("/resume", http.HandlerFunc(handlers.Resume))

	return mux
}
