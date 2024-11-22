package main

import (
	handler "go-crud-mysql/test/handler/rootHandler"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// Create an instance of the RootHandler
	rootHandler := &handler.RootHandler{}


	r.Get("/", rootHandler.ServeHTTP)
	r.Get("/about", rootHandler.AboutData)

	log.Fatal(http.ListenAndServe(":5050", r))
}
