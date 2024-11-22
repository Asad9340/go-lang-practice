package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/", HomeHandler)
	r.Get("/about", AboutHandler)
	r.Get("/contact", ContactHandler)
	return r
}
